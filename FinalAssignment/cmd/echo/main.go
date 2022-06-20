package main

import (
	handlers "FinalAssignment/Controllers"
	models "FinalAssignment/Repository/DatabaseContext"
	"github.com/labstack/echo/v4/middleware"

	//auth "FinalAssignment/Authorization"
	repo "FinalAssignment/Repository/DatabaseContext"
	"FinalAssignment/cmd"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func main() {
	models.ConnectDatabase()
	apiTask := &handlers.APIEnvTaskEcho{DB: repo.GetDB()}

	router := echo.New()
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())
	router.Use(middleware.CORS())
	router.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		// This is a sample demonstration of how to attach middlewares in Echo
		return func(ctx echo.Context) error {
			log.Println("Echo middleware was called")
			return next(ctx)
		}
	})

	// Add your handler (API endpoint) registrations here
	router.GET("/api/", func(ctx echo.Context) error {
		return ctx.JSON(200, "Hello, World!")
	})

	router.GET("/api/alltasks", apiTask.GetAllTasks)

	router.GET("/api/lists/:id/tasks", apiTask.GetOneTask)

	router.POST("/api/lists/:id/tasks", apiTask.CreateTask)

	// Do not touch this line!
	log.Fatal(http.ListenAndServe(":3000", cmd.CreateCommonMux(router)))
}
