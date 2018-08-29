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
var myVisitor string

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
  removeDublicatesOfVisitors = makeUniqArrayOfVisitors(events)
  removeDublicatesOfItems = makeUniqArrayOfItems(events)
  visitors = make([] Visitor, len(removeDublicatesOfVisitors))
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

func GetPerson (c *gin.Context) {
  myVisitor = c.GetString("myVisitor")
  println(myVisitor)
}
func GetRecommends (c *gin.Context) {
  matrixOfSales := makeMatrixOfSales(visitors, removeDublicatesOfVisitors, removeDublicatesOfItems)

  /* init array of sales to get it into CA */
  arrayOfSales := makeArrayOfSales(matrixOfSales, len(removeDublicatesOfVisitors), len(removeDublicatesOfItems) )

  /* CA algorithm*/
  prefs := MakeRatingMatrix(arrayOfSales, len(removeDublicatesOfVisitors), len(removeDublicatesOfItems))
  //products := removeDublicatesOfItems
  products := make([]string, 0)
  for i := 0; i < len(removeDublicatesOfItems); i++ {
    products = append(products,strconv.Itoa(i))
  }
  //for i := 0; i < /*len(removeDublicatesOfVisitors)*/ 2; i++ {

  /*
  scanner := bufio.NewScanner(os.Stdin)
  scanner.Scan()
  myVisitor = scanner.Text()
  */

  indexOfVisitor := getIndVisitor(visitors, myVisitor)
  if (indexOfVisitor == -1) {
    fmt.Println("Error: visitor doesn't found!")
   // os.Exit(-1)
  }
  recommendations, err := GetRecommendations(prefs, /*getIndVisitor(visitors, myVisitor)*/ indexOfVisitor, products)
  if err != nil {
    fmt.Println("WHAT!?")
  }
  /*if len(recommendations) > 0 {
    fmt.Println("For user ", myVisitor, " recommended products are x with scores y (x --> y)")
    for i := 0; i < len(recommendations); i++ {
      fmt.Println(recommendations[i].Product, "-->", recommendations[i].MpRating)
    }
  } else {
    fmt.Println("There are no recommendations for user ", myVisitor)
  }*/

  c.JSON(200, recommendations)
}
