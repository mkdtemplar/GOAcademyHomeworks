package Controlers

import (
	models "final/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GET /api
// Get all Task
func FindTasks(c *gin.Context) {
	var task []models.Task
	models.DB.Find(&task)

	c.JSON(http.StatusOK, gin.H{"data": task})
}

func CreateTask(c *gin.Context) {
	// Validate input
	var input models.CreateTask
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task := models.Task{
		Text: input.Text,
		//ListId:    input.ListId,
		//Completed: input.Completed,
	}
	models.DB.Create(&task)

	c.JSON(http.StatusOK, gin.H{"data": task})
}

func FindSingleTask(c *gin.Context) {
	var task models.Task

	if err := models.DB.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": task})
}

func UpdateTask(c *gin.Context) {
	var task models.Task
	if err := models.DB.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input models.UpdateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&task).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": task})
}

func DeleteTask(c *gin.Context) {

	var task models.Task
	if err := models.DB.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&task)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
