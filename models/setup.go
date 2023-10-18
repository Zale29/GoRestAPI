//Koneksi DB

package models

import {
	"gorm.io/gorm"
	"gorm.io/drivers/mysql"
}

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/gorestapi"))
	if err := nil {
		panic(err)
	}

	database.AutoMigrate(&Product{})

	DB = database
}