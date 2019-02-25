package rpcsupport

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func RpcServer(port int, service interface{}) error {
	rpc.Register(service)
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
        return err
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Accept error %v", err)
		}
		go jsonrpc.ServeConn(conn)
	}
	return nil
}

func NewClient(port int) (*rpc.Client, error) {
	conn, err := net.Dial("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
        return nil, err
	}
	return jsonrpc.NewClient(conn), nil
}