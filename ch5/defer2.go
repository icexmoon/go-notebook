package main

import (
	"fmt"
	"log"
	"net/http"
)

func myFunc8(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	contentType := resp.Header.Get("Content-Type")
	fmt.Println(contentType)
}

func main() {
	myFunc8("http://bing.com")
	// text/html; charset=utf-8
	myFunc8("http://baidu.com")
	// 2021/11/17 17:52:16 Get "http://baidu.com": read tcp 192.168.1.13:11584->220.181.38.148:80: wsarecv: An existing connection was forcibly closed by the remote host.
}
