package core

import (
	"github.com/Catlordx/CampusTrade/internal/router"
	"github.com/gin-gonic/gin"
)

type Server struct {
	*gin.Engine
}

func New() *Server {
	appContext, err := NewAppContext()
	if err != nil {
		panic("Failed to create app context: " + err.Error())
	}
	engine := gin.New()
	engine.Use(gin.Recovery(), gin.Logger())
	engine.Use(func(context *gin.Context) {
		context.Set("appContext", appContext)
	})
	router.InitRoutes(engine)
	return &Server{Engine: engine}
}
