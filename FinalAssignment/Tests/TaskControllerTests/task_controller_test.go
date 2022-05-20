package TaskControllerTests

import (
	apiTask "FinalAssignment/Controllers"
	repo "FinalAssignment/Repository/DatabaseContext"
	models "FinalAssignment/Repository/Models"
	taskRepo "FinalAssignment/Repository/TaskRepository"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_GetTask_OK(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	a := assert.New(t)
	repo.ConnectDatabase()
	db := repo.GetDB()

	task, err := insertTestTask()
	if err != nil {
		a.Error(err)
	}

	req, w := setGetTasksRouter(db, "/api/lists/1/tasks")

	a.Equal(http.MethodGet, req.Method, "HTTP request method error")
	a.Equal(http.StatusOK, w.Code, "HTTP request status code error")
	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		a.Error(err)
	}

	actual, _, _ := taskRepo.FindTaskById(1, db)
	if err := json.Unmarshal(body, &actual); err != nil {
		a.Error(err)
	}

	expected := task

	a.Equal(expected, actual)
}

func TestCreateTask(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	a := assert.New(t)
	repo.ConnectDatabase()
	db := repo.GetDB()

	task := models.Tasks{
		Id:        8,
		Text:      "Task 6",
		ListId:    5,
		Completed: false,
	}

	reBody, err := json.Marshal(task)
	if err != nil {
		a.Error(err)
	}

	req, w, err := setCreateTaskRouter(db, bytes.NewBuffer(reBody))
	if err != nil {
		a.Error(err)
	}

	a.Equal(http.MethodPost, req.Method, "HTTP request method error")
	a.Equal(http.StatusOK, w.Code, "HTTP request status code error")

	body, err := ioutil.ReadAll(w.Body)

	if err != nil {
		a.Error(err)
	}

	actual := models.Tasks{}
	if err := json.Unmarshal(body, &actual); err != nil {
		a.Error(err)
	}

	expected := task

	a.Equal(expected, actual)
}

func setGetTasksRouter(db *gorm.DB, url string) (*http.Request, *httptest.ResponseRecorder) {
	r := gin.New()
	gin.SetMode(gin.TestMode)
	api := &apiTask.APIEnvTask{DB: db}
	r.GET("/api/lists/:id/tasks", api.GetTasks)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	return req, w
}

func setCreateTaskRouter(db *gorm.DB, body *bytes.Buffer) (*http.Request, *httptest.ResponseRecorder, error) {
	r := gin.New()
	gin.SetMode(gin.TestMode)
	api := &apiTask.APIEnvTask{DB: db}
	r.POST("/api/lists/:id/tasks", api.CreateTask)
	req, err := http.NewRequest(http.MethodPost, "/api/lists/:id/tasks", body)
	if err != nil {
		return req, httptest.NewRecorder(), err
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	return req, w, nil
}

func insertTestTask() (models.Tasks, error) {
	t := models.Tasks{
		Id:        1,
		Text:      "Task 1",
		ListId:    1,
		Completed: true,
	}
	return t, nil
}
