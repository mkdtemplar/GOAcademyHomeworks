package main

import (
	CSV "FinalAssignment/CSV"
	controllers "FinalAssignment/Controllers"
	models "FinalAssignment/Models"
	"FinalAssignment/cmd"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {

	models.ConnectDatabase()

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.Use(func(ctx *gin.Context) {
		// This is a sample demonstration of how to attach middlewares in Gin
		log.Println("Gin middleware was called")
		ctx.Next()
	})

	// Tasks endpoints
	router.GET("/api/alltasks", controllers.FindTasks)

	router.GET("/api/lists/:id/tasks", controllers.FindSingleTask)

	router.POST("/api/lists/:id/tasks", controllers.CreateTask)

	router.PATCH("/api/tasks/:id", controllers.UpdateTask)

	router.DELETE("/api/tasks/:id", controllers.DeleteTask)

	router.DELETE("/api/DeleteAllTasks", controllers.DeleteAllTasks)

	// Lists endpoints
	router.GET("/api/lists", controllers.FindLists)

	router.POST("/api/lists", controllers.CreateList)

	router.DELETE("/api/lists/:id", controllers.DeleteList)

	router.GET("/api/list/export", CSV.ReadListRow)

	// Weather endpoint
	router.GET("/api/weather/:lat/:lon", controllers.GetWeather)

	// Do not touch this line!
	log.Fatal(http.ListenAndServe(":3000", cmd.CreateCommonMux(router)))
}
