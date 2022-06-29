package main

import (
	"sixstar/shopstar-micro/client/bootstrap"
	coreRouter "sixstar/shopstar-micro/client/core/router"
	"sixstar/shopstar-micro/client/global"

	"sixstar/shopstar-micro/client/app/router"
)

var Config = "conf.yml"

func main()  {
	// 启动系统启动
	bootstrap.Init(Config)

	router.InitRouter()
	// 启动系统服务
	coreRouter.InitRoutes().Run(global.Config.App.Address)
}
