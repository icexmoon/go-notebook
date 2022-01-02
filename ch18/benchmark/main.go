package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Article struct {
	Id       int       `json:"id"`
	Content  string    `json:"content"`
	Comments []Comment `json:"contents"`
	Uid      int       `json:"uid"`
}

func (a *Article) String() string {
	var comments []string
	for _, c := range a.Comments {
		comments = append(comments, c.String())
	}
	scs := strings.Join(comments, ",")
	return fmt.Sprintf("Article(Id:%d,Content:'%s',Comments:[%s],Uid:%d)", a.Id, a.Content, scs, a.Uid)
}

type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Uid     int    `json:"uid"`
}

func (c *Comment) String() string {
	return fmt.Sprintf("Comment(Id:%d,Content:'%s',Uid:%d)", c.Id, c.Content, c.Uid)
}

func StreamDecode(fileName string) Article {
	fopen, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer fopen.Close()
	decoder := json.NewDecoder(fopen)
	art := Article{}
	err = decoder.Decode(&art)
	if err != nil {
		panic(err)
	}
	return art
}

func MemoryDecode(fileName string) Article {
	fopen, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer fopen.Close()
	content, err := ioutil.ReadAll(fopen)
	if err != nil {
		panic(err)
	}
	art := Article{}
	err = json.Unmarshal(content, &art)
	if err != nil {
		panic(err)
	}
	return art
}

func main() {
}
