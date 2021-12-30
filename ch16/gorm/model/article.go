package model

import (
	"time"
)

type Article struct {
	Id      int
	Content string `sql:"not null"`
	UserId  int    `sql:"index"`
	Ctime   time.Time
}

func (a *Article) Add() (err error) {
	// stmt, err := Db.Prepare(`insert into
	// article (content, user_id)
	// values ($1,$2) returning id`)
	// if err != nil {
	// 	return
	// }
	// defer stmt.Close()
	// stmt.QueryRow(a.Content, a.UserId).Scan(&a.Id)
	Db.Create(a)
	return
}

func (a *Article) Delete() (err error) {
	// _, err = Db.Exec(`delete from article
	// where id=$1`, a.Id)
	Db.Delete(a, a.Id)
	return
}

func (a *Article) Get() (err error) {
	// Db.QueryRowx(`select content,user_id,ctime
	// from article
	// where id=$1`, a.Id).StructScan(a)
	// fmt.Println(*a)
	Db.First(a, a.Id)
	return
}

func (a *Article) Update() (err error) {
	// _, err = Db.Exec(`UPDATE article
	// SET content=$2,user_id=$3
	// WHERE id=$1`, a.Id, a.Content, a.UserId)
	Db.Model(a).Updates(Article{Content: a.Content, UserId: a.UserId})
	return
}

func GetAllArticles() (arts []Article, err error) {
	// rows, err := Db.Query("select id,content,user_id,ctime from article")
	// if err != nil {
	// 	return
	// }
	// defer rows.Close()
	// for rows.Next() {
	// 	art := Article{}
	// 	rows.Scan(&art.Id, &art.Content, &art.UserId, &art.Ctime)
	// 	arts = append(arts, art)
	// }
	Db.Find(&arts)
	return
}
