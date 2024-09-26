package router

import (
	"net/http"
	"simple_go/config"
	"simple_go/http/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter(handler *handler.Handler, config *config.Config) *gin.Engine {
	router := gin.New()
	// cors middleware
	c := cors.New(cors.Config{
		AllowOrigins:     config.AllowedOrigins,
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodDelete},
		AllowHeaders:     []string{"Origin"},
		AllowCredentials: true,
	})
	router.Use(c)

	// TODO: add routes to router
	routes(router, handler)
	return router
}
