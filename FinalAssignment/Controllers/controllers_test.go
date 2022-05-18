package Controllers

import (
	models "FinalAssignment/Models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestFindTasks(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	FindTasks(c)
	assert.Equal(t, 200, w.Code) // or what value you need it to be

	var got gin.H
	err := json.Unmarshal(w.Body.Bytes(), &got)
	if err != nil {
		t.Fatal(err)
	}

	var want models.Tasks

	assert.Equal(t, want, got)
}
