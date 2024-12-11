// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameGoodsInfo = "goods_info"

// GoodsInfo 商品表
type GoodsInfo struct {
	ID               int32          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	PicURL           string         `gorm:"column:pic_url;not null;comment:图片" json:"pic_url"`                           // 图片
	Name             string         `gorm:"column:name;not null;comment:商品名称" json:"name"`                               // 商品名称
	Price            int32          `gorm:"column:price;not null;default:1;comment:价格 单位分" json:"price"`                 // 价格 单位分
	Level1CategoryID int32          `gorm:"column:level1_category_id;not null;comment:1级分类id" json:"level1_category_id"` // 1级分类id
	Level2CategoryID int32          `gorm:"column:level2_category_id;not null;comment:2级分类id" json:"level2_category_id"` // 2级分类id
	Level3CategoryID int32          `gorm:"column:level3_category_id;not null;comment:3级分类id" json:"level3_category_id"` // 3级分类id
	Brand            string         `gorm:"column:brand;not null;comment:品牌" json:"brand"`                               // 品牌
	Stock            int32          `gorm:"column:stock;not null;comment:库存" json:"stock"`                               // 库存
	Sale             int32          `gorm:"column:sale;not null;comment:销量" json:"sale"`                                 // 销量
	Tags             string         `gorm:"column:tags;not null;comment:标签" json:"tags"`                                 // 标签
	DetailInfo       string         `gorm:"column:detail_info;comment:商品详情" json:"detail_info"`                          // 商品详情
	CreatedAt        time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt        time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName GoodsInfo's table name
func (*GoodsInfo) TableName() string {
	return TableNameGoodsInfo
}