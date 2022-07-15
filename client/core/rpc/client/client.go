package client

import (
	"net/rpc"
	"net/rpc/jsonrpc"
)

type RpcClient struct {
	cfg *Config
}

func NewClient(cfg *Config) *RpcClient {
	return &RpcClient{cfg: cfg}
}

type Conn interface {
	Call(serverName string, req interface{}, resp interface{}) error
	Close() error
}

//连接服务
type Connect struct {
	conn *rpc.Client
}

//根据服务名称,创建新的服务
func (c *RpcClient) NewConnect(serverName string) (*Connect, error) {
	if len(c.cfg.Servers) > 0 {
		server, ok := c.cfg.Servers[serverName]
		if !ok {
			return nil, ErrNotServerCfg
		}
		conn, err := jsonrpc.Dial(server.Network, server.Address)
		if err != nil {
			return nil, err
		}
		return &Connect{conn: conn}, err
	}
	return &Connect{}, nil
}

//调用
func (c *Connect) Call(serverMethod string, req interface{}, resp interface{}) error {
	return c.conn.Call(serverMethod, req, resp)
}

//关闭
func (c *Connect) Close() error {
	return c.Close()
}
