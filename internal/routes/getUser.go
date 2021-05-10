package routes

import (
	"net/http"

	"github.com/00mrx00/slaves3.0_back/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func (r *Router) getUser(c *gin.Context) {
	userVkInfo, _ := c.MustGet("user").(domain.UserVkInfo)

	userInfo, err := r.services.User.GetUserFull(userVkInfo.Id)

	if err == pgx.ErrNoRows {
		userInfo, err = r.services.User.CreateUser(userVkInfo.Id, "simp", userVkInfo.Fio, userVkInfo.Photo)
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
