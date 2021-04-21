package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Router) getUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
