package global

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"toboefa/config"
)

const ConfigFile = "./conf.yml"

const JwtKey = "shopstar"

var (
	Config config.Config
	Logs   *zap.Logger
	DB     *gorm.DB
)
var SaveDbOmitTime = *gorm.DB.Omit("created_time").Omit("created_user")
