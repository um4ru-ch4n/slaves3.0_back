package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func (r *Router) getOtherUser(c *gin.Context) {
	userId, _ := c.MustGet("userId").(int32)

	userToken := c.MustGet("userToken").(string)

	userInfo, err := r.services.User.GetUserFull(userId)

	if errors.Cause(err) == pgx.ErrNoRows {
		userVkInfo, err := r.services.User.GetUsersVkInfo(userToken, []int32{userId})
		if err != nil {
			r.logger.Error("getOtherUser r.services.User.GetUsersVkInfo Router: ", zap.Error(err))
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		userInfo, err = r.services.User.CreateUser(userId, "simp", userVkInfo[0].Fio, userVkInfo[0].Photo)

		if err != nil {
			r.logger.Error("getOtherUser r.services.User.CreateUser Router: ", zap.Error(err))
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
	} else if err != nil {
		r.logger.Error("getOtherUser r.services.User.GetUserFull Router: ", zap.Error(err))
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, userInfo)
}
