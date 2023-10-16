package config

import (
  "fmt"
  "log"
  "database/sql"
  
  _ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
  dbInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta", ENV.DB_HOST, ENV.DB_PORT, ENV.DB_USER, ENV.DB_PASSWORD, ENV.DB_NAME)
  db, err := sql.Open(ENV.DB_DRIVER, dbInfo)
  
  if err != nil {
    panic(err)
  }
  
  log.Printf("You are now connected to database '%s' as user '%s'\n", ENV.DB_NAME, ENV.DB_USER)
  DB = db
}
