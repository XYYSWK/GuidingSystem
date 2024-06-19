package routers

import (
	"GuidingSystem/controller"
	"github.com/gin-gonic/gin"
)

type user struct {
}

func (user) Init(router *gin.RouterGroup) {
	r := router.Group("/user")
	{
		r.POST("register", controller.Apis.User.Register)
		r.POST("login", controller.Apis.User.Login)
	}
}
