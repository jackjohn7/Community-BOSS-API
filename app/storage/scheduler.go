package storage

import (
	"database/sql"
	"fmt"
  _ "github.com/lib/pq"
)

var db *sql.DB

func InitBossDataDB(connectionString string) *sql.DB {
  newDB, err := sql.Open("postgres", connectionString)

  if err != nil {
    panic(err)
  }

  if err = newDB.Ping(); err != nil {
    panic(err)
  }

  fmt.Println("The database is connected")
  db = newDB
  return newDB
}

func GetBossDB() *sql.DB {
  return db
}

