package data

import (
	"context"
	"gorm.io/gorm"
	"kratos-admin/api"
	v1 "kratos-admin/api/product/v1"
	"kratos-admin/internal/biz"
	"kratos-admin/internal/data/po"
	"strconv"
)

type ProductAttributeRepo struct {
	data *Data
}

func (p *ProductAttributeRepo) FindProductAttributeById(ctx context.Context, id int64) (*v1.ProductAttribute, error) {
	var productAttribute *po.PmsProductAttribute
	if err := p.data.DB.WithContext(ctx).First(&productAttribute, id).Error; err != nil {
		return nil, api.ErrorDbError("Failed to find product attribute").WithCause(err)
	}

	return p.productAttributePoToDto(productAttribute), nil
}

func (p *ProductAttributeRepo) FindProductAttributeList(ctx context.Context, param *v1.FindProductAttributeListRequest) (reply *v1.FindProductAttributeListReply, err error) {
	reply = new(v1.FindProductAttributeListReply)

	var (
		count int64
		list  []*po.PmsProductAttribute
	)
	tx := p.data.DB.WithContext(ctx).
		Model(&po.PmsProductAttribute{}).
		Scopes(func(db *gorm.DB) *gorm.DB {
			if param.Name != "" {
				db = db.Where("name like ?", "%"+param.Name+"%")
			}
			if param.Type != v1.ProductAttributeType_ProductAttributeTypeUnknown {
				db = db.Where("select_type = ?", param.Type)
			}
			return db
		})
	err = tx.Count(&count).Error
	if err != nil {
		return nil, api.ErrorDbError("Failed to find product attribute list").WithCause(err)
	}
	err = tx.Offset(int((param.GetPageNumber() - 1) * param.PageSize)).Limit(int(param.PageSize)).Find(&list).Error
	if err != nil {
		return nil, api.ErrorDbError("Failed to find product attribute list").WithCause(err)
	}

	reply.ProductAttributes = make([]*v1.ProductAttribute, 0, len(list))
	for _, v := range list {
		reply.ProductAttributes = append(reply.ProductAttributes, p.productAttributePoToDto(v))
	}
	reply.Total = int32(count)
	return
}

func (p *ProductAttributeRepo) productAttributeDtoToPo(attribute *v1.ProductAttribute) *po.PmsProductAttribute {
	categoryId, _ := strconv.ParseInt(attribute.ProductAttributeCategoryId, 10, 64)
	id, _ := strconv.ParseInt(attribute.Id, 10, 64)
	return &po.PmsProductAttribute{
		Id:                         id,
		ProductAttributeCategoryId: categoryId,
		Name:                       attribute.Name,
		SelectType:                 int32(attribute.SelectType),
		InputType:                  int32(attribute.InputType),
		InputList:                  attribute.InputList,
		Sort:                       attribute.Sort,
		FilterType:                 int32(attribute.FilterType),
		SearchType:                 int32(attribute.SearchType),
		RelatedStatus:              int32(attribute.RelatedStatus),
		HandAddStatus:              int32(attribute.HandAddStatus),
		Type:                       int32(attribute.Type),
	}
}
func (p *ProductAttributeRepo) CreateProductAttribute(ctx context.Context, param *v1.CreateProductAttributeRequest) (int64, error) {
	productAttribute := p.productAttributeDtoToPo(param.ProductAttribute)
	if err := p.data.DB.WithContext(ctx).Create(productAttribute).Error; err != nil {
		return 0, api.ErrorDbError("Failed to create product attribute").WithCause(err)
	}
	return productAttribute.Id, nil
}

func (p *ProductAttributeRepo) UpdateProductAttribute(ctx context.Context, param *v1.UpdateProductAttributeRequest) error {
	id, err := strconv.ParseInt(param.ProductAttributeId, 10, 64)
	if err != nil {
		return api.ErrorInvalidParam("Invalid product attribute id").WithCause(err)
	}
	if err := p.data.DB.WithContext(ctx).Model(&po.PmsProductAttribute{}).Where("id = ?", id).Updates(p.productAttributeDtoToPo(param.ProductAttribute)).Error; err != nil {
		return api.ErrorDbError("Failed to update product attribute").WithCause(err)
	}
	return nil
}

func (p *ProductAttributeRepo) DeleteProductAttribute(ctx context.Context, id int64) error {
	if err := p.data.DB.WithContext(ctx).Delete(&po.PmsProductAttribute{}, id).Error; err != nil {
		return api.ErrorDbError("Failed to delete product attribute").WithCause(err)
	}
	return nil
}

func (p *ProductAttributeRepo) productAttributePoToDto(attribute *po.PmsProductAttribute) *v1.ProductAttribute {
	return &v1.ProductAttribute{
		Id:                         strconv.FormatInt(attribute.Id, 10),
		ProductAttributeCategoryId: strconv.FormatInt(attribute.ProductAttributeCategoryId, 10),
		Name:                       attribute.Name,
		InputType:                  v1.ProductAttributeInputType(attribute.InputType),
		InputList:                  attribute.InputList,
		Sort:                       attribute.Sort,
		SelectType:                 v1.ProductAttributeSelectType(attribute.SelectType),
		FilterType:                 v1.ProductAttributeFilterType(attribute.FilterType),
		SearchType:                 v1.ProductAttributeSearchType(attribute.SearchType),
		Type:                       v1.ProductAttributeType(attribute.Type),
		HandAddStatus:              v1.ProductAttributeHandAddStatus(attribute.HandAddStatus),
	}
}

func NewProductAttributeRepo(data *Data) biz.ProductAttributeRepo {
	return &ProductAttributeRepo{data: data}
}
