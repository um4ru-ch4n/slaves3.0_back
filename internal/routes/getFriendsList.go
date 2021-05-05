package routes

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (r *Router) getFriendsList(c *gin.Context) {
	token := strings.Split(c.GetHeader(authorizationHeader), " ")[1]

	friends, err := r.services.User.GetFriendsList(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, friends)
}
