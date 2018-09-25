package routes

import (
  "github.com/gin-gonic/gin"
  "fmt"
  "strconv"
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
  matrixOfSales := makeMatrixOfSales(visitors, removeDublicatesOfVisitors, removeDublicatesOfItems)

  /* init array of sales to get it into CA */
  arrayOfSales := makeArrayOfSales(matrixOfSales, len(removeDublicatesOfVisitors), len(removeDublicatesOfItems))

  /* CA algorithm*/
  prefs := MakeRatingMatrix(arrayOfSales, len(removeDublicatesOfVisitors), len(removeDublicatesOfItems))
  //products := removeDublicatesOfItems
  products := make([]string, 0)
  for i := 0; i < len(removeDublicatesOfItems); i++ {
    products = append(products, strconv.Itoa(i))
  }

  indexOfVisitor := getIndVisitor(visitors, myVisitor)
  if (indexOfVisitor == -1) {
    fmt.Println("Error: visitor doesn't found!")
    //os.Exit(-1)
    c.JSON(400, ApiMessage{"User doesn't found"})
  }
  var err error
  recommendations, err = GetRecommendations(prefs, /*getIndVisitor(visitors, myVisitor)*/ indexOfVisitor, products)
  if err != nil {
    fmt.Println("WHAT!?")
  }
  //fmt.Println(recommendations)
  c.JSON(200,recommendations)
}