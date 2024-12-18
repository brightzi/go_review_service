// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNamePraiseInfo = "praise_info"

// PraiseInfo mapped from table <praise_info>
type PraiseInfo struct {
	ID        int32     `gorm:"column:id;primaryKey;autoIncrement:true;comment:点赞表" json:"id"` // 点赞表
	UserID    int32     `gorm:"column:user_id;not null" json:"user_id"`
	Type      bool      `gorm:"column:type;not null;comment:点赞类型 1商品 2文章" json:"type"`            // 点赞类型 1商品 2文章
	ObjectID  int32     `gorm:"column:object_id;not null;comment:点赞对象id 方便后期扩展" json:"object_id"` // 点赞对象id 方便后期扩展
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName PraiseInfo's table name
func (*PraiseInfo) TableName() string {
	return TableNamePraiseInfo
}
