package handlers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *handler) AuthMiddleware(c *gin.Context) {
	header := c.GetHeader("Authorization")

	if header == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"reason": "empty authorization header"})
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"reason": "invalid authorization header"})
		return
	}

	id, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"reason": err.Error()})
		return
	}
	c.Set("userId", id)
}

func getUserId(c *gin.Context) (uint, error) {
	userId, ok := c.Get("userId")
	if !ok {
		return 0, errors.New("user id not found")
	}
	id, ok := userId.(uint)
	if !ok {
		return 0, errors.New("error type conversion")
	}
	return id, nil
}
