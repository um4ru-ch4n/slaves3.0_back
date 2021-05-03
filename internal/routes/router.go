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
		user.GET("/", r.getUser)
		// user.GET("/friends", r.getFriendsList)
		// user.POST("/buyslave", r.buySlave)
	}
	fellow := router.Group("/fellow", r.hasAuth)
	{
		fellow.POST("/", r.getOtherUser)
	}

	return router
}
