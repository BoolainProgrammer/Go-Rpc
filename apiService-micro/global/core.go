package global

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"toboefa/config"
)

const ConfigFile = "./conf.yml"

const JwtKey = "shopstar"

var (
	Config config.Config
	Viper  *viper.Viper
	Logs   *zap.Logger
	DB     *gorm.DB
)
//var SaveDbOmitTime = DB.Omit("created_time").Omit("created_user").Save
