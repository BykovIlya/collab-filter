package models

import (
  "fmt"
  "log"
  "github.com/lib/pq"
  "math/rand"
  "time"
)

/**
	the struct of visitors
 */
type Visitor struct {
  Visitorid_string string `json:"visitorid_string"`
  Items [] Items          `json:"items_array"`
}

type Person struct {
  id string                `json:"person_id"`
  name string             `json:"person_name"`
  surname string          `json:"person_surname"`
  age int64               `json:"person_age"`
  gender bool             `json:"person_gender"`
  properties string       `json:"person_properties"`
}

func ImportVisitorsToDB(visitor []Visitor) bool {
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

func InitPersons(ids []string) []Person {
  fmt.Println("INIT PERSONS")
  rand.Seed(time.Now().UnixNano())
  ps := []Person{}
  for i := 0; i < len(ids); i++ {
    p := Person{}
    p.id = ids[i]
    p.name = "user_name"
    p.surname = "user_surname"
    p.age = rand.Int63n(56) + 14
    gender := rand.Intn(2)
    if (gender == 1) {
      p.gender = false
    } else {
      p.gender = true;
    }
    p.properties = "nothing"
    ps = append(ps, p)
  }
  return ps
}

func ImportPersonsToDB(ps []Person) bool {
  fmt.Println(ps[2])
  fmt.Println("IMPORT PERSONS")
  db, err := DB.Begin()
  if err != nil {
    fmt.Println("Input Error: ", err)
    log.Fatal(err)
    return false
  }
  stmt, err := db.Prepare(pq.CopyIn("persons","id","name","surname","age","gender","properties"))
  for _, ev := range ps {
    _, err = stmt.Exec(ev.id,ev.name,ev.surname,ev.age,ev.gender,ev.properties)
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