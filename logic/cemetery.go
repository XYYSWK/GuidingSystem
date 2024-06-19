package logic

import (
	"GuidingSystem/dao"
	"GuidingSystem/errcodes"
	"GuidingSystem/global"
	"GuidingSystem/middlewares"
	"GuidingSystem/model/reply"
	"GuidingSystem/model/request"
	"github.com/XYYSWK/Rutils/pkg/app/errcode"
	"github.com/gin-gonic/gin"
)

type cemetery struct {
}

func (cemetery) GetCemeteries(ctx *gin.Context) (*reply.Cemeteries, errcode.Err) {
	result, err := dao.DB.GetCemeteries()
	if err != nil {
		global.Logger.Error(err.Error(), middlewares.ErrLogMsg(ctx)...)
		return nil, errcode.ErrServer
	}
	res := make([]*reply.CemeteryInfo, 0, len(result))
	for i, _ := range result {
		res = append(res, &reply.CemeteryInfo{
			ID:    result[i].ID,
			Name:  result[i].Name,
			Image: result[i].Image,
		})
	}
	return &reply.Cemeteries{CemeteriesInfo: res}, nil
}

func (cemetery) GetCemeteriesByName(ctx *gin.Context, name string) (*reply.Cemeteries, errcode.Err) {
	result, err := dao.DB.GetCemeteriesByName(name)
	if err != nil {
		global.Logger.Error(err.Error(), middlewares.ErrLogMsg(ctx)...)
		return nil, errcode.ErrServer
	}
	res := make([]*reply.CemeteryInfo, 0, len(result))
	for i, _ := range result {
		res = append(res, &reply.CemeteryInfo{
			ID:    result[i].ID,
			Name:  result[i].Name,
			Image: result[i].Image,
		})
	}
	return &reply.Cemeteries{CemeteriesInfo: res}, nil
}

func (cemetery) GetCemeteryByID(ctx *gin.Context, id int64) (*reply.GetCemeteryByID, errcode.Err) {
	result, err := dao.DB.GetCemeteryByID(id)
	if err != nil {
		global.Logger.Error(err.Error(), middlewares.ErrLogMsg(ctx)...)
		return nil, errcode.ErrServer
	}
	if result.ID <= 0 {
		return nil, errcodes.CemeteryNotFound
	}
	attractions, err := dao.DB.GetAttractionsByCemeteryID(result.ID)
	if err != nil {
		global.Logger.Error(err.Error(), middlewares.ErrLogMsg(ctx)...)
		return nil, errcode.ErrServer
	}
	return &reply.GetCemeteryByID{
		CemeteryInfo: result,
		Attraction:   attractions,
	}, nil
}

func (cemetery) CreateCemetery(ctx *gin.Context, params *request.Cemetery) (*reply.Cemetery, errcode.Err) {
	result, err := dao.DB.CreateCemetery(params)
	if err != nil {
		global.Logger.Error(err.Error(), middlewares.ErrLogMsg(ctx)...)
		return nil, errcode.ErrServer
	}
	return &result, nil
}
