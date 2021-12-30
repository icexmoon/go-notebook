package model

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func init() {
	var err error
	connStr := `user=bbs dbname=bbs password=bbs_admin port=5433 host=localhost sslmode=disable`
	Db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
}
