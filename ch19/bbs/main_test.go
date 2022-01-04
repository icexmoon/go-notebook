package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/icexmoon/go-notebook/ch19/bbs/api"
	"github.com/julienschmidt/httprouter"
	. "gopkg.in/check.v1"
)

type LoginTestSuite struct {
	Recoder *httptest.ResponseRecorder
	Router  *httprouter.Router
}

func (s *LoginTestSuite) SetUpTest(c *C) {
	s.Router = httprouter.New()
	s.Recoder = httptest.NewRecorder()
	s.Router.POST("/api/login", api.ApiLogin)
}

func (s *LoginTestSuite) TestLogin(c *C) {
	reader := strings.NewReader(`
	{
		"data": {
			"name": "111",
			"password": "111"
		}
	}
	`)
	r, err := http.NewRequest("POST", "/api/login", reader)
	c.Check(err, Equals, nil)
	s.Router.ServeHTTP(s.Recoder, r)
	c.Check(s.Recoder.Code, Equals, 200)
	data := struct {
		Data struct {
			Token string `json:"token"`
		} `json:"data"`
	}{}
	err = json.Unmarshal(s.Recoder.Body.Bytes(), &data)
	c.Check(err, Equals, nil)
	c.Check(len(data.Data.Token) > 0, Equals, true)
}

func (s *LoginTestSuite) TestLoginFail(c *C) {
	reader := strings.NewReader(`
	{
		"data": {
			"name": "111",
			"password": "222"
		}
	}
	`)
	r, err := http.NewRequest("POST", "/api/login", reader)
	c.Check(err, Equals, nil)
	s.Router.ServeHTTP(s.Recoder, r)
	c.Check(s.Recoder.Code, Not(Equals), 200)
	// c.Log(s.Recoder.Result().Status)
}

func init() {
	Suite(&LoginTestSuite{})
}

func Test(t *testing.T) {
	TestingT(t)
}
