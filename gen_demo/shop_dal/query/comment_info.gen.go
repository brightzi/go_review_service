// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"gen_demo/dal/model"
)

func newCommentInfo(db *gorm.DB, opts ...gen.DOOption) commentInfo {
	_commentInfo := commentInfo{}

	_commentInfo.commentInfoDo.UseDB(db, opts...)
	_commentInfo.commentInfoDo.UseModel(&model.CommentInfo{})

	tableName := _commentInfo.commentInfoDo.TableName()
	_commentInfo.ALL = field.NewAsterisk(tableName)
	_commentInfo.ID = field.NewInt32(tableName, "id")
	_commentInfo.ParentID = field.NewInt32(tableName, "parent_id")
	_commentInfo.UserID = field.NewInt32(tableName, "user_id")
	_commentInfo.ObjectID = field.NewInt32(tableName, "object_id")
	_commentInfo.Type = field.NewBool(tableName, "type")
	_commentInfo.Content = field.NewString(tableName, "content")
	_commentInfo.CreatedAt = field.NewTime(tableName, "created_at")
	_commentInfo.UpdatedAt = field.NewTime(tableName, "updated_at")
	_commentInfo.DeletedAt = field.NewField(tableName, "deleted_at")

	_commentInfo.fillFieldMap()

	return _commentInfo
}

type commentInfo struct {
	commentInfoDo commentInfoDo

	ALL       field.Asterisk
	ID        field.Int32
	ParentID  field.Int32 // 父级评论id
	UserID    field.Int32
	ObjectID  field.Int32
	Type      field.Bool   // 评论类型：1商品 2文章
	Content   field.String // 评论内容
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field

	fieldMap map[string]field.Expr
}

func (c commentInfo) Table(newTableName string) *commentInfo {
	c.commentInfoDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c commentInfo) As(alias string) *commentInfo {
	c.commentInfoDo.DO = *(c.commentInfoDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *commentInfo) updateTableName(table string) *commentInfo {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewInt32(table, "id")
	c.ParentID = field.NewInt32(table, "parent_id")
	c.UserID = field.NewInt32(table, "user_id")
	c.ObjectID = field.NewInt32(table, "object_id")
	c.Type = field.NewBool(table, "type")
	c.Content = field.NewString(table, "content")
	c.CreatedAt = field.NewTime(table, "created_at")
	c.UpdatedAt = field.NewTime(table, "updated_at")
	c.DeletedAt = field.NewField(table, "deleted_at")

	c.fillFieldMap()

	return c
}

func (c *commentInfo) WithContext(ctx context.Context) ICommentInfoDo {
	return c.commentInfoDo.WithContext(ctx)
}

func (c commentInfo) TableName() string { return c.commentInfoDo.TableName() }

func (c commentInfo) Alias() string { return c.commentInfoDo.Alias() }

func (c commentInfo) Columns(cols ...field.Expr) gen.Columns { return c.commentInfoDo.Columns(cols...) }

func (c *commentInfo) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *commentInfo) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 9)
	c.fieldMap["id"] = c.ID
	c.fieldMap["parent_id"] = c.ParentID
	c.fieldMap["user_id"] = c.UserID
	c.fieldMap["object_id"] = c.ObjectID
	c.fieldMap["type"] = c.Type
	c.fieldMap["content"] = c.Content
	c.fieldMap["created_at"] = c.CreatedAt
	c.fieldMap["updated_at"] = c.UpdatedAt
	c.fieldMap["deleted_at"] = c.DeletedAt
}

func (c commentInfo) clone(db *gorm.DB) commentInfo {
	c.commentInfoDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c commentInfo) replaceDB(db *gorm.DB) commentInfo {
	c.commentInfoDo.ReplaceDB(db)
	return c
}

type commentInfoDo struct{ gen.DO }

type ICommentInfoDo interface {
	gen.SubQuery
	Debug() ICommentInfoDo
	WithContext(ctx context.Context) ICommentInfoDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ICommentInfoDo
	WriteDB() ICommentInfoDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ICommentInfoDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ICommentInfoDo
	Not(conds ...gen.Condition) ICommentInfoDo
	Or(conds ...gen.Condition) ICommentInfoDo
	Select(conds ...field.Expr) ICommentInfoDo
	Where(conds ...gen.Condition) ICommentInfoDo
	Order(conds ...field.Expr) ICommentInfoDo
	Distinct(cols ...field.Expr) ICommentInfoDo
	Omit(cols ...field.Expr) ICommentInfoDo
	Join(table schema.Tabler, on ...field.Expr) ICommentInfoDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ICommentInfoDo
	RightJoin(table schema.Tabler, on ...field.Expr) ICommentInfoDo
	Group(cols ...field.Expr) ICommentInfoDo
	Having(conds ...gen.Condition) ICommentInfoDo
	Limit(limit int) ICommentInfoDo
	Offset(offset int) ICommentInfoDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ICommentInfoDo
	Unscoped() ICommentInfoDo
	Create(values ...*model.CommentInfo) error
	CreateInBatches(values []*model.CommentInfo, batchSize int) error
	Save(values ...*model.CommentInfo) error
	First() (*model.CommentInfo, error)
	Take() (*model.CommentInfo, error)
	Last() (*model.CommentInfo, error)
	Find() ([]*model.CommentInfo, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CommentInfo, err error)
	FindInBatches(result *[]*model.CommentInfo, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.CommentInfo) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ICommentInfoDo
	Assign(attrs ...field.AssignExpr) ICommentInfoDo
	Joins(fields ...field.RelationField) ICommentInfoDo
	Preload(fields ...field.RelationField) ICommentInfoDo
	FirstOrInit() (*model.CommentInfo, error)
	FirstOrCreate() (*model.CommentInfo, error)
	FindByPage(offset int, limit int) (result []*model.CommentInfo, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ICommentInfoDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c commentInfoDo) Debug() ICommentInfoDo {
	return c.withDO(c.DO.Debug())
}

func (c commentInfoDo) WithContext(ctx context.Context) ICommentInfoDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c commentInfoDo) ReadDB() ICommentInfoDo {
	return c.Clauses(dbresolver.Read)
}

func (c commentInfoDo) WriteDB() ICommentInfoDo {
	return c.Clauses(dbresolver.Write)
}

func (c commentInfoDo) Session(config *gorm.Session) ICommentInfoDo {
	return c.withDO(c.DO.Session(config))
}

func (c commentInfoDo) Clauses(conds ...clause.Expression) ICommentInfoDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c commentInfoDo) Returning(value interface{}, columns ...string) ICommentInfoDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c commentInfoDo) Not(conds ...gen.Condition) ICommentInfoDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c commentInfoDo) Or(conds ...gen.Condition) ICommentInfoDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c commentInfoDo) Select(conds ...field.Expr) ICommentInfoDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c commentInfoDo) Where(conds ...gen.Condition) ICommentInfoDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c commentInfoDo) Order(conds ...field.Expr) ICommentInfoDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c commentInfoDo) Distinct(cols ...field.Expr) ICommentInfoDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c commentInfoDo) Omit(cols ...field.Expr) ICommentInfoDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c commentInfoDo) Join(table schema.Tabler, on ...field.Expr) ICommentInfoDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c commentInfoDo) LeftJoin(table schema.Tabler, on ...field.Expr) ICommentInfoDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c commentInfoDo) RightJoin(table schema.Tabler, on ...field.Expr) ICommentInfoDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c commentInfoDo) Group(cols ...field.Expr) ICommentInfoDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c commentInfoDo) Having(conds ...gen.Condition) ICommentInfoDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c commentInfoDo) Limit(limit int) ICommentInfoDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c commentInfoDo) Offset(offset int) ICommentInfoDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c commentInfoDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ICommentInfoDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c commentInfoDo) Unscoped() ICommentInfoDo {
	return c.withDO(c.DO.Unscoped())
}

func (c commentInfoDo) Create(values ...*model.CommentInfo) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c commentInfoDo) CreateInBatches(values []*model.CommentInfo, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c commentInfoDo) Save(values ...*model.CommentInfo) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c commentInfoDo) First() (*model.CommentInfo, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.CommentInfo), nil
	}
}

func (c commentInfoDo) Take() (*model.CommentInfo, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.CommentInfo), nil
	}
}

func (c commentInfoDo) Last() (*model.CommentInfo, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.CommentInfo), nil
	}
}

func (c commentInfoDo) Find() ([]*model.CommentInfo, error) {
	result, err := c.DO.Find()
	return result.([]*model.CommentInfo), err
}

func (c commentInfoDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CommentInfo, err error) {
	buf := make([]*model.CommentInfo, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c commentInfoDo) FindInBatches(result *[]*model.CommentInfo, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c commentInfoDo) Attrs(attrs ...field.AssignExpr) ICommentInfoDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c commentInfoDo) Assign(attrs ...field.AssignExpr) ICommentInfoDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c commentInfoDo) Joins(fields ...field.RelationField) ICommentInfoDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c commentInfoDo) Preload(fields ...field.RelationField) ICommentInfoDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c commentInfoDo) FirstOrInit() (*model.CommentInfo, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.CommentInfo), nil
	}
}

func (c commentInfoDo) FirstOrCreate() (*model.CommentInfo, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.CommentInfo), nil
	}
}

func (c commentInfoDo) FindByPage(offset int, limit int) (result []*model.CommentInfo, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c commentInfoDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c commentInfoDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c commentInfoDo) Delete(models ...*model.CommentInfo) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *commentInfoDo) withDO(do gen.Dao) *commentInfoDo {
	c.DO = *do.(*gen.DO)
	return c
}