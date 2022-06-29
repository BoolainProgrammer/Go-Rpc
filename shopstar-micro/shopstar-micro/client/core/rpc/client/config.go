package client

type Config struct{
	// name => server => []serverCfg // 作为作业去设计
	// 记录服务的地址配置信息,根据服务名定义
	Servers map[string]serverCfg `server`
}

type serverCfg struct {
	Network string
	Address string
}
