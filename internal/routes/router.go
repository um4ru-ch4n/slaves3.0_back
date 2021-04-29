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

	auth := router.Group("/auth", r.hasAuth, r.userIdentity)
	{
		auth.GET("/user", r.getUser)
		auth.GET("/friends", r.getFriendsList)
	}
	people := router.Group("/people", r.hasAuth)
	{
		people.GET("/user", r.getOtherUser)
	}

	return router
}
