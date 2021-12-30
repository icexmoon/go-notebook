package model

import "time"

type Comment struct {
	Id      int
	Content string `sql:"not null"`
	UserId  int    `sql:"index"`
	ArtId   int    `sql:"index"`
	Ctime   time.Time
}

func (c *Comment) Add() {
	Db.Create(c)
}

func GetCommentsByArtId(artId int) (cmts []Comment, err error) {
	Db.Where("art_id=?", artId).Find(&cmts)
	return
}
