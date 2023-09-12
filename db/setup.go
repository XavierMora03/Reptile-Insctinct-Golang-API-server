package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "storeadmin"
	password = "storepass"
	dbname   = "reptileinstinct"
)

var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
var Db *sql.DB = nil

func PrintPsqlInfo() {
	fmt.Println(psqlInfo)
}

func ConnectDb() {
	Db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	if err := Db.Ping(); err != nil {
		panic(err)
	}
	defer Db.Close()

	fmt.Println("Successfully connected!")
}
