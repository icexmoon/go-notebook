package model

import (
	"time"
)

type Article struct {
	Id      int       `json:"id"`
	Content string    `sql:"not null" json:"content"`
	UserId  int       `sql:"index" json:"user_id"`
	Ctime   time.Time `json:"ctime"`
}

func (a *Article) Add() (err error) {
	Db.Create(a)
	return
}

func (a *Article) Delete() (err error) {
	Db.Delete(a, a.Id)
	return
}

func (a *Article) Get() (err error) {
	Db.First(a, a.Id)
	return
}

func (a *Article) Update() (err error) {
	Db.Model(a).Updates(Article{Content: a.Content, UserId: a.UserId})
	return
}

func GetAllArticles() (arts []Article, err error) {
	Db.Find(&arts)
	return
}
