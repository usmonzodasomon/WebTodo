package handlers

import (
	"net/http"
	"webtodo/models"

	"github.com/gin-gonic/gin"
)

func (h *handler) SignUp(c *gin.Context) {
	h.logs.Info("Signing up")
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		h.logs.Error("Failed to bind JSON: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input body"})
		return
	}

	id, err := h.services.Authorization.AddUser(&user)
	if err != nil {
		h.logs.Error("Failed to authenticate: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error to add user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
	h.logs.Infof("User with Id %v added succesfully", id)
}

func (h *handler) SignIn(c *gin.Context) {
	h.logs.Info("Signing in")

	var user models.User
	if err := c.BindJSON(&user); err != nil {
		h.logs.Error("Failed to bind JSON: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect user data"})
		return
	}

	token, err := h.services.Authorization.GenerateToken(user.Username, user.Password)
	if err != nil {
		h.logs.Error("Failed to generate token: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
