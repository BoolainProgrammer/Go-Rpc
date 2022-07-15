package client

import (
	"net/rpc"
	"net/rpc/jsonrpc"
)

type RpcClient struct {
	cfg *Config
}

func NewClient(cfg *Config) *RpcClient {
	return &RpcClient{
		cfg: cfg,
	}
}

type Conn interface {
	Call(serverMethod string, req interface{}, resp interface{}) error
	Close() error
}

type Connect struct {
	conn *rpc.Client
	//conn Conn
}

// 根据服务名,创建新的连接
func (c *RpcClient) NewConnect(serverName string) (*Connect, error) {
	if len(c.cfg.Servers) > 0 {
		server, ok := c.cfg.Servers[serverName]

		if !ok {
			return nil, ErrNotServerCfg
		}

		// 建立连接
		conn, err := jsonrpc.Dial(server.Network, server.Address)

		if err != nil {
			return nil, err
		}

		return &Connect{
			conn: conn,
		}, nil
	}
	return &Connect{}, nil
}

// 调用
func (c *Connect) Call(serverMethod string, req interface{}, resp interface{}) error {
	return c.conn.Call(serverMethod, req, resp)
}

// 关闭连接
func (c *Connect) Close() error {
	return c.conn.Close()
}
