package main

import (
	"context"
	"fmt"
	"log"

	"github.com/changyuan/grpc-demo/services"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	prodClient := services.NewProductServiceClient(conn)
	prodRes, err := prodClient.GetProdStock(context.Background(), &services.ProductRequest{ProdId: 2})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(prodRes.ProdStock)
}
