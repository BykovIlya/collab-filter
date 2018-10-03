package models

import (
	"fmt"
	"log"
	"github.com/lib/pq"
	"ColabFilter/colab-filter/backend/algorithm"
)

type Recommend struct {
	user_id string;
	recommendations []algorithm.Recommendation;
}

func ImportRecommendsToDB(visitor string, recommendations []algorithm.Recommendation) bool {
	db, err := DB.Begin()
	if err != nil {
		fmt.Println("Input Error: ", err)
		log.Fatal(err)
		return false
	}
	stmt, err := db.Prepare(pq.CopyIn("recommends","user_id","recommend","score"))
	for _, ev := range recommendations {
		_, err = stmt.Exec(visitor, ev.Product, ev.MpRating)
		if err != nil {
			fmt.Println("Input Error: ", err)
			log.Fatal(err)
		}
	}

	_, err = stmt.Exec()
	if err != nil {
		fmt.Println("Error Import ",  err)
		log.Fatal(err)
		return false
	}

	err = stmt.Close()
	if err != nil {
		fmt.Println("Error Import ",  err)
		log.Fatal(err)
		return false
	}

	err = db.Commit()
	if err != nil {
		fmt.Println("Error Import ",  err)
		log.Fatal(err)
		return false
	}

	return true
}

func GetRecommendsFromBD(user string) []algorithm.Recommendation{
	rows, err := DB.Query("SELECT recommend,score FROM recommends WHERE user_id=$1",user)
	recs := []algorithm.Recommendation{}
	fmt.Println(rows)
	fmt.Println("ROW BEGIN!!!")
	if (rows == nil) {
		fmt.Println("ERROR!!")
		return []algorithm.Recommendation{}
	}
	fmt.Println("ROW BEGIN222!!!")
	for rows.Next() {
		rec := algorithm.Recommendation{}
		err = rows.Scan(&rec.Product, &rec.MpRating)
		if err != nil {
			fmt.Println("Scan Error: ", err)
			log.Fatal(err)
		}
		recs = append(recs, rec)
	}
	if err != nil {
		fmt.Println("Reading Error: ", err)
		log.Fatal(err)
	}
	if (len(recs) > 0) {
		return recs
	} else {
		return []algorithm.Recommendation{}
	}
}
