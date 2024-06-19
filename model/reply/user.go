package reply

import (
	"GuidingSystem/model/common"
	"github.com/jinzhu/gorm"
)

type User struct {
	Name     string `json:"name,omitempty" binding:"required" gorm:"UNIQUE"`
	Password string `json:"password,omitempty" binding:"required,gte=6,lte=50"`
	gorm.Model
}

type Register struct {
	User  User         `json:"user"`
	Token common.Token `json:"token"`
}

type Login struct {
	User  User         `json:"user"`
	Token common.Token `json:"token"`
}
