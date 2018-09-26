package models

import (
  "database/sql"
  "log"
  _ "github.com/lib/pq"
  "fmt"
)

var DB *sql.DB

func InitDB() bool{
  connection := "host=localhost port=5432 user=postgres password=postgres dbname=ColabFilter sslmode=disable"
  db, err := sql.Open("postgres", connection)
  if err != nil {
    log.Fatal(err)
    return false
  }
  DB = db
  return true
}

func CreateDB(db *sql.DB) {
  sql := `
  CREATE TABLE IF NOT EXISTS events(
    id BIGSERIAL PRIMARY KEY NOT NULL,
    timestamp VARCHAR(255) NOT NULL,
	  visitorid VARCHAR(255) NOT NULL,
		event VARCHAR(255) NOT NULL,
	  itemid VARCHAR(255) NOT NULL,
		transactionid VARCHAR(255)
  );
  CREATE TABLE IF NOT EXISTS persons(
    id integer PRIMARY KEY NOT NULL,
    name VARCHAR(255),
    surname VARCHAR(255),
    age integer,
    gender boolean
  );
  CREATE TABLE IF NOT EXISTS products(
    id integer PRIMARY KEY NOT NULL,
    name VARCHAR(255),
    cathegory VARCHAR(255),
    price real
  );
  CREATE TABLE IF NOT EXISTS recommends(
    user_id VARCHAR(255),
    recommendations float
  );
  CREATE TABLE IF NOT EXISTS visitors(
    visitor_id VARCHAR(255)
  );
  `
  _, err := db.Exec(sql)
  if err != nil {
    panic(err)
  }
}

func ClearDB( db *sql.DB) bool {
  stmt, err := db.Prepare("delete from events")
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

