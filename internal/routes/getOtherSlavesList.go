package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type UserId struct {
	UserId int32 `json:"user_id" binding:"required"`
}

func (r *Router) getOtherSlavesList(c *gin.Context) {
	var userId UserId

	if err := c.ShouldBindJSON(&userId); err != nil {
		r.logger.Error("getOtherSlavesList c.ShouldBindJSON Router: ", zap.Error(err))
		c.JSON(http.StatusBadRequest, errors.Cause(err).Error())
		return
	}

	slavesList, err := r.services.User.GetSlavesList(userId.UserId)
	if err != nil {
		r.logger.Error("getOtherSlavesList r.services.User.GetSlavesList Router: ", zap.Error(err))
		c.JSON(http.StatusConflict, errors.Cause(err).Error())
		return
	}

	c.JSON(http.StatusOK, slavesList)
}
