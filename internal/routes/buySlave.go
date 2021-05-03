package routes

import (
	"fmt"
	"net/http"

	"github.com/00mrx00/slaves3.0_back/internal/domain"
	"github.com/gin-gonic/gin"
)

type slaveId struct {
	SlaveId int32 `json:"slave_id" binding:"required"`
}

func (r *Router) buySlave(c *gin.Context) {
	userVkInfo, _ := c.MustGet("user").(domain.UserVkInfo)
	// temp string
	fmt.Println(userVkInfo)
	var slaveId slaveId

	if err := c.ShouldBindJSON(&slaveId); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	// if err := r.services.User.BuySlave(userVkInfo.Id, slaveId.SlaveId); err != nil {
	// 	c.JSON(http.StatusInternalServerError, err)
	// 	return
	// }

	c.Status(http.StatusOK)
}
