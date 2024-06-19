package controller

import (
	"GuidingSystem/logic"
	"GuidingSystem/model/request"
	"github.com/XYYSWK/Rutils/pkg/app"
	"github.com/XYYSWK/Rutils/pkg/app/errcode"
	"github.com/gin-gonic/gin"
)

type attraction struct {
}

// GetAttractionsByCemeteryID 根据墓陵ID获取该墓陵所有景点
// @Tags application
// @Summary 根据墓陵ID获取该墓陵所有景点
// @accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 账户令牌"
// @Param data body request.GetAttractions true "墓陵名称"
// @Success 200 {object} common.State{data=[]*reply.Attraction} "1003:系统错误 2008:身份验证失败 2010:账号不存在"
// @Router /api/attraction/list [get]
func (attraction) GetAttractionsByCemeteryID(ctx *gin.Context) {
	reply := app.NewResponse(ctx)
	params := new(request.GetAttractions)
	if err := ctx.ShouldBind(params); err != nil {
		reply.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	result, err := logic.Logics.Attraction.GetAttractionsByCemeteryID(ctx, params.CemeteryID)
	reply.Reply(err, result)
}
