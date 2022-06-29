package bootstrap

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	"sixstar/shopstar-micro/member-server/core/model"
	"sixstar/shopstar-micro/member-server/core/cache"
	"sixstar/shopstar-micro/member-server/core/log"
	"sixstar/shopstar-micro/member-server/core/validate"
	"sixstar/shopstar-micro/member-server/global"
)

func Init(cfgPath string) {
	global.Viper = Viper(&global.Config, cfgPath)
	global.Logs = log.InitLogger(global.Config.Log)

	initCore()

	// 加载 -行不行
}
func initCore() {
	// 初始化缓存
	if global.Config.Cache.Default == "freecache" {
		cache.CacheManager = cache.NewCache(cache.NewFreeCache(global.Config.Cache))
	}
	// 初始化模型
	v, _ := binding.Validator.Engine().(*validator.Validate)
	validate.InitValidate(v, validators, "zh")

	// 初识mysql
	global.DB = model.InitDb(global.Config.Mysql)
}
