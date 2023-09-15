package db

import (
	_ "database/sql"
	"fmt"
	"log"
)

func RetriveReptiles() {

	retriveReptiles := `SELECT * FROM productos.reptiles`
	data, err := db.Query(retriveReptiles)

	if err != nil {
		panic(err)
	}
	log.Println("esta es la rata ", data)
	fmt.Println(data.Columns())
	fmt.Println(data.Scan())
	// return data.LastInsertId()
}
