package Controllers

import (
	listRepo "FinalAssignment/Repository/ListRepository"
	models "FinalAssignment/Repository/Models"
	taskRepo "FinalAssignment/Repository/TaskRepository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type APIEnvList struct {
	DB *gorm.DB
}

func (l APIEnvList) CreateList(c *gin.Context) {
	list := models.Lists{}
	err := c.BindJSON(&list)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err = l.DB.Create(&list).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, list)
}

func (l APIEnvList) GetAllLists(c *gin.Context) {
	lists, err := listRepo.FindAllLists(l.DB)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, lists)
}

func (l APIEnvList) DeleteList(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	_, exists, err := listRepo.GetListById(id, l.DB)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !exists {
		c.JSON(http.StatusNotFound, "There is no list in the db")
		return
	}

	_, existsTask, errTask := taskRepo.FindTaskById(id, l.DB)

	if errTask != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !existsTask {
		c.JSON(http.StatusNotFound, "There is no task in the db")
		return
	}

	err = listRepo.DeleteList(id, l.DB)
	errTask = taskRepo.DeleteTask(id, l.DB)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if errTask != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "Record deleted successfully")
}
