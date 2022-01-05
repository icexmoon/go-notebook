package model

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func init() {
	var err error
	dsn := "host=localhost user=bbs password=bbs_admin dbname=bbs2 port=5433 sslmode=disable TimeZone=Asia/Shanghai"
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	Db.AutoMigrate(&Article{}, &User{}, &Comment{})
}
