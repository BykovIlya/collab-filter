package algorithm

import (
  "fmt"
  "ColabFilter/colab-filter/backend/models"
)

type PersonRecommends struct {
  id int64
  recommends []Recommendation
}

func InsertRecommendToDB(/*user models.Visitor,*/ recommends [] Recommendation) {
  row:=models.DB.QueryRow("INSERT INTO recommends(user_id,recommendations) VALUES($1,$2)", recommends[0].Product, recommends[0].MpRating)
  fmt.Println(row)
}