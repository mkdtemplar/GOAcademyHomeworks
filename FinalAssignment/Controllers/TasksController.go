package Controllers

import (
	models "FinalAssignment/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FindTasks(c *gin.Context) {
	var task []models.Tasks
	models.DB.Find(&task)

	c.JSON(http.StatusOK, task)
}

func CreateTask(c *gin.Context) {
	var list models.Lists
	var input models.CreateTask

	if err := models.DB.Where("id = ?", c.Param("id")).First(&list).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task := models.Tasks{
		Text:      input.Text,
		ListId:    input.ListId,
		Completed: input.Completed,
	}
	models.DB.Create(&task)

	c.JSON(http.StatusOK, task)
}

func FindSingleTask(c *gin.Context) {
	var task models.Tasks

	if err := models.DB.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, task)
}

func UpdateTask(c *gin.Context) {
	var task models.Tasks
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

	c.JSON(http.StatusOK, task)
}

func DeleteTask(c *gin.Context) {

	var task models.Tasks
	if err := models.DB.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&task)

	c.JSON(http.StatusOK, task)
}

func DeleteAllTasks(c *gin.Context) {
	models.DB.Exec(`DELETE FROM tasks`)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
