package services

import (
	context "context"
)

//GetProdStock service
type ProdService struct {
}

// GetProdStock service implments
func (this *ProdService) GetProdStock(context.Context, *ProdRequest) (*ProdResponse, error) {

	return &ProdResponse{ProdStock: 222}, nil
}
