package model

import "time"

type Comment struct {
	Id      int
	Content string
	UserId  int
	ArtId   int
	Ctime   time.Time
}

func (c *Comment) Add() {
	sqlStr := "insert into comment(content,user_id,art_id) values ($1,$2,$3) returning id"
	Db.QueryRow(sqlStr, c.Content, c.UserId, c.ArtId).Scan(&c.Id)
}

func GetCommentsByArtId(artId int) (cmts []Comment, err error) {
	rows, err := Db.Query("select id,content,user_id,art_id,ctime from comment where art_id=$1", artId)
	if err != nil {
		return
	}
	for rows.Next() {
		cmt := Comment{}
		rows.Scan(&cmt.Id, &cmt.Content, &cmt.UserId, &cmt.ArtId, &cmt.Ctime)
		cmts = append(cmts, cmt)
	}
	return
}
