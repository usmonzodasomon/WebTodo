package handlers

import (
	"github.com/gin-gonic/gin"
)

func (app *app) Routes() *gin.Engine {
	router := gin.Default()
	tasks := router.Group("/task", app.AuthMiddleware)
	{
		tasks.GET("/", app.GetTasks)
		tasks.GET("/:id", app.GetTaskById)
		tasks.GET("/expired_tasks_by_user", app.GetExpiredTasksByUser)
		tasks.POST("/", app.AddTask)
		tasks.POST("/reassign_task", app.ReassignTask)
		// tasks.PUT("/:id", app.UpdateTask)
		tasks.DELETE("/:id", app.DeleteTask)
	}

	// users := router.Group("/user")
	// {
	// 	users.GET("/", app.GetUsers)
	// 	// users.GET("/user_tasks/:id", app.GetUserTasks)
	// 	// users.GET("/user/:id", app.GetUserById)
	// 	// users.PUT("/user/:id", app.UpdateUser)
	// 	// users.DELETE("/user/:id", app.DeleteUser)
	// }
	router.POST("/sign-up", app.SignUp)
	router.POST("/sign-in", app.SignIn)

	return router
}
