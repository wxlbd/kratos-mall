package biz

import "context"

// CartRepo 购物车仓储接口
type CartRepo interface {
	AddItem(ctx context.Context, param *AddCartItemParam) error
	DeleteItem(ctx context.Context, param *DeleteCartItemParam) error
	UpdateItemQuantity(ctx context.Context, param *UpdateCartItemParam) error
	GetCartList(ctx context.Context, param *GetCartListParam) ([]*CartItem, error)
	ClearCart(ctx context.Context, param *ClearCartParam) error
}

type Product struct {
	ProductCategoryId          []int64       // 产品分类列表
	BrandId                    int64         // 品牌id
	FlashPromotionId           int64         // 闪购活动
	ProductAttributeCategoryId int64         // 属性分类
	Name                       string        // 商品名称
	Pic                        string        // 主图
	ProductSn                  string        // 商品编码
	PublishStatus              int64         // 上下架状态
	NewStatus                  int64         // 是否新品
	RecommendStatus            int64         // 是否推荐
	VerifyStatus               int64         // 审核状态
	Sort                       int64         // 排序
	TotalSales                 int64         // 销量
	Price                      float64       // 售卖价(sku里的最低价格)
	PromotionPrice             float64       // 促销价
	GiftGrowth                 int64         // 赠送的成长值
	GiftPoint                  int64         // 赠送的积分
	UsePointLimit              int64         // 限制使用的积分
	SubTitle                   string        // 副标题
	Description                string        // 描述
	OriginalPrice              float64       // 原价
	Stock                      int64         // 库存
	StockWarn                  int64         // 库存预警值
	Unit                       string        // 单位
	Weight                     float64       // 重量
	PreviewStatus              int64         // 是否为预览商品
	ServiceIds                 []int64       // 产品服务id数组
	Keywords                   []string      // 关键字
	Note                       string        // 商品备注
	AlbumPics                  []string      // 画册图片,连产品图片限制为5张
	DetailTitle                string        // 产品详述标题
	DetailDesc                 string        // 产品详述描述
	DetailHtml                 string        // 产品详述
	DetailMobileHtml           string        // 产品详述移动端
	PromotionStartTime         string        // 促销开始时间
	PromotionEndTime           string        // 促销结束时间
	PromotionPerLimit          int64         // 活动限购数量
	PromotionType              int64         // 促销类型
	BrandName                  string        // 品牌名称
	ProductCategoryName        string        // 产品分类名称
	SkuList                    []*ProductSku // sku列表
}

type ProductSku struct {
	Id             int64   // sku id
	SkuCode        string  // sku编码
	Price          float32 // 售价
	Stock          int64   // 库存
	StockWarn      int64   // 库存预警值
	Pic            string  // 主图
	Sales          int64   // 销量
	PromotionPrice float32 // 促销价
	GiftBlockStock int64   // 赠送的库存
	SkuName        string  // sku名称
	AttributeData  string  // 属性数据json
}
type CreateProductDo struct {
	*Product
}

type UpdateProductDo struct {
	Id int64
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
	// DeleteProduct 删除商品
	DeleteProduct(ctx context.Context, productId int64) error
	// DeleteProductSku 删除商品sku
	DeleteProductSku(ctx context.Context, skuId int64) error
	// FindProductList 查询商品列表
	FindProductList(ctx context.Context, req *ListProductParam) (total int64, list []*Product, err error)
}
