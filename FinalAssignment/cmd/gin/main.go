package main

import (
	controlers "final/Controlers"
	models "final/Models"
	"final/cmd"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {

	models.ConnectDatabase()
	router := gin.Default()

	router.Use(func(ctx *gin.Context) {
		// This is a sample demonstration of how to attach middlewares in Gin
		log.Println("Gin middleware was called")
		ctx.Next()
	})

	// Task endpoints
	router.GET("/api/tasks", controlers.FindTasks)

	router.GET("/api/lists/:id/tasks", controlers.FindSingleTask)

	router.POST("/api/lists/:id/tasks", controlers.CreateTask)

	router.PATCH("/api/tasks/:id", controlers.UpdateTask)

	router.DELETE("/api/tasks/:id", controlers.DeleteTask)

	// List endpoints
	router.GET("/api/lists", controlers.FindLists)

	router.POST("/api/lists", controlers.CreateList)

	router.DELETE("/api/lists/:id", controlers.DeleteList)

	// Do not touch this line!
	log.Fatal(http.ListenAndServe(":3000", cmd.CreateCommonMux(router)))
}
