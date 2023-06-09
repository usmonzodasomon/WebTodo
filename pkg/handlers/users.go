package handlers

// import (
// 	"database/sql"
// 	"errors"
// 	"net/http"
// 	"strconv"
// 	"webtodo/repository"

// 	"github.com/gin-gonic/gin"
// )

// func (h *handler) GetUserById(c *gin.Context) {
// 	h.l.Println("Get User By Id")

// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
// 		return
// 	}

// 	user, err := repository.GetUserById(id)

// 	if err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"reason": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, user)
// }

// func (h *handler) GetUsers(c *gin.Context) {
// 	h.l.Println("Get Users")
// 	userId, ok := c.Get("userId")
// 	if !ok {
// 		c.JSON(http.StatusInternalServerError, gin.H{"reason": "user id not found in auth header"})
// 		return
// 	}
// 	users, err := service.GetAllUsers(userId.(uint))
// 	if err != nil {
// 		if errors.Is(err, sql.ErrNoRows) {
// 			c.JSON(http.StatusOK, gin.H{"reason": err.Error()})
// 		} else {
// 			c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
// 		}
// 	}
// 	c.JSON(http.StatusOK, users)
// }

// func (h *handler) GetUserTasks(c *gin.Context) {
// 	h.l.Println("Get User Task")
// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
// 		return
// 	}
// 	tasks, err := repository.GetUserTasks(id)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, tasks)
// }

// // func (h *handler) AddUser(c *gin.Context) {
// // 	h.l.Println("Add User")
// // 	var user models.User
// // 	if err := c.BindJSON(&user); err != nil {
// // 		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
// // 		return
// // 	}
// // 	user.Password = service.generatePasswordHash(user.Password)
// // 	id, err := repository.AddUser(&user)
// // 	if err != nil {
// // 		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
// // 	}
// // 	c.JSON(http.StatusCreated, gin.H{"id": id})
// // }

// func (h *handler) UpdateUser(c *gin.Context) {
// 	h.l.Println("Update User")
// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
// 		return
// 	}

// 	user, err := repository.GetUserById(id)
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"reason": err.Error()})
// 		return
// 	}
// 	if err := c.BindJSON(&user); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
// 		return
// 	}

// 	err = repository.UpdateUser(id, user)
// 	if err == repository.ErrUserNotFound {
// 		c.JSON(http.StatusNotFound, gin.H{"reason": err.Error()})
// 		return
// 	}
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"reason": "Задача успешно обновлена!"})
// }

// func (h *handler) DeleteUser(c *gin.Context) {
// 	h.l.Println("Delete User")
// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
// 		return
// 	}
// 	err = repository.DeleteUser(id)
// 	if err == repository.ErrUserNotFound {
// 		c.JSON(http.StatusNotFound, gin.H{"reason": err.Error()})
// 		return
// 	}
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"reason": err.Error()})
// 		return

// 	}
// 	c.JSON(http.StatusOK, gin.H{"reason": "User успешно удален"})
// }
