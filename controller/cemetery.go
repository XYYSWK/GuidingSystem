package controller

import (
	"GuidingSystem/logic"
	"GuidingSystem/model/request"
	"github.com/XYYSWK/Rutils/pkg/app"
	"github.com/XYYSWK/Rutils/pkg/app/errcode"
	"github.com/gin-gonic/gin"
)

type cemetery struct {
}

// GetCemeteries 获取所有的墓陵
// @Tags application
// @Summary 获取所有的墓陵
// @accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 账户令牌"
// @Success 200 {object} common.State{data=[]*reply.Cemetery} "1003:系统错误 2008:身份验证失败 2010:账号不存在"
// @Router /api/cemetery/list [get]
func (cemetery) GetCemeteries(ctx *gin.Context) {
	reply := app.NewResponse(ctx)
	result, err := logic.Logics.Cemetery.GetCemeteries(ctx)
	reply.Reply(err, result)
}

// GetCemeteriesByName 根据名称对墓陵进行模糊查找
// @Tags application
// @Summary 根据名称对墓陵进行模糊查找
// @accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 账户令牌"
// @Param data body request.GetCemeteriesByName true "墓陵名称"
// @Success 200 {object} common.State{data=[]*reply.Cemetery} "1003:系统错误 2008:身份验证失败 2010:账号不存在"
// @Router /api/cemetery/name [get]
func (cemetery) GetCemeteriesByName(ctx *gin.Context) {
	reply := app.NewResponse(ctx)
	params := new(request.GetCemeteriesByName)
	result, err := logic.Logics.Cemetery.GetCemeteriesByName(ctx, params.Name)
	reply.Reply(err, result)
}

// GetCemeteryByID 根据墓陵 id 获取墓陵详细信息
// @Tags application
// @Summary 根据墓陵 id 获取墓陵详细信息
// @accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 账户令牌"
// @Param data body request.GetCemeteryByID true "墓陵名称"
// @Success 200 {object} common.State{data=[]*reply.GetCemeteryByID} "1003:系统错误 2008:身份验证失败 2010:账号不存在"
// @Router /api/cemetery/id [get]
func (cemetery) GetCemeteryByID(ctx *gin.Context) {
	reply := app.NewResponse(ctx)
	params := new(request.GetCemeteryByID)
	if err := ctx.ShouldBind(params); err != nil {
		reply.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	result, err := logic.Logics.Cemetery.GetCemeteryByID(ctx, params.ID)
	reply.Reply(err, result)
}

// CreateCemetery 创建新的墓陵
// @Tags application
// @Summary 创建新的墓陵
// @accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 账户令牌"
// @Param data body request.Cemetery true "墓陵名称"
// @Success 200 {object} common.State{data=reply.Cemetery} "1003:系统错误 2008:身份验证失败 2010:账号不存在"
// @Router /api/cemetery/create [post]
func (cemetery) CreateCemetery(ctx *gin.Context) {
	reply := app.NewResponse(ctx)
	params := new(request.Cemetery)
	if err := ctx.ShouldBind(params); err != nil {
		reply.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	result, err := logic.Logics.Cemetery.CreateCemetery(ctx, params)
	reply.Reply(err, result)
}
