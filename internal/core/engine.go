package core

import (
	"github.com/Catlordx/CampusTrade/internal/router"
	"github.com/gin-gonic/gin"
)

type Server struct {
	*gin.Engine
}

func NewServer() *Server {
	engine := gin.New()
	engine.Use(gin.Recovery())
	router.InitRoutes(engine)
	return &Server{Engine: engine}
}
