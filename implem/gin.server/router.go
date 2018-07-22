package server

import (
	"github.com/gin-gonic/gin"
	
	"github.com/err0r500/go-realworld-clean/uc"
)

type Router struct {
	handler uc.Handler
}

func NewRouter(handler uc.Handler) Router {
	return Router{
		handler: handler,
	}
}

func (rH Router) SetRoutes(router *gin.Engine) {
	router.GET("/health", rH.health)
}