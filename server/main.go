package main

import (
	"net"

	"github.com/changyuan/grpc-demo/services"
	"google.golang.org/grpc"
)

func main() {
	rpcServer := grpc.NewServer()
	services.RegisterProdServiceServer(rpcServer, new(services.ProdService))

	listen, _ := net.Listen("tcp", ":8081")

	rpcServer.Serve(listen)

}
