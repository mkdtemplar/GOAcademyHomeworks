package Controllers

import (
	models "FinalAssignment/Repository/Models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
)

type APIEnvUser struct {
	DB *gorm.DB
}

func (u APIEnvUser) CreateUser(c *gin.Context) {
	user := models.User{}
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	user.Password = string(hashed)
	err = c.BindJSON(&user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := u.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}
