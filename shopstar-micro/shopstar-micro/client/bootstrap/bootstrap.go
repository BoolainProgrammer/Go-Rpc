package bootstrap

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	"sixstar/shopstar-micro/client/core/rpc/client"
	"sixstar/shopstar-micro/client/core/cache"
	"sixstar/shopstar-micro/client/core/log"
	"sixstar/shopstar-micro/client/core/model"
	"sixstar/shopstar-micro/client/core/validate"
	"sixstar/shopstar-micro/client/global"
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
	// 初始化rpc client
	global.RpcClient = client.NewClient(global.Config.RpcClient)

}
