package main

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sync"
	"time"
)

type Article struct {
	Id      int
	Content string
	Author  string
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

func main() {
	http.HandleFunc("/login", login)
	http.HandleFunc("/login_check", loginCheck)
	http.HandleFunc("/all_articles", allArticles)
	http.HandleFunc("/add_article", addArticleHandle)
	http.HandleFunc("/exit", exitHandle)
	http.ListenAndServe(":8080", nil)
}
