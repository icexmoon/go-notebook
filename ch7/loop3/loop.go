package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sync"
	"time"
)

type RespObj struct {
	Status string `json:status`
	Result Person `json:result`
}

type Person struct {
	Name   string `json:name`
	Age    int    `json:"age,string"`
	Career string `json:career`
}

func (p Person) String() string {
	return fmt.Sprintf("Person(name:%s,age:%d,career:%s)", p.Name, p.Age, p.Career)
}

type ChanResult struct {
	err error
	p   Person
}

func main() {
	start := time.Now()
	var names = []string{"Han Meimei", "Brus Lee", "Jack Chen"}
	var results = make(chan ChanResult, len(names))
	var gNums sync.WaitGroup
	for _, name := range names {
		name := name
		gNums.Add(1)
		go func() {
			defer gNums.Done()
			p, err := getPerson(name)
			results <- ChanResult{err: err, p: p}
		}()
	}
	go func() {
		gNums.Wait()
		close(results)
	}()
	for result := range results {
		// result := <-results
		if result.err != nil {
			log.Println(result.err)
			continue
		}
		fmt.Println(result.p)
	}
	end := time.Now()
	usedTimes := end.Sub(start)
	fmt.Printf("used time is %.2fs\n", usedTimes.Seconds())
	// Person(name:Han Meimei,age:20,career:student)
	// Person(name:Jack Chen,age:50,career:actor)
	// Person(name:Brus Lee,age:30,career:engineer)
	// used time is 10.04s
}

func getPerson(name string) (p Person, err error) {
	hr, err := http.NewRequest("GET", "http://myweb.com/index.php", nil)
	params := make(url.Values)
	params.Add("name", name)
	hr.URL.RawQuery = params.Encode()
	resp, err := http.DefaultClient.Do(hr)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	respText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	var ro RespObj
	err = json.Unmarshal(respText, &ro)
	if err != nil {
		return
	}
	if ro.Status == "success" {
		p = ro.Result
	}
	return
}
