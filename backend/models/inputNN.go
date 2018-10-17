package models

import (
	"fmt"
	"log"
	"github.com/lib/pq"
)

func ImportInputNNToDB(arr [][]float64) bool {
	db, err := DB.Begin()
	rows, err := db.Query("SELECT COUNT (gender) FROM inputNN")
	if err != nil {
		fmt.Println("Error Count: ", err)
		log.Fatal(err)
		return false
	}
	defer rows.Close();
	var count int64
	for rows.Next() {
		rows.Scan(&count)
		if (count > 0) {
			fmt.Println("db INPUTNN is not empty")
			return  false
		}
	}

	if err != nil {
		fmt.Println("Input Error 1: ", err)
		log.Fatal(err)
		return false
	}
	stmt, err := db.Prepare(pq.CopyIn("inputNN","gender", "age","category","price"))
	for i := 0; i  < len(arr); i++ {
			_, err = stmt.Exec(arr[i][0], arr[i][1], arr[i][2], arr[i][3])
			if err != nil {
				fmt.Println("Input Error 2: ", err)
				log.Fatal(err)
			}
	}
	_, err = stmt.Exec()
	if err != nil {
		fmt.Println("Error Import ", err)
		log.Fatal(err)
		return false
	}

	err = stmt.Close()
	if err != nil {
		fmt.Println("Error Import ", err)
		log.Fatal(err)
		return false
	}

	err = db.Commit()
	if err != nil {
		fmt.Println("Error Import ", err)
		log.Fatal(err)
		return false
	}

	return true
}
