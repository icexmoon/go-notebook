package main

import (
	"encoding/xml"
	"os"
)

type Article struct {
	XMLName  xml.Name  `xml:"article"`
	Id       int       `xml:"id,attr"`
	Content  string    `xml:content`
	Comments []Comment `xml:"comments>comment"`
	Uid      int       `xml:"uid,attr"`
}

type Comment struct {
	XMLName xml.Name `xml:"comment"`
	Id      int      `xml:"id,attr"`
	Content string   `xml:",chardata"`
	Uid     int      `xml:"uid,attr"`
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
	fopen, err := os.OpenFile("art.xml", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	defer fopen.Close()
	fopen.Write([]byte(xml.Header))
	encoder := xml.NewEncoder(fopen)
	encoder.Indent("", "\t")
	encoder.Encode(art)
}
