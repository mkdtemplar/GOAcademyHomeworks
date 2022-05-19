package Controllers

import (
	"FinalAssignment/Repository/DatabaseContext"
	"FinalAssignment/Repository/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FindLists(c *gin.Context) {
	var list []Models.Lists
	DatabaseContext.DB.Find(&list)

	c.JSON(http.StatusOK, list)
}

func CreateList(c *gin.Context) {
	// Validate input
	var input Models.CreateLists
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create list
	list := Models.Lists{
		Name: input.Name,
	}
	DatabaseContext.DB.Create(&list)

	c.JSON(http.StatusOK, list)
}

func DeleteList(c *gin.Context) {

	var list Models.Lists
	var task Models.Tasks

	if err := DatabaseContext.DB.Where("id = ?", c.Param("id")).First(&list).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	if err := DatabaseContext.DB.Where("list_id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	DatabaseContext.DB.Delete(&task)
	DatabaseContext.DB.Delete(&list)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

func FindSingleListItem(c *gin.Context) {
	var list Models.Tasks

	if err := DatabaseContext.DB.Where("id = ?", c.Param("id")).First(&list).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, list)
}
