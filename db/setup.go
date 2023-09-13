package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "storeadmin"
	password = "storepass"
	dbname   = "reptileinstinct"
)

var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
var db *sql.db = nil

func Connectdb() {
	db, err := sql.Open("postgres", psqlInfo)
	DoesdbPointerWork()
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	log.Println("Successfully connected!")
}
