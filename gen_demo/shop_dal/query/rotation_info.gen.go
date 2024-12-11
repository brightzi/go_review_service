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

func newRotationInfo(db *gorm.DB, opts ...gen.DOOption) rotationInfo {
	_rotationInfo := rotationInfo{}

	_rotationInfo.rotationInfoDo.UseDB(db, opts...)
	_rotationInfo.rotationInfoDo.UseModel(&model.RotationInfo{})

	tableName := _rotationInfo.rotationInfoDo.TableName()
	_rotationInfo.ALL = field.NewAsterisk(tableName)
	_rotationInfo.ID = field.NewInt32(tableName, "id")
	_rotationInfo.PicURL = field.NewString(tableName, "pic_url")
	_rotationInfo.Link = field.NewString(tableName, "link")
	_rotationInfo.Sort = field.NewBool(tableName, "sort")
	_rotationInfo.CreatedAt = field.NewTime(tableName, "created_at")
	_rotationInfo.UpdatedAt = field.NewTime(tableName, "updated_at")
	_rotationInfo.DeletedAt = field.NewField(tableName, "deleted_at")

	_rotationInfo.fillFieldMap()

	return _rotationInfo
}

// rotationInfo 轮播图表

type rotationInfo struct {
	rotationInfoDo rotationInfoDo

	ALL       field.Asterisk
	ID        field.Int32
	PicURL    field.String // 轮播图片
	Link      field.String // 跳转链接
	Sort      field.Bool   // 排序字段
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field

	fieldMap map[string]field.Expr
}

func (r rotationInfo) Table(newTableName string) *rotationInfo {
	r.rotationInfoDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r rotationInfo) As(alias string) *rotationInfo {
	r.rotationInfoDo.DO = *(r.rotationInfoDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *rotationInfo) updateTableName(table string) *rotationInfo {
	r.ALL = field.NewAsterisk(table)
	r.ID = field.NewInt32(table, "id")
	r.PicURL = field.NewString(table, "pic_url")
	r.Link = field.NewString(table, "link")
	r.Sort = field.NewBool(table, "sort")
	r.CreatedAt = field.NewTime(table, "created_at")
	r.UpdatedAt = field.NewTime(table, "updated_at")
	r.DeletedAt = field.NewField(table, "deleted_at")

	r.fillFieldMap()

	return r
}

func (r *rotationInfo) WithContext(ctx context.Context) IRotationInfoDo {
	return r.rotationInfoDo.WithContext(ctx)
}

func (r rotationInfo) TableName() string { return r.rotationInfoDo.TableName() }

func (r rotationInfo) Alias() string { return r.rotationInfoDo.Alias() }

func (r rotationInfo) Columns(cols ...field.Expr) gen.Columns {
	return r.rotationInfoDo.Columns(cols...)
}

func (r *rotationInfo) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *rotationInfo) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 7)
	r.fieldMap["id"] = r.ID
	r.fieldMap["pic_url"] = r.PicURL
	r.fieldMap["link"] = r.Link
	r.fieldMap["sort"] = r.Sort
	r.fieldMap["created_at"] = r.CreatedAt
	r.fieldMap["updated_at"] = r.UpdatedAt
	r.fieldMap["deleted_at"] = r.DeletedAt
}

func (r rotationInfo) clone(db *gorm.DB) rotationInfo {
	r.rotationInfoDo.ReplaceConnPool(db.Statement.ConnPool)
	return r
}

func (r rotationInfo) replaceDB(db *gorm.DB) rotationInfo {
	r.rotationInfoDo.ReplaceDB(db)
	return r
}

type rotationInfoDo struct{ gen.DO }

type IRotationInfoDo interface {
	gen.SubQuery
	Debug() IRotationInfoDo
	WithContext(ctx context.Context) IRotationInfoDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IRotationInfoDo
	WriteDB() IRotationInfoDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IRotationInfoDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IRotationInfoDo
	Not(conds ...gen.Condition) IRotationInfoDo
	Or(conds ...gen.Condition) IRotationInfoDo
	Select(conds ...field.Expr) IRotationInfoDo
	Where(conds ...gen.Condition) IRotationInfoDo
	Order(conds ...field.Expr) IRotationInfoDo
	Distinct(cols ...field.Expr) IRotationInfoDo
	Omit(cols ...field.Expr) IRotationInfoDo
	Join(table schema.Tabler, on ...field.Expr) IRotationInfoDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IRotationInfoDo
	RightJoin(table schema.Tabler, on ...field.Expr) IRotationInfoDo
	Group(cols ...field.Expr) IRotationInfoDo
	Having(conds ...gen.Condition) IRotationInfoDo
	Limit(limit int) IRotationInfoDo
	Offset(offset int) IRotationInfoDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IRotationInfoDo
	Unscoped() IRotationInfoDo
	Create(values ...*model.RotationInfo) error
	CreateInBatches(values []*model.RotationInfo, batchSize int) error
	Save(values ...*model.RotationInfo) error
	First() (*model.RotationInfo, error)
	Take() (*model.RotationInfo, error)
	Last() (*model.RotationInfo, error)
	Find() ([]*model.RotationInfo, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.RotationInfo, err error)
	FindInBatches(result *[]*model.RotationInfo, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.RotationInfo) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IRotationInfoDo
	Assign(attrs ...field.AssignExpr) IRotationInfoDo
	Joins(fields ...field.RelationField) IRotationInfoDo
	Preload(fields ...field.RelationField) IRotationInfoDo
	FirstOrInit() (*model.RotationInfo, error)
	FirstOrCreate() (*model.RotationInfo, error)
	FindByPage(offset int, limit int) (result []*model.RotationInfo, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IRotationInfoDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (r rotationInfoDo) Debug() IRotationInfoDo {
	return r.withDO(r.DO.Debug())
}

func (r rotationInfoDo) WithContext(ctx context.Context) IRotationInfoDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r rotationInfoDo) ReadDB() IRotationInfoDo {
	return r.Clauses(dbresolver.Read)
}

func (r rotationInfoDo) WriteDB() IRotationInfoDo {
	return r.Clauses(dbresolver.Write)
}

func (r rotationInfoDo) Session(config *gorm.Session) IRotationInfoDo {
	return r.withDO(r.DO.Session(config))
}

func (r rotationInfoDo) Clauses(conds ...clause.Expression) IRotationInfoDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r rotationInfoDo) Returning(value interface{}, columns ...string) IRotationInfoDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r rotationInfoDo) Not(conds ...gen.Condition) IRotationInfoDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r rotationInfoDo) Or(conds ...gen.Condition) IRotationInfoDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r rotationInfoDo) Select(conds ...field.Expr) IRotationInfoDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r rotationInfoDo) Where(conds ...gen.Condition) IRotationInfoDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r rotationInfoDo) Order(conds ...field.Expr) IRotationInfoDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r rotationInfoDo) Distinct(cols ...field.Expr) IRotationInfoDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r rotationInfoDo) Omit(cols ...field.Expr) IRotationInfoDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r rotationInfoDo) Join(table schema.Tabler, on ...field.Expr) IRotationInfoDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r rotationInfoDo) LeftJoin(table schema.Tabler, on ...field.Expr) IRotationInfoDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r rotationInfoDo) RightJoin(table schema.Tabler, on ...field.Expr) IRotationInfoDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r rotationInfoDo) Group(cols ...field.Expr) IRotationInfoDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r rotationInfoDo) Having(conds ...gen.Condition) IRotationInfoDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r rotationInfoDo) Limit(limit int) IRotationInfoDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r rotationInfoDo) Offset(offset int) IRotationInfoDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r rotationInfoDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IRotationInfoDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r rotationInfoDo) Unscoped() IRotationInfoDo {
	return r.withDO(r.DO.Unscoped())
}

func (r rotationInfoDo) Create(values ...*model.RotationInfo) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r rotationInfoDo) CreateInBatches(values []*model.RotationInfo, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r rotationInfoDo) Save(values ...*model.RotationInfo) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r rotationInfoDo) First() (*model.RotationInfo, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.RotationInfo), nil
	}
}

func (r rotationInfoDo) Take() (*model.RotationInfo, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.RotationInfo), nil
	}
}

func (r rotationInfoDo) Last() (*model.RotationInfo, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.RotationInfo), nil
	}
}

func (r rotationInfoDo) Find() ([]*model.RotationInfo, error) {
	result, err := r.DO.Find()
	return result.([]*model.RotationInfo), err
}

func (r rotationInfoDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.RotationInfo, err error) {
	buf := make([]*model.RotationInfo, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r rotationInfoDo) FindInBatches(result *[]*model.RotationInfo, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r rotationInfoDo) Attrs(attrs ...field.AssignExpr) IRotationInfoDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r rotationInfoDo) Assign(attrs ...field.AssignExpr) IRotationInfoDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r rotationInfoDo) Joins(fields ...field.RelationField) IRotationInfoDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r rotationInfoDo) Preload(fields ...field.RelationField) IRotationInfoDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r rotationInfoDo) FirstOrInit() (*model.RotationInfo, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.RotationInfo), nil
	}
}

func (r rotationInfoDo) FirstOrCreate() (*model.RotationInfo, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.RotationInfo), nil
	}
}

func (r rotationInfoDo) FindByPage(offset int, limit int) (result []*model.RotationInfo, count int64, err error) {
	result, err = r.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = r.Offset(-1).Limit(-1).Count()
	return
}

func (r rotationInfoDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r rotationInfoDo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r rotationInfoDo) Delete(models ...*model.RotationInfo) (result gen.ResultInfo, err error) {
	return r.DO.Delete(models)
}

func (r *rotationInfoDo) withDO(do gen.Dao) *rotationInfoDo {
	r.DO = *do.(*gen.DO)
	return r
}
