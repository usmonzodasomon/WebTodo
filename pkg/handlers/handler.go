package handlers

import (
	"webtodo/pkg/service"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type handler struct {
	services *service.Service
	logs     *logrus.Logger
}

func NewHandler(services *service.Service, logs *logrus.Logger) *handler {
	return &handler{services, logs}
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
		tasks.PUT("/:id", h.UpdateTask)
		tasks.DELETE("/:id", h.DeleteTask)
	}

	return router
}
