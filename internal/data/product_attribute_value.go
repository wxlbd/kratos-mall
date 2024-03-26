package data

import (
	"context"
	"strconv"

	"github.com/wxlbd/kratos-pms/internal/biz"

	"github.com/wxlbd/kratos-pms/api"
	v1 "github.com/wxlbd/kratos-pms/api/product/v1"
	"github.com/wxlbd/kratos-pms/internal/data/po"
)

type ProductAttributeValueRepo struct {
	data *Data
}

func NewProductAttributeValueRepo(data *Data) biz.ProductAttributeValueRepo {
	return &ProductAttributeValueRepo{data: data}
}

func (p *ProductAttributeValueRepo) CreateOrUpdateProductAttributeValue(ctx context.Context, param *v1.CreateOrUpdateProductAttributeValueRequest) error {
	id, _ := strconv.ParseInt(param.ProductAttributeValue.Id, 10, 64)
	productAttributeId, _ := strconv.ParseInt(param.ProductAttributeValue.ProductAttributeId, 10, 64)
	err := p.data.DB.WithContext(ctx).Save(&po.PmsProductAttributeValue{
		Id:                 id,
		ProductAttributeId: productAttributeId,
		Value:              param.ProductAttributeValue.Value,
	}).Error
	if err != nil {
		return api.ErrorDbError("Failed to create or update product Attribute value").WithCause(err)
	}
	return nil
}

func (p *ProductAttributeValueRepo) DeleteProductAttributeValue(ctx context.Context, id int64) error {
	err := p.data.DB.WithContext(ctx).Delete(&po.PmsProductAttributeValue{}, id).Error
	if err != nil {
		return api.ErrorDbError("Failed to delete product Attribute value").WithCause(err)
	}
	return nil
}

func (p *ProductAttributeValueRepo) FindProductAttributeValueByAttributeId(ctx context.Context, attributeId int64) ([]*v1.ProductAttributeValue, error) {
	var list []*v1.ProductAttributeValue
	err := p.data.DB.WithContext(ctx).Where("product_attribute_id = ?", attributeId).Find(&list).Error
	if err != nil {
		return nil, api.ErrorDbError("Failed to find product Attribute value by Attribute id").WithCause(err)
	}
	return list, nil
}
