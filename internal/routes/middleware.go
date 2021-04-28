package routes

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "user"
	tokenType           = "AccessToken"
)

func (r *Router) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Empty 'Authorization' header",
		})
		return
	}

	headerParts := strings.Split(header, " ")

	if len(headerParts) != 2 || headerParts[0] != tokenType {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid 'Authorization' header",
		})
		return
	}

	if len(headerParts[1]) == 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Access token is empty",
		})
		return
	}

	userVkInfo, err := r.services.Authorization.GetUserVkInfo(headerParts[1])

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.Set(userCtx, userVkInfo)
}
