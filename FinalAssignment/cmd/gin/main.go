package main

import (
	models "FinalAssignment/Repository/DatabaseContext"
	routers "FinalAssignment/RoutersSetup"
	"FinalAssignment/cmd"
	"log"
	"net/http"
)

func main() {

	models.ConnectDatabase()
	router := routers.Setup()
	/*
		gin.SetMode(gin.ReleaseMode)

		router := gin.Default()

		router.Use(func(ctx *gin.Context) {
			// This is a sample demonstration of how to attach middlewares in Gin
			log.Println("Gin middleware was called")
			ctx.Next()
		})

		// Tasks endpoints
		router.GET("/api/alltasks", auth.BasicAuth(), controllers.FindTasks)

		router.GET("/api/lists/:id/tasks", auth.BasicAuth(), controllers.FindSingleTask)

		router.POST("/api/lists/:id/tasks", auth.BasicAuth(), controllers.CreateTask)

		router.PATCH("/api/tasks/:id", auth.BasicAuth(), controllers.UpdateTask)

		router.DELETE("/api/tasks/:id", auth.BasicAuth(), controllers.DeleteTask)

		router.DELETE("/api/DeleteAllTasks", auth.BasicAuth(), controllers.DeleteAllTasks)

		// Lists endpoints
		router.GET("/api/lists", auth.BasicAuth(), controllers.FindLists)

		router.POST("/api/lists", auth.BasicAuth(), controllers.CreateList)

		router.DELETE("/api/lists/:id", auth.BasicAuth(), controllers.DeleteList)

		router.GET("/api/list/export", auth.BasicAuth(), CSV.ReadListRow)

		// Weather endpoint
		router.GET("/api/weather/:lat/:lon", auth.BasicAuth(), controllers.GetWeather)
	*/
	// Do not touch this line!
	log.Fatal(http.ListenAndServe(":3000", cmd.CreateCommonMux(router)))

}
