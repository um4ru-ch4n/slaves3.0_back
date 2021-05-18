package routes

import (
	"net/http"

	"github.com/00mrx00/slaves3.0_back/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func (r *Router) getSlavesList(c *gin.Context) {
	userVkInfo, _ := c.MustGet("user").(domain.UserVkInfo)

	slavesList, err := r.services.User.GetSlavesList(userVkInfo.Id)
	if err != nil {
		r.logger.Error("getSlavesList r.services.User.GetSlavesList Router: ", zap.Error(err))
		c.JSON(http.StatusInternalServerError, errors.Cause(err).Error())
		return
	}

	c.JSON(http.StatusOK, slavesList)
}
