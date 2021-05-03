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

func (r *Router) hasAuth(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Empty 'Authorization' header")
		return
	}

	headerParts := strings.Split(header, " ")

	if len(headerParts) != 2 || headerParts[0] != tokenType {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Invalid 'Authorization' header")
		return
	}

	if len(headerParts[1]) == 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Access token is empty")
		return
	}
}

func (r *Router) userIdentity(c *gin.Context) {
	headerParts := strings.Split(c.GetHeader(authorizationHeader), " ")

	userVkInfo, err := r.services.User.GetUserVkInfo(headerParts[1])

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, err)
		return
	}

	c.Set(userCtx, userVkInfo)
}
