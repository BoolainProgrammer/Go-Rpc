package rpc

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Server struct {
	
}

func NewServer(server interface{}) *Server {
	rpc.Register(server)
	return &Server{}
}

func (Ser *Server) Run(addr ...string) (err error)  {
	address := resolveAddress(addr)
	debugPrint("Listen and serving TCP on %s\n",address)
	list,err := net.Listen("tcp",address)
	if err!=nil {
		return err
	}
	for  {
		conn,err := list.Accept()
		if err!= nil {
			continue
		}
		go func(conn net.Conn) {
			jsonrpc.ServeConn(conn)
		}(conn)
	}
	return
}