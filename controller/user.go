package controller

import (
	"GuidingSystem/logic"
	"GuidingSystem/model/request"
	"github.com/XYYSWK/Rutils/pkg/app"
	"github.com/XYYSWK/Rutils/pkg/app/errcode"
	"github.com/gin-gonic/gin"
)

type user struct {
}

// Register 注册用户
// @Tags user
// @Summary 用户注册
// @accept application/json
// @Produce application/json
// @Param data body request.User true "用户注册信息"
// @Success 200 {object} common.State{data=reply.Register} "1001:参数有误 1003:系统错误 2003:用户已存在"
// @Router /api/user/register [post]
func (user) Register(ctx *gin.Context) {
	// 获取参数和参数校验
	reply := app.NewResponse(ctx)
	params := new(request.User)
	if err := ctx.ShouldBind(params); err != nil {
		reply.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	// 业务处理
	result, err := logic.Logics.User.Register(ctx, params.Name, params.Password)
	// 返回响应
	reply.Reply(err, result)
}

// Login 用户登录
// @Tags user
// @Summary 用户登录
// @accept application/json
// @Produce application/json
// @Param data body request.User true "用户登录信息"
// @Success 200 {object} common.State{data=reply.Login} "1001:参数错误 1003:系统错误 2001:用户不存在 2002:密码错误"
// @Router /api/user/login [post]
func (user) Login(ctx *gin.Context) {
	// 获取参数和参数校验
	reply := app.NewResponse(ctx)
	params := new(request.User)
	if err := ctx.ShouldBind(params); err != nil {
		reply.Reply(errcode.ErrParamsNotValid.WithDetails(err.Error()))
		return
	}
	result, err := logic.Logics.User.Login(ctx, params.Name, params.Password)
	reply.Reply(err, result)
}
