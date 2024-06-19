package model

import (
	"encoding/json"
	"github.com/XYYSWK/Rutils/pkg/token"
)

type Token struct {
	AccessToken string
	Payload     *token.Payload
	Content     *Content
}

type Content struct {
	ID int64 `json:"id"`
}

// NewTokenContent 新建一种类型的 token
func NewTokenContent(id int64) *Content {
	return &Content{
		ID: id,
	}
}

// Marshal 将 Content 结构体序列化为 json 序列
func (c *Content) Marshal() ([]byte, error) {
	return json.Marshal(c)
}

// Unmarshal 将 json 序列解析为 Content 结构体
func (c *Content) Unmarshal(data []byte) error {
	return json.Unmarshal(data, c)
}
