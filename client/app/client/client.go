package client

import (
	"go.uber.org/zap"
	"toboefa/global"
)

type rpcClientStruct struct {
	ApiServiceClient
}

var RpaClient = new(rpcClientStruct)
//服务创建于调度
func call(serverName, serverMethod string, req interface{}, resp interface{}) error {
	conn, err := global.RpcClient.NewConnect(serverName)
	if err != nil {
		global.Logs.Error("创建服务： "+serverName+" err:", zap.Error(err))
		return err
	}
	if err := conn.Call(serverMethod, req, resp); err != nil {
		global.Logs.Error("调度服务： "+serverName+" err:", zap.Error(err))
		return err
	}
	return nil
}
