package biz

import (
	"context"

	v1 "github.com/wxlbd/kratos-pms/api/product/v1"
)

type ProductAttributeUseCase struct {
	productAttributeRepo ProductAttributeRepo
}

func NewProductAttributeUseCase(productAttributeRepo ProductAttributeRepo) *ProductAttributeUseCase {
	return &ProductAttributeUseCase{
		productAttributeRepo: productAttributeRepo,
	}
}

func (u *ProductAttributeUseCase) FindProductAttributeList(ctx context.Context, param *v1.FindProductAttributeListRequest) (reply *v1.FindProductAttributeListReply, err error) {
	list, err := u.productAttributeRepo.FindProductAttributeList(ctx, param)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (u *ProductAttributeUseCase) CreateProductAttribute(ctx context.Context, param *v1.CreateProductAttributeRequest) error {
	if _, err := u.productAttributeRepo.CreateProductAttribute(ctx, param); err != nil {
		return err
	}
	return nil
}

func (u *ProductAttributeUseCase) UpdateProductAttribute(ctx context.Context, param *v1.UpdateProductAttributeRequest) error {
	if err := u.productAttributeRepo.UpdateProductAttribute(ctx, param); err != nil {
		return err
	}
	return nil
}

func (u *ProductAttributeUseCase) DeleteProductAttribute(ctx context.Context, id int64) error {
	if err := u.productAttributeRepo.DeleteProductAttribute(ctx, id); err != nil {
		return err
	}
	return nil
}

func (u *ProductAttributeUseCase) FindProductAttributeById(ctx context.Context, id int64) (*v1.ProductAttribute, error) {
	attribute, err := u.productAttributeRepo.FindProductAttributeById(ctx, id)
	if err != nil {
		return nil, err
	}
	return attribute, nil
}
