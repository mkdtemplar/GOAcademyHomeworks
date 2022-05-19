package DatabaseContext

import (
	"FinalAssignment/Repository/Models"
	_ "github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("D:\\GO\\GOAcademyHomeworks\\FinalAssignment\\Repository\\DatabaseContext\\Database\\test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	database.AutoMigrate(&Models.Tasks{}, &Models.Lists{}, &Models.User{})

	DB = database
}

func GetDB() *gorm.DB {
	return DB
}
