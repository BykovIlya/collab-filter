package models

import (
  "fmt"
  "log"
  "github.com/lib/pq"
)

/**
	the struct of visitors
 */
type Visitor struct {
  Visitorid_string string `json:"visitorid_string"`
  Items [] Items          `json:"items_array"`
}

type Person struct {
  id int64                `json:"person_id"`
  name string             `json:"person_name"`
  surname string          `json:"person_surname"`
  age int64               `json:"person_age"`
  gender bool             `json:"person_gender"`
}

func ImportPersonsToDB(visitor []Visitor) bool {
  db, err := DB.Begin()
  if err != nil {
    fmt.Println("Input Error: ", err)
    log.Fatal(err)
    return false
  }
  stmt, err := db.Prepare(pq.CopyIn("visitors","visitor_id"))
  for _, ev := range visitor {
    _, err = stmt.Exec(ev.Visitorid_string)
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