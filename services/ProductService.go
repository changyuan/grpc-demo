package services

import (
	context "context"
)

//GetProdStock service
type ProductService struct {
}

// GetProdStock service implments
func (this *ProductService) GetProdStock(context.Context, *ProductRequest) (*ProductResponse, error) {

	return &ProductResponse{ProdStock: 100}, nil
}
