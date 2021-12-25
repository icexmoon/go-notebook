package main

import (
	"encoding/base64"
	"log"
	"net/http"
	"text/template"
	"time"
)

const FLASH_MSG_COOKIE_NAME = "flash"

func writeFlashMsg(rw http.ResponseWriter, msg string) {
	flashCookie := http.Cookie{
		Name:  FLASH_MSG_COOKIE_NAME,
		Value: base64.StdEncoding.EncodeToString([]byte(msg)),
	}
	http.SetCookie(rw, &flashCookie)
}

func readFlashMsg(rw http.ResponseWriter, r *http.Request) (msg string, err error) {
	flashCookie, err := r.Cookie(FLASH_MSG_COOKIE_NAME)
	if err != nil {
		if err == http.ErrNoCookie {
			return "", nil
		}
		return
	}
	bytes, err := base64.StdEncoding.DecodeString(flashCookie.Value)
	if err != nil {
		return
	}
	msg = string(bytes)
	//清除http客户端cookie
	flashCookie = &http.Cookie{
		Name:    FLASH_MSG_COOKIE_NAME,
		Expires: time.Unix(1, 0),
		MaxAge:  -1,
	}
	http.SetCookie(rw, flashCookie)
	return
}

func checkLogin(rw http.ResponseWriter, r *http.Request) {
	name := r.PostFormValue("name")
	pasword := r.PostFormValue("password")
	if name == "icexmoon" && pasword == "12345" {
		//写入闪回消息
		writeFlashMsg(rw, "hello! welcome to our web site!")
		rw.Header().Set("location", "/index")
		rw.WriteHeader(301)
	} else {
		writeFlashMsg(rw, "username or password error!")
		rw.Header().Set("location", "/login")
		rw.WriteHeader(301)
	}
}

func login(rw http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("login.html")
	if err != nil {
		log.Fatal(err)
	}
	msg, err := readFlashMsg(rw, r)
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(rw, msg)
}

func index(rw http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatal(err)
	}
	msg, err := readFlashMsg(rw, r)
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(rw, msg)
}

func main() {
	serverMux := http.NewServeMux()
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: serverMux,
	}
	serverMux.HandleFunc("/index", index)
	serverMux.HandleFunc("/login", login)
	serverMux.HandleFunc("/check_login", checkLogin)
	server.ListenAndServe()
}
