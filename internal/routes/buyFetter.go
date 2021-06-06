package routes

import (
	"net/http"

	"github.com/00mrx00/slaves3.0_back/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type buyFetterInfo struct {
	SlaveId    int32  `json:"slave_id" binding:"required"`
	FetterType string `json:"fetter_type" binding:"required"`
}

func (r *Router) buyFetter(c *gin.Context) {
	userVkInfo, _ := c.MustGet("user").(domain.UserVkInfo)

	var bFInfo buyFetterInfo

	if err := c.ShouldBindJSON(&bFInfo); err != nil {
		r.logger.Error("buyFetter c.ShouldBindJSON Router: ", zap.Error(err))
		c.JSON(http.StatusBadRequest, errors.Cause(err).Error())
		return
	}

	if err := r.services.User.BuyFetter(userVkInfo.Id, bFInfo.SlaveId, bFInfo.FetterType); err != nil {
		r.logger.Error("buyFetter r.services.User.BuySlave Router: ", zap.Error(err))
		c.JSON(http.StatusConflict, errors.Cause(err).Error())
		return
	}

	c.Status(http.StatusOK)
}
