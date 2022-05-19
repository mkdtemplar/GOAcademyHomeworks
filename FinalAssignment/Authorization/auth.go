package Authorization

import (
	"FinalAssignment/Repository/DatabaseContext"
	"FinalAssignment/Repository/Models"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"strings"
)

func BasicAuth() gin.HandlerFunc {

	return func(c *gin.Context) {
		auth := strings.SplitN(c.Request.Header.Get("Authorization"), " ", 2)

		if len(auth) != 2 || auth[0] != "Basic" {
			respondWithError(401, "Unauthorized", c)
			return
		}
		payload, _ := base64.StdEncoding.DecodeString(auth[1])
		pair := strings.SplitN(string(payload), ":", 2)

		if len(pair) != 2 || !checkUser(pair[0], pair[1]) {
			respondWithError(401, "Unauthorized", c)
			return
		}

		c.Next()
	}
}

func checkUser(username string, password string) bool {
	var user Models.User

	err := DatabaseContext.DB.Where(Models.User{
		Username: username,
		Password: password,
	}).First(&user)

	if err.Error != nil {
		return false
	}

	return true
}

func respondWithError(code int, message string, c *gin.Context) {
	resp := map[string]string{"error": message}

	c.JSON(code, resp)
	c.Abort()
}
