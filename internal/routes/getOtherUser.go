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
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	userToken := c.MustGet("userToken").(string)

	userInfo, err := r.services.User.GetUserFull(oUserId.UserId)

	if err == pgx.ErrNoRows {
		userVkInfo, err := r.services.User.GetUsersVkInfo(userToken, []int32{oUserId.UserId})
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		userInfo, err = r.services.User.CreateUser(oUserId.UserId, "simp", userVkInfo[0].Fio, userVkInfo[0].Photo)

		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, userInfo)
}
