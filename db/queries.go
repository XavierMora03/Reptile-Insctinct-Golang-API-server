package db

import (
	"ReptileApi/models"
	"database/sql"
	"log"
)

func RetriveReptiles() {

	retriveReptiles := `SELECT * FROM productos.reptiles`
	rows, err := db.Query(retriveReptiles)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

}

func AddReptile(reptile models.Reptile) int {

	addReptilesStatement := `INSERT INTO productos.reptiles (name, RegularPrice , price,age,description,genre)
													VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	id := 0
	err := db.QueryRow(addReptilesStatement, reptile.Name, reptile.RegularPrice, reptile.Price,
		reptile.AgeCategory, reptile.Description, reptile.Genre).Scan(&id)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No rows returned, error while inserting reptile")
			return -1
		}
		panic(err)
	}

	return id
}
