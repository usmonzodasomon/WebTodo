package handlers

import (
	"errors"
	"net/http"
	"strconv"
	"webtodo/models"
	"webtodo/pkg/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (h *handler) AddTask(c *gin.Context) {
	h.l.Println("Add Task")
	var task models.Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}
	userId, err := getUserId(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}
	task.UserId = userId
	id, err := h.services.Todo.Add(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (h *handler) GetTaskById(c *gin.Context) {
	h.l.Println("Get Task By Id")

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	userId, err := getUserId(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}

	task, err := h.services.Todo.GetTaskById(uint(id), userId)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"reason": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}

func (h *handler) GetExpiredTasksByUser(c *gin.Context) {
	h.l.Println("Get Expired Tasks By User")
	userId, err := getUserId(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}
	tasks, err := h.services.Todo.GetExpiredTasksByUser(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

func (h *handler) GetTasks(c *gin.Context) {
	h.l.Println("Get Tasks")
	userId, err := getUserId(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}
	tasks, err := h.services.Todo.GetAllTasks(userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"reason": "task not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		}
	}
	c.JSON(http.StatusOK, tasks)
}

type Numbers struct {
	TaskID    uint `json:"task_id"`
	NewUserId uint `json:"new_user_id"`
}

func (h *handler) ReassignTask(c *gin.Context) {
	h.l.Println("Reassign Task")
	var numbers Numbers
	if err := c.BindJSON(&numbers); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	err := h.services.Todo.ReassignTask(numbers.TaskID, numbers.NewUserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"reason": "success"})
}

func (h *handler) UpdateTask(c *gin.Context) {
	h.l.Println("Update Task")
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

	var input models.Update
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	err = h.services.Todo.UpdateTask(input, uint(id), userId.(uint))
	if err == repository.ErrTaskNotFound {
		c.JSON(http.StatusNotFound, gin.H{"reason": err.Error()})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"reason": "Задача успешно обновлена!"})
}

func (h *handler) DeleteTask(c *gin.Context) {
	h.l.Println("Delete Task")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}
	userId, err := getUserId(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}
	// task, err := repository.GetTaskById(uint(id), userId.(uint))
	// if err != nil {
	// 	c.JSON(http.StatusNotFound, gin.H{"reason": err.Error()})
	// 	return
	// }
	err = h.services.DeleteTask(uint(id), userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return

	}
	c.JSON(http.StatusOK, gin.H{"reason": "Задача успешно удалена"})
}
