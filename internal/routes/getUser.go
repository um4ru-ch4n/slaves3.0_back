package routes

import (
	"net/http"

	"github.com/00mrx00/slaves3.0_back/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func (r *Router) getUser(c *gin.Context) {
	userVkInfo, _ := c.MustGet("user").(domain.UserVkInfo)

	userInfo, err := r.services.User.GetUserFull(userVkInfo.Id)

	if errors.Cause(err) == pgx.ErrNoRows {
		userInfo, err = r.services.User.CreateUser(userVkInfo.Id, "simp", userVkInfo.Fio, userVkInfo.Photo)
		if err != nil {
			r.logger.Error("getUser r.services.User.CreateUser Router: ", zap.Error(err))
			c.JSON(http.StatusConflict, errors.Cause(err).Error())
			return
		}

	} else if err != nil {
		r.logger.Error("getUser r.services.User.GetUserFull Router: ", zap.Error(err))
		c.JSON(http.StatusConflict, errors.Cause(err).Error())
		return
	}

	c.JSON(http.StatusOK, userInfo)
}
