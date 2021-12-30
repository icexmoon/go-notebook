package model

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var Db *sqlx.DB

func init() {
	var err error
	connStr := `user=bbs dbname=bbs password=bbs_admin port=5433 host=localhost sslmode=disable`
	Db, err = sqlx.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
}
