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

func newCategoryInfo(db *gorm.DB, opts ...gen.DOOption) categoryInfo {
	_categoryInfo := categoryInfo{}

	_categoryInfo.categoryInfoDo.UseDB(db, opts...)
	_categoryInfo.categoryInfoDo.UseModel(&model.CategoryInfo{})

	tableName := _categoryInfo.categoryInfoDo.TableName()
	_categoryInfo.ALL = field.NewAsterisk(tableName)
	_categoryInfo.ID = field.NewInt32(tableName, "id")
	_categoryInfo.ParentID = field.NewInt32(tableName, "parent_id")
	_categoryInfo.Name = field.NewString(tableName, "name")
	_categoryInfo.PicURL = field.NewString(tableName, "pic_url")
	_categoryInfo.Level = field.NewBool(tableName, "level")
	_categoryInfo.Sort = field.NewBool(tableName, "sort")
	_categoryInfo.CreatedAt = field.NewTime(tableName, "created_at")
	_categoryInfo.UpdatedAt = field.NewTime(tableName, "updated_at")
	_categoryInfo.DeletedAt = field.NewField(tableName, "deleted_at")

	_categoryInfo.fillFieldMap()

	return _categoryInfo
}

// categoryInfo 轮播图表

type categoryInfo struct {
	categoryInfoDo categoryInfoDo

	ALL       field.Asterisk
	ID        field.Int32
	ParentID  field.Int32 // 父级id
	Name      field.String
	PicURL    field.String // icon
	Level     field.Bool   // 等级 默认1级分类
	Sort      field.Bool
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field

	fieldMap map[string]field.Expr
}

func (c categoryInfo) Table(newTableName string) *categoryInfo {
	c.categoryInfoDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c categoryInfo) As(alias string) *categoryInfo {
	c.categoryInfoDo.DO = *(c.categoryInfoDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *categoryInfo) updateTableName(table string) *categoryInfo {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewInt32(table, "id")
	c.ParentID = field.NewInt32(table, "parent_id")
	c.Name = field.NewString(table, "name")
	c.PicURL = field.NewString(table, "pic_url")
	c.Level = field.NewBool(table, "level")
	c.Sort = field.NewBool(table, "sort")
	c.CreatedAt = field.NewTime(table, "created_at")
	c.UpdatedAt = field.NewTime(table, "updated_at")
	c.DeletedAt = field.NewField(table, "deleted_at")

	c.fillFieldMap()

	return c
}

func (c *categoryInfo) WithContext(ctx context.Context) ICategoryInfoDo {
	return c.categoryInfoDo.WithContext(ctx)
}

func (c categoryInfo) TableName() string { return c.categoryInfoDo.TableName() }

func (c categoryInfo) Alias() string { return c.categoryInfoDo.Alias() }

func (c categoryInfo) Columns(cols ...field.Expr) gen.Columns {
	return c.categoryInfoDo.Columns(cols...)
}

func (c *categoryInfo) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *categoryInfo) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 9)
	c.fieldMap["id"] = c.ID
	c.fieldMap["parent_id"] = c.ParentID
	c.fieldMap["name"] = c.Name
	c.fieldMap["pic_url"] = c.PicURL
	c.fieldMap["level"] = c.Level
	c.fieldMap["sort"] = c.Sort
	c.fieldMap["created_at"] = c.CreatedAt
	c.fieldMap["updated_at"] = c.UpdatedAt
	c.fieldMap["deleted_at"] = c.DeletedAt
}

func (c categoryInfo) clone(db *gorm.DB) categoryInfo {
	c.categoryInfoDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c categoryInfo) replaceDB(db *gorm.DB) categoryInfo {
	c.categoryInfoDo.ReplaceDB(db)
	return c
}

type categoryInfoDo struct{ gen.DO }

type ICategoryInfoDo interface {
	gen.SubQuery
	Debug() ICategoryInfoDo
	WithContext(ctx context.Context) ICategoryInfoDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ICategoryInfoDo
	WriteDB() ICategoryInfoDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ICategoryInfoDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ICategoryInfoDo
	Not(conds ...gen.Condition) ICategoryInfoDo
	Or(conds ...gen.Condition) ICategoryInfoDo
	Select(conds ...field.Expr) ICategoryInfoDo
	Where(conds ...gen.Condition) ICategoryInfoDo
	Order(conds ...field.Expr) ICategoryInfoDo
	Distinct(cols ...field.Expr) ICategoryInfoDo
	Omit(cols ...field.Expr) ICategoryInfoDo
	Join(table schema.Tabler, on ...field.Expr) ICategoryInfoDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ICategoryInfoDo
	RightJoin(table schema.Tabler, on ...field.Expr) ICategoryInfoDo
	Group(cols ...field.Expr) ICategoryInfoDo
	Having(conds ...gen.Condition) ICategoryInfoDo
	Limit(limit int) ICategoryInfoDo
	Offset(offset int) ICategoryInfoDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ICategoryInfoDo
	Unscoped() ICategoryInfoDo
	Create(values ...*model.CategoryInfo) error
	CreateInBatches(values []*model.CategoryInfo, batchSize int) error
	Save(values ...*model.CategoryInfo) error
	First() (*model.CategoryInfo, error)
	Take() (*model.CategoryInfo, error)
	Last() (*model.CategoryInfo, error)
	Find() ([]*model.CategoryInfo, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CategoryInfo, err error)
	FindInBatches(result *[]*model.CategoryInfo, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.CategoryInfo) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ICategoryInfoDo
	Assign(attrs ...field.AssignExpr) ICategoryInfoDo
	Joins(fields ...field.RelationField) ICategoryInfoDo
	Preload(fields ...field.RelationField) ICategoryInfoDo
	FirstOrInit() (*model.CategoryInfo, error)
	FirstOrCreate() (*model.CategoryInfo, error)
	FindByPage(offset int, limit int) (result []*model.CategoryInfo, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ICategoryInfoDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c categoryInfoDo) Debug() ICategoryInfoDo {
	return c.withDO(c.DO.Debug())
}

func (c categoryInfoDo) WithContext(ctx context.Context) ICategoryInfoDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c categoryInfoDo) ReadDB() ICategoryInfoDo {
	return c.Clauses(dbresolver.Read)
}

func (c categoryInfoDo) WriteDB() ICategoryInfoDo {
	return c.Clauses(dbresolver.Write)
}

func (c categoryInfoDo) Session(config *gorm.Session) ICategoryInfoDo {
	return c.withDO(c.DO.Session(config))
}

func (c categoryInfoDo) Clauses(conds ...clause.Expression) ICategoryInfoDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c categoryInfoDo) Returning(value interface{}, columns ...string) ICategoryInfoDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c categoryInfoDo) Not(conds ...gen.Condition) ICategoryInfoDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c categoryInfoDo) Or(conds ...gen.Condition) ICategoryInfoDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c categoryInfoDo) Select(conds ...field.Expr) ICategoryInfoDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c categoryInfoDo) Where(conds ...gen.Condition) ICategoryInfoDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c categoryInfoDo) Order(conds ...field.Expr) ICategoryInfoDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c categoryInfoDo) Distinct(cols ...field.Expr) ICategoryInfoDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c categoryInfoDo) Omit(cols ...field.Expr) ICategoryInfoDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c categoryInfoDo) Join(table schema.Tabler, on ...field.Expr) ICategoryInfoDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c categoryInfoDo) LeftJoin(table schema.Tabler, on ...field.Expr) ICategoryInfoDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c categoryInfoDo) RightJoin(table schema.Tabler, on ...field.Expr) ICategoryInfoDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c categoryInfoDo) Group(cols ...field.Expr) ICategoryInfoDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c categoryInfoDo) Having(conds ...gen.Condition) ICategoryInfoDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c categoryInfoDo) Limit(limit int) ICategoryInfoDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c categoryInfoDo) Offset(offset int) ICategoryInfoDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c categoryInfoDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ICategoryInfoDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c categoryInfoDo) Unscoped() ICategoryInfoDo {
	return c.withDO(c.DO.Unscoped())
}

func (c categoryInfoDo) Create(values ...*model.CategoryInfo) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c categoryInfoDo) CreateInBatches(values []*model.CategoryInfo, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c categoryInfoDo) Save(values ...*model.CategoryInfo) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c categoryInfoDo) First() (*model.CategoryInfo, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.CategoryInfo), nil
	}
}

func (c categoryInfoDo) Take() (*model.CategoryInfo, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.CategoryInfo), nil
	}
}

func (c categoryInfoDo) Last() (*model.CategoryInfo, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.CategoryInfo), nil
	}
}

func (c categoryInfoDo) Find() ([]*model.CategoryInfo, error) {
	result, err := c.DO.Find()
	return result.([]*model.CategoryInfo), err
}

func (c categoryInfoDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CategoryInfo, err error) {
	buf := make([]*model.CategoryInfo, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c categoryInfoDo) FindInBatches(result *[]*model.CategoryInfo, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c categoryInfoDo) Attrs(attrs ...field.AssignExpr) ICategoryInfoDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c categoryInfoDo) Assign(attrs ...field.AssignExpr) ICategoryInfoDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c categoryInfoDo) Joins(fields ...field.RelationField) ICategoryInfoDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c categoryInfoDo) Preload(fields ...field.RelationField) ICategoryInfoDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c categoryInfoDo) FirstOrInit() (*model.CategoryInfo, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.CategoryInfo), nil
	}
}

func (c categoryInfoDo) FirstOrCreate() (*model.CategoryInfo, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.CategoryInfo), nil
	}
}

func (c categoryInfoDo) FindByPage(offset int, limit int) (result []*model.CategoryInfo, count int64, err error) {
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

func (c categoryInfoDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c categoryInfoDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c categoryInfoDo) Delete(models ...*model.CategoryInfo) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *categoryInfoDo) withDO(do gen.Dao) *categoryInfoDo {
	c.DO = *do.(*gen.DO)
	return c
}
