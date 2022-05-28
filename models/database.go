package models

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("CONECTADO---------------")
	}
	db.AutoMigrate(&User{})

	/*
	db.Create(&User{Username: "Anton", Password: "1234", Email: "anton@gmail.es"})
	db.Create(&User{Username: "Dani", Password: "1234", Email: "Dani@gmail.es"})
	*/
	DB = db
}
