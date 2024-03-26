package po

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

// OmsOrder 订单
type OmsOrder struct {
	Id                    int64                 `gorm:"column:id;type:bigint;comment:订单id;primaryKey;" json:"id"`                                                  // 订单id
	MemberId              int64                 `gorm:"column:member_id;type:bigint;comment:会员id;not null;" json:"member_id"`                                      // 会员id
	CouponId              int64                 `gorm:"column:coupon_id;type:bigint;comment:优惠券id;not null;" json:"coupon_id"`                                     // 优惠券id
	OrderSn               string                `gorm:"column:order_sn;type:varchar(64);comment:订单编号;not null;" json:"order_sn"`                                   // 订单编号
	MemberUsername        string                `gorm:"column:member_username;type:varchar(64);comment:用户帐号;not null;" json:"member_username"`                     // 用户帐号
	TotalAmount           float64               `gorm:"column:total_amount;type:decimal(10, 2);comment:订单总金额;not null;" json:"total_amount"`                       // 订单总金额
	PayAmount             float64               `gorm:"column:pay_amount;type:decimal(10, 2);comment:应付金额（实际支付金额）;not null;" json:"pay_amount"`                    // 应付金额（实际支付金额）
	FreightAmount         float64               `gorm:"column:freight_amount;type:decimal(10, 2);comment:运费金额;not null;" json:"freight_amount"`                    // 运费金额
	PromotionAmount       float64               `gorm:"column:promotion_amount;type:decimal(10, 2);comment:促销优化金额（促销价、满减、阶梯价）;not null;" json:"promotion_amount"`  // 促销优化金额（促销价、满减、阶梯价）
	IntegrationAmount     float64               `gorm:"column:integration_amount;type:decimal(10, 2);comment:积分抵扣金额;not null;" json:"integration_amount"`          // 积分抵扣金额
	CouponAmount          float64               `gorm:"column:coupon_amount;type:decimal(10, 2);comment:优惠券抵扣金额;not null;" json:"coupon_amount"`                   // 优惠券抵扣金额
	DiscountAmount        float64               `gorm:"column:discount_amount;type:decimal(10, 2);comment:管理员后台调整订单使用的折扣金额;not null;" json:"discount_amount"`      // 管理员后台调整订单使用的折扣金额
	PayType               int32                 `gorm:"column:pay_type;type:int(1);comment:支付方式：0->未支付；1->支付宝；2->微信;not null;" json:"pay_type"`                    // 支付方式：0->未支付；1->支付宝；2->微信
	SourceType            int32                 `gorm:"column:source_type;type:int(1);comment:订单来源：0->PC订单；1->app订单;not null;" json:"source_type"`                 // 订单来源：0->PC订单；1->app订单
	Status                int32                 `gorm:"column:status;type:int(1);comment:订单状态：0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单;not null;" json:"status"` // 订单状态：0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单
	OrderType             int32                 `gorm:"column:order_type;type:int(1);comment:订单类型：0->正常订单；1->秒杀订单;not null;" json:"order_type"`                    // 订单类型：0->正常订单；1->秒杀订单
	DeliveryCompany       string                `gorm:"column:delivery_company;type:varchar(64);comment:物流公司(配送方式);not null;" json:"delivery_company"`             // 物流公司(配送方式)
	DeliverySn            string                `gorm:"column:delivery_sn;type:varchar(64);comment:物流单号;not null;" json:"delivery_sn"`                             // 物流单号
	AutoConfirmDay        int32                 `gorm:"column:auto_confirm_day;type:int;comment:自动确认时间（天）;not null;" json:"auto_confirm_day"`                      // 自动确认时间（天）
	Integration           int32                 `gorm:"column:integration;type:int;comment:可以获得的积分;not null;" json:"integration"`                                  // 可以获得的积分
	Growth                int32                 `gorm:"column:growth;type:int;comment:可以活动的成长值;not null;" json:"growth"`                                           // 可以活动的成长值
	PromotionInfo         string                `gorm:"column:promotion_info;type:varchar(100);comment:活动信息;not null;" json:"promotion_info"`                      // 活动信息
	BillType              int32                 `gorm:"column:bill_type;type:int(1);comment:发票类型：0->不开发票；1->电子发票；2->纸质发票;not null;" json:"bill_type"`              // 发票类型：0->不开发票；1->电子发票；2->纸质发票
	BillHeader            string                `gorm:"column:bill_header;type:varchar(200);comment:发票抬头;not null;" json:"bill_header"`                            // 发票抬头
	BillContent           string                `gorm:"column:bill_content;type:varchar(200);comment:发票内容;not null;" json:"bill_content"`                          // 发票内容
	BillReceiverPhone     string                `gorm:"column:bill_receiver_phone;type:varchar(32);comment:收票人电话;not null;" json:"bill_receiver_phone"`            // 收票人电话
	BillReceiverEmail     string                `gorm:"column:bill_receiver_email;type:varchar(64);comment:收票人邮箱;not null;" json:"bill_receiver_email"`            // 收票人邮箱
	ReceiverName          string                `gorm:"column:receiver_name;type:varchar(100);comment:收货人姓名;not null;" json:"receiver_name"`                       // 收货人姓名
	ReceiverPhone         string                `gorm:"column:receiver_phone;type:varchar(32);comment:收货人电话;not null;" json:"receiver_phone"`                      // 收货人电话
	ReceiverPostCode      string                `gorm:"column:receiver_post_code;type:varchar(32);comment:收货人邮编;not null;" json:"receiver_post_code"`              // 收货人邮编
	ReceiverProvince      string                `gorm:"column:receiver_province;type:varchar(32);comment:省份/直辖市;not null;" json:"receiver_province"`               // 省份/直辖市
	ReceiverCity          string                `gorm:"column:receiver_city;type:varchar(32);comment:城市;not null;" json:"receiver_city"`                           // 城市
	ReceiverRegion        string                `gorm:"column:receiver_region;type:varchar(32);comment:区;not null;" json:"receiver_region"`                        // 区
	ReceiverDetailAddress string                `gorm:"column:receiver_detail_address;type:varchar(200);comment:详细地址;not null;" json:"receiver_detail_address"`    // 详细地址
	Note                  string                `gorm:"column:note;type:varchar(500);comment:订单备注;not null;" json:"note"`                                          // 订单备注
	ConfirmStatus         int32                 `gorm:"column:confirm_status;type:int(1);comment:确认收货状态：0->未确认；1->已确认;not null;" json:"confirm_status"`            // 确认收货状态：0->未确认；1->已确认
	DeleteStatus          int32                 `gorm:"column:delete_status;type:int(1);comment:删除状态：0->未删除；1->已删除;not null;default:0;" json:"delete_status"`      // 删除状态：0->未删除；1->已删除
	UseIntegration        int32                 `gorm:"column:use_integration;type:int;comment:下单时使用的积分;not null;" json:"use_integration"`                         // 下单时使用的积分
	PaymentTime           time.Time             `gorm:"column:payment_time;type:datetime;comment:支付时间;not null;" json:"payment_time"`                              // 支付时间
	DeliveryTime          time.Time             `gorm:"column:delivery_time;type:datetime;comment:发货时间;not null;" json:"delivery_time"`                            // 发货时间
	ReceiveTime           time.Time             `gorm:"column:receive_time;type:datetime;comment:确认收货时间;not null;" json:"receive_time"`                            // 确认收货时间
	CommentTime           time.Time             `gorm:"column:comment_time;type:datetime;comment:评价时间;not null;" json:"comment_time"`                              // 评价时间
	CreatedAt             time.Time             `gorm:"column:created_at;type:datetime;comment:提交时间;not null;" json:"created_at"`                                  // 提交时间
	UpdatedAt             time.Time             `gorm:"column:updated_at;type:datetime;comment:修改时间;not null;" json:"updated_at"`                                  // 修改时间
	DeletedAt             soft_delete.DeletedAt `gorm:"column:deleted_at;comment:删除时间;not null;" json:"deleted_at"`                                                // 删除时间
}

type OmsOrderItem struct {
	Id                int64   `gorm:"column:id;type:bigint;primaryKey;" json:"id"`
	OrderId           int64   `gorm:"column:order_id;type:bigint;comment:订单id;not null;" json:"order_id"`                                 // 订单id
	OrderSn           string  `gorm:"column:order_sn;type:varchar(64);comment:订单编号;not null;" json:"order_sn"`                            // 订单编号
	ProductId         int64   `gorm:"column:product_id;type:bigint;comment:商品id;not null;" json:"product_id"`                             // 商品id
	ProductPic        string  `gorm:"column:product_pic;type:varchar(500);comment:商品图片;not null;" json:"product_pic"`                     // 商品图片
	ProductName       string  `gorm:"column:product_name;type:varchar(200);comment:商品名称;not null;" json:"product_name"`                   // 商品名称
	ProductBrand      string  `gorm:"column:product_brand;type:varchar(200);comment:商品品牌;not null;" json:"product_brand"`                 // 商品品牌
	ProductSn         string  `gorm:"column:product_sn;type:varchar(64);comment:货号;not null;" json:"product_sn"`                          // 货号
	ProductPrice      float64 `gorm:"column:product_price;type:decimal(10, 2);comment:销售价格;not null;" json:"product_price"`               // 销售价格
	ProductQuantity   int32   `gorm:"column:product_quantity;type:int;comment:购买数量;not null;" json:"product_quantity"`                    // 购买数量
	ProductSkuId      int64   `gorm:"column:product_sku_id;type:bigint;comment:商品sku编号;not null;" json:"product_sku_id"`                  // 商品sku编号
	ProductSkuCode    string  `gorm:"column:product_sku_code;type:varchar(50);comment:商品sku条码;not null;" json:"product_sku_code"`         // 商品sku条码
	ProductCategoryId int64   `gorm:"column:product_category_id;type:bigint;comment:商品分类id;not null;" json:"product_category_id"`         // 商品分类id
	PromotionName     string  `gorm:"column:promotion_name;type:varchar(200);comment:商品促销名称;not null;" json:"promotion_name"`             // 商品促销名称
	PromotionAmount   float64 `gorm:"column:promotion_amount;type:decimal(10, 2);comment:商品促销分解金额;not null;" json:"promotion_amount"`     // 商品促销分解金额
	CouponAmount      float64 `gorm:"column:coupon_amount;type:decimal(10, 2);comment:优惠券优惠分解金额;not null;" json:"coupon_amount"`          // 优惠券优惠分解金额
	IntegrationAmount float64 `gorm:"column:integration_amount;type:decimal(10, 2);comment:积分优惠分解金额;not null;" json:"integration_amount"` // 积分优惠分解金额
	RealAmount        float64 `gorm:"column:real_amount;type:decimal(10, 2);comment:该商品经过优惠后的分解金额;not null;" json:"real_amount"`          // 该商品经过优惠后的分解金额
	GiftIntegration   int32   `gorm:"column:gift_integration;type:int;not null;default:0;" json:"gift_integration"`
	GiftGrowth        int32   `gorm:"column:gift_growth;type:int;not null;default:0;" json:"gift_growth"`
	ProductAttr       string  `gorm:"column:product_attr;type:varchar(500);comment:商品销售属性:[{\"key\":\"颜色\",\"value\":\"颜色\"},{\"key\":\"容量\",\"value\":\"4G\"}];not null;" json:"product_attr"` // 商品销售属性:[{"key":"颜色","value":"颜色"},{"key":"容量","value":"4G"}]
}

// OmsOrderOperateHistory 订单操作历史
type OmsOrderOperateHistory struct {
	Id          int64     `gorm:"column:id;type:bigint;primaryKey;" json:"id"`
	OrderId     int64     `gorm:"column:order_id;type:bigint;comment:订单id;not null;" json:"order_id"`                                                    // 订单id
	OperateMan  string    `gorm:"column:operate_man;type:varchar(100);comment:操作人：用户；系统；后台管理员;not null;" json:"operate_man"`                             // 操作人：用户；系统；后台管理员
	CreateTime  time.Time `gorm:"column:create_time;type:datetime;comment:操作时间;not null;" json:"create_time"`                                            // 操作时间
	OrderStatus int32     `gorm:"column:order_status;type:int(1);comment:订单状态：0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单;not null;" json:"order_status"` // 订单状态：0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单
	Note        string    `gorm:"column:note;type:varchar(500);comment:备注;" json:"note"`                                                                 // 备注
}

// OmsOrderReturnApply 退货申请
type OmsOrderReturnApply struct {
	Id               int64     `gorm:"column:id;type:bigint;primaryKey;" json:"id"`
	OrderId          int64     `gorm:"column:order_id;type:bigint;comment:订单id;not null;" json:"order_id"`                                 // 订单id
	CompanyAddressId int64     `gorm:"column:company_address_id;type:bigint;comment:收货地址表id;not null;" json:"company_address_id"`          // 收货地址表id
	ProductId        int64     `gorm:"column:product_id;type:bigint;comment:退货商品id;not null;" json:"product_id"`                           // 退货商品id
	OrderSn          string    `gorm:"column:order_sn;type:varchar(64);comment:订单编号;not null;" json:"order_sn"`                            // 订单编号
	MemberUsername   string    `gorm:"column:member_username;type:varchar(64);comment:会员用户名;not null;" json:"member_username"`             // 会员用户名
	ReturnAmount     float64   `gorm:"column:return_amount;type:decimal(10, 2);comment:退款金额;not null;" json:"return_amount"`               // 退款金额
	ReturnName       string    `gorm:"column:return_name;type:varchar(100);comment:退货人姓名;not null;" json:"return_name"`                    // 退货人姓名
	ReturnPhone      string    `gorm:"column:return_phone;type:varchar(100);comment:退货人电话;not null;" json:"return_phone"`                  // 退货人电话
	Status           int32     `gorm:"column:status;type:int(1);comment:申请状态：0->待处理；1->退货中；2->已完成；3->已拒绝;not null;" json:"status"`         // 申请状态：0->待处理；1->退货中；2->已完成；3->已拒绝
	HandleTime       time.Time `gorm:"column:handle_time;type:datetime;comment:处理时间;not null;" json:"handle_time"`                         // 处理时间
	ProductPic       string    `gorm:"column:product_pic;type:varchar(500);comment:商品图片;not null;" json:"product_pic"`                     // 商品图片
	ProductName      string    `gorm:"column:product_name;type:varchar(200);comment:商品名称;not null;" json:"product_name"`                   // 商品名称
	ProductBrand     string    `gorm:"column:product_brand;type:varchar(200);comment:商品品牌;not null;" json:"product_brand"`                 // 商品品牌
	ProductAttr      string    `gorm:"column:product_attr;type:varchar(500);comment:商品销售属性：颜色：红色；尺码：xl;;not null;" json:"product_attr"`    // 商品销售属性：颜色：红色；尺码：xl;
	ProductCount     int32     `gorm:"column:product_count;type:int;comment:退货数量;not null;" json:"product_count"`                          // 退货数量
	ProductPrice     float64   `gorm:"column:product_price;type:decimal(10, 2);comment:商品单价;not null;" json:"product_price"`               // 商品单价
	ProductRealPrice float64   `gorm:"column:product_real_price;type:decimal(10, 2);comment:商品实际支付单价;not null;" json:"product_real_price"` // 商品实际支付单价
	Reason           string    `gorm:"column:reason;type:varchar(200);comment:原因;not null;" json:"reason"`                                 // 原因
	Description      string    `gorm:"column:description;type:varchar(500);comment:描述;not null;" json:"description"`                       // 描述
	ProofPics        string    `gorm:"column:proof_pics;type:varchar(1000);comment:凭证图片，以逗号隔开;not null;" json:"proof_pics"`                // 凭证图片，以逗号隔开
	HandleNote       string    `gorm:"column:handle_note;type:varchar(500);comment:处理备注;not null;" json:"handle_note"`                     // 处理备注
	HandleMan        string    `gorm:"column:handle_man;type:varchar(100);comment:处理人员;not null;" json:"handle_man"`                       // 处理人员
	ReceiveMan       string    `gorm:"column:receive_man;type:varchar(100);comment:收货人;not null;" json:"receive_man"`                      // 收货人
	ReceivedTime     time.Time `gorm:"column:receive_time;type:datetime;comment:收货时间;not null;" json:"receive_time"`                       // 收货时间
	ReceiveNote      string    `gorm:"column:receive_note;type:varchar(500);comment:收货备注;" json:"receive_note"`                            // 收货备注
	CreatedAt        time.Time `gorm:"column:created_at;type:datetime;comment:申请时间;not null;" json:"created_at"`                           // 申请时间
}

// OmsOrderReturnReason 退货原因
type OmsOrderReturnReason struct {
	Id        int64     `gorm:"column:id;type:bigint;primaryKey;" json:"id"`
	Name      string    `gorm:"column:name;type:varchar(100);comment:退货类型;not null;" json:"name"`          // 退货类型
	Sort      int32     `gorm:"column:sort;type:int;comment:排序;not null;" json:"sort"`                     // 排序
	Status    int32     `gorm:"column:status;type:int(1);comment:状态：0->不启用；1->启用;not null;" json:"status"` // 状态：0->不启用；1->启用
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;comment:添加时间;not null;" json:"created_at"`  // 添加时间
}

// OmsCompanyAddress 公司收发货地址表
type OmsCompanyAddress struct {
	Id            int64  `gorm:"column:id;type:bigint;primaryKey;" json:"id"`
	AddressName   string `gorm:"column:address_name;type:varchar(200);comment:地址名称;not null;" json:"address_name"`             // 地址名称
	SendStatus    int32  `gorm:"column:send_status;type:int(1);comment:默认发货地址：0->否；1->是;not null;" json:"send_status"`         // 默认发货地址：0->否；1->是
	ReceiveStatus int32  `gorm:"column:receive_status;type:int(1);comment:是否默认收货地址：0->否；1->是;not null;" json:"receive_status"` // 是否默认收货地址：0->否；1->是
	Name          string `gorm:"column:name;type:varchar(64);comment:收发货人姓名;not null;" json:"name"`                            // 收发货人姓名
	Phone         string `gorm:"column:phone;type:varchar(64);comment:收货人电话;not null;" json:"phone"`                           // 收货人电话
	Province      string `gorm:"column:province;type:varchar(64);comment:省/直辖市;not null;" json:"province"`                     // 省/直辖市
	City          string `gorm:"column:city;type:varchar(64);comment:市;not null;" json:"city"`                                 // 市
	Region        string `gorm:"column:region;type:varchar(64);comment:区;not null;" json:"region"`                             // 区
	DetailAddress string `gorm:"column:detail_address;type:varchar(200);comment:详细地址;not null;" json:"detail_address"`         // 详细地址
}

// OmsCartItem 购物车
type OmsCartItem struct {
	Id                int64     `gorm:"column:id;type:bigint;primaryKey;" json:"id"`
	ProductId         int64     `gorm:"column:product_id;type:bigint;comment:商品id;not null;" json:"product_id"`                                                                   // 商品id
	ProductSkuId      int64     `gorm:"column:product_sku_id;type:bigint;comment:商品库存id;not null;" json:"product_sku_id"`                                                         // 商品库存id
	MemberId          int64     `gorm:"column:member_id;type:bigint;comment:会员id;not null;" json:"member_id"`                                                                     // 会员id
	Quantity          int32     `gorm:"column:quantity;type:int;comment:购买数量;not null;" json:"quantity"`                                                                          // 购买数量
	Price             float64   `gorm:"column:price;type:decimal(10, 2);comment:添加到购物车的价格;not null;" json:"price"`                                                                // 添加到购物车的价格
	ProductPic        string    `gorm:"column:product_pic;type:varchar(1000);comment:商品主图;not null;" json:"product_pic"`                                                          // 商品主图
	ProductName       string    `gorm:"column:product_name;type:varchar(500);comment:商品名称;not null;" json:"product_name"`                                                         // 商品名称
	ProductSubTitle   string    `gorm:"column:product_sub_title;type:varchar(500);comment:商品副标题（卖点）;not null;" json:"product_sub_title"`                                          // 商品副标题（卖点）
	ProductSkuCode    string    `gorm:"column:product_sku_code;type:varchar(200);comment:商品sku条码;not null;" json:"product_sku_code"`                                              // 商品sku条码
	MemberNickname    string    `gorm:"column:member_nickname;type:varchar(500);comment:会员昵称;not null;" json:"member_nickname"`                                                   // 会员昵称
	CreateDate        time.Time `gorm:"column:create_date;type:datetime;comment:创建时间;not null;" json:"create_date"`                                                               // 创建时间
	ModifyDate        time.Time `gorm:"column:modify_date;type:datetime;comment:修改时间;not null;" json:"modify_date"`                                                               // 修改时间
	DeleteStatus      int32     `gorm:"column:delete_status;type:int(1);comment:是否删除;not null;default:0;" json:"delete_status"`                                                   // 是否删除
	ProductCategoryId int64     `gorm:"column:product_category_id;type:bigint;comment:商品分类;not null;" json:"product_category_id"`                                                 // 商品分类
	ProductBrand      string    `gorm:"column:product_brand;type:varchar(200);comment:商品品牌;not null;" json:"product_brand"`                                                       // 商品品牌
	ProductSn         string    `gorm:"column:product_sn;type:varchar(200);comment:货号;not null;" json:"product_sn"`                                                               // 货号
	ProductAttr       string    `gorm:"column:product_attr;type:varchar(500);comment:商品销售属性:[{"key":"颜色","value":"颜色"},{"key":"容量","value":"4G"}];not null;" json:"product_attr"` // 商品销售属性:[{"key":"颜色","value":"颜色"},{"key":"容量","value":"4G"}]
}

// OmsOrderSetting 订单设置
type OmsOrderSetting struct {
	Id                  int64 `gorm:"column:id;type:bigint;primaryKey;" json:"id"`
	FlashOrderOvertime  int32 `gorm:"column:flash_order_overtime;type:int;comment:秒杀订单超时关闭时间(分);not null;" json:"flash_order_overtime"` // 秒杀订单超时关闭时间(分)
	NormalOrderOvertime int32 `gorm:"column:normal_order_overtime;type:int;comment:正常订单超时时间(分);not null;" json:"normal_order_overtime"` // 正常订单超时时间(分)
	ConfirmOvertime     int32 `gorm:"column:confirm_overtime;type:int;comment:发货后自动确认收货时间（天）;not null;" json:"confirm_overtime"`        // 发货后自动确认收货时间（天）
	FinishOvertime      int32 `gorm:"column:finish_overtime;type:int;comment:自动完成交易时间，不能申请售后（天）;not null;" json:"finish_overtime"`      // 自动完成交易时间，不能申请售后（天）
	CommentOvertime     int32 `gorm:"column:comment_overtime;type:int;comment:订单完成后自动好评时间（天）;not null;" json:"comment_overtime"`        // 订单完成后自动好评时间（天）
}
