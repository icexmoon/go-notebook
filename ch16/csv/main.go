package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

type Article struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

var aMutex sync.RWMutex
var articles []*Article

const COOKIE_USER_NAME = "user_name"

func addArticle(content string, author string) {
	aMutex.Lock()
	defer aMutex.Unlock()
	aLen := len(articles)
	newArticle := Article{
		Id:      aLen,
		Content: content,
		Author:  author,
	}
	articles = append(articles, &newArticle)
	//如果当前文章条目是5的整数倍，进行备份
	lenA := len(articles)
	if lenA%5 == 0 && lenA != 0 {
		go backupArticles()
	}
}

func getArticles() []Article {
	aMutex.RLock()
	defer aMutex.RUnlock()
	var aCopy []Article
	for _, a := range articles {
		aCopy = append(aCopy, *a)
	}
	return aCopy
}

func getUserName(rw http.ResponseWriter, r *http.Request) (userName string, err error) {
	lc, err := r.Cookie(COOKIE_USER_NAME)
	if err != nil || lc.Value == "" {
		//去登录
		rw.Header().Set("Location", "/login")
		rw.WriteHeader(301)
		err = errors.New("need login")
		return
	}
	userName = lc.Value
	return
}

func doLogin(rw http.ResponseWriter, user string) {
	lc := http.Cookie{
		Name:  COOKIE_USER_NAME,
		Value: user,
	}
	rw.Header().Set("Set-Cookie", lc.String())
}

func unLogin(rw http.ResponseWriter) {
	lc := http.Cookie{
		Name:    COOKIE_USER_NAME,
		Value:   "",
		Expires: time.Unix(1, 0),
		MaxAge:  -1,
	}
	rw.Header().Set("Set-Cookie", lc.String())
}

func login(rw http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("login.html"))
	t.Execute(rw, nil)
}

func loginCheck(rw http.ResponseWriter, r *http.Request) {
	un := r.PostFormValue("user_name")
	p := r.PostFormValue("password")
	if un == "" && p == "" {
		log.Fatal(errors.New("user name and password is empty"))
	}
	doLogin(rw, un)
	rw.Header().Set("Location", "/all_articles")
	rw.WriteHeader(301)
}

func allArticles(rw http.ResponseWriter, r *http.Request) {
	data := struct {
		UserName string
		Articles []Article
	}{}
	data.Articles = getArticles()
	un, err := getUserName(rw, r)
	if err != nil {
		fmt.Println(err)
		rw.Header().Set("Location", "/login")
		rw.WriteHeader(301)
		return
	}
	data.UserName = un
	t := template.Must(template.ParseFiles("articles.html"))
	t.Execute(rw, data)
}

func addArticleHandle(rw http.ResponseWriter, r *http.Request) {
	content := r.PostFormValue("content")
	author, err := getUserName(rw, r)
	if err != nil {
		fmt.Println(err)
		return
	}
	addArticle(content, author)
	rw.Header().Set("Location", "/all_articles")
	rw.WriteHeader(301)
}

func exitHandle(rw http.ResponseWriter, r *http.Request) {
	unLogin(rw)
	rw.Header().Set("Location", "/login")
	rw.WriteHeader(301)
}

const DATA_FILE = "data.csv"

//备份文章数据
func backupArticles() {
	as := getArticles()
	//如果存在备份文件，打开，不存在，创建。
	f, err := os.OpenFile(DATA_FILE, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	r := csv.NewWriter(f)
	for _, a := range as {
		record := []string{strconv.Itoa(a.Id), a.Content, a.Author}
		r.Write(record)
	}
	r.Flush()
}

func getArticlesByBackup() []Article {
	as := make([]Article, 0)
	f, err := os.Open(DATA_FILE)
	if err != nil {
		//文件不存在，返回空切片
		return as
	}
	defer f.Close()
	reader := csv.NewReader(f)
	reader.FieldsPerRecord = 0
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}
	for _, rec := range records {
		id, err := strconv.Atoi(rec[0])
		if err != nil {
			panic(err)
		}
		newA := Article{
			Id:      id,
			Content: rec[1],
			Author:  rec[2],
		}
		as = append(as, newA)
	}
	return as
}

func main() {
	//如果存在备份数据，加载
	as := getArticlesByBackup()
	if len(as) > 0 {
		for _, a := range as {
			a := a
			articles = append(articles, &a)
		}
	}
	http.HandleFunc("/login", login)
	http.HandleFunc("/login_check", loginCheck)
	http.HandleFunc("/all_articles", allArticles)
	http.HandleFunc("/add_article", addArticleHandle)
	http.HandleFunc("/exit", exitHandle)
	http.ListenAndServe(":8080", nil)
}
