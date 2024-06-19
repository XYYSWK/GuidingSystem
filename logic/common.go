package logic

import (
	"GuidingSystem/global"
	"GuidingSystem/model"
	"github.com/XYYSWK/Rutils/pkg/token"
)

// newToken token
// 成功：返回 token，*token.Payload
// 失败：返回 nil, error
func newToken(id int64) (string, *token.Payload, error) {
	duration := global.PrivateSetting.Token.UserTokenDuration
	data, err := model.NewTokenContent(id).Marshal()
	if err != nil {
		return "", nil, err
	}
	result, payload, err := global.TokenMaker.CreateToken(data, duration)
	if err != nil {
		return "", nil, err
	}
	return result, payload, nil
}

// 将 id 从小到大排序返回
func sortID(id1, id2 int64) (_, _ int64) {
	if id1 > id2 {
		return id2, id1
	}
	return id1, id2
}
