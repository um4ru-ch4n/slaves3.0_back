package routes

import (
	"net/http"
	"strings"
	"time"

	"github.com/00mrx00/slaves3.0_back/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "user"
	tokenType           = "AccessToken"
	tokenCtx            = "userToken"
)

func (r *Router) hasAuth(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Empty 'Authorization' header")
		return
	}

	headerParts := strings.Split(header, " ")

	if len(headerParts) != 2 || headerParts[0] != tokenType {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Invalid 'Authorization' header")
		return
	}

	if len(headerParts[1]) == 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Access token is empty")
		return
	}

	c.Set(tokenCtx, headerParts[1])
}

func (r *Router) userIdentity(c *gin.Context) {
	token := c.MustGet(tokenCtx).(string)

	userVkInfo, err := r.services.User.GetUserVkInfo(token)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, err)
		return
	}

	c.Set(userCtx, userVkInfo)
}

func (r *Router) updateStats(c *gin.Context) {
	userVkInfo, _ := c.MustGet(userCtx).(domain.UserVkInfo)

	lastUpdate, err := r.services.User.GetLastUpdate(userVkInfo.Id)
	if err != nil {
		if err == pgx.ErrNoRows {
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		}
		return
	}

	if time.Since(lastUpdate).Minutes() < 1 {
		c.Next()
		return
	}

	if err := r.services.User.UpdateUserInfo(userVkInfo.Id); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
}
