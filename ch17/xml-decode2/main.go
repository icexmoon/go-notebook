package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

type Article struct {
	XMLName  xml.Name  `xml:"article"`
	Id       int       `xml:"id,attr"`
	Content  string    `xml:content`
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
	d := xml.NewDecoder(fopen)
	var comments []Comment
	for {
		token, err := d.Token()
		if err == io.EOF {
			//xml解析完毕
			break
		}
		if err != nil {
			//解析出错
			panic(err)
		}
		switch node := token.(type) {
		case xml.StartElement:
			if node.Name.Local == "comment" {
				cmmt := Comment{}
				d.DecodeElement(&cmmt, &node)
				comments = append(comments, cmmt)
			}
		}
	}
	art := Article{}
	art.Comments = comments
	fmt.Println(art.String())
	// Article(Id:0,Content:'',Comments:[Comment(Id:1,Content:'first comment content.',Uid:1),Comment(Id:2,Content:'second comment content.',Uid:1),Comment(Id:3,Content:'third comment content.',Uid:2)],Uid:0)
}
