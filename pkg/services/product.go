package services

import (
	"context"
	"github.com/vandong9/go-grpc-product-svc/pkg/db"
	"github.com/vandong9/go-grpc-product-svc/pkg/models"
	"github.com/vandong9/go-grpc-product-svc/pkg/pb"
	"net/http"
)

type Server struct {
	H db.Handler
}

//func (s *Server) mustEmbedUnimplementedProductServiceServer() {
//	//TODO implement me
//	panic("implement me")
//}

func (s *Server) CreateProduct(ctx context.Context, request pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	var product = models.Product{}

	product.Name = request.Name
	product.Stock = request.Stock
	product.Price = request.Price

	if result := s.H.DB.Create(&product); result.Error != nil {
		return &pb.CreateProductResponse{Status: http.StatusConflict, Error: result.Error.Error()}, nil
	}

	return &pb.CreateProductResponse{Status: http.StatusCreated, Id: product.Id}, nil
}

func (s *Server) FindOne(ctx context.Context, request pb.FindOneRequest) (*pb.FindOneResponse, error) {
	return &pb.FindOneResponse{Status: http.StatusCreated}, nil
}

func (s *Server) DecreaseStock(ctx context.Context, request pb.DecreaseStockRequest) (*pb.DecreaseStockResponse, error) {
	return &pb.DecreaseStockResponse{Status: http.StatusCreated}, nil
}
