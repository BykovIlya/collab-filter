package routes

import (
"github.com/gin-gonic/gin"
"net/http"
"fmt"
  "ColabFilter/colab-filter/backend/models"
  "ColabFilter/colab-filter/backend/algorithm"
  "strconv"
  . "github.com/skelterjohn/go.matrix"
)

var events []models.Events
var visitors []models.Visitor
var removeDublicatesOfVisitors []string
var removeDublicatesOfItems []string
var items []models.ItemsGlobal
var recommendations []algorithm.Recommendation
//var csvFileName ="api/upload/"+"File.csv"
//var myVisitor string
var matrixOfSales [][]float64
var arrayOfSales []float64
var prefs *DenseMatrix
var products []string

type EventsList struct {
  Events []models.Events `json:"events"`
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
  models.ClearDB(models.DB)
  events = algorithm.ReadingTransactionsFromFile(csvFileName)
  models.ImportEventsToDB(events)
  removeDublicatesOfVisitors = algorithm.MakeUniqArrayOfVisitors(events)
  removeDublicatesOfItems = algorithm.MakeUniqArrayOfItems(events)
  visitors = make([] models.Visitor, len(removeDublicatesOfVisitors))
  /* make struct of visitors */
  algorithm.InitVisitors(visitors, removeDublicatesOfVisitors)
  /* add items to each visitor */
  algorithm.AddItemsToVisitor(visitors,events)
  algorithm.AddCountToEachProductOfEachVisitor(visitors)
  items = make ([]models.ItemsGlobal, len(events))
  for i := 0; i < len(events); i++ {
    items[i].Itemid = events[i].Itemid
    items[i].Count = 1
  }
  matrixOfSales = algorithm.MakeMatrixOfSales(visitors, removeDublicatesOfVisitors, removeDublicatesOfItems)

  /* init array of sales to get it into CA */
  arrayOfSales = algorithm.MakeArrayOfSales(matrixOfSales, len(removeDublicatesOfVisitors), len(removeDublicatesOfItems))

  /* CA algorithm*/
  prefs = algorithm.MakeRatingMatrix(arrayOfSales, len(removeDublicatesOfVisitors), len(removeDublicatesOfItems))
  //products := removeDublicatesOfItems
  products = make([]string, 0)
  for i := 0; i < len(removeDublicatesOfItems); i++ {
    products = append(products, strconv.Itoa(i))
  }
  models.ImportPersonsToDB(visitors)

}

func GetEvents (c *gin.Context) {
  c.JSON(200,events)
}

