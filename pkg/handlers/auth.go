package handlers

import (
	"net/http"
	"webtodo/models"
	"webtodo/service"

	"github.com/gin-gonic/gin"
)

func (h *handler) SignUp(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": "invalid input body"})
		return
	}

	id, err := service.AddUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *handler) SignIn(c *gin.Context) {
	h.l.Println("Signing in")

	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": err.Error()})
		return
	}

	token, err := service.GenerateToken(user.Username, user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
