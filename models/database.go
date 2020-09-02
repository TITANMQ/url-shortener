package models

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	connection, err := gorm.Open(sqlite.Open("./database.db"), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
	}

	db = connection
	db.AutoMigrate(&URL{})
	fmt.Println("Connected to database")
	
	// var url []URL
	// err = db.Find(&url).Error
	// if err != nil{
	// 	fmt.Println(err)
	// }
	// fmt.Println(url)
}

//GetDB returns the database
func GetDB() *gorm.DB {
	return db
}
