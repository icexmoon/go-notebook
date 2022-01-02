//pacakge api,rest接口
package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/icexmoon/go-notebook/ch18/web-test/model"
	"github.com/julienschmidt/httprouter"
)

var Uid int

func getParam(r *http.Request, param interface{}) error {
	len := r.ContentLength
	bodyBytes := make([]byte, len)
	_, err := r.Body.Read(bodyBytes)
	// fmt.Println(string(bodyBytes))
	if err != nil && err != io.EOF {
		return err
	}
	err = json.Unmarshal(bodyBytes, param)
	if err != nil {
		return err
	}
	return nil
}

func ApiLogin(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	param := struct {
		Data struct {
			Name     string `json:"name"`
			Password string `json:"password"`
		} `json:"data"`
	}{}
	err := getParam(r, &param)
	if err != nil {
		fmt.Println(err)
		http.Error(rw, "login error", http.StatusInternalServerError)
		return
	}
	n := param.Data.Name
	pwd := param.Data.Password
	user, ok := model.CheckLogin(n, pwd)
	if !ok {
		http.Error(rw, "login error", http.StatusInternalServerError)
		return
	}
	t := NewToken(user.Id)
	data := struct {
		Success bool `json:"success"`
		Data    struct {
			Token string `json:"token"`
		} `json:"data"`
	}{}
	data.Success = true
	data.Data.Token, err = t.String()
	if err != nil {
		fmt.Println(err)
		http.Error(rw, "login error", http.StatusInternalServerError)
		return
	}
	jBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		http.Error(rw, "login error", http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(jBytes)
}

func responseData(rw http.ResponseWriter, d interface{}) {
	data := struct {
		Success bool        `json:"success"`
		Data    interface{} `json:"data"`
	}{}
	data.Success = true
	data.Data = d
	jBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		http.Error(rw, "login error", http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(jBytes)
}

func ApiLoginCheck(h httprouter.Handle) httprouter.Handle {
	return func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		contents, err := ioutil.ReadAll(r.Body)
		r.Body = ioutil.NopCloser(bytes.NewBuffer(contents))
		if err != nil {
			fmt.Println(err)
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		param := map[string]interface{}{}
		json.Unmarshal(contents, &param)
		t := Token{}
		var token string
		switch v := param["token"].(type) {
		case string:
			token = v
		default:
			err = errors.New("token format is error")
			fmt.Println(err)
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		err = t.Parse(token)
		if err != nil {
			fmt.Println(err)
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		err = t.Validate()
		if err != nil {
			fmt.Println(err)
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		Uid = t.Id
		fmt.Println("login check is success.")
		h(rw, r, p)
	}
}

func ApiAllArticles(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	arts, err := model.GetAllArticles()
	if err != nil {
		fmt.Println(err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	data := struct {
		Success bool `json:"success"`
		Data    struct {
			Articles []model.Article `json:"articles"`
		} `json:"data"`
	}{}
	data.Success = true
	data.Data.Articles = arts
	jBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(jBytes)
}

func ApiArticleDetail(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	sid := p.ByName("id")
	id, err := strconv.Atoi(sid)
	if err != nil {
		fmt.Println(err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	art := model.Article{Id: id}
	if err := art.Get(); err != nil {
		fmt.Println(err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	responseData(rw, art)
}

func ApiDelArticle(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	sid := p.ByName("id")
	id, err := strconv.Atoi(sid)
	if err != nil {
		fmt.Println(err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	art := model.Article{Id: id}
	err = art.Get()
	if err != nil {
		fmt.Println(err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	if art.UserId != Uid {
		err = errors.New("now user is not the article's owner")
		fmt.Println(err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	err = art.Delete()
	if err != nil {
		fmt.Println(err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	responseData(rw, nil)
}

func ApiAddArticle(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	param := struct {
		Token string `json:"token"`
		Data  struct {
			Content string `json:"content"`
		} `json:"data"`
	}{}
	err := getParam(r, &param)
	if err != nil {
		fmt.Println(err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	art := model.Article{
		UserId:  Uid,
		Content: param.Data.Content,
	}
	err = art.Add()
	if err != nil {
		fmt.Println(err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	rd := struct {
		ArtId int `json:"art_id"`
	}{}
	rd.ArtId = art.Id
	responseData(rw, rd)
}
