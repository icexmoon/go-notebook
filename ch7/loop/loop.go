package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
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

func main() {
	start := time.Now()
	var names = []string{"Han Meimei", "Brus Lee", "Jack Chen"}
	var results = make(map[string]Person)
	for _, name := range names {
		p, err := getPerson(name)
		if err != nil {
			log.Println(err)
			continue
		}
		results[name] = p
	}
	for _, result := range results {
		fmt.Println(result)
	}
	end := time.Now()
	usedTimes := end.Sub(start)
	fmt.Printf("used time is %.2fs\n", usedTimes.Seconds())
	// Person(name:Brus Lee,age:30,career:engineer)
	// Person(name:Jack Chen,age:50,career:actor)
	// Person(name:Han Meimei,age:20,career:student)
	// used time is 30.07s
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
