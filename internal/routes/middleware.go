package routes

import (
	"net/http"
	"strings"
	"time"

	"github.com/00mrx00/slaves3.0_back/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "user"
	tokenType           = "AccessToken"
	tokenCtx            = "userToken"
	userId              = "userId"
)

func (r *Router) hasAuth(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		r.logger.Error("hasAuth c.GetHeader middleware: Empty 'Authorization' header")
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Empty 'Authorization' header")
		return
	}

	headerParts := strings.Split(header, " ")

	if len(headerParts) != 2 || headerParts[0] != tokenType {
		r.logger.Error("hasAuth middleware: Invalid 'Authorization' header")
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Invalid 'Authorization' header")
		return
	}

	if len(headerParts[1]) == 0 {
		r.logger.Error("hasAuth middleware: Access token is empty")
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Access token is empty")
		return
	}

	c.Set(tokenCtx, headerParts[1])
}

func (r *Router) userIdentity(c *gin.Context) {
	token := c.MustGet(tokenCtx).(string)

	userVkInfo, err := r.services.User.GetUserVkInfo(token)

	if err != nil {
		r.logger.Error("userIdentity r.services.User.GetUserVkInfo middleware: ", zap.Error(err))
		c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, userVkInfo)
}

func (r *Router) updateStatsHour(c *gin.Context) {
	userVkInfo, _ := c.MustGet(userCtx).(domain.UserVkInfo)

	lastUpdate, err := r.services.User.GetLastUpdate(userVkInfo.Id)
	if err != nil {
		if errors.Cause(err) == pgx.ErrNoRows {
			c.Next()
		} else {
			r.logger.Error("updateStatsHour r.services.User.GetLastUpdate middleware: ", zap.Error(err))
			c.AbortWithStatusJSON(http.StatusConflict, err.Error())
		}
		return
	}

	if time.Since(lastUpdate).Minutes() < 1 {
		c.Next()
		return
	}

	if err := r.services.User.UpdateUserInfo(userVkInfo.Id); err != nil {
		r.logger.Error("updateStatsHour r.services.User.UpdateUserInfo middleware: ", zap.Error(err))
		c.AbortWithStatusJSON(http.StatusConflict, err.Error())
		return
	}
}

type oUserIdType struct {
	UserId int32 `json:"user_id" binding:"required"`
}

func (r *Router) updateStatsHourOther(c *gin.Context) {
	var oUserId oUserIdType

	if err := c.ShouldBindJSON(&oUserId); err != nil {
		r.logger.Error("updateStatsHourOther c.ShouldBindJSON middleware: ", zap.Error(err))
		c.JSON(http.StatusBadRequest, errors.Cause(err).Error())
		return
	}

	c.Set(userId, oUserId.UserId)

	lastUpdate, err := r.services.User.GetLastUpdate(oUserId.UserId)
	if err != nil {
		if errors.Cause(err) == pgx.ErrNoRows {
			c.Next()
		} else {
			r.logger.Error("updateStatsHourOther r.services.User.GetLastUpdate middleware: ", zap.Error(err))
			c.AbortWithStatusJSON(http.StatusConflict, errors.Cause(err).Error())
		}
		return
	}

	if time.Since(lastUpdate).Minutes() < 1 {
		c.Next()
		return
	}

	if err := r.services.User.UpdateUserInfo(oUserId.UserId); err != nil {
		r.logger.Error("updateStatsHourOther r.services.User.UpdateUserInfo middleware: ", zap.Error(err))
		c.AbortWithStatusJSON(http.StatusConflict, errors.Cause(err).Error())
		return
	}
}
