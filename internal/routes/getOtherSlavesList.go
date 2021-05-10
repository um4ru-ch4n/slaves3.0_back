package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserId struct {
	UserId int32 `json:"user_id" binding:"required"`
}

func (r *Router) getOtherSlavesList(c *gin.Context) {
	var userId UserId

	if err := c.ShouldBindJSON(&userId); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	slavesList, err := r.services.User.GetSlavesList(userId.UserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, slavesList)
}
