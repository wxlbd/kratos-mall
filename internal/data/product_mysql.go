package data

import (
	"context"
	"gorm.io/gorm"
	"kratos-admin/api"
	"kratos-admin/internal/biz"
	"kratos-admin/internal/data/po"
)

type ProductRepo struct {
	data *Data
}

func (p *ProductRepo) FindProductById(ctx context.Context, id int64) (*biz.Product, error) {
	var product po.PmsProduct
	if err := p.data.DB.First(&product, id).Error; err != nil {
		return nil, api.ErrorDbError("Failed to find product %d", id).WithCause(err)
	}
	return p.productPoToDo(&product), nil
}
func (p *ProductRepo) productPoToDo(po *po.PmsProduct) *biz.Product {
	return &biz.Product{
		Id:                         po.Id,
		ProductCategoryIds:         po.ProductCategoryIds,
		BrandId:                    po.BrandId,
		ProductAttributeCategoryId: po.ProductAttributeCategoryId,
		Name:                       po.Name,
		Pic:                        po.Pic,
		ProductSn:                  po.ProductSn,
		PublishStatus:              po.PublishStatus,
		NewStatus:                  po.NewStatus,
		RecommendStatus:            po.RecommendStatus,
		VerifyStatus:               po.VerifyStatus,
		Sort:                       po.Sort,
		TotalSales:                 po.TotalSales,
		Price:                      po.Price,
		PromotionPrice:             po.PromotionPrice,
		GiftGrowth:                 po.GiftGrowth,
		GiftPoint:                  po.GiftPoint,
		UsePointLimit:              po.UsePointLimit,
	}
}

func (p *ProductRepo) productSkuPoToDo(sku *po.ProductSku) *biz.ProductSku {
	attributes := make([]biz.Attribute, 0, len(sku.Attributes))
	for _, attr := range sku.Attributes {
		attributes = append(attributes, biz.Attribute{
			AttributeId:        attr.AttributeId,
			AttributeName:      attr.AttributeName,
			AttributeValueId:   attr.AttributeValueId,
			AttributeValueName: attr.AttributeValueName,
		})
	}
	return &biz.ProductSku{
		Id:             sku.Id,
		SkuCode:        sku.SkuCode,
		Price:          sku.Price,
		Stock:          sku.Stock,
		StockWarn:      sku.StockWarn,
		Pic:            sku.Pic,
		Sales:          sku.Sales,
		PromotionPrice: sku.PromotionPrice,
		GiftBlockStock: sku.GiftBlockStock,
		Name:           sku.Name,
		Attributes:     attributes,
	}
}

func (p *ProductRepo) FindProductSkuBySkuId(ctx context.Context, skuId int64) (*biz.ProductSku, error) {
	var productSku po.ProductSku
	if err := p.data.DB.First(&productSku, skuId).Error; err != nil {
		return nil, api.ErrorDbError("Failed to find product sku %d", skuId).WithCause(err)
	}
	return p.productSkuPoToDo(&productSku), nil
}

func (p *ProductRepo) FindProductSkusBySkuIdList(ctx context.Context, skuIds []int64) ([]*biz.ProductSku, error) {
	var productSkus []po.ProductSku
	if err := p.data.DB.Where("id in ?", skuIds).Find(&productSkus).Error; err != nil {
		return nil, api.ErrorDbError("Failed to find product skus %v", skuIds).WithCause(err)
	}
	skus := make([]*biz.ProductSku, 0, len(productSkus))
	for _, sku := range productSkus {
		skus = append(skus, p.productSkuPoToDo(&sku))
	}
	return skus, nil
}

func (p *ProductRepo) FindProductSkusByProductId(ctx context.Context, productId int64) ([]*biz.ProductSku, error) {
	var productSkus []po.ProductSku
	if err := p.data.DB.Where("product_id = ?", productId).Find(&productSkus).Error; err != nil {
		return nil, api.ErrorDbError("Failed to find product skus %d", productId).WithCause(err)
	}
	skus := make([]*biz.ProductSku, 0, len(productSkus))
	for _, sku := range productSkus {
		skus = append(skus, p.productSkuPoToDo(&sku))
	}
	return skus, nil
}

func (p *ProductRepo) productDoToPo(param *biz.Product) *po.PmsProduct {
	product := &po.PmsProduct{
		BrandId:                    param.BrandId,
		ProductCategoryIds:         param.ProductCategoryIds,
		FreightTemplateId:          param.FreightTemplateId,
		ProductAttributeCategoryId: param.ProductAttributeCategoryId,
		Name:                       param.Name,
		Pic:                        param.Pic,
		ProductSn:                  param.ProductSn,
		Sort:                       param.Sort,
		TotalSales:                 param.TotalSales,
		Price:                      param.Price,
		PromotionPrice:             param.PromotionPrice,
		GiftGrowth:                 param.GiftGrowth,
		GiftPoint:                  param.GiftPoint,
		UsePointLimit:              param.UsePointLimit,
		SubTitle:                   param.SubTitle,
		Description:                param.Description,
		OriginalPrice:              param.OriginalPrice,
		TotalStock:                 param.TotalStock,
		Unit:                       param.Unit,
		Weight:                     param.Weight,
		PreviewStatus:              param.PreviewStatus,
		PublishStatus:              param.PublishStatus,
		NewStatus:                  param.NewStatus,
		RecommendStatus:            param.RecommendStatus,
		VerifyStatus:               param.VerifyStatus,
		ServiceIds:                 param.ServiceIds,
		Keywords:                   param.Keywords,
		Note:                       param.Note,
		AlbumPics:                  param.AlbumPics,
		DetailTitle:                param.DetailTitle,
		DetailDesc:                 param.DetailDesc,
		DetailHtml:                 param.DetailHtml,
		DetailMobileHtml:           param.DetailMobileHtml,
		PromotionStartTime:         param.PromotionStartTime,
		PromotionEndTime:           param.PromotionEndTime,
		PromotionPerLimit:          param.PromotionPerLimit,
		PromotionType:              param.PromotionType,
		BrandName:                  param.BrandName,
		ProductCategoryName:        param.ProductCategoryName,
	}
	product.Skus = make([]*po.ProductSku, 0, len(param.SkuList))
	for _, sku := range param.SkuList {
		attributes := make([]po.Attribute, 0, len(sku.Attributes))
		for _, attr := range sku.Attributes {
			attributes = append(attributes, po.Attribute{
				AttributeId:        attr.AttributeId,
				AttributeName:      attr.AttributeName,
				AttributeValueId:   attr.AttributeValueId,
				AttributeValueName: attr.AttributeValueName,
			})
		}
		product.Skus = append(product.Skus, &po.ProductSku{
			ProductId:      sku.ProductId,
			SkuCode:        sku.SkuCode,
			Name:           sku.Name,
			Attributes:     attributes,
			Price:          sku.Price,
			PromotionPrice: sku.PromotionPrice,
			Pic:            sku.Pic,
			Stock:          sku.Stock,
			StockWarn:      sku.StockWarn,
			Sales:          sku.Sales,
			GiftBlockStock: sku.GiftBlockStock,
		})
	}
	return product
}

func (p *ProductRepo) productSkuDoToPo(sku *biz.ProductSku) *po.ProductSku {
	return &po.ProductSku{
		Id:             sku.Id,
		ProductId:      sku.ProductId,
		SkuCode:        sku.SkuCode,
		Name:           sku.Name,
		Price:          sku.Price,
		PromotionPrice: sku.PromotionPrice,
		Pic:            sku.Pic,
		Stock:          sku.Stock,
		StockWarn:      sku.StockWarn,
		Sales:          sku.Sales,
		GiftBlockStock: sku.GiftBlockStock,
	}
}

func (p *ProductRepo) UpdateProduct(ctx context.Context, param *biz.UpdateProductDo) error {
	product := p.productDoToPo(param.Product)
	if err := p.data.DB.WithContext(ctx).Session(&gorm.Session{FullSaveAssociations: true}).Model(&po.PmsProduct{}).Where("id = ?", param.Id).Updates(&product).Error; err != nil {
		return api.ErrorDbError("Failed to update product %d", param.Id).WithCause(err)
	}
	return nil
}

func (p *ProductRepo) UpdateProductSku(ctx context.Context, param *biz.UpdateProductSkuDo) error {
	productSku := p.productSkuDoToPo(param.ProductSku)
	if err := p.data.DB.WithContext(ctx).Model(&po.ProductSku{}).Where("id = ?", param.Id).Updates(&productSku).Error; err != nil {
		return api.ErrorDbError("Failed to update product sku %d", param.Id).WithCause(err)
	}
	return nil
}

func (p *ProductRepo) CreateProduct(ctx context.Context, param *biz.CreateProductDo) (int64, error) {
	product := p.productDoToPo(param.Product)
	if err := p.data.DB.WithContext(ctx).Create(&product).Error; err != nil {
		return 0, api.ErrorDbError("Failed to create product").WithCause(err)
	}
	return product.Id, nil
}

func (p *ProductRepo) DeleteProduct(ctx context.Context, productId int64) error {
	err := p.data.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("product_id = ?", productId).Delete(&po.ProductSku{}).Error; err != nil {
			return api.ErrorDbError("Failed to delete product skus %d", productId).WithCause(err)
		}
		if err := tx.Delete(&po.PmsProduct{}, productId).Error; err != nil {
			return api.ErrorDbError("Failed to delete product %d", productId).WithCause(err)
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (p *ProductRepo) DeleteProductSku(ctx context.Context, skuId int64) error {
	if err := p.data.DB.WithContext(ctx).Delete(&po.ProductSku{}, skuId).Error; err != nil {
		return api.ErrorDbError("Failed to delete product sku %d", skuId).WithCause(err)
	}
	return nil
}

func (p *ProductRepo) FindProductList(ctx context.Context, req *biz.ListProductParam) (total int64, list []*biz.Product, err error) {
	if err = p.data.DB.WithContext(ctx).Model(&po.PmsProduct{}).Where("").Count(&total).Error; err != nil {
		return 0, nil, api.ErrorDbError("Failed to find product list").WithCause(err)
	}
	if err = p.data.DB.WithContext(ctx).Model(&po.PmsProduct{}).Scopes(
		func(db *gorm.DB) *gorm.DB {
			db = db.Offset((req.Number - 1) * req.Size)
			return db
		},
	).Scan(&list).Error; err != nil {
		return 0, nil, api.ErrorDbError("Failed to find product list").WithCause(err)
	}
	return total, list, nil
}
