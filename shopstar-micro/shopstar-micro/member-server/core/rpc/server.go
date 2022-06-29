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

func (Ser *Server) Run(addr ...string) (err error) {
	address := resolveAddress(addr)
	debugPrint("Listening and serving TCP on %s\n", address)
	lis, err := net.Listen("tcp", address)

	if err != nil {
		return
	}

	for {
		conn, err := lis.Accept()

		if err != nil {
			continue
		}

		go func(conn net.Conn) {
			jsonrpc.ServeConn(conn)
		}(conn)

	}

	return
}

