package RoutersSetup

import (
	auth "FinalAssignment/Authorization"
	handlers "FinalAssignment/Controllers"
	repo "FinalAssignment/Repository/DatabaseContext"
	"github.com/gin-gonic/gin"
	"log"
)

func Setup() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	api := &handlers.APIEnv{DB: repo.GetDB()}

	router.Use(func(ctx *gin.Context) {
		// This is a sample demonstration of how to attach middlewares in Gin
		log.Println("Gin middleware was called")
		ctx.Next()
	})

	router.GET("/api/alltasks", auth.BasicAuth(), api.GetTasks)

	router.GET("/api/lists/:id/tasks", auth.BasicAuth(), api.FindOneTask)

	router.POST("/api/lists/:id/tasks", auth.BasicAuth(), api.CreateTask)

	router.PATCH("/api/tasks/:id", auth.BasicAuth(), api.UpdatesTask)

	router.DELETE("/api/tasks/:id", auth.BasicAuth(), api.DeleteTask)

	router.DELETE("/api/DeleteAllTasks", auth.BasicAuth(), api.DeleteAll)
	return router
}
