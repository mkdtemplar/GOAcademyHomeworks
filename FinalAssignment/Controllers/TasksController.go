package Controllers

import (
	"FinalAssignment/Repository/DatabaseContext"
	"FinalAssignment/Repository/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FindTasks(c *gin.Context) {
	var task []Models.Tasks
	DatabaseContext.DB.Find(&task)

	c.JSON(http.StatusOK, task)
}

func CreateTask(c *gin.Context) {
	var list Models.Lists
	var input Models.CreateTask

	if err := DatabaseContext.DB.Where("id = ?", c.Param("id")).First(&list).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task := Models.Tasks{
		Text:      input.Text,
		ListId:    input.ListId,
		Completed: input.Completed,
	}
	DatabaseContext.DB.Create(&task)

	c.JSON(http.StatusOK, task)
}

func FindSingleTask(c *gin.Context) {
	var task Models.Tasks

	if err := DatabaseContext.DB.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, task)
}

func UpdateTask(c *gin.Context) {
	var task Models.Tasks
	if err := DatabaseContext.DB.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input Models.UpdateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	DatabaseContext.DB.Model(&task).Updates(Models.Tasks{
		Completed: input.Completed,
	})

	DatabaseContext.DB.Save(&task)
	c.JSON(http.StatusOK, task)
}

func DeleteTask(c *gin.Context) {

	var task Models.Tasks
	if err := DatabaseContext.DB.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	DatabaseContext.DB.Delete(&task)

	c.JSON(http.StatusOK, task)
}

func DeleteAllTasks(c *gin.Context) {
	DatabaseContext.DB.Exec(`DELETE FROM tasks`)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
