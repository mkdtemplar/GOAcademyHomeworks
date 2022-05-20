package ListControllerTests

import (
	apiList "FinalAssignment/Controllers"
	repo "FinalAssignment/Repository/DatabaseContext"
	models "FinalAssignment/Repository/Models"
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

func TestCreateList(t *testing.T) {
	gin.SetMode(gin.TestMode)
	a := assert.New(t)
	repo.ConnectDatabase()
	db := repo.GetDB()

	list := models.Lists{
		Id:   9,
		Name: "Test list",
	}

	reqBody, err := json.Marshal(list)
	if err != nil {
		a.Error(err)
	}

	req, w, err := setCreateListRouter(db, bytes.NewBuffer(reqBody))

	if err != nil {
		a.Error(err)
	}

	a.Equal(http.MethodPost, req.Method, "HTTP request error")
	a.Equal(http.StatusOK, w.Code, "HTTP status code error")

	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		a.Error(err)
	}

	actual := models.Lists{}
	if err := json.Unmarshal(body, &actual); err != nil {
		a.Error(err)
	}

	expected := list

	a.Equal(expected, actual)
}

func setCreateListRouter(db *gorm.DB, body *bytes.Buffer) (*http.Request, *httptest.ResponseRecorder, error) {
	r := gin.New()
	gin.SetMode(gin.TestMode)
	api := &apiList.APIEnvList{DB: db}

	r.POST("/api/lists", api.CreateList)

	req, err := http.NewRequest(http.MethodPost, "/api/lists", body)

	if err != nil {
		return req, httptest.NewRecorder(), err
	}

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	return req, w, nil
}
