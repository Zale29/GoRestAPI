//Koneksi DB

package models

import (
	"github.com/gorm.io/gorm"
	"github.com/gorm.io/drivers/mysql"

)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/GoRestApi"))
	if err := nil {
		panic(err)
	}

	database.AutoMigrate(&Product{})

	DB = database
}