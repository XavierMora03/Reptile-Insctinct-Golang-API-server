package db

import (
	"ReptileApi/models"
	"database/sql"
	"log"
)

func RetriveReptiles(reptileList *[]models.Reptile) {

	retriveReptiles := `SELECT * FROM productos.reptiles;`
	rows, err := db.Query(retriveReptiles)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var reptile models.Reptile
		err := rows.Scan(&reptile.ID, &reptile.Name, &reptile.RegularPrice, &reptile.Price, &reptile.AgeCategory, &reptile.Description, &reptile.Genre)

		if err != nil {
			panic(err)
		}

		*reptileList = append(*reptileList, reptile)
	}
}

func DeleteReptile(id_delete int) {
	deleteReptileById := `DELETE FROM productos.reptiles WHERE id = $1;`
	_, err := db.Exec(deleteReptileById, id_delete)
	if err != nil {
		log.Println("Error while deleting reptile with id: ", id_delete)
		panic(err)
	}
}

func AddReptile(reptile models.Reptile) int {

	addReptilesStatement := `INSERT INTO productos.reptiles (name, RegularPrice , price,age,description,genre)
													VALUES ($1, $2, $3, $4, $5, $6) RETURNING id;`

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
