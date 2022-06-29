package global

import (
	"github.com/mojocn/base64Captcha"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"sixstar/shopstar-micro/member-server/config"
)

const ConfigFile = "./conf.yml"

const JwtKey = "shopstar"

var (
	Config config.Config
	Viper *viper.Viper // 后面可能会对配置文件操作，可以通过它来实现
	Logs *zap.Logger
	DB *gorm.DB

	TIMESTR = "20060102150405"
	TimeDataFormat = "2016-01-02"
)

var CaptchaStore = base64Captcha.DefaultMemStore


