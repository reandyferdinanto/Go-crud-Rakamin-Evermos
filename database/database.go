package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func DatabaseInit() {
	const MYSQL = "root:Iamagreatman8@tcp(127.0.0.1:3306)/trygofiber?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := MYSQL
  DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

  if err != nil {
	panic("Cannot connect database")
  }

  fmt.Println("Connected to database")
  
}