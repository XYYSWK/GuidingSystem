package dao

import (
	"GuidingSystem/global"
	"GuidingSystem/model/config"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //MySQL 驱动库的初始化（不同类型的 DBType 需要引入不同的驱动库）
)

type db struct {
	DBEngine *gorm.DB
}

var DB = new(db)

func Init(config config.MysqlConfig) *db {
	s := "%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local"
	db1, err := gorm.Open(config.DBType, fmt.Sprintf(s,
		config.Username,
		config.Password,
		config.Host,
		config.DBName,
		config.Charset,
		config.ParseTime,
	))
	if err != nil {
		return nil
	}
	if global.PublicSetting.Server.RunMode == "debug" {
		db1.LogMode(true)
	}

	// 修改 gorm 的命名策略
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return config.TablePrefix + defaultTableName
	}

	db1.DB().SetMaxIdleConns(config.MaxIdle)
	db1.DB().SetMaxOpenConns(config.MaxOpen)

	return &db{DBEngine: db1}
}
