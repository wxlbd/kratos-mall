package biz

import (
	"context"
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

func (puc *ProductUseCase) ListProduct(ctx context.Context, req *ListProductParam) (total int64, list []*Product, err error) {
	total, list, err = puc.productRepo.FindProductList(ctx, req)
	if err != nil {
		return 0, nil, err
	}
	return
}

func (puc *ProductUseCase) ListProductSkus(ctx context.Context, req *ListProductSkuParam) (list []*ProductSku, err error) {
	skus, err := puc.productRepo.FindProductSkusByProductId(ctx, req.ProductId)
	if err != nil {
		return nil, err
	}
	return skus, nil
}
