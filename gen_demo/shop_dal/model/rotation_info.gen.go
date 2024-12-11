// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameRotationInfo = "rotation_info"

// RotationInfo 轮播图表

type RotationInfo struct {
	ID        int32          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	PicURL    string         `gorm:"column:pic_url;not null;comment:轮播图片" json:"pic_url"` // 轮播图片
	Link      string         `gorm:"column:link;not null;comment:跳转链接" json:"link"`       // 跳转链接
	Sort      bool           `gorm:"column:sort;not null;comment:排序字段" json:"sort"`       // 排序字段
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName RotationInfo's table name
func (*RotationInfo) TableName() string {
	return TableNameRotationInfo
}