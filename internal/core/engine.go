package core

import (
	"github.com/Catlordx/CampusTrade/internal/router"
	"github.com/gin-contrib/cors"
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

	//BUG 跨域策略最后需要删掉！
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type"}
	engine.Use(cors.New(corsConfig))
	engine.Use(gin.Recovery(), gin.Logger())
	engine.Use(func(context *gin.Context) {
		context.Set("appContext", appContext)
	})
	router.InitRoutes(engine)
	return &Server{Engine: engine}
}
