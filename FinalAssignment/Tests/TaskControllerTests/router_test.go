package TaskControllerTests

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETTasksRoute(t *testing.T) {

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

func TestGetTaskIdRoute(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.GET("/api/lists/:id/tasks")

	req, err := http.NewRequest(http.MethodGet, "/api/lists/:id/tasks", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	// Create a response recorder so you can inspect the response
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	// Check to see if the response was what you expected
	if w.Code == http.StatusOK {
		t.Logf("Expected to get status %d is same as %d\n", http.StatusOK, w.Code)
	} else {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}
}

func TestPostTaskRoute(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.POST("/api/lists/:id/tasks")

	req, err := http.NewRequest(http.MethodPost, "/api/lists/:id/tasks", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	// Create a response recorder so you can inspect the response
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	// Check to see if the response was what you expected
	if w.Code == http.StatusOK {
		t.Logf("Expected to get status %d is same as %d\n", http.StatusOK, w.Code)
	} else {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}
}

func TestPatchTaskRoute(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.PATCH("/api/tasks/:id")

	req, err := http.NewRequest(http.MethodPatch, "/api/tasks/:id", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	// Create a response recorder so you can inspect the response
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	// Check to see if the response was what you expected
	if w.Code == http.StatusOK {
		t.Logf("Expected to get status %d is same as %d\n", http.StatusOK, w.Code)
	} else {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}
}

func TestDeleteRouteTaskById(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.DELETE("/api/tasks/:id")

	req, err := http.NewRequest(http.MethodDelete, "/api/tasks/:id", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	// Create a response recorder so you can inspect the response
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	// Check to see if the response was what you expected
	if w.Code == http.StatusOK {
		t.Logf("Expected to get status %d is same as %d\n", http.StatusOK, w.Code)
	} else {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}
}

func TestDeleteAllTasksRoute(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.DELETE("/api/DeleteAllTasks")

	req, err := http.NewRequest(http.MethodDelete, "/api/DeleteAllTasks", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	// Create a response recorder so you can inspect the response
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	// Check to see if the response was what you expected
	if w.Code == http.StatusOK {
		t.Logf("Expected to get status %d is same as %d\n", http.StatusOK, w.Code)
	} else {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}
}
