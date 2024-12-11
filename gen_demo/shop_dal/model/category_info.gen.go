// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameCategoryInfo = "category_info"

// CategoryInfo 轮播图表

type CategoryInfo struct {
	ID        int32          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	ParentID  int32          `gorm:"column:parent_id;not null;comment:父级id" json:"parent_id"` // 父级id
	Name      string         `gorm:"column:name;not null" json:"name"`
	PicURL    string         `gorm:"column:pic_url;not null;comment:icon" json:"pic_url"`            // icon
	Level     bool           `gorm:"column:level;not null;default:1;comment:等级 默认1级分类" json:"level"` // 等级 默认1级分类
	Sort      bool           `gorm:"column:sort;not null;default:1" json:"sort"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName CategoryInfo's table name
func (*CategoryInfo) TableName() string {
	return TableNameCategoryInfo
}
