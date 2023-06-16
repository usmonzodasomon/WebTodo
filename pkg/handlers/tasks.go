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
	h.logs.Info("Adding Task")
	var task models.Task
	if err := c.BindJSON(&task); err != nil {
		h.logs.Error("Failed to bind JSON: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userId, err := getUserId(c)
	if err != nil {
		h.logs.Error("Failed to get User Id: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	task.UserId = userId
	id, err := h.services.Todo.Add(&task)
	if err != nil {
		h.logs.Error("Failed to add task: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
	h.logs.Infof("Task with Id %v added succesfully", id)
}

func (h *handler) GetTaskById(c *gin.Context) {
	h.logs.Info("Getting Task By Id")

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		h.logs.Error("Failed to get id: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, err := getUserId(c)
	if err != nil {
		h.logs.Error("Failed to get User by Id: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	task, err := h.services.Todo.GetTaskById(uint(id), userId)

	if err != nil {
		h.logs.Error("Failed to get Task by Id: ", err)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
	h.logs.Infof("Task with Id %v received succesfully", id)
}

func (h *handler) GetExpiredTasksByUser(c *gin.Context) {
	h.logs.Info("Getting Expired Tasks By User")
	userId, err := getUserId(c)
	if err != nil {
		h.logs.Error("Failed to get User Id: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	tasks, err := h.services.Todo.GetExpiredTasksByUser(userId)
	if err != nil {
		h.logs.Error("Failed to get Expired Tasks by User: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
	h.logs.Info("Expired tasks received succesfully")
}

func (h *handler) GetTasks(c *gin.Context) {
	h.logs.Info("Getting Tasks")
	userId, err := getUserId(c)
	if err != nil {
		h.logs.Error("Failed to get User Id: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		return
	}
	tasks, err := h.services.Todo.GetAllTasks(userId)
	if err != nil {
		h.logs.Error("Failed to get all tasks: ", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"reason": "task not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
		}
	}
	c.JSON(http.StatusOK, tasks)
	h.logs.Info("All tasks received succesfully")
}

type Numbers struct {
	TaskID    uint `json:"task_id"`
	NewUserId uint `json:"new_user_id"`
}

// func (h *handler) ReassignTask(c *gin.Context) {
// 	h.l.Println("Reassign Task")
// 	var numbers Numbers
// 	if err := c.BindJSON(&numbers); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
// 		return
// 	}

// 	err := h.services.Todo.ReassignTask(numbers.TaskID, numbers.NewUserId)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"reason": "success"})
// }

func (h *handler) UpdateTask(c *gin.Context) {
	h.logs.Info("Updating Task")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logs.Error("Failed to get Id: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userId, err := getUserId(c)
	if err != nil {
		h.logs.Error("Failed to get User Id: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var input models.Update
	if err := c.BindJSON(&input); err != nil {
		h.logs.Error("Failed to bind JSON: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.services.Todo.UpdateTask(input, uint(id), userId)
	if err == repository.ErrTaskNotFound {
		h.logs.Error("Failed to update Task: ", err)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task Updated succesfully"})
	h.logs.Infof("Task with Id %v updated succesfully", id)
}

func (h *handler) DeleteTask(c *gin.Context) {
	h.logs.Info("Deleting Task")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logs.Error("Failed to get Id: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userId, err := getUserId(c)
	if err != nil {
		h.logs.Error("Failed to get User Id: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = h.services.DeleteTask(uint(id), userId)
	if err != nil {
		h.logs.Error("Failed to delete Task: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	c.JSON(http.StatusOK, gin.H{"reason": "Task deleted succesfully"})
	h.logs.Infof("Task with Id %v deleted succesfully", id)
}
