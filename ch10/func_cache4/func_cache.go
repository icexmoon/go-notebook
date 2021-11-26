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
	ready    chan struct{} //结果是否已经准备好了
	respBody interface{}   //报文内容
	err      error         //错误
}

// 向FuncCache发起请求的结构体
type FCRequest struct {
	url      string          //请求的url
	respChan chan CallResult //返回结果的通道
}

type FuncCache struct {
	cf         CachedFunc             //缓存的函数
	results    map[string]*CallResult //缓存的函数调用结果
	requetChan chan FCRequest
}

func (fc *FuncCache) Get(url string) (interface{}, error) {
	respChan := make(chan CallResult)
	fc.requetChan <- FCRequest{url: url, respChan: respChan}
	result := <-respChan
	return result.respBody, result.err
}

//启动FC服务
func (fc *FuncCache) Service() {
	go func() {
		for fcr := range fc.requetChan {
			fcr := fcr
			result, ok := fc.results[fcr.url]
			if !ok {
				result = &CallResult{ready: make(chan struct{})}
				fc.results[fcr.url] = result
				go func() {
					result.respBody, result.err = fc.cf(fcr.url)
					close(result.ready)
				}()
			}
			go func() {
				<-result.ready
				fcr.respChan <- *result
			}()
		}
	}()
}

//关闭FC服务
func (fc *FuncCache) Close() {
	close(fc.requetChan)
}

func NewFuncCache(cf CachedFunc) *FuncCache {
	var fc FuncCache
	fc.cf = cf
	fc.results = make(map[string]*CallResult)
	fc.requetChan = make(chan FCRequest)
	return &fc
}

func main() {
	fc := NewFuncCache(getUrlBody)
	urls := []string{"http://baidu.com", "http://bing.com", "http://google.com", "http://baidu.com", "http://bing.com", "http://google.com"}
	var funcCallWG sync.WaitGroup
	fc.Service()
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
	fc.Close()
	// request http->http://baidu.com
	// request http->http://bing.com
	// request http->http://google.com
	// url:http://baidu.com, used time:0.12s, resp length:81
	// url:http://baidu.com, used time:0.12s, resp length:81
	// url:http://bing.com, used time:0.52s, resp length:75947
	// url:http://bing.com, used time:0.52s, resp length:75947
	// url:http://google.com, used time:21.07s, error:Get "http://google.com": dial tcp 172.217.160.78:80: connectex: A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond.
	// url:http://google.com, used time:21.07s, error:Get "http://google.com": dial tcp 172.217.160.78:80: connectex: A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond.
}
