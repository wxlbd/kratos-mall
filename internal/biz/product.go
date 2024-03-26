package biz

import (
	"context"

	v1 "github.com/wxlbd/kratos-pms/api"

	pb "github.com/wxlbd/kratos-pms/api/product/v1"
)

type ProductUseCase struct {
	productRepo ProductRepo
}

func NewProductUseCase(productRepo ProductRepo) *ProductUseCase {
	return &ProductUseCase{
		productRepo: productRepo,
	}
}

func (puc *ProductUseCase) CreateProduct(ctx context.Context, req *CreateProductDo) (int64, error) {
	id, err := puc.productRepo.CreateProduct(ctx, req)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (puc *ProductUseCase) UpdateProduct(ctx context.Context, req *UpdateProductDo) error {
	if err := puc.productRepo.UpdateProduct(ctx, req); err != nil {
		return err
	}
	return nil
}

func (puc *ProductUseCase) DeleteProduct(ctx context.Context, productId int64) error {
	if err := puc.productRepo.DeleteProduct(ctx, productId); err != nil {
		return err
	}
	return nil
}

func (puc *ProductUseCase) GetProduct(ctx context.Context, productId int64) (*Product, error) {
	product, err := puc.productRepo.FindProductById(ctx, productId)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (puc *ProductUseCase) ListProduct(ctx context.Context, req *pb.ListProductRequest) (total int64, list []*Product, err error) {
	total, list, err = puc.productRepo.FindProductList(ctx, req)
	if err != nil {
		return 0, nil, err
	}
	return
}

func (puc *ProductUseCase) ListProductSkus(ctx context.Context, req *pb.ListProductSkuRequest) (list []*ProductSku, err error) {
	skus, err := puc.productRepo.FindProductSkusByProductId(ctx, req.ProductId)
	if err != nil {
		return nil, err
	}
	return skus, nil
}

func (puc *ProductUseCase) DeleteProductSku(ctx context.Context, skuId int64) error {
	if err := puc.productRepo.DeleteProductSku(ctx, skuId); err != nil {
		return err
	}
	return nil
}

func (puc *ProductUseCase) UpdateProductSku(ctx context.Context, sku *ProductSku) error {
	if err := puc.productRepo.UpdateProductSku(ctx, &UpdateProductSkuDo{
		ProductSku: sku,
	}); err != nil {
		return v1.ErrorDbError("Failed to update product sku %d", sku.Id).WithCause(err)
	}
	return nil
}

func (puc *ProductUseCase) GetProductSku(ctx context.Context, skuId int64) (*ProductSku, error) {
	sku, err := puc.productRepo.FindProductSkuById(ctx, skuId)
	if err != nil {
		return nil, err
	}
	return sku, nil
}

func (puc *ProductUseCase) CreateProductSku(ctx context.Context, req *ProductSku) (int64, error) {
	id, err := puc.productRepo.CreateProductSku(ctx, req)
	if err != nil {
		return 0, err
	}
	return id, nil
}
