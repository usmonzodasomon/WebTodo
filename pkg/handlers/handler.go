package handlers

import (
	"log"
	"webtodo/pkg/service"

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

	router.POST("/sign-up", h.SignUp)
	router.POST("/sign-in", h.SignIn)

	tasks := router.Group("/task", h.AuthMiddleware)
	{
		tasks.GET("/", h.GetTasks)
		tasks.GET("/:id", h.GetTaskById)
		tasks.GET("/expired_tasks_by_user", h.GetExpiredTasksByUser)
		tasks.POST("/", h.AddTask)
		tasks.POST("/reassign_task", h.ReassignTask)
		tasks.PUT("/:id", h.UpdateTask)
		tasks.DELETE("/:id", h.DeleteTask)
	}

	return router
}
