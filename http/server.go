package http

import (
	"simple_go/config"
	"simple_go/database/db"
	"simple_go/http/handler"
	"simple_go/http/router"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func NewServer(store db.Store, config *config.Config) *Server {
	handler := handler.NewHandler(store)
	server := &Server{
		router: router.NewRouter(handler, config),
	}

	return server
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}
