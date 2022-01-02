package model

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var Db *gorm.DB

func init() {
	var err error
	connStr := `user=bbs dbname=bbs2 password=bbs_admin port=5433 host=localhost sslmode=disable`
	Db, err = gorm.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	Db.AutoMigrate(&Article{}, &User{}, &Comment{})
}
