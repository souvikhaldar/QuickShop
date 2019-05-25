package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "quickshop"
)

var DB *sql.DB

func init() {
	var err error
	psql := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	DB, err = sql.Open("postgres", psql)
	if err != nil {
		panic(err)
	}
	if er := DB.Ping(); er != nil {
		panic(er)
	}
}
