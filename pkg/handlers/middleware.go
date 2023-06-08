package handlers

import (
	"net/http"
	"strings"
	"webtodo/service"

	"github.com/gin-gonic/gin"
)

func (t *app) AuthMiddleware(c *gin.Context) {
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

	id, err := service.ParseToken(headerParts[1])
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"reason": err.Error()})
		return
	}
	c.Set("userId", id)
}
