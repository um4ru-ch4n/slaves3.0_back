package routes

import (
	"net/http"
	"time"

	"github.com/00mrx00/slaves3.0_back/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func (r *Router) redeem(c *gin.Context) {
	userVkInfo, _ := c.MustGet("user").(domain.UserVkInfo)

	masterId, err := r.services.User.GetMasterId(userVkInfo.Id)
	if err != nil {
		r.logger.Error("redeem r.services.User.GetMasterId: ", zap.Error(err))
		c.JSON(http.StatusConflict, errors.Cause(err).Error())
		return
	}

	lastUpdate, err := r.services.User.GetLastUpdate(masterId)
	if err != nil {
		r.logger.Error("redeem r.services.User.GetLastUpdate: ", zap.Error(err))
		c.AbortWithStatusJSON(http.StatusConflict, errors.Cause(err).Error())
		return
	}

	if time.Since(lastUpdate).Minutes() >= 1 {
		if err := r.services.User.UpdateUserInfo(masterId); err != nil {
			r.logger.Error("redeem r.services.User.UpdateUserInfo: ", zap.Error(err))
			c.JSON(http.StatusConflict, errors.Cause(err).Error())
			return
		}
	}

	if err := r.services.User.Redeem(userVkInfo.Id); err != nil {
		r.logger.Error("redeem r.services.User.Redeem: ", zap.Error(err))
		c.JSON(http.StatusConflict, errors.Cause(err).Error())
		return
	}

	c.Status(http.StatusOK)
}
