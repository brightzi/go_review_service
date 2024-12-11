// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNamePositionInfo = "position_info"

// PositionInfo mapped from table <position_info>
type PositionInfo struct {
	ID        int32          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	PicURL    string         `gorm:"column:pic_url;not null;comment:图片链接" json:"pic_url"`       // 图片链接
	GoodsName string         `gorm:"column:goods_name;not null;comment:商品名称" json:"goods_name"` // 商品名称
	Link      string         `gorm:"column:link;not null;comment:跳转链接" json:"link"`             // 跳转链接
	Sort      int32          `gorm:"column:sort;not null;comment:排序" json:"sort"`               // 排序
	GoodsID   int32          `gorm:"column:goods_id;not null;comment:商品id" json:"goods_id"`     // 商品id
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName PositionInfo's table name
func (*PositionInfo) TableName() string {
	return TableNamePositionInfo
}
