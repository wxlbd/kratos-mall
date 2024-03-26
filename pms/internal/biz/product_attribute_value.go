package biz

import (
	"context"

	v1 "github.com/wxlbd/kratos-pms/api/product/v1"
)

type ProductAttributeValueUseCase struct {
	productAttributeValueRepo ProductAttributeValueRepo
}

func NewProductAttributeValueUseCase(productAttributeValueRepo ProductAttributeValueRepo) *ProductAttributeValueUseCase {
	return &ProductAttributeValueUseCase{
		productAttributeValueRepo: productAttributeValueRepo,
	}
}

func (u *ProductAttributeValueUseCase) CreateOrUpdateProductAttributeValue(ctx context.Context, param *v1.CreateOrUpdateProductAttributeValueRequest) error {
	return u.productAttributeValueRepo.CreateOrUpdateProductAttributeValue(ctx, param)
}

func (u *ProductAttributeValueUseCase) DeleteProductAttributeValue(ctx context.Context, id int64) error {
	return u.productAttributeValueRepo.DeleteProductAttributeValue(ctx, id)
}

func (u *ProductAttributeValueUseCase) FindProductAttributeValueByAttributeId(ctx context.Context, attributeId int64) (*v1.FindProductAttributeValueByAttributeIdReply, error) {
	values, err := u.productAttributeValueRepo.FindProductAttributeValueByAttributeId(ctx, attributeId)
	if err != nil {
		return nil, err
	}
	return &v1.FindProductAttributeValueByAttributeIdReply{
		ProductAttributeValues: values,
	}, nil
}
