package TaskControllerTests

import (
	"github.com/gin-gonic/gin"
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
