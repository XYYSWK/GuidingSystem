package logic

import (
	"GuidingSystem/dao"
	"GuidingSystem/errcodes"
	"GuidingSystem/global"
	"GuidingSystem/middlewares"
	"GuidingSystem/model/common"
	"GuidingSystem/model/reply"
	"errors"
	"fmt"
	"github.com/XYYSWK/Rutils/pkg/app/errcode"
	"github.com/XYYSWK/Rutils/pkg/password"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type user struct{}

func (user) Register(ctx *gin.Context, name, pwd string) (*reply.Register, errcode.Err) {
	u, err := dao.DB.GetUserByName(name)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		global.Logger.Error(err.Error(), middlewares.ErrLogMsg(ctx)...)
		return nil, errcode.ErrServer
	}
	if u.ID > 0 {
		fmt.Println(u.Name)
		return nil, errcodes.UserNameIsExist
	}
	hashPassword, err := password.HashPassword(pwd)
	if err != nil {
		global.Logger.Error(err.Error(), middlewares.ErrLogMsg(ctx)...)
		return nil, errcode.ErrServer
	}
	userInfo := reply.User{
		Name:     name,
		Password: hashPassword,
	}
	resUser := reply.User{}
	if resUser, err = dao.DB.CreateUser(&userInfo); err != nil {
		global.Logger.Error(err.Error(), middlewares.ErrLogMsg(ctx)...)
		return nil, errcode.ErrServer
	}
	// 创建 Token
	userToken, payload, err := newToken(int64(resUser.ID))
	if err != nil {
		global.Logger.Error(err.Error(), middlewares.ErrLogMsg(ctx)...)
		return nil, errcode.ErrServer
	}
	return &reply.Register{
		User: resUser,
		Token: common.Token{
			Token:    userToken,
			ExpireAt: payload.ExpiredAt,
		},
	}, nil
}

func (user) Login(ctx *gin.Context, name, pwd string) (*reply.Login, errcode.Err) {
	u, err := dao.DB.GetUserByName(name)
	if err != nil {
		global.Logger.Error(err.Error(), middlewares.ErrLogMsg(ctx)...)
		return nil, errcode.ErrServer
	}
	if u.ID <= 0 {
		return nil, errcodes.UserNotFound
	}
	if err := password.CheckPassword(pwd, u.Password); err != nil {
		return nil, errcodes.PasswordNotValid
	}
	// 创建 token
	userToken, payload, err := newToken(int64(u.ID))
	if err != nil {
		global.Logger.Error(err.Error(), middlewares.ErrLogMsg(ctx)...)
		return nil, errcode.ErrServer
	}
	return &reply.Login{
		User: u,
		Token: common.Token{
			Token:    userToken,
			ExpireAt: payload.ExpiredAt,
		},
	}, nil
}
