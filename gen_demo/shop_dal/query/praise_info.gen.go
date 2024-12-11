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

func newPraiseInfo(db *gorm.DB, opts ...gen.DOOption) praiseInfo {
	_praiseInfo := praiseInfo{}

	_praiseInfo.praiseInfoDo.UseDB(db, opts...)
	_praiseInfo.praiseInfoDo.UseModel(&model.PraiseInfo{})

	tableName := _praiseInfo.praiseInfoDo.TableName()
	_praiseInfo.ALL = field.NewAsterisk(tableName)
	_praiseInfo.ID = field.NewInt32(tableName, "id")
	_praiseInfo.UserID = field.NewInt32(tableName, "user_id")
	_praiseInfo.Type = field.NewBool(tableName, "type")
	_praiseInfo.ObjectID = field.NewInt32(tableName, "object_id")
	_praiseInfo.CreatedAt = field.NewTime(tableName, "created_at")
	_praiseInfo.UpdatedAt = field.NewTime(tableName, "updated_at")

	_praiseInfo.fillFieldMap()

	return _praiseInfo
}

type praiseInfo struct {
	praiseInfoDo praiseInfoDo

	ALL       field.Asterisk
	ID        field.Int32 // 点赞表
	UserID    field.Int32
	Type      field.Bool  // 点赞类型 1商品 2文章
	ObjectID  field.Int32 // 点赞对象id 方便后期扩展
	CreatedAt field.Time
	UpdatedAt field.Time

	fieldMap map[string]field.Expr
}

func (p praiseInfo) Table(newTableName string) *praiseInfo {
	p.praiseInfoDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p praiseInfo) As(alias string) *praiseInfo {
	p.praiseInfoDo.DO = *(p.praiseInfoDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *praiseInfo) updateTableName(table string) *praiseInfo {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewInt32(table, "id")
	p.UserID = field.NewInt32(table, "user_id")
	p.Type = field.NewBool(table, "type")
	p.ObjectID = field.NewInt32(table, "object_id")
	p.CreatedAt = field.NewTime(table, "created_at")
	p.UpdatedAt = field.NewTime(table, "updated_at")

	p.fillFieldMap()

	return p
}

func (p *praiseInfo) WithContext(ctx context.Context) IPraiseInfoDo {
	return p.praiseInfoDo.WithContext(ctx)
}

func (p praiseInfo) TableName() string { return p.praiseInfoDo.TableName() }

func (p praiseInfo) Alias() string { return p.praiseInfoDo.Alias() }

func (p praiseInfo) Columns(cols ...field.Expr) gen.Columns { return p.praiseInfoDo.Columns(cols...) }

func (p *praiseInfo) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *praiseInfo) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 6)
	p.fieldMap["id"] = p.ID
	p.fieldMap["user_id"] = p.UserID
	p.fieldMap["type"] = p.Type
	p.fieldMap["object_id"] = p.ObjectID
	p.fieldMap["created_at"] = p.CreatedAt
	p.fieldMap["updated_at"] = p.UpdatedAt
}

func (p praiseInfo) clone(db *gorm.DB) praiseInfo {
	p.praiseInfoDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p praiseInfo) replaceDB(db *gorm.DB) praiseInfo {
	p.praiseInfoDo.ReplaceDB(db)
	return p
}

type praiseInfoDo struct{ gen.DO }

type IPraiseInfoDo interface {
	gen.SubQuery
	Debug() IPraiseInfoDo
	WithContext(ctx context.Context) IPraiseInfoDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPraiseInfoDo
	WriteDB() IPraiseInfoDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPraiseInfoDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPraiseInfoDo
	Not(conds ...gen.Condition) IPraiseInfoDo
	Or(conds ...gen.Condition) IPraiseInfoDo
	Select(conds ...field.Expr) IPraiseInfoDo
	Where(conds ...gen.Condition) IPraiseInfoDo
	Order(conds ...field.Expr) IPraiseInfoDo
	Distinct(cols ...field.Expr) IPraiseInfoDo
	Omit(cols ...field.Expr) IPraiseInfoDo
	Join(table schema.Tabler, on ...field.Expr) IPraiseInfoDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPraiseInfoDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPraiseInfoDo
	Group(cols ...field.Expr) IPraiseInfoDo
	Having(conds ...gen.Condition) IPraiseInfoDo
	Limit(limit int) IPraiseInfoDo
	Offset(offset int) IPraiseInfoDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPraiseInfoDo
	Unscoped() IPraiseInfoDo
	Create(values ...*model.PraiseInfo) error
	CreateInBatches(values []*model.PraiseInfo, batchSize int) error
	Save(values ...*model.PraiseInfo) error
	First() (*model.PraiseInfo, error)
	Take() (*model.PraiseInfo, error)
	Last() (*model.PraiseInfo, error)
	Find() ([]*model.PraiseInfo, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PraiseInfo, err error)
	FindInBatches(result *[]*model.PraiseInfo, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.PraiseInfo) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPraiseInfoDo
	Assign(attrs ...field.AssignExpr) IPraiseInfoDo
	Joins(fields ...field.RelationField) IPraiseInfoDo
	Preload(fields ...field.RelationField) IPraiseInfoDo
	FirstOrInit() (*model.PraiseInfo, error)
	FirstOrCreate() (*model.PraiseInfo, error)
	FindByPage(offset int, limit int) (result []*model.PraiseInfo, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPraiseInfoDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p praiseInfoDo) Debug() IPraiseInfoDo {
	return p.withDO(p.DO.Debug())
}

func (p praiseInfoDo) WithContext(ctx context.Context) IPraiseInfoDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p praiseInfoDo) ReadDB() IPraiseInfoDo {
	return p.Clauses(dbresolver.Read)
}

func (p praiseInfoDo) WriteDB() IPraiseInfoDo {
	return p.Clauses(dbresolver.Write)
}

func (p praiseInfoDo) Session(config *gorm.Session) IPraiseInfoDo {
	return p.withDO(p.DO.Session(config))
}

func (p praiseInfoDo) Clauses(conds ...clause.Expression) IPraiseInfoDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p praiseInfoDo) Returning(value interface{}, columns ...string) IPraiseInfoDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p praiseInfoDo) Not(conds ...gen.Condition) IPraiseInfoDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p praiseInfoDo) Or(conds ...gen.Condition) IPraiseInfoDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p praiseInfoDo) Select(conds ...field.Expr) IPraiseInfoDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p praiseInfoDo) Where(conds ...gen.Condition) IPraiseInfoDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p praiseInfoDo) Order(conds ...field.Expr) IPraiseInfoDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p praiseInfoDo) Distinct(cols ...field.Expr) IPraiseInfoDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p praiseInfoDo) Omit(cols ...field.Expr) IPraiseInfoDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p praiseInfoDo) Join(table schema.Tabler, on ...field.Expr) IPraiseInfoDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p praiseInfoDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPraiseInfoDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p praiseInfoDo) RightJoin(table schema.Tabler, on ...field.Expr) IPraiseInfoDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p praiseInfoDo) Group(cols ...field.Expr) IPraiseInfoDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p praiseInfoDo) Having(conds ...gen.Condition) IPraiseInfoDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p praiseInfoDo) Limit(limit int) IPraiseInfoDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p praiseInfoDo) Offset(offset int) IPraiseInfoDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p praiseInfoDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPraiseInfoDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p praiseInfoDo) Unscoped() IPraiseInfoDo {
	return p.withDO(p.DO.Unscoped())
}

func (p praiseInfoDo) Create(values ...*model.PraiseInfo) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p praiseInfoDo) CreateInBatches(values []*model.PraiseInfo, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p praiseInfoDo) Save(values ...*model.PraiseInfo) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p praiseInfoDo) First() (*model.PraiseInfo, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.PraiseInfo), nil
	}
}

func (p praiseInfoDo) Take() (*model.PraiseInfo, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.PraiseInfo), nil
	}
}

func (p praiseInfoDo) Last() (*model.PraiseInfo, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.PraiseInfo), nil
	}
}

func (p praiseInfoDo) Find() ([]*model.PraiseInfo, error) {
	result, err := p.DO.Find()
	return result.([]*model.PraiseInfo), err
}

func (p praiseInfoDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PraiseInfo, err error) {
	buf := make([]*model.PraiseInfo, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p praiseInfoDo) FindInBatches(result *[]*model.PraiseInfo, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p praiseInfoDo) Attrs(attrs ...field.AssignExpr) IPraiseInfoDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p praiseInfoDo) Assign(attrs ...field.AssignExpr) IPraiseInfoDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p praiseInfoDo) Joins(fields ...field.RelationField) IPraiseInfoDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p praiseInfoDo) Preload(fields ...field.RelationField) IPraiseInfoDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p praiseInfoDo) FirstOrInit() (*model.PraiseInfo, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.PraiseInfo), nil
	}
}

func (p praiseInfoDo) FirstOrCreate() (*model.PraiseInfo, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.PraiseInfo), nil
	}
}

func (p praiseInfoDo) FindByPage(offset int, limit int) (result []*model.PraiseInfo, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p praiseInfoDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p praiseInfoDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p praiseInfoDo) Delete(models ...*model.PraiseInfo) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *praiseInfoDo) withDO(do gen.Dao) *praiseInfoDo {
	p.DO = *do.(*gen.DO)
	return p
}