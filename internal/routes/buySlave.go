package routes

import (
	"net/http"

	"github.com/00mrx00/slaves3.0_back/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func (r *Router) buySlave(c *gin.Context) {
	userVkInfo, _ := c.MustGet("user").(domain.UserVkInfo)

	var slaveId domain.SlaveId

	if err := c.ShouldBindJSON(&slaveId); err != nil {
		r.logger.Error("buySlave c.ShouldBindJSON Router: ", zap.Error(err))
		c.JSON(http.StatusBadRequest, errors.Cause(err).Error())
		return
	}

	if err := r.services.User.BuySlave(userVkInfo.Id, slaveId.SlaveId); err != nil {
		r.logger.Error("buySlave r.services.User.BuySlave Router: ", zap.Error(err))
		c.JSON(http.StatusInternalServerError, errors.Cause(err).Error())
		return
	}

	c.Status(http.StatusOK)
}
