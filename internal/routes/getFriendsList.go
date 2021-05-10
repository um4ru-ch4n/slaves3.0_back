package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Router) getFriendsList(c *gin.Context) {
	token := c.MustGet("userToken").(string)

	friends, err := r.services.User.GetFriendsList(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, friends)
}
