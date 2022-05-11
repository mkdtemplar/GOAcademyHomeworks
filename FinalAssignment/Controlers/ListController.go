package Controlers

import (
	models "FinalAssignment/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FindLists(c *gin.Context) {
	var list []models.Lists
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
	list := models.Lists{
		Name: input.Name,
	}
	models.DB.Create(&list)

	c.JSON(http.StatusOK, gin.H{"data": list})
}

func DeleteList(c *gin.Context) {

	var list models.Lists
	var task models.Tasks

	if err := models.DB.Where("id = ?", c.Param("id")).First(&list).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	if err := models.DB.Where("list_id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	models.DB.Delete(&task)
	models.DB.Delete(&list)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

func FindSingleListItem(c *gin.Context) {
	var list models.Tasks

	if err := models.DB.Where("id = ?", c.Param("id")).First(&list).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": list})
}
