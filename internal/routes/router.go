package routes

import (
	"github.com/00mrx00/slaves3.0_back/internal/service"
	"github.com/gin-gonic/gin"
)

type Router struct {
	services *service.Service
}

func NewRouter(services *service.Service) *Router {
	return &Router{
		services: services,
	}
}

func (r *Router) InitRoutes() *gin.Engine {
	router := gin.Default()

	user := router.Group("/user", r.hasAuth, r.userIdentity)
	{
		user.GET("/", r.updateStatsHour, r.getUser)
		user.POST("/buyslave", r.updateStatsHour, r.buySlave)
		user.POST("/saleslave", r.updateStatsHour, r.saleSlave)
		user.GET("/slaves", r.getSlavesList)
		user.POST("/setjobname", r.setJobName)
	}
	fellow := router.Group("/fellow", r.hasAuth)
	{
		fellow.POST("/", r.updateStatsHourOther, r.getOtherUser)
		fellow.GET("/friends", r.getFriendsList)
		fellow.GET("/rating/slavescount", r.getRatingSlavesCount)
		fellow.POST("/slaves", r.getOtherSlavesList)
	}

	return router
}
