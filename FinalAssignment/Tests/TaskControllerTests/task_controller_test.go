package TaskControllerTests

import (
	apiTask "FinalAssignment/Controllers"
	repo "FinalAssignment/Repository/DatabaseContext"
	models "FinalAssignment/Repository/Models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETTasks(t *testing.T) {

	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.GET("/api/alltasks")

	req, err := http.NewRequest(http.MethodGet, "/api/alltasks", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	// Create a response recorder so you can inspect the response
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	// Check to see if the response was what you expected
	if w.Code == http.StatusOK {
		t.Logf("Expected to get status %d is same ast %d\n", http.StatusOK, w.Code)
	} else {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}

}

func Test_GetTask_OK(t *testing.T) {
	a := assert.New(t)
	repo.ConnectDatabase()
	db := repo.GetDB()

	task, err := insertTestTask(db)
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

	actual := models.Tasks{}
	if err := json.Unmarshal(body, &actual); err != nil {
		a.Error(err)
	}

	expected := task

	a.Equal(expected, actual)
	ClearTable()
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

func ClearTable() {
	repo.DB.Exec(`DELETE FROM tasks`)
	repo.DB.Exec(`ALTER SEQUENCE id RESTART WITH 1`)
}

func insertTestTask(db *gorm.DB) (models.Tasks, error) {
	t := models.Tasks{
		Text:      "Test task",
		ListId:    1,
		Completed: false,
	}

	if err := db.Create(&t).Error; err != nil {
		return t, err
	}
	return t, nil
}
