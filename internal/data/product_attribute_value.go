package data

import (
	"context"
	"kratos-admin/api"
	v1 "kratos-admin/api/product/v1"
	"kratos-admin/internal/data/po"
	"strconv"
)

type ProductAttributeValueRepo struct {
	data *Data
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
		return api.ErrorDbError("Failed to create or update product attribute value").WithCause(err)
	}
	return nil
}

func (p *ProductAttributeValueRepo) DeleteProductAttributeValue(ctx context.Context, id int64) error {
	err := p.data.DB.WithContext(ctx).Delete(&po.PmsProductAttributeValue{}, id).Error
	if err != nil {
		return api.ErrorDbError("Failed to delete product attribute value").WithCause(err)
	}
	return nil
}

func (p *ProductAttributeValueRepo) FindProductAttributeValueByAttributeId(ctx context.Context, attributeId int64) ([]*v1.ProductAttributeValue, error) {
	var list []*v1.ProductAttributeValue
	err := p.data.DB.WithContext(ctx).Where("product_attribute_id = ?", attributeId).Find(&list).Error
	if err != nil {
		return nil, api.ErrorDbError("Failed to find product attribute value by attribute id").WithCause(err)
	}
	return list, nil
}
