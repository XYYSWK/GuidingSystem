package logic

import (
	"GuidingSystem/dao"
	"GuidingSystem/global"
	"GuidingSystem/middlewares"
	"GuidingSystem/model/reply"
	"github.com/XYYSWK/Rutils/pkg/app/errcode"
	"github.com/gin-gonic/gin"
)

type attraction struct{}

func (attraction) GetAttractionsByCemeteryID(ctx *gin.Context, cemeteryID int64) ([]*reply.Attraction, errcode.Err) {
	result, err := dao.DB.GetAttractionsByCemeteryID(cemeteryID)
	if err != nil {
		global.Logger.Error(err.Error(), middlewares.ErrLogMsg(ctx)...)
		return nil, errcode.ErrServer
	}
	return result, nil
}
