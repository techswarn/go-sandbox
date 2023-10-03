package main

import (
	"fmt"
//	"server/app/model"
	"server/app/db"
)

func main() {
//  model.CreateSchema()
  dbhandle, err := db.Connect()
  if err != nil {
	panic("Database connection failed")
  }
  db := db.Database{
	SqlDb: dbhandle,
 }
  
  db.Query()
  fmt.Println("Hello")
}
