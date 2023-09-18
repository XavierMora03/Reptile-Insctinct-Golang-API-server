package db

import (
	"ReptileApi/models"
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
	log.Println("esta es la data ", data)
	fmt.Println(data.Columns())
	fmt.Println(data.Scan())
	// return data.LastInsertId()
}

func AddReptiles(reptilelist []models.Reptile) {
	// var newIdList []int64
	for i := 0; i < len(reptilelist); i++ {
	}

}

func AddReptile(reptile models.Reptile) int {

	addReptilesStatement := `INSERT INTO productos.reptiles (name, RegularPrice , price,age,description,genre)
													VALUES ($1, $2, $3, $4, $5, $6)`
	res, err := db.Exec(addReptilesStatement, reptile.Name, reptile.RegularPrice, reptile.Price,
		reptile.AgeCategory, reptile.Description, reptile.Genre)

	if err != nil {
		panic(err)
	}

	lastId, err := res.RowsAffected()

	if err != nil {
		panic(err)
	}
	return int(lastId)
}
