package models

import (
  "database/sql"
  "log"
  _ "github.com/lib/pq"
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
    timestamp VARCHAR(255) NOT NULL,
	  visitorid VARCHAR(255) NOT NULL,
		event VARCHAR(255) NOT NULL,
	  itemid VARCHAR(255) NOT NULL,
		transactionid VARCHAR(255)
  );
  `
  _, err := db.Exec(sql)
  if err != nil {
    panic(err)
  }
}
