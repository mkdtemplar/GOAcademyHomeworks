package Controllers

import (
	taskRepo "FinalAssignment/Repository/TaskRepository"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type APIEnvTaskEcho struct {
	DB *gorm.DB
}

func (e APIEnvTaskEcho) GetAllTasks(ech echo.Context) error {
	tasks, err := taskRepo.GetAllTasks(e.DB)

	if err != nil {
		err = ech.JSON(http.StatusInternalServerError, err.Error())
		if err != nil {
			return err
		}
		return err
	}

	err = ech.JSON(http.StatusOK, tasks)
	if err != nil {
		return err
	}
	return nil
}

func (e APIEnvTaskEcho) GetOneTask(ech echo.Context) error {
	id, _ := strconv.Atoi(ech.Param("id"))

	task, exists, err := taskRepo.FindTaskById(id, e.DB)

	if err != nil {
		err = ech.JSON(http.StatusInternalServerError, err.Error())
		if err != nil {
			return err
		}
		return err
	}
	if !exists {
		err = ech.JSON(http.StatusNotFound, "There is no task in the db")
		if err != nil {
			return err
		}
		return err
	}

	err = ech.JSON(http.StatusOK, task)
	if err != nil {
		return err
	}
	return nil
}
