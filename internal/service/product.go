package service

import (
	"context"
	"kratos-admin/internal/biz"

	pb "kratos-admin/api/product/v1"
)

type ProductService struct {
	pb.UnimplementedProductServer
	productUseCase *biz.ProductUseCase
}

func NewProductService(productUseCase *biz.ProductUseCase) *ProductService {
	return &ProductService{
		productUseCase: productUseCase,
	}
}

func (s *ProductService) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductReply, error) {
	ctx.Value("adminId")
	return &pb.CreateProductReply{}, nil
}
func (s *ProductService) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.UpdateProductReply, error) {
	return &pb.UpdateProductReply{}, nil
}
func (s *ProductService) DeleteProduct(ctx context.Context, req *pb.DeleteProductRequest) (*pb.DeleteProductReply, error) {
	return &pb.DeleteProductReply{}, nil
}
func (s *ProductService) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.GetProductReply, error) {
	return &pb.GetProductReply{}, nil
}
func (s *ProductService) ListProduct(ctx context.Context, req *pb.ListProductRequest) (*pb.ListProductReply, error) {
	return &pb.ListProductReply{}, nil
}
func (s *ProductService) ListProductSku(ctx context.Context, req *pb.ListProductSkuRequest) (*pb.ListProductSkuReply, error) {
	return &pb.ListProductSkuReply{}, nil
}
