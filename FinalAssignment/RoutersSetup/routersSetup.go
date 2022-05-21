package RoutersSetup

import (
	auth "FinalAssignment/Authorization"
	"FinalAssignment/CSV"
	controllers "FinalAssignment/Controllers"
	handlers "FinalAssignment/Controllers"
	repo "FinalAssignment/Repository/DatabaseContext"
	"github.com/gin-gonic/gin"
	"log"
)

func Setup() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	apiTask := &handlers.APIEnvTask{DB: repo.GetDB()}
	apiList := &handlers.APIEnvList{DB: repo.GetDB()}
	apiUser := &handlers.APIEnvUser{DB: repo.GetDB()}

	router.Use(func(ctx *gin.Context) {
		// This is a sample demonstration of how to attach middlewares in Gin
		log.Println("Gin middleware was called")
		ctx.Next()
	})

	// Tasks endpoints
	router.GET("/api/alltasks", auth.BasicAuth(), apiTask.GetTasks)

	router.GET("/api/lists/:id/tasks", auth.BasicAuth(), apiTask.FindOneTask)

	router.POST("/api/lists/:id/tasks", auth.BasicAuth(), apiTask.CreateTask)

	router.PATCH("/api/tasks/:id", auth.BasicAuth(), apiTask.UpdatesTask)

	router.DELETE("/api/tasks/:id", auth.BasicAuth(), apiTask.DeleteTask)

	router.DELETE("/api/DeleteAllTasks", auth.BasicAuth(), apiTask.DeleteAll)

	//List endpoints
	router.GET("/api/list/export", auth.BasicAuth(), CSV.ReadListRow)

	router.GET("/api/lists", auth.BasicAuth(), apiList.GetAllLists)

	router.POST("/api/lists", auth.BasicAuth(), apiList.CreateList)

	router.DELETE("/api/lists/:id", auth.BasicAuth(), apiList.DeleteList)

	// Weather endpoint
	router.GET("/api/weather/:lat/:lon", auth.BasicAuth(), controllers.GetWeather)

	router.POST("/api/CreateUser", auth.BasicAuth(), apiUser.CreateUser)

	return router
}
