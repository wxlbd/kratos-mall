package service

import (
	"context"
	"strconv"
	"time"

	"github.com/wxlbd/kratos-pms/api"
	"github.com/wxlbd/kratos-pms/internal/biz"

	pb "github.com/wxlbd/kratos-pms/api/product/v1"
)

type ProductService struct {
	pb.UnimplementedProductServiceServer
	productUseCase          *biz.ProductUseCase
	productAttrUseCase      *biz.ProductAttributeUseCase
	productCategoryUseCase  *biz.ProductCategoryUseCase
	productAttrValueUseCase *biz.ProductAttributeValueUseCase
}

func NewProductService(productUseCase *biz.ProductUseCase, productAttrUseCase *biz.ProductAttributeUseCase, productCategoryUseCase *biz.ProductCategoryUseCase, productAttrValueUseCase *biz.ProductAttributeValueUseCase) *ProductService {
	return &ProductService{
		productUseCase:          productUseCase,
		productAttrUseCase:      productAttrUseCase,
		productCategoryUseCase:  productCategoryUseCase,
		productAttrValueUseCase: productAttrValueUseCase,
	}
}

func (s *ProductService) CreateOrUpdateProductAttributeValue(ctx context.Context, req *pb.CreateOrUpdateProductAttributeValueRequest) (*pb.CreateOrUpdateProductAttributeValueReply, error) {
	if err := s.productAttrValueUseCase.CreateOrUpdateProductAttributeValue(ctx, req); err != nil {
		return nil, err
	}
	return &pb.CreateOrUpdateProductAttributeValueReply{}, nil
}

func (s *ProductService) DeleteProductAttributeValue(ctx context.Context, req *pb.DeleteProductAttributeValueRequest) (*pb.DeleteProductAttributeValueReply, error) {
	id, err := strconv.ParseInt(req.GetProductAttributeValueId(), 10, 64)
	if err != nil {
		return nil, api.ErrorInvalidParam("invalid product Attribute value id")
	}
	if err := s.productAttrValueUseCase.DeleteProductAttributeValue(ctx, id); err != nil {
		return nil, err
	}
	return &pb.DeleteProductAttributeValueReply{}, nil
}

func (s *ProductService) FindProductAttributeValueByAttributeId(ctx context.Context, req *pb.FindProductAttributeValueByAttributeIdRequest) (*pb.FindProductAttributeValueByAttributeIdReply, error) {
	id, err := strconv.ParseInt(req.GetProductAttributeId(), 10, 64)
	if err != nil {
		return nil, api.ErrorInvalidParam("invalid product Attribute id")
	}
	value, err := s.productAttrValueUseCase.FindProductAttributeValueByAttributeId(ctx, id)
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (s *ProductService) CreateProductAttribute(ctx context.Context, req *pb.CreateProductAttributeRequest) (*pb.CreateProductAttributeReply, error) {
	if err := s.productAttrUseCase.CreateProductAttribute(ctx, req); err != nil {
		return nil, err
	}
	return &pb.CreateProductAttributeReply{}, nil
}

func (s *ProductService) UpdateProductAttribute(ctx context.Context, req *pb.UpdateProductAttributeRequest) (*pb.UpdateProductAttributeReply, error) {
	if err := s.productAttrUseCase.UpdateProductAttribute(ctx, req); err != nil {
		return nil, err
	}
	return &pb.UpdateProductAttributeReply{}, nil
}

func (s *ProductService) DeleteProductAttribute(ctx context.Context, req *pb.DeleteProductAttributeRequest) (*pb.DeleteProductAttributeReply, error) {
	id, err := strconv.ParseInt(req.GetProductAttributeId(), 10, 64)
	if err != nil {
		return nil, api.ErrorInvalidParam("invalid product Attribute id")
	}
	if err := s.productAttrUseCase.DeleteProductAttribute(ctx, id); err != nil {
		return nil, err
	}
	return &pb.DeleteProductAttributeReply{}, nil
}

func (s *ProductService) FindProductAttribute(ctx context.Context, req *pb.FindProductAttributeRequest) (*pb.FindProductAttributeReply, error) {
	id, err := strconv.ParseInt(req.GetProductAttributeId(), 10, 64)
	if err != nil {
		return nil, api.ErrorInvalidParam("invalid product Attribute id")
	}
	value, err := s.productAttrUseCase.FindProductAttributeById(ctx, id)
	if err != nil {
		return nil, err
	}
	return &pb.FindProductAttributeReply{
		ProductAttribute: value,
	}, nil
}

func (s *ProductService) FindProductAttributeList(ctx context.Context, req *pb.FindProductAttributeListRequest) (*pb.FindProductAttributeListReply, error) {
	list, err := s.productAttrUseCase.FindProductAttributeList(ctx, req)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (s *ProductService) CreateProductCategory(ctx context.Context, req *pb.CreateProductCategoryRequest) (*pb.CreateProductCategoryReply, error) {
	if err := s.productCategoryUseCase.CreateProductCategory(ctx, req); err != nil {
		return nil, err
	}
	return &pb.CreateProductCategoryReply{}, nil
}

func (s *ProductService) UpdateProductCategory(ctx context.Context, req *pb.UpdateProductCategoryRequest) (*pb.UpdateProductCategoryReply, error) {
	if err := s.productCategoryUseCase.UpdateProductCategory(ctx, req); err != nil {
		return nil, err
	}
	return &pb.UpdateProductCategoryReply{}, nil
}

func (s *ProductService) DeleteProductCategory(ctx context.Context, req *pb.DeleteProductCategoryRequest) (*pb.DeleteProductCategoryReply, error) {
	id, err := strconv.ParseInt(req.GetProductCategoryId(), 10, 64)
	if err != nil {
		return nil, api.ErrorInvalidParam("invalid product category id")
	}
	if err := s.productCategoryUseCase.DeleteProductCategory(ctx, id); err != nil {
		return nil, err
	}
	return &pb.DeleteProductCategoryReply{}, nil
}

func (s *ProductService) FindProductCategory(ctx context.Context, req *pb.FindProductCategoryRequest) (*pb.FindProductCategoryReply, error) {
	id, err := strconv.ParseInt(req.GetProductCategoryId(), 10, 64)
	if err != nil {
		return nil, api.ErrorInvalidParam("invalid product category id")
	}
	value, err := s.productCategoryUseCase.FindProductCategoryById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &pb.FindProductCategoryReply{
		ProductCategory: value,
	}, nil
}

func (s *ProductService) FindProductCategoryList(ctx context.Context, req *pb.FindProductCategoryListRequest) (*pb.FindProductCategoryListReply, error) {
	list, err := s.productCategoryUseCase.FindProductCategoryList(ctx, req)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (s *ProductService) FindProductCategoryTree(ctx context.Context, req *pb.FindProductCategoryTreeRequest) (*pb.FindProductCategoryTreeReply, error) {
	list, err := s.productCategoryUseCase.FindProductCategoryTree(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.FindProductCategoryTreeReply{
		Tree: list,
	}, nil
}

func (s *ProductService) productDTOToDO(product *pb.Product) (*biz.Product, error) {
	productCateIds := make([]int64, 0, len(product.CategoryIds))
	for _, v := range product.CategoryIds {
		id, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil, api.ErrorInvalidParam("invalid product category id")
		}
		productCateIds = append(productCateIds, id)
	}
	freightTemplateId, err := strconv.ParseInt(product.FreightTemplateId, 10, 64)
	if err != nil {
		return nil, api.ErrorInvalidParam("invalid freight template id")
	}
	productAttributeCategoryId, err := strconv.ParseInt(product.ProductAttributeCategoryId, 10, 64)
	if err != nil {
		return nil, api.ErrorInvalidParam("invalid product Attribute category id")
	}
	brandId, err := strconv.ParseInt(product.BrandId, 10, 64)
	if err != nil {
		return nil, api.ErrorInvalidParam("invalid brand id")
	}
	promotionStartTime, err := time.Parse(time.DateTime, product.PromotionStartTime)
	if err != nil {
		return nil, api.ErrorInvalidParam("invalid promotion start time")
	}
	promotionEndTime, err := time.Parse(time.DateTime, product.PromotionEndTime)
	if err != nil {
		return nil, api.ErrorInvalidParam("invalid promotion end time")
	}
	productSkus := make([]*biz.ProductSku, 0, len(product.SkuList))
	for _, v := range product.SkuList {
		productSku, err := s.skuDTOToDO(v)
		if err != nil {
			return nil, err
		}
		productSkus = append(productSkus, productSku)
	}
	id, err := strconv.ParseInt(product.Id, 10, 64)
	if err != nil {
		return nil, api.ErrorInvalidParam("invalid product id")
	}
	return &biz.Product{
		Id:                         id,
		ProductCategoryIds:         productCateIds,
		FreightTemplateId:          freightTemplateId,
		BrandId:                    brandId,
		ProductAttributeCategoryId: productAttributeCategoryId,
		Name:                       product.Name,
		Pic:                        product.Pic,
		ProductSn:                  product.ProductSn,
		Sort:                       product.Sort,
		TotalSales:                 product.TotalSales,
		Price:                      product.Price,
		PromotionPrice:             product.PromotionPrice,
		GiftGrowth:                 product.GiftGrowth,
		GiftPoint:                  product.GiftPoint,
		UsePointLimit:              product.UsePointLimit,
		SubTitle:                   product.SubTitle,
		Description:                product.Description,
		OriginalPrice:              product.OriginalPrice,
		TotalStock:                 product.TotalStock,
		TotalWarnStock:             product.TotalWarnStock,
		Unit:                       product.Unit,
		Weight:                     product.Weight,
		PreviewStatus:              product.PreviewStatus,
		ListingStatus:              product.ListingStatus,
		NewStatus:                  product.NewStatus,
		RecommendStatus:            product.RecommendStatus,
		VerifyStatus:               product.VerifyStatus,
		ServiceIds:                 product.ServiceIds,
		Keywords:                   product.Keywords,
		Note:                       product.Note,
		AlbumPics:                  product.AlbumPics,
		DetailTitle:                product.DetailTitle,
		DetailDesc:                 product.DetailDesc,
		DetailHtml:                 product.DetailHtml,
		DetailMobileHtml:           product.DetailMobileHtml,
		PromotionStartTime:         promotionStartTime,
		PromotionEndTime:           promotionEndTime,
		PromotionPerLimit:          product.PromotionPerLimit,
		PromotionType:              product.PromotionType,
		BrandName:                  product.BrandName,
		ProductCategoryName:        product.CategoryName,
		SkuList:                    productSkus,
	}, nil
}

func (s *ProductService) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductReply, error) {
	product, err := s.productDTOToDO(req.Product)
	if err != nil {
		return nil, err
	}
	id, err := s.productUseCase.CreateProduct(ctx, &biz.CreateProductDo{Product: product})
	if err != nil {
		return nil, err
	}
	return &pb.CreateProductReply{
		ProductId: strconv.FormatInt(id, 10),
	}, nil
}

func (s *ProductService) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.UpdateProductReply, error) {
	req.Product.Id = req.GetProductId()
	product, err := s.productDTOToDO(req.Product)
	if err != nil {
		return nil, err
	}
	err = s.productUseCase.UpdateProduct(ctx, &biz.UpdateProductDo{Product: product})
	return &pb.UpdateProductReply{}, nil
}

func (s *ProductService) DeleteProduct(ctx context.Context, req *pb.DeleteProductRequest) (*pb.DeleteProductReply, error) {
	id, err := strconv.ParseInt(req.GetProductId(), 10, 64)
	if err != nil {
		return nil, err
	}
	if err := s.productUseCase.DeleteProduct(ctx, id); err != nil {
		return nil, err
	}
	return &pb.DeleteProductReply{}, nil
}

func (s *ProductService) ProductDOToDTO(product *biz.Product) *pb.Product {
	productCateIds := make([]string, 0, len(product.ProductCategoryIds))
	for _, v := range product.ProductCategoryIds {
		productCateIds = append(productCateIds, strconv.FormatInt(v, 10))
	}

	p := &pb.Product{
		Id:          strconv.FormatInt(product.Id, 10),
		CategoryIds: productCateIds,
		BrandId:     strconv.FormatInt(product.BrandId, 10),
		// FlashPromotionId:           strconv.FormatInt(product.FreightTemplateId, 10),
		ProductAttributeCategoryId: strconv.FormatInt(product.ProductAttributeCategoryId, 10),
		Name:                       product.Name,
		Pic:                        product.Pic,
		ProductSn:                  product.ProductSn,
		ListingStatus:              product.ListingStatus,
		NewStatus:                  product.NewStatus,
		RecommendStatus:            product.RecommendStatus,
		VerifyStatus:               product.VerifyStatus,
		Sort:                       product.Sort,
		TotalSales:                 product.TotalSales,
		Price:                      product.Price,
		PromotionPrice:             product.PromotionPrice,
		GiftGrowth:                 product.GiftGrowth,
		GiftPoint:                  product.GiftPoint,
		UsePointLimit:              product.UsePointLimit,
		SubTitle:                   product.SubTitle,
		Description:                product.Description,
		OriginalPrice:              product.OriginalPrice,
		TotalStock:                 product.TotalStock,
		TotalWarnStock:             product.TotalWarnStock,
		Unit:                       product.Unit,
		Weight:                     product.Weight,
		PreviewStatus:              product.PreviewStatus,
		ServiceIds:                 product.ServiceIds,
		Keywords:                   product.Keywords,
		Note:                       product.Note,
		AlbumPics:                  product.AlbumPics,
		DetailTitle:                product.DetailTitle,
		DetailDesc:                 product.DetailDesc,
		DetailHtml:                 product.DetailHtml,
		DetailMobileHtml:           product.DetailMobileHtml,
		PromotionStartTime:         product.PromotionStartTime.Format(time.DateTime),
		PromotionEndTime:           product.PromotionEndTime.Format(time.DateTime),
		PromotionPerLimit:          product.PromotionPerLimit,
		PromotionType:              product.PromotionType,
		BrandName:                  product.BrandName,
		CategoryName:               product.ProductCategoryName,
		FreightTemplateId:          strconv.FormatInt(product.FreightTemplateId, 10),
	}
	return p
}

func (s *ProductService) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.GetProductReply, error) {
	id, err := strconv.ParseInt(req.GetProductId(), 10, 64)
	if err != nil {
		return nil, api.ErrorInvalidParam("invalid product id")
	}
	product, err := s.productUseCase.GetProduct(ctx, id)
	if err != nil {
		return nil, err
	}
	return &pb.GetProductReply{Product: s.ProductDOToDTO(product)}, nil
}

func (s *ProductService) ListProduct(ctx context.Context, req *pb.ListProductRequest) (*pb.ListProductReply, error) {
	total, list, err := s.productUseCase.ListProduct(ctx, req)
	if err != nil {
		return nil, err
	}
	products := make([]*pb.Product, 0, len(list))
	for _, v := range list {
		products = append(products, s.ProductDOToDTO(v))
	}
	return &pb.ListProductReply{Total: total, Products: products}, nil
}

func (s *ProductService) ListProductSku(ctx context.Context, req *pb.ListProductSkuRequest) (*pb.ListProductSkuReply, error) {
	list, err := s.productUseCase.ListProductSkus(ctx, req)
	if err != nil {
		return nil, err
	}
	skus := make([]*pb.ProductSku, 0, len(list))
	for _, v := range list {
		skus = append(skus, s.skuDOToDTO(v))
	}
	return &pb.ListProductSkuReply{ProductSkus: skus}, nil
}

func (s *ProductService) skuDOToDTO(v *biz.ProductSku) *pb.ProductSku {
	attributes := make([]*pb.Attribute, 0, len(v.Attributes))
	for _, attribute := range v.Attributes {
		attributes = append(attributes, &pb.Attribute{
			AttributeId:      strconv.FormatInt(attribute.AttributeId, 10),
			AttributeName:    attribute.AttributeName,
			AttributeValue:   attribute.AttributeValue,
			AttributeValueId: strconv.FormatInt(attribute.AttributeValueId, 10),
		})
	}
	return &pb.ProductSku{
		Code:           v.SkuCode,
		Price:          v.Price,
		Stock:          v.Stock,
		StockWarn:      v.StockWarn,
		Pic:            v.Pic,
		Sales:          v.Sales,
		PromotionPrice: v.PromotionPrice,
		Name:           v.Name,
		Attributes:     attributes,
	}
}

func (s *ProductService) DeleteProductSku(ctx context.Context, req *pb.DeleteProductSkuRequest) (*pb.DeleteProductSkuReply, error) {
	skuId, err := strconv.ParseInt(req.GetProductSkuId(), 10, 64)
	if err != nil {
		return nil, api.ErrorInvalidParam("invalid sku id")
	}
	if err := s.productUseCase.DeleteProductSku(ctx, skuId); err != nil {
	}
	return &pb.DeleteProductSkuReply{}, nil
}

func (s *ProductService) CreateProductSku(ctx context.Context, req *pb.CreateProductSkuRequest) (*pb.CreateProductSkuReply, error) {
	return &pb.CreateProductSkuReply{}, nil
}

func (s *ProductService) UpdateProductSku(ctx context.Context, req *pb.UpdateProductSkuRequest) (*pb.UpdateProductSkuReply, error) {
	req.ProductSku.Id = req.ProductSkuId
	sku, err := s.skuDTOToDO(req.ProductSku)
	if err != nil {
		return nil, err
	}
	if err := s.productUseCase.UpdateProductSku(ctx, sku); err != nil {
		return nil, err
	}
	return &pb.UpdateProductSkuReply{}, nil
}

func (s *ProductService) skuDTOToDO(req *pb.ProductSku) (*biz.ProductSku, error) {
	skuId, err := strconv.ParseInt(req.GetId(), 10, 64)
	if err != nil {
		return nil, api.ErrorInvalidParam("invalid sku id")
	}
	productId, err := strconv.ParseInt(req.GetProductId(), 10, 64)
	if err != nil {
		return nil, api.ErrorInvalidParam("invalid product id")
	}
	attributs := make([]biz.Attribute, 0, len(req.Attributes))
	for _, attribute := range req.Attributes {
		attributeId, err := strconv.ParseInt(attribute.AttributeId, 10, 64)
		if err != nil {
			return nil, api.ErrorInvalidParam("invalid attribute id")
		}
		attributeValueId, err := strconv.ParseInt(attribute.AttributeValueId, 10, 64)
		if err != nil {
			return nil, api.ErrorInvalidParam("invalid attribute value id")
		}
		attributs = append(attributs, biz.Attribute{
			AttributeId:      attributeId,
			AttributeName:    attribute.AttributeName,
			AttributeValueId: attributeValueId,
			AttributeValue:   attribute.AttributeValue,
		})
	}
	sku := &biz.ProductSku{
		Id:             skuId,
		SkuCode:        req.Code,
		Name:           req.Name,
		Price:          req.Price,
		PromotionPrice: req.PromotionPrice,
		Stock:          req.Stock,
		StockWarn:      req.StockWarn,
		Pic:            req.Pic,
		Sales:          req.Sales,
		Attributes:     attributs,
		ProductId:      productId,
	}
	return sku, nil
}
