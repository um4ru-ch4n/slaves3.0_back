package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func (r *Router) getFriendsList(c *gin.Context) {
	token := c.MustGet("userToken").(string)

	friends, err := r.services.User.GetFriendsList(token)
	if err != nil {
		r.logger.Error("getFriendsList r.services.User.GetFriendsList Router: ", zap.Error(err))
		c.JSON(http.StatusInternalServerError, errors.Cause(err).Error())
		return
	}

	c.JSON(http.StatusOK, friends)
}
