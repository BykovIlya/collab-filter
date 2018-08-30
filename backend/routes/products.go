package routes

import (
  "github.com/gin-gonic/gin"
  "net/http"
  "fmt"
  "strconv"
)

var events []Events
var visitors []Visitor
var removeDublicatesOfVisitors []string
var removeDublicatesOfItems []string
var items []ItemsGlobal
var myVisitor string
var recommendations []Recommendation
//var csvFileName ="api/upload/"+"File.csv"
//var myVisitor string

type EventsList struct {
  Events []Events `json:"events"`
  Total int `json:"Total"`
}

func ImportEvents(c *gin.Context)  {
  file, err := c.FormFile("file")
  if err != nil {
    c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
    return
  }
  csvFileName:="api/upload/"+"File.csv"
  if err := c.SaveUploadedFile(file, csvFileName); err != nil {
    c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
    return
  }
  Algorithm(csvFileName)
}

func Algorithm(csvFileName string)  {
  events = readingTransactionsFromFile(csvFileName)
  // fmt.Println(events)
  removeDublicatesOfVisitors = makeUniqArrayOfVisitors(events)
  removeDublicatesOfItems = makeUniqArrayOfItems(events)
  visitors = make([] Visitor, len(removeDublicatesOfVisitors))
  /* make struct of visitors */
  initVisitors(visitors, removeDublicatesOfVisitors)
  /* add items to each visitor */
  addItemsToVisitor(visitors,events)
  addCountToEachProductOfEachVisitor(visitors)
  items = make ([]ItemsGlobal, len(events))
  for i := 0; i < len(events); i++ {
    items[i].Itemid = events[i].Itemid
    items[i].Count = 1
  }

}
func GetEvents (c *gin.Context) {
  c.JSON(200,events)
}

func GetUsers (c *gin.Context) {
  c.JSON(200,visitors)
}

func GetProducts (c *gin.Context) {
  c.JSON(200, items)
}

func GetRecommends (c *gin.Context) {
  if myVisitor != "" {
    c.JSON(200, recommendations)
  }
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
}

