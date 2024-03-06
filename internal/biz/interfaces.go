package biz

import (
	"context"
	v1 "kratos-admin/api/product/v1"
	"time"
)

// CartRepo 购物车仓储接口
type CartRepo interface {
	AddItem(ctx context.Context, param *AddCartItemParam) error
	DeleteItem(ctx context.Context, param *DeleteCartItemParam) error
	UpdateItemQuantity(ctx context.Context, param *UpdateCartItemParam) error
	GetCartList(ctx context.Context, param *GetCartListParam) ([]*CartItem, error)
	ClearCart(ctx context.Context, param *ClearCartParam) error
}

type Product struct {
	Id                         int64         // 商品id
	ProductCategoryIds         []int64       // 产品分类列表
	FreightTemplateId          int64         // 运费模版
	BrandId                    int64         // 品牌id
	ProductAttributeCategoryId int64         // 属性分类
	Name                       string        // 商品名称
	Pic                        string        // 主图
	ProductSn                  string        // 商品编码
	Sort                       int32         // 排序
	TotalSales                 int32         // 销量
	Price                      float64       // 售卖价(sku里的最低价格)
	PromotionPrice             float64       // 促销价
	GiftGrowth                 int32         // 赠送的成长值
	GiftPoint                  int32         // 赠送的积分
	UsePointLimit              int32         // 限制使用的积分
	SubTitle                   string        // 副标题
	Description                string        // 描述
	OriginalPrice              float64       // 原价
	TotalStock                 int32         // 库存
	Unit                       string        // 单位
	Weight                     float64       // 重量
	PreviewStatus              int8          // 是否为预览商品
	PublishStatus              int8          // 上下架状态
	NewStatus                  int8          // 是否新品
	RecommendStatus            int8          // 是否推荐
	VerifyStatus               int8          // 审核状态
	ServiceIds                 []int64       // 产品服务id数组
	Keywords                   []string      // 关键字
	Note                       string        // 商品备注
	AlbumPics                  []string      // 画册图片,连产品图片限制为5张
	DetailTitle                string        // 产品详述标题
	DetailDesc                 string        // 产品详述描述
	DetailHtml                 string        // 产品详述
	DetailMobileHtml           string        // 产品详述移动端
	PromotionStartTime         time.Time     // 促销开始时间
	PromotionEndTime           time.Time     // 促销结束时间
	PromotionPerLimit          int32         // 活动限购数量
	PromotionType              int8          // 促销类型
	BrandName                  string        // 品牌名称
	ProductCategoryName        string        // 产品分类名称
	SkuList                    []*ProductSku // sku列表
}

type ProductSku struct {
	Id             int64       // sku id
	SkuCode        string      // sku编码
	Name           string      // sku名称
	Price          float64     // 售价
	PromotionPrice float64     // 促销价
	Stock          int32       // 库存
	StockWarn      int32       // 库存预警值
	Pic            string      // 主图
	Sales          int32       // 销量
	GiftBlockStock int32       // 赠送的库存
	Attributes     []Attribute // 属性数据json
	ProductId      int64       // 产品id
}

type Attribute struct {
	AttributeId        int    `json:"attribute_id"`
	AttributeName      string `json:"attribute_name"`
	AttributeValueId   int    `json:"attribute_value_id"`
	AttributeValueName string `json:"attribute_value_name"`
}

type CreateProductDo struct {
	*Product
}

type UpdateProductDo struct {
	*Product
}

type Pagination[T any] struct {
	Total int64
	Items []T
}
type Page struct {
	Number int // 页码
	Size   int // 每页条数
}

type ListProductParam struct {
	*Page
}

type ListProductSkuParam struct {
	ProductId int64
}

type UpdateProductSkuDo struct {
	*ProductSku
}

// ProductRepo 商品仓储接口
type ProductRepo interface {
	// FindProductById 查询商品
	FindProductById(ctx context.Context, id int64) (*Product, error)
	// FindProductSkuBySkuId 查询商品sku
	FindProductSkuBySkuId(ctx context.Context, skuId int64) (*ProductSku, error)
	// FindProductSkusBySkuIdList 查询商品sku
	FindProductSkusBySkuIdList(ctx context.Context, skuIds []int64) ([]*ProductSku, error)
	// FindProductSkusByProductId 查询商品sku列表
	FindProductSkusByProductId(ctx context.Context, productId int64) ([]*ProductSku, error)
	// CreateProduct 创建商品
	CreateProduct(ctx context.Context, param *CreateProductDo) (int64, error)
	// UpdateProduct 更新商品
	UpdateProduct(ctx context.Context, param *UpdateProductDo) error
	// UpdateProductSku 更新商品sku
	UpdateProductSku(ctx context.Context, param *UpdateProductSkuDo) error
	// DeleteProduct 删除商品
	DeleteProduct(ctx context.Context, productId int64) error
	// DeleteProductSku 删除商品sku
	DeleteProductSku(ctx context.Context, skuId int64) error
	// FindProductList 查询商品列表
	FindProductList(ctx context.Context, req *ListProductParam) (total int64, list []*Product, err error)
}

// ProductCategoryRepo 商品分类仓储接口
type ProductCategoryRepo interface {
	// FindProductCategoryById 查询商品分类
	FindProductCategoryById(ctx context.Context, id int64) (*v1.ProductCategory, error)
	// FindProductCategoryList 查询商品分类列表
	FindProductCategoryList(ctx context.Context, param *v1.FindProductCategoryListRequest) (reply *v1.FindProductCategoryListReply, err error)
	// CreateProductCategory 创建商品分类
	CreateProductCategory(ctx context.Context, param *v1.CreateProductCategoryRequest) (int64, error)
	// UpdateProductCategory 更新商品分类
	UpdateProductCategory(ctx context.Context, param *v1.UpdateProductCategoryRequest) error
	// DeleteProductCategory 删除商品分类
	DeleteProductCategory(ctx context.Context, id int64) error
	// FindProductCategoryListByParentId 查询商品分类列表
	FindProductCategoryListByParentId(ctx context.Context, parentId int64) ([]*v1.ProductCategory, error)
	// FindProductCategoryTree 获取商品分类树形结构
	FindProductCategoryTree(ctx context.Context) ([]*v1.ProductCategoryTreeNode, error)
}
