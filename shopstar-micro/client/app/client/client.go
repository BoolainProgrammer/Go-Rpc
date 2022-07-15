package client

import (
	"go.uber.org/zap"
	"sixstar/shopstar-micro/client/global"
)

type rpcClient struct {
	MemberClient
}

var RpcClient = new(rpcClient)

func call(serverName, serverMethod string, req interface{}, resp interface{}) error {
	conn, err := global.RpcClient.NewConnect(serverName)

	if err != nil {
		global.Logs.Error("创建服务："+serverName+" error ", zap.Error(err))
		return err
	}

	if err := conn.Call(serverMethod, req, resp); err != nil {
		global.Logs.Error("调度服务："+serverName+" error ", zap.Error(err))
		return err
	}

	return nil
}