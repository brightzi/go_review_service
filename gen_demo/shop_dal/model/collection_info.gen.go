// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameCollectionInfo = "collection_info"

// CollectionInfo mapped from table <collection_info>
type CollectionInfo struct {
	ID        int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UserID    int32     `gorm:"column:user_id;not null;comment:用户id" json:"user_id"`     // 用户id
	ObjectID  int32     `gorm:"column:object_id;not null;comment:对象id" json:"object_id"` // 对象id
	Type      bool      `gorm:"column:type;not null;comment:收藏类型：1商品 2文章" json:"type"`   // 收藏类型：1商品 2文章
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName CollectionInfo's table name
func (*CollectionInfo) TableName() string {
	return TableNameCollectionInfo
}
