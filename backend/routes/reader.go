package routes

import (
	"math"

	"os"
	"encoding/csv"
	"bufio"
	"io"
	"log"
	"sort"
  "fmt"
)

/**
	the struct of events
 */
type Events struct {
		Timestamp string /*int64*/ `json:"timestamp"`
	  Visitorid string /*int64*/  `json:"visitorid"`
		Event_ string /*object*/ `json:"event"`
	  Itemid string /*int64*/  `json:"itemid"`
		Transactionid string /*float64*/ `json:"transactionid"`
}

/**
	the struct of items
 */
type Items struct{
	Itemid_string string  `json:"itemid_string"`
	Itemid_count float64  `json:"itemid_count"`
}

type ItemsGlobal struct {
  Itemid string `json:"itemid"`
  Count int64 `json:"count"`
}
/**
	the struct of visitors
 */
type Visitor struct {
	Visitorid_string string `json:"visitorid_string"`
	Items [] Items          `json:"items"`
}

/**
	reading data from .csv
 */
func readingTransactionsFromFile(csvFileName string) []Events {
	csvFile, err := os.Open(csvFileName)
  if err != nil {
    //error
  }
  fmt.Println("success open file!")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var events []Events
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		//if line[2] == "transaction" {
		var event = Events{}
    event.Timestamp = line[0]
		event.Visitorid = line[1]
    event.Event_ = line[2]
    event.Itemid = line[3]
    event.Transactionid = line[4]

    events = append(events, event)

			//events = append(events, &Events{
			//  timestamp:     line[0],
			//	visitorid:     line[1],
			//	event_ :       line[2],
			//	itemid:        line[3],
			//	transactionid: line[4],
			//})
	//	}

	}

	return events
}

func makeUniqArrayOfVisitors(events []Events) []string {
	bufOfVisitors := make ([] string, len(events))
	for i := 0; i < len(events); i++ {
		bufOfVisitors[i] = events[i].Visitorid
	}
	sort.Strings(bufOfVisitors)
	removeDublicatesOfVisitors := removeDuplicates(bufOfVisitors)
	return removeDublicatesOfVisitors
}

func makeUniqArrayOfItems(events []Events) [] string {
	bufOfItems := make ([] string, len(events))
	for i := 0; i < len(events); i++ {
		bufOfItems[i] = events[i].Itemid
	}
	sort.Strings(bufOfItems)
	removeDublicatesOfItems := removeDuplicates(bufOfItems)
	return removeDublicatesOfItems
}

func makeMatrixOfSales (visitors [] Visitor, removeDublicatesOfVisitors [] string, removeDublicatesOfItems [] string) [][] float64{
	/*
		init matrix
	 */
	matrixOfSales := make([][] float64, len(removeDublicatesOfVisitors))
	for i := 0; i < len(removeDublicatesOfVisitors); i++  {
		matrixOfSales[i] = make([] float64, len(removeDublicatesOfItems))
	}
	/*
	make matrix
	 */
	for i := 0; i < len(removeDublicatesOfVisitors); i++ {
		for j := 0; j < len(visitors[i].Items); j++ {
			//if visitors[i].items[j].itemid_count > 0 {
			matrixOfSales[i][getIndItem(removeDublicatesOfItems,visitors[i].Items[j].Itemid_string)] = visitors[i].Items[j].Itemid_count;
			//}
		}
	}
	return matrixOfSales
}

func makeArrayOfSales (matrixOfSales [][] float64, n int, m int) [] float64 {
	arrayOfSales := make ([]float64, 0)
	arrayOfSales = toArray(matrixOfSales, n, m, arrayOfSales)
	return arrayOfSales
}
func addCountToEachProductOfEachVisitor (visitors [] Visitor) {
	for i := 0; i < len(visitors); i++  {
		sort.Slice(visitors[i].Items, func(j, k int) bool { return visitors[i].Items[j].Itemid_string < visitors[i].Items[k].Itemid_string })
	}
	for i := 0; i < len(visitors); i++ {
		visitors[i].Items = findCount(visitors[i].Items)
	}
}
/**
	get index of visitor
 */
func getIndVisitor (visitor [] Visitor, finder string) int {
	for i := 0; i < len(visitor); i++ {
		if visitor[i].Visitorid_string == finder {
			return i
		}
	}
	return -1
}

/**
	get index of item
 */
func getIndItem (items [] string, finder string) int {
	for i := 0; i < len(items); i++ {
		if items[i] == finder {
			return i
		}
	}
	return -1
}

/**
	set the field visitorid_strnig of the structure Visitor to the value of unique visitors from the array buffer
 */
func initVisitors (visitor [] Visitor, buffer [] string) {
	for i := 0; i < len(buffer); i++ {
		visitor[i].Visitorid_string =  buffer[i]
	}
}

/**
	set each visitor an array of items
 */
func addItemsToVisitor (visitor [] Visitor, events []Events){
	for i := 0; i < len(visitor); i++ {
		for j := 0; j < len(events); j++ {
			if visitor[i].Visitorid_string == events[j].Visitorid {
				visitor[i].Items = append(visitor[i].Items, Items{
					Itemid_string: events[j].Itemid,
					Itemid_count: 1,
				})
			}
		}
	}
}

/*
func findVisitorInEvents(events []*Events, finder string) int {
	for i := 0; i < len(events); i++ {
		if events[i].visitorid == finder {
			return i
		}
	}
	return -1
}

func findItemsInEvents (events []*Events, finder string) int {
	for i := 0; i < len(events); i++ {
		if events[i].itemid == finder {
			return i
		}
	}
	return -1
}
*/

/**
	remove dublicates from visitors and itmes for make uniq arrays
 */
func removeDuplicates(array [] string) [] string{
	if len(array) == 1 || len(array) == 0 {
		return array
	}
	unique := 1
	for i := 1; i < len(array); i++{
		if array[i] != array[i - 1] {
			unique++;
		}
	}
	result := make([] string, unique)
	k := 0;
	if len(result) > 0 {
		result[k] = array[0]
		k++
	}
	for i := 1; i < len(array); i++ {
		if array[i] != array[i - 1] {
			result[k] = array[i];
			k++
		}
	}
	return result;
}

/**
	convert matrix to array
 */

func toArray (matrix [][] float64, n int, m int, array [] float64) []float64 {
	for i := 0; i < n ; i++  {
		for j := 0; j < m; j++ {
			array = append(array, matrix[i][j])
		}
	}
	return array
}

/*
func initCountToResult (item []Items) {
	for i := 0; i < len(item); i++ {
		item[i].itemid_count = 1
	}
}
*/

/**
	find count of each items in array of items for each visitor
 */
func findCount (item []Items) [] Items{
	buffer := make( [] Items, 0);
	var prev string
	for i := 0; i < len(item); i++ {
		if (item[i].Itemid_string != prev) {
			buffer = append(buffer, Items {
				item[i].Itemid_string,
				1,
			})
		} else {
			buffer[len(buffer) - 1].Itemid_count++
		}
		prev = item[i].Itemid_string
	}
	return buffer
}

/*
func removeDuplicatesInItems(item []Items) []Items {
	unique := 1
	for i := 1; i < len(item); i++{
		if item[i] != item[i - 1] {
			unique++;
		}
	}
	result := make([]Items, unique)
	initCountToResult(result)
	k := 0;
	if len(result) > 0 {
		result[k].itemid_string = item[0].itemid_string
		k++
	}
	for i := 1; i < len(item); i++ {
		if item[i].itemid_string != item[i - 1].itemid_string {
			result[k].itemid_string = item[i].itemid_string;
			result[k].itemid_count++
			k++
		}
	}
	return result;
}
*/
/**
find element from array
 */
/*func find(buf []*Events, events []*Events, visitor []*Visitor) {
	for i := 1; i < len(buf); i++ {
		if buf[i].visitorid == buf[0].visitorid {
			resultInd := findInEvents(events, buf[i].visitorid)
			var itemsBuf []*Items
			itemsBuf = append(itemsBuf,&Items{
				itemid_string: buf[i].itemid,
				itemid_int: resultInd,
			})
			visitor = append(visitor, &Visitor{
				buf[i].visitorid,
				resultInd,
				itemsBuf,
			})
			remove(buf, buf[i])
		}
	}
	remove(buf,buf[0])
}
*/


/**
delete elem from array
 */
/*
func remove(list []*Events, item *Events) []*Events {
   for i, v := range list {
	   if v == item {
		   copy(list[i:], list[i+1:])
		   list[len(list)-1] = nil
		   list = list[:len(list)-1]
	   }
   }
   return list
}
*/


/*
	remove unnecessary elements from score array
 */
func optimizeScores(scores [] float64, good [] float64) []float64{
	for i := 0; i < len(scores); i++ {
		if !math.IsNaN(scores[i]) {
			good = append(good, scores[i])
		}
	}
	return good
}
