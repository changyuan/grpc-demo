package main

import (
	"log"
	"net"

	"github.com/changyuan/grpc-demo/services"
	"google.golang.org/grpc"
)

func main() {
	// 创建 gRPC Server 对象
	rpcServer := grpc.NewServer()
	// 将 SearchService（其包含需要被调用的服务端接口）注册到 gRPC Server 的内部注册中心
	services.RegisterProductServiceServer(rpcServer, new(services.ProductServiceServer))

	listen, err := net.Listen("tcp", ":8081")

	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	rpcServer.Serve(listen)

}
