package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

func getUrlBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return make([]byte, 0), err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

type CachedFunc func(string) (interface{}, error)

type CallResult struct {
	respBody interface{} //报文内容
	err      error       //错误
}

type FuncCache struct {
	cf           CachedFunc            //缓存的函数
	resultsMutex sync.RWMutex          //保护results
	results      map[string]CallResult //缓存的函数调用结果
}

func (fc *FuncCache) Get(url string) (interface{}, error) {
	fc.resultsMutex.RLock()
	result, ok := fc.results[url]
	fc.resultsMutex.RUnlock()
	if !ok {
		respBody, err := fc.cf(url)
		result = CallResult{respBody: respBody, err: err}
		fc.resultsMutex.Lock()
		if _, ok := fc.results[url]; !ok {
			fc.results[url] = result
		}
		fc.resultsMutex.Unlock()
	}
	return result.respBody, result.err
}

func NewFuncCache(cf CachedFunc) *FuncCache {
	var fc FuncCache
	fc.cf = cf
	fc.results = make(map[string]CallResult)
	return &fc
}

func main() {
	fc := NewFuncCache(getUrlBody)
	urls := []string{"http://baidu.com", "http://bing.com", "http://google.com", "http://baidu.com", "http://bing.com", "http://google.com"}
	var funcCallWG sync.WaitGroup
	for _, url := range urls {
		url := url
		funcCallWG.Add(1)
		go func() {
			defer funcCallWG.Done()
			start := time.Now()
			respBody, err := fc.Get(url)
			usedTime := time.Since(start).Seconds()
			if err == nil {
				fmt.Printf("url:%s, used time:%.2fs, resp length:%d\n", url, usedTime, len(respBody.([]byte)))
			} else {
				fmt.Printf("url:%s, used time:%.2fs, error:%s\n", url, usedTime, err.Error())
			}
		}()
	}
	funcCallWG.Wait()
	// url:http://baidu.com, used time:0.09s, error:Get "http://baidu.com": read tcp 192.168.1.13:3102->220.181.38.148:80: wsarecv: An existing connection was forcibly closed by the remote host.
	// url:http://baidu.com, used time:0.11s, resp length:81
	// url:http://bing.com, used time:0.55s, resp length:75947
	// url:http://bing.com, used time:0.57s, resp length:75947
	// url:http://google.com, used time:21.08s, error:Get "http://google.com": dial tcp 172.217.163.46:80: connectex: A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond.
	// url:http://google.com, used time:21.08s, error:Get "http://google.com": dial tcp 172.217.163.46:80: connectex: A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond.
}
