package Models

import (
	_ "github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Tasks{}, &Lists{})

	DB = database
}
