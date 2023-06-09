package handlers

import (
	"log"
	"webtodo/service"

	"github.com/gin-gonic/gin"
)

type handler struct {
	services *service.Service
	l        *log.Logger
}

func NewHandler(services *service.Service, l *log.Logger) *handler {
	return &handler{services, l}
}

func (h *handler) Routes() *gin.Engine {
	router := gin.New()
	tasks := router.Group("/task", h.AuthMiddleware)
	{
		tasks.GET("/", h.GetTasks)
		tasks.GET("/:id", h.GetTaskById)
		tasks.GET("/expired_tasks_by_user", h.GetExpiredTasksByUser)
		tasks.POST("/", h.AddTask)
		tasks.POST("/reassign_task", h.ReassignTask)
		// tasks.PUT("/:id", h.UpdateTask)
		tasks.DELETE("/:id", h.DeleteTask)
	}

	// users := router.Group("/user")
	// {
	// 	users.GET("/", h.GetUsers)
	// 	// users.GET("/user_tasks/:id", h.GetUserTasks)
	// 	// users.GET("/user/:id", h.GetUserById)
	// 	// users.PUT("/user/:id", h.UpdateUser)
	// 	// users.DELETE("/user/:id", h.DeleteUser)
	// }
	router.POST("/sign-up", h.SignUp)
	router.POST("/sign-in", h.SignIn)

	return router
}
