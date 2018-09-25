package routes

import (
"github.com/gin-gonic/gin"
"net/http"
"fmt"
  "ColabFilter/colab-filter/backend/models"
  "log"
  "github.com/lib/pq"
)

var events []Events
var visitors []Visitor
var removeDublicatesOfVisitors []string
var removeDublicatesOfItems []string
var items []ItemsGlobal
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
  clearEventsDB()
  events = readingTransactionsFromFile(csvFileName)
  ImportEventsToDB(events)
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

/*----------------------------------------------------------------*/

func InsertOneEventToDB(event Events) {
  fmt.Println(event)

  row:=models.DB.QueryRow("INSERT INTO events(timestamp,visitorid,event,itemid,transactionid) VALUES($1,$2,$3,$4,$5)",
    event.Timestamp, event.Visitorid, event.Event_, event.Itemid, event.Transactionid)
  fmt.Println(row)
}

func ImportEventsToDB(events []Events) bool {
  db, err := models.DB.Begin()
  if err != nil {
    fmt.Println("Input Error: ", err)
    log.Fatal(err)
    return false
  }
  stmt, err := db.Prepare(pq.CopyIn("events","timestamp","visitorid","event","itemid","transactionid"))
  for _, ev := range events {
    _, err = stmt.Exec(ev.Timestamp, ev.Visitorid, ev.Event_, ev.Itemid, ev.Transactionid)
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

func clearEventsDB() bool {
  stmt, err := models.DB.Prepare("delete from events")
  if (err != nil) {
    panic(err)
  }
  defer stmt.Close()
  res, err := stmt.Exec()
  if (err != nil) {
    panic(err)
  }
  affect, err := res.RowsAffected()
  if (err != nil) {
    panic(err)
  }
  fmt.Println(affect," rows deleted")
  return true
}
