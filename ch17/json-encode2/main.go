package main

import (
	"encoding/json"
	"os"
)

type Article struct {
	Id       int       `json:"id"`
	Content  string    `json:"content"`
	Comments []Comment `json:"contents"`
	Uid      int       `json:"uid"`
}

type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Uid     int    `json:"uid"`
}

func main() {
	art := Article{
		Id:      1,
		Content: "this is a art's content.",
		Uid:     1,
		Comments: []Comment{
			{
				Id:      1,
				Content: "first comment content.",
				Uid:     1,
			},
			{
				Id:      2,
				Content: "second comment content.",
				Uid:     1,
			},
			{
				Id:      3,
				Content: "third comment content.",
				Uid:     2,
			},
		},
	}
	rest, _ := json.MarshalIndent(art, "", "\t")
	fopen, err := os.OpenFile("art.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	defer fopen.Close()
	fopen.Write(rest)
}
