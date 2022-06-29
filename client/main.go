package main

import (
	"toboefa/bootstrap"
	coreRouter "toboefa/core/router"
	"toboefa/global"

	"toboefa/app/router"
)

var Config = "conf.yml"

func main()  {
	//panic("p")
	// 启动系统启动
	bootstrap.Init(Config)

	router.InitRouter()
	// 启动系统服务
	// fmt.Println(global.Config.App.Host + ":" + global.Config.App.Port)
	coreRouter.InitRoutes().Run(global.Config.App.Address)
}
