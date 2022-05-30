package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Connect() {
	database, err := gorm.Open("mysql", "root:quddus1916@tcp(127.0.0.1:3306)/testdatabase?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println("error connecting to db")
		panic(err)
	} else {
		fmt.Println("connected to db")
	}
	db = database
}

func GetDB() *gorm.DB {
	Connect()

	return db
}
