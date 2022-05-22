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

	err := c.BindJSON(&user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	password := user.Password

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 500,
			"MSG":  "encryption error",
		})
		return
	}

	newUser := models.User{
		Username: user.Username,
		Password: string(hashPassword),
	}

	if err = u.DB.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Registration is successfully")
}
