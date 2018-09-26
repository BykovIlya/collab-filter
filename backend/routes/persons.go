package routes

import (
  "github.com/gin-gonic/gin"
  "fmt"
  "ColabFilter/colab-filter/backend/algorithm"
  "ColabFilter/colab-filter/backend/utils"
)

var myVisitor string

func GetUsers (c *gin.Context) {
  count := 0
  for i := 0; i < len(visitors); i++ {
    if len(visitors[i].Items) > 15 {
      fmt.Println(visitors[i].Visitorid_string)
      count++
    }
  }
  fmt.Println(count)
  c.JSON(200,visitors)
}

func GetPerson(c *gin.Context)  {
  myVisitor = c.Param("id")

  indexOfVisitor := algorithm.GetIndVisitor(visitors, myVisitor)
  if (indexOfVisitor == -1) {
    fmt.Println("Error: visitor doesn't found!")
    c.JSON(400, utils.ApiMessage{"User doesn't found"})
  }
  var err error
  recommendations, err = algorithm.GetRecommendations(prefs, /*getIndVisitor(visitors, myVisitor)*/ indexOfVisitor, products)
  if err != nil {
    fmt.Println("WHAT!?")
  }
  if (len(recommendations) > 0) {
    algorithm.InsertRecommendToDB(recommendations)
    c.JSON(200, recommendations)
  } else {
    fmt.Println("No recommendations!")
  }
}
