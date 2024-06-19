package router

import (
	"GuidingSystem/global"
	"GuidingSystem/middlewares"
	"GuidingSystem/routers"
	"github.com/XYYSWK/Rutils/pkg/app"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(middlewares.Cors(), middlewares.GinLogger(), middlewares.Recovery(true))

	root := r.Group("/api", middlewares.LogBody(), middlewares.PasetoAuth())
	{
		root.GET("swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
		root.GET("ping", func(ctx *gin.Context) {
			reply := app.NewResponse(ctx)
			global.Logger.Info("ping", middlewares.ErrLogMsg(ctx)...)
			reply.Reply(nil, "pong")
		})

		rg := routers.Routers
		rg.User.Init(root)
		rg.Cemetery.Init(root)
		rg.Attraction.Init(root)
	}
	return r
}
