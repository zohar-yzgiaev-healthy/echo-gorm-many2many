package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"gorm-many2many/models"
)

var database *gorm.DB

func Init() {
	database, connectionError := gorm.Open("postgres", "host=localhost port=5432 user=zoharyzgeav dbname=many2many sslmode=disable")
	if connectionError != nil {
		panic(connectionError)
	}

	database.AutoMigrate(&models.Feature{})
	database.AutoMigrate(&models.Photo{})
	database.AutoMigrate(&models.Place{})
}

// Get TODO: save database instance in context to eliminate the Get() method
func Get() *gorm.DB {
	database, _ := gorm.Open("postgres", "host=localhost port=5432 user=zoharyzgeav dbname=many2many sslmode=disable")
	return database
}