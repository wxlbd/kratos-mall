package po

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

// PmsProductCategory 商品分类
type PmsProductCategory struct {
	Id           int64  `gorm:"column:id;type:bigint;primaryKey;" json:"id"`
	ParentId     int64  `gorm:"column:parent_id;type:bigint;comment:上机分类的编号：0表示一级分类;not null;" json:"parent_id"` // 上机分类的编号：0表示一级分类
	Name         string `gorm:"column:name;type:varchar(64);not null;" json:"name"`
	Level        int32  `gorm:"column:level;type:int(1);comment:分类级别：0->1级；1->2级;not null;" json:"level"` // 分类级别：0->1级；1->2级
	ProductCount int32  `gorm:"column:product_count;type:int;not null;" json:"product_count"`
	ProductUnit  string `gorm:"column:product_unit;type:varchar(64);not null;" json:"product_unit"`
	NavStatus    int32  `gorm:"column:nav_status;type:int(1);comment:是否显示在导航栏：0->不显示；1->显示;not null;" json:"nav_status"` // 是否显示在导航栏：0->不显示；1->显示
	ShowStatus   int32  `gorm:"column:show_status;type:int(1);comment:显示状态：0->不显示；1->显示;not null;" json:"show_status"`   // 显示状态：0->不显示；1->显示
	Sort         int32  `gorm:"column:sort;type:int;not null;" json:"sort"`
	Icon         string `gorm:"column:icon;type:varchar(255);comment:图标;not null;" json:"icon"` // 图标
	Keywords     string `gorm:"column:keywords;type:varchar(255);not null;" json:"keywords"`
	Description  string `gorm:"column:description;type:text;comment:描述;" json:"description"` // 描述
}

// PmsProduct 商品
type PmsProduct struct {
	Id                         int64                 `gorm:"column:id;type:bigint;primaryKey;" json:"id"`
	BrandId                    int64                 `gorm:"column:brand_id;type:bigint;comment:品牌id;not null;" json:"brand_id"`                                                                      // 品牌id
	ProductCategoryId          int64                 `gorm:"column:product_category_id;type:bigint;comment:商品分类id;not null;" json:"product_category_id"`                                              // 商品分类id
	FreightTemplateId          int64                 `gorm:"column:freight_template_id;type:bigint;comment:商品运费模板id;not null;" json:"freight_template_id"`                                            // 商品运费模板id
	ProductAttributeCategoryId int64                 `gorm:"column:product_attribute_category_id;type:bigint;comment:商品属性分类id;not null;" json:"product_attribute_category_id"`                        // 商品属性分类id
	Name                       string                `gorm:"column:name;type:varchar(64);comment:商品名称;not null;" json:"name"`                                                                         // 商品名称
	Pic                        string                `gorm:"column:pic;type:varchar(255);comment:商品图片;not null;" json:"pic"`                                                                          // 商品图片
	ProductSn                  string                `gorm:"column:product_sn;type:varchar(64);comment:货号;not null;" json:"product_sn"`                                                               // 货号
	Sort                       int32                 `gorm:"column:sort;type:int;comment:排序;not null;" json:"sort"`                                                                                   // 排序
	TotalSales                 int32                 `gorm:"column:total_sales;type:int;comment:销量;not null;" json:"total_sales"`                                                                     // 销量
	Price                      float64               `gorm:"column:price;type:decimal(10, 2);comment:商品价格;not null;" json:"price"`                                                                    // 商品价格
	PromotionPrice             float64               `gorm:"column:promotion_price;type:decimal(10, 2);comment:促销价格;not null;" json:"promotion_price"`                                                // 促销价格
	GiftGrowth                 int32                 `gorm:"column:gift_growth;type:int;comment:赠送的成长值;not null;default:0;" json:"gift_growth"`                                                       // 赠送的成长值
	GiftPoint                  int32                 `gorm:"column:gift_point;type:int;comment:赠送的积分;not null;default:0;" json:"gift_point"`                                                          // 赠送的积分
	UsePointLimit              int32                 `gorm:"column:use_point_limit;type:int;comment:限制使用的积分数;not null;" json:"use_point_limit"`                                                       // 限制使用的积分数
	SubTitle                   string                `gorm:"column:sub_title;type:varchar(255);comment:副标题;not null;" json:"sub_title"`                                                               // 副标题
	Description                string                `gorm:"column:description;type:text;comment:商品描述;not null;" json:"description"`                                                                  // 商品描述
	OriginalPrice              float64               `gorm:"column:original_price;type:decimal(10, 2);comment:市场价;not null;" json:"original_price"`                                                   // 市场价
	TotalStock                 int32                 `gorm:"column:total_stock;type:int;comment:库存;not null;" json:"total_stock"`                                                                     // 库存
	Unit                       string                `gorm:"column:unit;type:varchar(16);comment:单位;not null;" json:"unit"`                                                                           // 单位
	Weight                     float64               `gorm:"column:weight;type:decimal(10, 2);comment:商品重量，默认为克;not null;" json:"weight"`                                                             // 商品重量，默认为克
	PreviewStatus              int32                 `gorm:"column:preview_status;type:int(1);comment:是否为预告商品：0->不是；1->是;not null;" json:"preview_status"`                                            // 是否为预告商品：0->不是；1->是
	PublishStatus              int32                 `gorm:"column:publish_status;type:int(1);comment:上架状态：0->下架；1->上架;not null;" json:"publish_status"`                                              // 上架状态：0->下架；1->上架
	NewStatus                  int32                 `gorm:"column:new_status;type:int(1);comment:新品状态:0->不是新品；1->新品;not null;" json:"new_status"`                                                    // 新品状态:0->不是新品；1->新品
	RecommendStatus            int32                 `gorm:"column:recommend_status;type:int(1);comment:推荐状态；0->不推荐；1->推荐;not null;" json:"recommend_status"`                                         // 推荐状态；0->不推荐；1->推荐
	VerifyStatus               int32                 `gorm:"column:verify_status;type:int(1);comment:审核状态：0->未审核；1->审核通过;not null;" json:"verify_status"`                                             // 审核状态：0->未审核；1->审核通过
	ServiceIds                 string                `gorm:"column:service_ids;type:varchar(64);comment:以逗号分割的产品服务：1->无忧退货；2->快速退款；3->免费包邮;not null;" json:"service_ids"`                             // 以逗号分割的产品服务：1->无忧退货；2->快速退款；3->免费包邮
	Keywords                   string                `gorm:"column:keywords;type:varchar(255);comment:搜索关键字;not null;" json:"keywords"`                                                               // 搜索关键字
	Note                       string                `gorm:"column:note;type:varchar(255);comment:备注;not null;" json:"note"`                                                                          // 备注
	AlbumPics                  string                `gorm:"column:album_pics;type:varchar(255);comment:画册图片，连产品图片限制为5张，以逗号分割;not null;" json:"album_pics"`                                           // 画册图片，连产品图片限制为5张，以逗号分割
	DetailTitle                string                `gorm:"column:detail_title;type:varchar(255);comment:详情标题;not null;" json:"detail_title"`                                                        // 详情标题
	DetailDesc                 string                `gorm:"column:detail_desc;type:text;comment:详情描述;not null;" json:"detail_desc"`                                                                  // 详情描述
	DetailHtml                 string                `gorm:"column:detail_html;type:text;comment:产品详情网页内容;not null;" json:"detail_html"`                                                              // 产品详情网页内容
	DetailMobileHtml           string                `gorm:"column:detail_mobile_html;type:text;comment:移动端网页详情;not null;" json:"detail_mobile_html"`                                                 // 移动端网页详情
	PromotionStartTime         time.Time             `gorm:"column:promotion_start_time;type:datetime;comment:促销开始时间;not null;" json:"promotion_start_time"`                                          // 促销开始时间
	PromotionEndTime           time.Time             `gorm:"column:promotion_end_time;type:datetime;comment:促销结束时间;not null;" json:"promotion_end_time"`                                              // 促销结束时间
	PromotionPerLimit          int32                 `gorm:"column:promotion_per_limit;type:int;comment:活动限购数量;not null;" json:"promotion_per_limit"`                                                 // 活动限购数量
	PromotionType              int32                 `gorm:"column:promotion_type;type:int(1);comment:促销类型：0->没有促销使用原价;1->使用促销价；2->使用会员价；3->使用阶梯价格；4->使用满减价格；5->限时购;not null;" json:"promotion_type"` // 促销类型：0->没有促销使用原价;1->使用促销价；2->使用会员价；3->使用阶梯价格；4->使用满减价格；5->限时购
	BrandName                  string                `gorm:"column:brand_name;type:varchar(255);comment:品牌名称;not null;" json:"brand_name"`                                                            // 品牌名称
	ProductCategoryName        string                `gorm:"column:product_category_name;type:varchar(255);comment:商品分类名称;not null;" json:"product_category_name"`                                    // 商品分类名称
	ProductCategoryIdArray     []int64               `gorm:"column:product_category_id_array;type:json;comment:商品分类id字符串;not null;" json:"product_category_id_array"`                                 // 商品分类id字符串
	CreatedAt                  time.Time             `gorm:"column:created_at;type:datetime;comment:创建时间;not null;" json:"created_at"`
	UpdatedAt                  time.Time             `gorm:"column:updated_at;type:datetime;comment:更新时间;not null;" json:"updated_at"`
	DeletedAt                  soft_delete.DeletedAt `gorm:"softDelete:nano,DeletedAtField:DeletedAt;column:deleted_at;comment:删除时间;not null;" json:"deleted_at"`
}

// ProductSku 商品sku
type ProductSku struct {
	Id        int64                 `gorm:"column:id;type:bigint;primaryKey;" json:"id"`                  // sku id
	ProductId int64                 `gorm:"column:product_id;type:bigint;not null;" json:"product_id"`    // 商品id
	SkuCode   string                `gorm:"column:sku_code;type:varchar(64);not null;" json:"sku_code"`   // sku编码
	Specs     string                `gorm:"column:specs;type:json;comment:规格;not null;" json:"specs"`     // 规格
	Price     float64               `gorm:"column:price;type:decimal(10,2);not null;" json:"price"`       // 价格
	Pic       string                `gorm:"column:pic;type:varchar(255);comment:图片;not null;" json:"pic"` // sku图片
	Stock     int32                 `gorm:"column:stock;type:int;comment:库存;not null;" json:"stock"`      // 库存
	StockWarn int32                 `gorm:"column:stock_warn;type:int;comment:库存预警值;not null;" json:"stock_warn"`
	Sales     int32                 `gorm:"column:sales;type:int;comment:销量;not null;" json:"sales"` // 销量
	CreatedAt time.Time             `gorm:"column:created_at;type:datetime;comment:创建时间;not null;" json:"created_at"`
	UpdatedAt time.Time             `gorm:"column:updated_at;type:datetime;comment:更新时间;not null;" json:"updated_at"`
	DeletedAt soft_delete.DeletedAt `gorm:"softDelete:nano,DeletedAtField:DeletedAt;column:deleted_at;comment:删除时间;not null;" json:"deleted_at"`
}

// PmsProductAttribute 商品属性参数表
type PmsProductAttribute struct {
	Id                         int64  `gorm:"column:id;type:bigint;primaryKey;" json:"id"`
	ProductAttributeCategoryId int64  `gorm:"column:product_attribute_category_id;type:bigint;not null;" json:"product_attribute_category_id"`
	Name                       string `gorm:"column:name;type:varchar(64);not null;" json:"name"`
	SelectType                 int32  `gorm:"column:select_type;type:int(1);comment:属性选择类型：0->唯一；1->单选；2->多选;not null;" json:"select_type"`         // 属性选择类型：0->唯一；1->单选；2->多选
	InputType                  int32  `gorm:"column:input_type;type:int(1);comment:属性录入方式：0->手工录入；1->从列表中选取;not null;" json:"input_type"`           // 属性录入方式：0->手工录入；1->从列表中选取
	InputList                  string `gorm:"column:input_list;type:varchar(255);comment:可选值列表，以逗号隔开;not null;" json:"input_list"`                  // 可选值列表，以逗号隔开
	Sort                       int32  `gorm:"column:sort;type:int;comment:排序字段：最高的可以单独上传图片;not null;" json:"sort"`                                  // 排序字段：最高的可以单独上传图片
	FilterType                 int32  `gorm:"column:filter_type;type:int(1);comment:分类筛选样式：1->普通；1->颜色;not null;" json:"filter_type"`               // 分类筛选样式：1->普通；1->颜色
	SearchType                 int32  `gorm:"column:search_type;type:int(1);comment:检索类型；0->不需要进行检索；1->关键字检索；2->范围检索;not null;" json:"search_type"` // 检索类型；0->不需要进行检索；1->关键字检索；2->范围检索
	RelatedStatus              int32  `gorm:"column:related_status;type:int(1);comment:相同属性产品是否关联；0->不关联；1->关联;not null;" json:"related_status"`    // 相同属性产品是否关联；0->不关联；1->关联
	HandAddStatus              int32  `gorm:"column:hand_add_status;type:int(1);comment:是否支持手动新增；0->不支持；1->支持;not null;" json:"hand_add_status"`    // 是否支持手动新增；0->不支持；1->支持
	Type                       int32  `gorm:"column:type;type:int(1);comment:属性的类型；0->规格；1->参数;not null;" json:"type"`                              // 属性的类型；0->规格；1->参数
}

// PmsProductAttributeCategory 商品属性分类
type PmsProductAttributeCategory struct {
	Id             int64  `gorm:"column:id;type:bigint;primaryKey;" json:"id"`
	Name           string `gorm:"column:name;type:varchar(64);not null;" json:"name"`
	AttributeCount int32  `gorm:"column:attribute_count;type:int;comment:属性数量;not null;default:0;" json:"attribute_count"` // 属性数量
	ParamCount     int32  `gorm:"column:param_count;type:int;comment:参数数量;not null;default:0;" json:"param_count"`         // 参数数量
}

// PmsProductAttributeValue 商品属性值
type PmsProductAttributeValue struct {
	Id                 int64  `gorm:"column:id;type:bigint;primaryKey;" json:"id"`
	ProductId          int64  `gorm:"column:product_id;type:bigint;not null;" json:"product_id"`
	ProductAttributeId int64  `gorm:"column:product_attribute_id;type:bigint;not null;" json:"product_attribute_id"`
	Value              string `gorm:"column:value;type:varchar(64);comment:手动添加规格或参数的值，参数单值，规格有多个时以逗号隔开;" json:"value"` // 手动添加规格或参数的值，参数单值，规格有多个时以逗号隔开
}

// PmsProductCategoryAttributeRelation 产品的分类和属性的关系表，用于设置分类筛选条件（只支持一级分类）
type PmsProductCategoryAttributeRelation struct {
	Id                 int64 `gorm:"column:id;type:bigint;primaryKey;" json:"id"`
	ProductCategoryId  int64 `gorm:"column:product_category_id;type:bigint;not null;" json:"product_category_id"`
	ProductAttributeId int64 `gorm:"column:product_attribute_id;type:bigint;not null;" json:"product_attribute_id"`
}
