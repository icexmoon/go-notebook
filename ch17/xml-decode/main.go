package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type Article struct {
	XMLName  xml.Name  `xml:"article"`
	Id       int       `xml:"id,attr"`
	Content  string    `xml:"content"`
	Comments []Comment `xml:"comments>comment"`
	Uid      int       `xml:"uid,attr"`
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
	XMLName xml.Name `xml:"comment"`
	Id      int      `xml:"id,attr"`
	Content string   `xml:",chardata"`
	Uid     int      `xml:"uid,attr"`
}

func (c *Comment) String() string {
	return fmt.Sprintf("Comment(Id:%d,Content:'%s',Uid:%d)", c.Id, c.Content, c.Uid)
}

func main() {
	fopen, err := os.Open("art.xml")
	if err != nil {
		panic(err)
	}
	defer fopen.Close()
	content, err := ioutil.ReadAll(fopen)
	if err != nil && err != io.EOF {
		panic(err)
	}
	art := Article{}
	err = xml.Unmarshal(content, &art)
	if err != nil {
		panic(err)
	}
	fmt.Println(art.String())
}
