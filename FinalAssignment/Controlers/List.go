package Controlers

import (
	models "final/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FindLists(c *gin.Context) {
	var list []models.List
	models.DB.Find(&list)

	c.JSON(http.StatusOK, gin.H{"data": list})
}

func CreateList(c *gin.Context) {
	// Validate input
	var input models.CreateLists
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create list
	list := models.List{
		Name: input.Name,
	}
	models.DB.Create(&list)

	c.JSON(http.StatusOK, gin.H{"data": list})
}

func FindSingleListItem(c *gin.Context) {
	var list models.Task

	if err := models.DB.Where("id = ?", c.Param("id")).First(&list).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": list})
}
