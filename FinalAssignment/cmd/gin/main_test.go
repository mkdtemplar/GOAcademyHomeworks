package main

import (
	controller "FinalAssignment/Controllers"
	modules "FinalAssignment/Models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFindTasks(t *testing.T) {
	_, err := gorm.Open(sqlite.Open("Models/Database/test.db"), &gorm.Config{})

	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.GET("/api/alltasks", controller.FindTasks)

	req, err := http.NewRequest(http.MethodGet, "/api/alltasks", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	// Create a response recorder so you can inspect the response
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	var task []modules.Tasks
	json.Unmarshal([]byte(w.Body.String()), &task)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, 492, len(task))
	// Check to see if the response was what you expected
	if w.Code == http.StatusOK {
		t.Logf("Expected to get status %d is same ast %d\n", http.StatusOK, w.Code)
	} else {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}

	//log.Fatal(http.ListenAndServe(":3000", cmd.CreateCommonMux(r)))
}
