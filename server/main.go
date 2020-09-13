package main

import (
	"net"

	"github.com/changyuan/grpc-demo/services"
	"google.golang.org/grpc"
)

func main() {
	rpcServer := grpc.NewServer()
	services.RegisterProductServiceServer(rpcServer, new(services.ProductService))

	listen, _ := net.Listen("tcp", ":8081")

	rpcServer.Serve(listen)

}
