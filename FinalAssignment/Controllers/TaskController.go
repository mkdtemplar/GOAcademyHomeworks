package Controllers

import (
	models "FinalAssignment/Repository/Models"
	taskRepo "FinalAssignment/Repository/TaskRepository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type APIEnvTask struct {
	DB *gorm.DB
}

func (a APIEnvTask) GetTasks(c *gin.Context) {
	tasks, err := taskRepo.GetAllTasks(a.DB)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func (a APIEnvTask) FindOneTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	task, exists, err := taskRepo.FindTaskById(id, a.DB)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	if !exists {
		c.JSON(http.StatusNotFound, "There is no task in the db")
		return
	}

	c.JSON(http.StatusOK, task)
}

func (a APIEnvTask) CreateTask(c *gin.Context) {
	task := models.Tasks{}

	err := c.BindJSON(&task)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := a.DB.Create(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, task)
}

func (a APIEnvTask) DeleteTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	_, exists, err := taskRepo.FindTaskById(id, a.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		c.JSON(http.StatusNotFound, "Record not found")
		return
	}

	err = taskRepo.DeleteTask(id, a.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Record deleted successfully")
}

func (a APIEnvTask) UpdatesTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	_, exists, err := taskRepo.FindTaskById(id, a.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		c.JSON(http.StatusNotFound, "Record not found")
		return
	}

	updatedTask := models.Tasks{}
	err = c.BindJSON(&updatedTask)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := taskRepo.UpdateTask(id, a.DB, &updatedTask); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	a.FindOneTask(c)
}

func (a APIEnvTask) DeleteAll(c *gin.Context) {
	taskRepo.DeleteAllTasks(a.DB)
	c.JSON(http.StatusOK, "Deleted all tasks")
}
