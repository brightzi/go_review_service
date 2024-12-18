// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameOrderInfo = "order_info"

// OrderInfo 文章（种草）表
type OrderInfo struct {
	ID               int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Number           string    `gorm:"column:number;not null;comment:订单编号" json:"number"`                                           // 订单编号
	UserID           int32     `gorm:"column:user_id;not null;comment:用户id" json:"user_id"`                                         // 用户id
	PayType          bool      `gorm:"column:pay_type;not null;comment:支付方式 1微信 2支付宝 3云闪付" json:"pay_type"`                         // 支付方式 1微信 2支付宝 3云闪付
	Remark           string    `gorm:"column:remark;not null;comment:备注" json:"remark"`                                             // 备注
	PayAt            time.Time `gorm:"column:pay_at;comment:支付时间" json:"pay_at"`                                                    // 支付时间
	Status           bool      `gorm:"column:status;not null;default:1;comment:订单状态： 1待支付 2已支付待发货 3已发货 4已收货待评价 5已评价" json:"status"` // 订单状态： 1待支付 2已支付待发货 3已发货 4已收货待评价 5已评价
	ConsigneeName    string    `gorm:"column:consignee_name;not null;comment:收货人姓名" json:"consignee_name"`                          // 收货人姓名
	ConsigneePhone   string    `gorm:"column:consignee_phone;not null;comment:收货人手机号" json:"consignee_phone"`                       // 收货人手机号
	ConsigneeAddress string    `gorm:"column:consignee_address;not null;comment:收货人详细地址" json:"consignee_address"`                  // 收货人详细地址
	Price            int32     `gorm:"column:price;not null;comment:订单金额 单位分" json:"price"`                                         // 订单金额 单位分
	CouponPrice      int32     `gorm:"column:coupon_price;not null;comment:优惠券金额 单位分" json:"coupon_price"`                          // 优惠券金额 单位分
	ActualPrice      int32     `gorm:"column:actual_price;not null;comment:实际支付金额 单位分" json:"actual_price"`                         // 实际支付金额 单位分
	CreatedAt        time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt        time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName OrderInfo's table name
func (*OrderInfo) TableName() string {
	return TableNameOrderInfo
}
