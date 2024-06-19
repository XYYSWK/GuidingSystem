package settings

import (
	"GuidingSystem/dao"
	"GuidingSystem/global"
)

type database struct {
}

func (database) Init() {
	dao.DB = dao.Init(global.PrivateSetting.Mysql)
	// 数据迁移
	dao.DB.Migration()
}
