package main

import (
	"toboefa/app/service"
	"toboefa/bootstrap"
	"toboefa/core/rpc"
	"toboefa/global"
)

var Config = "conf.yml"

func main()  {
	//panic("p")
	// 启动系统启动
	bootstrap.Init(Config)

	// 启动系统服务
	// fmt.Println(global.Config.App.Host + ":" + global.Config.App.Port)
	//coreRouter.InitRoutes().Run(global.Config.App.Address)
	rpc.NewServer(service.NewApiServiceServer).Run(global.Config.App.Address)
}

