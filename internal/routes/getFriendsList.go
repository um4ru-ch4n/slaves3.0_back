package routes

import (
	"net/http"

	"github.com/00mrx00/slaves3.0_back/internal/domain"
	"github.com/gin-gonic/gin"
)

func (r *Router) getFriendsList(c *gin.Context) {
	// userVkInfo, _ := c.MustGet("user").(domain.UserVkInfo)

	// token := strings.Split(c.GetHeader(authorizationHeader), " ")[1]

	// friends, err := r.services.User.GetFriendsList(token, userVkInfo.Id)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, err)
	// 	return
	// }

	c.JSON(http.StatusOK, []domain.UserVkInfo{})
}
