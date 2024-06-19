package routers

import (
	"GuidingSystem/controller"
	"github.com/gin-gonic/gin"
)

type attraction struct {
}

func (attraction) Init(router *gin.RouterGroup) {
	r := router.Group("/attraction")
	{
		r.GET("list", controller.Apis.Attraction.GetAttractionsByCemeteryID)
	}
}
