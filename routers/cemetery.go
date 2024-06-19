package routers

import (
	"GuidingSystem/controller"
	"github.com/gin-gonic/gin"
)

type cemetery struct {
}

func (cemetery) Init(router *gin.RouterGroup) {
	r := router.Group("/cemetery")
	{
		r.GET("list", controller.Apis.Cemetery.GetCemeteries)
		r.GET("name", controller.Apis.Cemetery.GetCemeteriesByName)
		r.GET("id", controller.Apis.Cemetery.GetCemeteryByID)
		// 导入数据时使用
		r.POST("create", controller.Apis.Cemetery.CreateCemetery)
	}
}
