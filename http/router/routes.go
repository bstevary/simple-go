package router

import (
	"simple_go/http/handler"

	"github.com/gin-gonic/gin"
)

func routes(router *gin.Engine, handler *handler.Handler) {
	v1 := router.Group("/v1/")
	v1.POST("user", handler.CreateUser)
	v1.GET("health", handler.Health)
}
