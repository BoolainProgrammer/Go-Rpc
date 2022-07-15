package main

import (
	"sixstar/shopstar-micro/member-server/app/server"
	"sixstar/shopstar-micro/member-server/bootstrap"
	"sixstar/shopstar-micro/member-server/core/rpc"

	//coreRouter "sixstar/shopstar-micro/member-server/core/router"
	"sixstar/shopstar-micro/member-server/global"

)

var Config = "conf.yml"

func main()  {
	// 启动系统启动
	bootstrap.Init(Config)


	rpc.NewServer(server.NewMemberServer()).Run(global.Config.App.Address)
	// 启动系统服务
	// fmt.Println(global.Config.App.Host + ":" + global.Config.App.Port)
	//coreRouter.InitRoutes().Run(global.Config.App.Address)
}
