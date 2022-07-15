package bootstrap

import (
	"toboefa/core/cache"
	"toboefa/core/log"
	"toboefa/core/model"
	"toboefa/core/rpc/client"
	"toboefa/global"
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
	global.RpcClient = client.NewClient(global.Config.RpcClient)
	// 初识mysql
	global.DB = model.InitDb(global.Config.Mysql)
}
