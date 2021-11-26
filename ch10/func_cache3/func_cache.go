package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

func getUrlBody(url string) (interface{}, error) {
	fmt.Printf("request http->%s\n", url)
	resp, err := http.Get(url)
	if err != nil {
		return make([]byte, 0), err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

type CachedFunc func(string) (interface{}, error)

type CallResult struct {
	ready    chan struct{}
	respBody interface{} //报文内容
	err      error       //错误
}

type FuncCache struct {
	cf           CachedFunc             //缓存的函数
	resultsMutex sync.Mutex             //保护results
	results      map[string]*CallResult //缓存的函数调用结果
}

func (fc *FuncCache) Get(url string) (interface{}, error) {
	fc.resultsMutex.Lock()
	result, ok := fc.results[url]
	if !ok {
		result = &CallResult{ready: make(chan struct{})}
		fc.results[url] = result
		fc.resultsMutex.Unlock()
		result.respBody, result.err = fc.cf(url)
		close(result.ready)
	} else {
		fc.resultsMutex.Unlock()
		<-result.ready
	}
	return result.respBody, result.err
}

func NewFuncCache(cf CachedFunc) *FuncCache {
	var fc FuncCache
	fc.cf = cf
	fc.results = make(map[string]*CallResult)
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
	// request http->http://baidu.com
	// request http->http://bing.com
	// request http->http://google.com
	// url:http://baidu.com, used time:0.10s, resp length:81
	// url:http://baidu.com, used time:0.10s, resp length:81
	// url:http://bing.com, used time:0.57s, resp length:75947
	// url:http://bing.com, used time:0.57s, resp length:75947
	// url:http://google.com, used time:21.08s, error:Get "http://google.com": dial tcp 172.217.163.46:80: connectex: A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond.
	// url:http://google.com, used time:21.08s, error:Get "http://google.com": dial tcp 172.217.163.46:80: connectex: A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond.
}
