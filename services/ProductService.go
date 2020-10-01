package services

import (
	context "context"
	"math/rand"
)

//GetProdStock service
type ProductService struct {
}

// GetProdStock service implments
func (this *ProductService) GetProdStock(context.Context, *ProductRequest) (*ProductResponse, error) {

	num:= rand.Int31n(10000)
	return &ProductResponse{ProdStock: num}, nil
}
