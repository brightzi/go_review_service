// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameArticleInfo = "article_info"

// ArticleInfo 文章（种草）表
type ArticleInfo struct {
	ID        int32          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UserID    int32          `gorm:"column:user_id;not null;comment:作者id" json:"user_id"`                         // 作者id
	Title     string         `gorm:"column:title;not null;comment:标题" json:"title"`                               // 标题
	Desc      string         `gorm:"column:desc;not null;comment:摘要" json:"desc"`                                 // 摘要
	PicURL    string         `gorm:"column:pic_url;not null;comment:封面图" json:"pic_url"`                          // 封面图
	IsAdmin   bool           `gorm:"column:is_admin;not null;default:2;comment:1后台管理员发布 2前台用户发布" json:"is_admin"` // 1后台管理员发布 2前台用户发布
	Praise    int32          `gorm:"column:praise;not null;comment:点赞数" json:"praise"`                            // 点赞数
	Detail    string         `gorm:"column:detail;comment:文章详情" json:"detail"`                                    // 文章详情
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName ArticleInfo's table name
func (*ArticleInfo) TableName() string {
	return TableNameArticleInfo
}
