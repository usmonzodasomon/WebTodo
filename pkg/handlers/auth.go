package handlers

import (
	"net/http"
	"webtodo/models"

	"github.com/gin-gonic/gin"
)

func (h *handler) SignUp(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": "invalid input body"})
		return
	}

	id, err := h.services.Authorization.AddUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
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

	token, err := h.services.Authorization.GenerateToken(user.Username, user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
