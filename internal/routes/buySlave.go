package routes

import (
	"net/http"

	"github.com/00mrx00/slaves3.0_back/internal/domain"
	"github.com/gin-gonic/gin"
)

type slaveId struct {
	SlaveId int32 `json:"slave_id" binding:"required"`
}

func (r *Router) buySlave(c *gin.Context) {
	userVkInfo, _ := c.MustGet("user").(domain.UserVkInfo)

	var slaveId slaveId

	if err := c.ShouldBindJSON(&slaveId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := r.services.Authorization.BuySlave(userVkInfo.Id, slaveId.SlaveId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
