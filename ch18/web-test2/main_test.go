package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/icexmoon/go-notebook/ch18/web-test/api"
	"github.com/julienschmidt/httprouter"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	tearDown()
	os.Exit(code)
}

var router *httprouter.Router
var recorder *httptest.ResponseRecorder

func setup() {
	router = httprouter.New()
	recorder = httptest.NewRecorder()
	router.POST("/api/login", api.ApiLogin)
}

func tearDown() {}

func TestApiLogin(t *testing.T) {
	reader := strings.NewReader(`
	{
		"data": {
			"name": "111",
			"password": "111"
		}
	}
	`)
	r, err := http.NewRequest("POST", "/api/login", reader)
	if err != nil {
		t.Fatal(err)
	}
	router.ServeHTTP(recorder, r)
	if recorder.Code != 200 {
		t.Fatal("http status is not 200")
	}
	data := struct {
		Data struct {
			Token string `json:"token"`
		} `json:"data"`
	}{}
	err = json.Unmarshal(recorder.Body.Bytes(), &data)
	if err != nil {
		t.Fatal(err)
	}
	if len(data.Data.Token) == 0 {
		t.Fatal("the token returned is empty")
	}
	t.Logf("the token returned is %s\n", data.Data.Token)
}

func TestLoginFail(t *testing.T) {
	reader := strings.NewReader(`
	{
		"data": {
			"name": "111",
			"password": "222"
		}
	}
	`)
	r, err := http.NewRequest("POST", "/api/login", reader)
	if err != nil {
		t.Fatal(err)
	}
	router.ServeHTTP(recorder, r)
	if recorder.Code == 200 {
		t.Fatal("http status is 200")
	}
	rs := recorder.Result().Status
	t.Logf("the status message is %s\n", rs)
}
