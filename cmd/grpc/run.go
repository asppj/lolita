package grpc

import (
	"log"
	"net"

	"github.com/asppj/t-mk-opentrace/config"
	g "github.com/asppj/t-mk-opentrace/ext/grpc-driver/grpc"
)

// // Host Host
// var Host = ""
//
// // Port Port
// var Port = "16005"
//
// // RPCAddr addr
// var RPCAddr = Host + ":" + Port

// app defaultRPCServer
var server *g.Server

// RPCServer rpc
func RPCServer() error {
	server = g.NewServer()
	RegisterRPC(server)
	conf := config.Get().RPC
	host := conf.Host
	Port := conf.Port
	RPCAddr := conf.Host + ":" + Port
	lis, err := net.Listen("tcp", RPCAddr)
	if err != nil {
		log.Fatalf("net.Listen rpc err: %v", err)
		return err
	}
	if host == "" {
		host = "0.0.0.0"
	}
	log.Printf("rpc start 监听:%s:%s", host, Port)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("rpc Server start err: %v", err)
		return err
	}
	return nil
}

// RPCShotDown 关闭
func RPCShotDown() {
	if server != nil {
		server.Stop()
		log.Println("rpc stop")
	}
}
