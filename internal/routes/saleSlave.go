package routes

import (
	"net/http"

	"github.com/00mrx00/slaves3.0_back/internal/domain"
	"github.com/gin-gonic/gin"
)

func (r *Router) saleSlave(c *gin.Context) {
	userVkInfo, _ := c.MustGet("user").(domain.UserVkInfo)

	var slaveId domain.SlaveId

	if err := c.ShouldBindJSON(&slaveId); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := r.services.User.SaleSlave(userVkInfo.Id, slaveId.SlaveId); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}
