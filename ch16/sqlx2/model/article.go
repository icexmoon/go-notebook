package model

import (
	"fmt"
	"time"
)

type Article struct {
	Id      int
	Content string
	UserId  int `db:"user_id"`
	Ctime   time.Time
}

func (a *Article) Add() (err error) {
	stmt, err := Db.Prepare(`insert into 
	article (content, user_id) 
	values ($1,$2) returning id`)
	if err != nil {
		return
	}
	defer stmt.Close()
	stmt.QueryRow(a.Content, a.UserId).Scan(&a.Id)
	return
}

func (a *Article) Delete() (err error) {
	_, err = Db.Exec(`delete from article
	where id=$1`, a.Id)
	return
}

func (a *Article) Get() (err error) {
	Db.QueryRowx(`select content,user_id,ctime
	from article
	where id=$1`, a.Id).StructScan(a)
	fmt.Println(*a)
	return
}

func (a *Article) Update() (err error) {
	_, err = Db.Exec(`UPDATE article 
	SET content=$2,user_id=$3
	WHERE id=$1`, a.Id, a.Content, a.UserId)
	return
}

func GetAllArticles() (arts []Article, err error) {
	rows, err := Db.Query("select id,content,user_id,ctime from article")
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		art := Article{}
		rows.Scan(&art.Id, &art.Content, &art.UserId, &art.Ctime)
		arts = append(arts, art)
	}
	return
}
