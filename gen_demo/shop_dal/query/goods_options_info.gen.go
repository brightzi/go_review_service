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

func newGoodsOptionsInfo(db *gorm.DB, opts ...gen.DOOption) goodsOptionsInfo {
	_goodsOptionsInfo := goodsOptionsInfo{}

	_goodsOptionsInfo.goodsOptionsInfoDo.UseDB(db, opts...)
	_goodsOptionsInfo.goodsOptionsInfoDo.UseModel(&model.GoodsOptionsInfo{})

	tableName := _goodsOptionsInfo.goodsOptionsInfoDo.TableName()
	_goodsOptionsInfo.ALL = field.NewAsterisk(tableName)
	_goodsOptionsInfo.ID = field.NewInt32(tableName, "id")
	_goodsOptionsInfo.GoodsID = field.NewInt32(tableName, "goods_id")
	_goodsOptionsInfo.PicURL = field.NewString(tableName, "pic_url")
	_goodsOptionsInfo.Name = field.NewString(tableName, "name")
	_goodsOptionsInfo.Price = field.NewInt32(tableName, "price")
	_goodsOptionsInfo.Stock = field.NewInt32(tableName, "stock")
	_goodsOptionsInfo.CreatedAt = field.NewTime(tableName, "created_at")
	_goodsOptionsInfo.UpdatedAt = field.NewTime(tableName, "updated_at")
	_goodsOptionsInfo.DeletedAt = field.NewField(tableName, "deleted_at")

	_goodsOptionsInfo.fillFieldMap()

	return _goodsOptionsInfo
}

// goodsOptionsInfo 商品规格表

type goodsOptionsInfo struct {
	goodsOptionsInfoDo goodsOptionsInfoDo

	ALL       field.Asterisk
	ID        field.Int32
	GoodsID   field.Int32  // 商品id
	PicURL    field.String // 图片
	Name      field.String // 商品名称
	Price     field.Int32  // 价格 单位分
	Stock     field.Int32  // 库存
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field

	fieldMap map[string]field.Expr
}

func (g goodsOptionsInfo) Table(newTableName string) *goodsOptionsInfo {
	g.goodsOptionsInfoDo.UseTable(newTableName)
	return g.updateTableName(newTableName)
}

func (g goodsOptionsInfo) As(alias string) *goodsOptionsInfo {
	g.goodsOptionsInfoDo.DO = *(g.goodsOptionsInfoDo.As(alias).(*gen.DO))
	return g.updateTableName(alias)
}

func (g *goodsOptionsInfo) updateTableName(table string) *goodsOptionsInfo {
	g.ALL = field.NewAsterisk(table)
	g.ID = field.NewInt32(table, "id")
	g.GoodsID = field.NewInt32(table, "goods_id")
	g.PicURL = field.NewString(table, "pic_url")
	g.Name = field.NewString(table, "name")
	g.Price = field.NewInt32(table, "price")
	g.Stock = field.NewInt32(table, "stock")
	g.CreatedAt = field.NewTime(table, "created_at")
	g.UpdatedAt = field.NewTime(table, "updated_at")
	g.DeletedAt = field.NewField(table, "deleted_at")

	g.fillFieldMap()

	return g
}

func (g *goodsOptionsInfo) WithContext(ctx context.Context) IGoodsOptionsInfoDo {
	return g.goodsOptionsInfoDo.WithContext(ctx)
}

func (g goodsOptionsInfo) TableName() string { return g.goodsOptionsInfoDo.TableName() }

func (g goodsOptionsInfo) Alias() string { return g.goodsOptionsInfoDo.Alias() }

func (g goodsOptionsInfo) Columns(cols ...field.Expr) gen.Columns {
	return g.goodsOptionsInfoDo.Columns(cols...)
}

func (g *goodsOptionsInfo) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := g.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (g *goodsOptionsInfo) fillFieldMap() {
	g.fieldMap = make(map[string]field.Expr, 9)
	g.fieldMap["id"] = g.ID
	g.fieldMap["goods_id"] = g.GoodsID
	g.fieldMap["pic_url"] = g.PicURL
	g.fieldMap["name"] = g.Name
	g.fieldMap["price"] = g.Price
	g.fieldMap["stock"] = g.Stock
	g.fieldMap["created_at"] = g.CreatedAt
	g.fieldMap["updated_at"] = g.UpdatedAt
	g.fieldMap["deleted_at"] = g.DeletedAt
}

func (g goodsOptionsInfo) clone(db *gorm.DB) goodsOptionsInfo {
	g.goodsOptionsInfoDo.ReplaceConnPool(db.Statement.ConnPool)
	return g
}

func (g goodsOptionsInfo) replaceDB(db *gorm.DB) goodsOptionsInfo {
	g.goodsOptionsInfoDo.ReplaceDB(db)
	return g
}

type goodsOptionsInfoDo struct{ gen.DO }

type IGoodsOptionsInfoDo interface {
	gen.SubQuery
	Debug() IGoodsOptionsInfoDo
	WithContext(ctx context.Context) IGoodsOptionsInfoDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IGoodsOptionsInfoDo
	WriteDB() IGoodsOptionsInfoDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IGoodsOptionsInfoDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IGoodsOptionsInfoDo
	Not(conds ...gen.Condition) IGoodsOptionsInfoDo
	Or(conds ...gen.Condition) IGoodsOptionsInfoDo
	Select(conds ...field.Expr) IGoodsOptionsInfoDo
	Where(conds ...gen.Condition) IGoodsOptionsInfoDo
	Order(conds ...field.Expr) IGoodsOptionsInfoDo
	Distinct(cols ...field.Expr) IGoodsOptionsInfoDo
	Omit(cols ...field.Expr) IGoodsOptionsInfoDo
	Join(table schema.Tabler, on ...field.Expr) IGoodsOptionsInfoDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IGoodsOptionsInfoDo
	RightJoin(table schema.Tabler, on ...field.Expr) IGoodsOptionsInfoDo
	Group(cols ...field.Expr) IGoodsOptionsInfoDo
	Having(conds ...gen.Condition) IGoodsOptionsInfoDo
	Limit(limit int) IGoodsOptionsInfoDo
	Offset(offset int) IGoodsOptionsInfoDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IGoodsOptionsInfoDo
	Unscoped() IGoodsOptionsInfoDo
	Create(values ...*model.GoodsOptionsInfo) error
	CreateInBatches(values []*model.GoodsOptionsInfo, batchSize int) error
	Save(values ...*model.GoodsOptionsInfo) error
	First() (*model.GoodsOptionsInfo, error)
	Take() (*model.GoodsOptionsInfo, error)
	Last() (*model.GoodsOptionsInfo, error)
	Find() ([]*model.GoodsOptionsInfo, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.GoodsOptionsInfo, err error)
	FindInBatches(result *[]*model.GoodsOptionsInfo, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.GoodsOptionsInfo) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IGoodsOptionsInfoDo
	Assign(attrs ...field.AssignExpr) IGoodsOptionsInfoDo
	Joins(fields ...field.RelationField) IGoodsOptionsInfoDo
	Preload(fields ...field.RelationField) IGoodsOptionsInfoDo
	FirstOrInit() (*model.GoodsOptionsInfo, error)
	FirstOrCreate() (*model.GoodsOptionsInfo, error)
	FindByPage(offset int, limit int) (result []*model.GoodsOptionsInfo, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IGoodsOptionsInfoDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (g goodsOptionsInfoDo) Debug() IGoodsOptionsInfoDo {
	return g.withDO(g.DO.Debug())
}

func (g goodsOptionsInfoDo) WithContext(ctx context.Context) IGoodsOptionsInfoDo {
	return g.withDO(g.DO.WithContext(ctx))
}

func (g goodsOptionsInfoDo) ReadDB() IGoodsOptionsInfoDo {
	return g.Clauses(dbresolver.Read)
}

func (g goodsOptionsInfoDo) WriteDB() IGoodsOptionsInfoDo {
	return g.Clauses(dbresolver.Write)
}

func (g goodsOptionsInfoDo) Session(config *gorm.Session) IGoodsOptionsInfoDo {
	return g.withDO(g.DO.Session(config))
}

func (g goodsOptionsInfoDo) Clauses(conds ...clause.Expression) IGoodsOptionsInfoDo {
	return g.withDO(g.DO.Clauses(conds...))
}

func (g goodsOptionsInfoDo) Returning(value interface{}, columns ...string) IGoodsOptionsInfoDo {
	return g.withDO(g.DO.Returning(value, columns...))
}

func (g goodsOptionsInfoDo) Not(conds ...gen.Condition) IGoodsOptionsInfoDo {
	return g.withDO(g.DO.Not(conds...))
}

func (g goodsOptionsInfoDo) Or(conds ...gen.Condition) IGoodsOptionsInfoDo {
	return g.withDO(g.DO.Or(conds...))
}

func (g goodsOptionsInfoDo) Select(conds ...field.Expr) IGoodsOptionsInfoDo {
	return g.withDO(g.DO.Select(conds...))
}

func (g goodsOptionsInfoDo) Where(conds ...gen.Condition) IGoodsOptionsInfoDo {
	return g.withDO(g.DO.Where(conds...))
}

func (g goodsOptionsInfoDo) Order(conds ...field.Expr) IGoodsOptionsInfoDo {
	return g.withDO(g.DO.Order(conds...))
}

func (g goodsOptionsInfoDo) Distinct(cols ...field.Expr) IGoodsOptionsInfoDo {
	return g.withDO(g.DO.Distinct(cols...))
}

func (g goodsOptionsInfoDo) Omit(cols ...field.Expr) IGoodsOptionsInfoDo {
	return g.withDO(g.DO.Omit(cols...))
}

func (g goodsOptionsInfoDo) Join(table schema.Tabler, on ...field.Expr) IGoodsOptionsInfoDo {
	return g.withDO(g.DO.Join(table, on...))
}

func (g goodsOptionsInfoDo) LeftJoin(table schema.Tabler, on ...field.Expr) IGoodsOptionsInfoDo {
	return g.withDO(g.DO.LeftJoin(table, on...))
}

func (g goodsOptionsInfoDo) RightJoin(table schema.Tabler, on ...field.Expr) IGoodsOptionsInfoDo {
	return g.withDO(g.DO.RightJoin(table, on...))
}

func (g goodsOptionsInfoDo) Group(cols ...field.Expr) IGoodsOptionsInfoDo {
	return g.withDO(g.DO.Group(cols...))
}

func (g goodsOptionsInfoDo) Having(conds ...gen.Condition) IGoodsOptionsInfoDo {
	return g.withDO(g.DO.Having(conds...))
}

func (g goodsOptionsInfoDo) Limit(limit int) IGoodsOptionsInfoDo {
	return g.withDO(g.DO.Limit(limit))
}

func (g goodsOptionsInfoDo) Offset(offset int) IGoodsOptionsInfoDo {
	return g.withDO(g.DO.Offset(offset))
}

func (g goodsOptionsInfoDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IGoodsOptionsInfoDo {
	return g.withDO(g.DO.Scopes(funcs...))
}

func (g goodsOptionsInfoDo) Unscoped() IGoodsOptionsInfoDo {
	return g.withDO(g.DO.Unscoped())
}

func (g goodsOptionsInfoDo) Create(values ...*model.GoodsOptionsInfo) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Create(values)
}

func (g goodsOptionsInfoDo) CreateInBatches(values []*model.GoodsOptionsInfo, batchSize int) error {
	return g.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (g goodsOptionsInfoDo) Save(values ...*model.GoodsOptionsInfo) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Save(values)
}

func (g goodsOptionsInfoDo) First() (*model.GoodsOptionsInfo, error) {
	if result, err := g.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.GoodsOptionsInfo), nil
	}
}

func (g goodsOptionsInfoDo) Take() (*model.GoodsOptionsInfo, error) {
	if result, err := g.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.GoodsOptionsInfo), nil
	}
}

func (g goodsOptionsInfoDo) Last() (*model.GoodsOptionsInfo, error) {
	if result, err := g.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.GoodsOptionsInfo), nil
	}
}

func (g goodsOptionsInfoDo) Find() ([]*model.GoodsOptionsInfo, error) {
	result, err := g.DO.Find()
	return result.([]*model.GoodsOptionsInfo), err
}

func (g goodsOptionsInfoDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.GoodsOptionsInfo, err error) {
	buf := make([]*model.GoodsOptionsInfo, 0, batchSize)
	err = g.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (g goodsOptionsInfoDo) FindInBatches(result *[]*model.GoodsOptionsInfo, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return g.DO.FindInBatches(result, batchSize, fc)
}

func (g goodsOptionsInfoDo) Attrs(attrs ...field.AssignExpr) IGoodsOptionsInfoDo {
	return g.withDO(g.DO.Attrs(attrs...))
}

func (g goodsOptionsInfoDo) Assign(attrs ...field.AssignExpr) IGoodsOptionsInfoDo {
	return g.withDO(g.DO.Assign(attrs...))
}

func (g goodsOptionsInfoDo) Joins(fields ...field.RelationField) IGoodsOptionsInfoDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Joins(_f))
	}
	return &g
}

func (g goodsOptionsInfoDo) Preload(fields ...field.RelationField) IGoodsOptionsInfoDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Preload(_f))
	}
	return &g
}

func (g goodsOptionsInfoDo) FirstOrInit() (*model.GoodsOptionsInfo, error) {
	if result, err := g.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.GoodsOptionsInfo), nil
	}
}

func (g goodsOptionsInfoDo) FirstOrCreate() (*model.GoodsOptionsInfo, error) {
	if result, err := g.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.GoodsOptionsInfo), nil
	}
}

func (g goodsOptionsInfoDo) FindByPage(offset int, limit int) (result []*model.GoodsOptionsInfo, count int64, err error) {
	result, err = g.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = g.Offset(-1).Limit(-1).Count()
	return
}

func (g goodsOptionsInfoDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = g.Count()
	if err != nil {
		return
	}

	err = g.Offset(offset).Limit(limit).Scan(result)
	return
}

func (g goodsOptionsInfoDo) Scan(result interface{}) (err error) {
	return g.DO.Scan(result)
}

func (g goodsOptionsInfoDo) Delete(models ...*model.GoodsOptionsInfo) (result gen.ResultInfo, err error) {
	return g.DO.Delete(models)
}

func (g *goodsOptionsInfoDo) withDO(do gen.Dao) *goodsOptionsInfoDo {
	g.DO = *do.(*gen.DO)
	return g
}
