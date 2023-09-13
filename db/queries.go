package db

import(
	"log"
	"fmt"
)

func DoesdbPointerWork() {
	log.Println("ESTA ES EL VALOR", db)
	return
	if err := db.Ping(); err != nil {
		panic(err)
	}
}

func retriveReptiles(){

	retriveReptiles := `SELECT * FROM productos.reptiles`
	data, err := db.db.Exec(retriveReptiles)

	if err != nil {
		panic(err)
	}
	log.Println(data)

}