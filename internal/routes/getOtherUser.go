package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

type oUserId struct {
	UserId int32 `json:"user_id" binding:"required"`
}

func (r *Router) getOtherUser(c *gin.Context) {
	var oUserId oUserId

	if err := c.ShouldBindJSON(&oUserId); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	userInfo, err := r.services.User.GetUserFull(oUserId.UserId)

	if err == pgx.ErrNoRows {
		userInfo, err = r.services.User.CreateUser(oUserId.UserId, "default")

		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, userInfo)
}
