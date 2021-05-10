package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *Router) getRatingSlavesCount(c *gin.Context) {
	ratingUsers, err := r.services.User.GetRatingBySlavesCount()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, ratingUsers)
}
