package global

import (
	"GuidingSystem/model/config"
	"github.com/XYYSWK/Rutils/pkg/logger"
	"github.com/XYYSWK/Rutils/pkg/token"
)

var (
	Logger         *logger.Log          // 日志
	PublicSetting  config.PublicConfig  // Public 配置
	PrivateSetting config.PrivateConfig // Private 配置
	TokenMaker     token.MakerToken     // token
)
