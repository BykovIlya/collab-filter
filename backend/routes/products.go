package routes

import (
  "github.com/gin-gonic/gin"
  "net/http"
  "fmt"
)

var events []Events
//var visitors []Visitor
//var
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

}

func GetEvents (c *gin.Context) {
  csvFileName:="api/upload/"+"File.csv"
  events = readingTransactionsFromFile(csvFileName)
 // fmt.Println(events)
  c.JSON(200,events)
}

func GetUsers (c *gin.Context) {
  removeDublicatesOfVisitors := makeUniqArrayOfVisitors(events)
  fmt.Println("Number of uniq visitors: ", len(removeDublicatesOfVisitors))
//  removeDublicatesOfItems := makeUniqArrayOfItems(events)
  visitors := make([] Visitor, len(removeDublicatesOfVisitors))

  /* make struct of visitors */
  initVisitors(visitors, removeDublicatesOfVisitors)

  /* add items to each visitor */
  addItemsToVisitor(visitors,events)
  addCountToEachProductOfEachVisitor(visitors)
  c.JSON(200,visitors)
}

func GetProducts (c *gin.Context) {
  items := make ([]ItemsGlobal, len(events))
  for i := 0; i < len(events); i++ {
    items[i].Itemid = events[i].Itemid
    items[i].Count = 1
  }
  c.JSON(200, items)
}
