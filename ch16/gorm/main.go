package main

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/icexmoon/go-notebook/ch16/gorm/model"
)

type Article struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
	ShowOpt bool
}

const COOKIE_USER_NAME = "user"

func getUser(rw http.ResponseWriter, r *http.Request) (user model.User, err error) {
	lc, err := r.Cookie(COOKIE_USER_NAME)
	if err != nil || lc.Value == "" {
		err = errors.New("need login")
		return
	}
	bts, err := base64.StdEncoding.DecodeString(lc.Value)
	if err != nil {
		return
	}
	buf := bytes.NewBuffer(bts)
	decoder := gob.NewDecoder(buf)
	user = model.User{}
	err = decoder.Decode(&user)
	return
}

func doLogin(rw http.ResponseWriter, user model.User) {
	buf := bytes.Buffer{}
	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(user)
	if err != nil {
		panic(err)
	}
	lc := http.Cookie{
		Name:  COOKIE_USER_NAME,
		Value: base64.StdEncoding.EncodeToString(buf.Bytes()),
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
	t := template.Must(template.ParseFiles("./template/login.html"))
	t.Execute(rw, nil)
}

func loginCheck(rw http.ResponseWriter, r *http.Request) {
	un := r.PostFormValue("user_name")
	p := r.PostFormValue("password")
	user, ok := model.CheckLogin(un, p)
	if !ok {
		panic("username and password error!")
	}
	doLogin(rw, user)
	redirect(rw, "/all_articles")
}

func allArticles(rw http.ResponseWriter, r *http.Request) {
	data := struct {
		UserName string
		Articles []Article
	}{}
	var err error

	marts, err := model.GetAllArticles()
	if err != nil {
		panic(err)
	}
	user, err := getUser(rw, r)
	if err != nil {
		fmt.Println(err)
		redirect(rw, "/login")
		return
	}
	for _, art := range marts {
		author := model.User{Id: art.UserId}
		author.Get()
		newArt := Article{
			Id:      art.Id,
			Content: art.Content,
			Author:  author.Name,
			ShowOpt: user.Id == art.UserId,
		}
		data.Articles = append(data.Articles, newArt)
	}
	data.UserName = user.Name
	t := template.Must(template.ParseFiles("./template/articles.html", "./template/header.html"))
	t.Execute(rw, data)
}

func addArticleHandle(rw http.ResponseWriter, r *http.Request) {
	content := r.PostFormValue("content")
	author, err := getUser(rw, r)
	if err != nil {
		panic(err)
	}
	art := model.Article{Content: content, UserId: author.Id}
	err = art.Add()
	if err != nil {
		panic(err)
	}
	redirect(rw, "/all_articles")
}

func exitHandle(rw http.ResponseWriter, r *http.Request) {
	unLogin(rw)
	redirect(rw, "/login")
}

func redirect(rw http.ResponseWriter, location string) {
	rw.Header().Set("Location", location)
	rw.WriteHeader(307)
}

func delArtHandle(rw http.ResponseWriter, r *http.Request) {
	ParamArtId := r.FormValue("art_id")
	ArtId, err := strconv.Atoi(ParamArtId)
	if err != nil {
		panic(err)
	}
	art := model.Article{Id: ArtId}
	if err := art.Get(); err != nil {
		panic(err)
	}
	//发帖人与当前用户一致才能删除
	user, err := getUser(rw, r)
	if err != nil {
		redirect(rw, "/login")
		return
	}
	if user.Id == art.UserId {
		if err := art.Delete(); err != nil {
			panic(err)
		}
	}
	redirect(rw, "/all_articles")
}

func articleHandle(rw http.ResponseWriter, r *http.Request) {
	paramId := r.FormValue("id")
	artId, err := strconv.Atoi(paramId)
	if err != nil {
		panic(err)
	}
	art := model.Article{Id: artId}
	if err := art.Get(); err != nil {
		panic(err)
	}
	cmts, err := model.GetCommentsByArtId(artId)
	if err != nil {
		panic(err)
	}
	data := struct {
		Article  model.Article
		Comments []model.Comment
		User     model.User
	}{}
	data.Comments = cmts
	data.Article = art
	user := model.User{Id: art.UserId}
	user.Get()
	data.User = user
	t := template.Must(template.ParseFiles("./template/article_detail.html", "./template/header.html"))
	t.Execute(rw, data)
}

func addCommentHandle(rw http.ResponseWriter, r *http.Request) {
	paramArtId := r.FormValue("art_id")
	content := r.FormValue("content")
	artId, err := strconv.Atoi(paramArtId)
	if err != nil {
		panic(err)
	}
	user, err := getUser(rw, r)
	if err != nil {
		redirect(rw, "/login")
		return
	}
	cmt := model.Comment{
		Content: content,
		ArtId:   artId,
		UserId:  user.Id,
	}
	cmt.Add()
	// fmt.Printf("Comment(id:%d) is created\n", cmt.Id)
	redirect(rw, "/article?id="+paramArtId)
}

func main() {
	fmt.Println("server started.")
	http.HandleFunc("/login", login)
	http.HandleFunc("/login_check", loginCheck)
	http.HandleFunc("/all_articles", allArticles)
	http.HandleFunc("/add_article", addArticleHandle)
	http.HandleFunc("/exit", exitHandle)
	http.HandleFunc("/del_art", delArtHandle)
	http.HandleFunc("/article", articleHandle)
	http.HandleFunc("/add_comment", addCommentHandle)
	http.ListenAndServe(":8080", nil)
}
