package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func (r *Router) getRatingSlavesCount(c *gin.Context) {
	ratingUsers, err := r.services.User.GetRatingBySlavesCount()
	if err != nil {
		r.logger.Error("getRatingSlavesCount r.services.User.GetRatingBySlavesCount Router: ", zap.Error(err))
		c.JSON(http.StatusInternalServerError, errors.Cause(err).Error())
		return
	}

	c.JSON(http.StatusOK, ratingUsers)
}
