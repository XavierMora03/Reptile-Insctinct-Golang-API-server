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
var db *sql.DB = nil

func Connectdb() {

	if dbPointer, err := sql.Open("postgres", psqlInfo); err != nil {
		panic(err)
	} else {
		db = dbPointer
	}

	log.Println("DATABASE SUCCESSFULLY CONNECTED!")
}
