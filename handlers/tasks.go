package handlers

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"webtodo/models"
	"webtodo/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// import (
// 	"database/sql"
// 	"errors"
// 	"log"
// 	"net/http"
// 	"strconv"
// 	"webtodo/models"
// 	"webtodo/repository"

// 	"github.com/gin-gonic/gin"
// 	"gorm.io/gorm"
// )

type app struct {
	l  *log.Logger
	db *gorm.DB
}

func NewTasks(l *log.Logger, db *gorm.DB) *app {
	return &app{l, db}
}

func (t *app) GetTaskById(c *gin.Context) {
	t.l.Println("Get Task By Id")

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}
	userId, ok := c.Get("userId")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": "error getting userId from header"})
		return
	}
	task, err := service.GetTaskById(uint(id), userId.(uint))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"reason": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}

func (t *app) GetExpiredTasksByUser(c *gin.Context) {
	t.l.Println("Get Expired Tasks By User")
	userId, ok := c.Get("userId")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": "error getting userId from header"})
		return
	}
	tasks, err := service.GetExpiredTasksByUser(userId.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

func (t *app) GetTasks(c *gin.Context) {
	t.l.Println("Get Tasks")
	userId, ok := c.Get("userId")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": "error getting userId from header"})
		return
	}
	tasks, err := service.GetAllTasks(userId.(uint))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"reason": "task not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		}
	}
	c.JSON(http.StatusOK, tasks)
}

func (t *app) AddTask(c *gin.Context) {
	t.l.Println("Add Task")
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}
	userId, ok := c.Get("userId")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": "user id not found in auth header"})
		return
	}
	task.UserId = userId.(uint)
	id, err := service.AddTask(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

type Numbers struct {
	TaskID    uint `json:"task_id"`
	NewUserId uint `json:"new_user_id"`
}

func (t *app) ReassignTask(c *gin.Context) {
	t.l.Println("Reassign Task")
	var numbers Numbers
	if err := c.BindJSON(&numbers); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	err := service.ReassignTask(numbers.TaskID, numbers.NewUserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"reason": "success"})
}

// func (t *app) UpdateTask(c *gin.Context) {
// 	t.l.Println("Update Task")
// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
// 		return
// 	}
// 	userId, ok := c.Get("userId")
// 	if !ok {
// 		c.JSON(http.StatusInternalServerError, gin.H{"reason": "error getting userId from header"})
// 		return
// 	}
// 	// task, err := service.GetTaskById(uint(id), userId.(uint))
// 	// if err != nil {
// 	// 	c.JSON(http.StatusNotFound, gin.H{"reason": err.Error()})
// 	// 	return
// 	// }
// 	// if err := c.BindJSON(&task); err != nil {
// 	// 	c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
// 	// 	return
// 	// }

// 	err = service.UpdateTask(uint(id), userId.(uint))
// 	if err == repository.ErrTaskNotFound {
// 		c.JSON(http.StatusNotFound, gin.H{"reason": err.Error()})
// 		return
// 	}
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"reason": "Задача успешно обновлена!"})
// }

func (t *app) DeleteTask(c *gin.Context) {
	t.l.Println("Delete Task")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}
	userId, ok := c.Get("userId")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": "error getting userId from header"})
		return
	}
	// task, err := repository.GetTaskById(uint(id), userId.(uint))
	// if err != nil {
	// 	c.JSON(http.StatusNotFound, gin.H{"reason": err.Error()})
	// 	return
	// }
	err = service.DeleteTask(uint(id), userId.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return

	}
	c.JSON(http.StatusOK, gin.H{"reason": "Задача успешно удалена"})
}
