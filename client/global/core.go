package global

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"toboefa/config"
	"toboefa/core/rpc/client"
)

const ConfigFile = "./conf.yml"

const JwtKey = "shopstar"

var (
	Config config.Config
	Logs   *zap.Logger
	Viper *viper.Viper // 后面可能会对配置文件操作，可以通过它来实现
	DB     *gorm.DB
	RpcClient *client.RpcClient
)
var SaveDbOmitTime = DB.Omit("created_time").Omit("created_user")
