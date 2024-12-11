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

func newRefundInfo(db *gorm.DB, opts ...gen.DOOption) refundInfo {
	_refundInfo := refundInfo{}

	_refundInfo.refundInfoDo.UseDB(db, opts...)
	_refundInfo.refundInfoDo.UseModel(&model.RefundInfo{})

	tableName := _refundInfo.refundInfoDo.TableName()
	_refundInfo.ALL = field.NewAsterisk(tableName)
	_refundInfo.ID = field.NewInt32(tableName, "id")
	_refundInfo.Number = field.NewString(tableName, "number")
	_refundInfo.OrderID = field.NewInt32(tableName, "order_id")
	_refundInfo.GoodsID = field.NewInt32(tableName, "goods_id")
	_refundInfo.Reason = field.NewString(tableName, "reason")
	_refundInfo.Status = field.NewBool(tableName, "status")
	_refundInfo.UserID = field.NewInt32(tableName, "user_id")
	_refundInfo.CreatedAt = field.NewTime(tableName, "created_at")
	_refundInfo.UpdatedAt = field.NewTime(tableName, "updated_at")
	_refundInfo.DeletedAt = field.NewField(tableName, "deleted_at")

	_refundInfo.fillFieldMap()

	return _refundInfo
}

type refundInfo struct {
	refundInfoDo refundInfoDo

	ALL     field.Asterisk
	ID      field.Int32  // 售后退款表
	Number  field.String // 售后订单号
	OrderID field.Int32  // 订单id
	/*
		要售后的商品id

	*/
	GoodsID field.Int32
	Reason  field.String // 退款原因
	/*
		状态 1待处理 2同意退款 3拒绝退款

	*/
	Status    field.Bool
	UserID    field.Int32 // 用户id
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field

	fieldMap map[string]field.Expr
}

func (r refundInfo) Table(newTableName string) *refundInfo {
	r.refundInfoDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r refundInfo) As(alias string) *refundInfo {
	r.refundInfoDo.DO = *(r.refundInfoDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *refundInfo) updateTableName(table string) *refundInfo {
	r.ALL = field.NewAsterisk(table)
	r.ID = field.NewInt32(table, "id")
	r.Number = field.NewString(table, "number")
	r.OrderID = field.NewInt32(table, "order_id")
	r.GoodsID = field.NewInt32(table, "goods_id")
	r.Reason = field.NewString(table, "reason")
	r.Status = field.NewBool(table, "status")
	r.UserID = field.NewInt32(table, "user_id")
	r.CreatedAt = field.NewTime(table, "created_at")
	r.UpdatedAt = field.NewTime(table, "updated_at")
	r.DeletedAt = field.NewField(table, "deleted_at")

	r.fillFieldMap()

	return r
}

func (r *refundInfo) WithContext(ctx context.Context) IRefundInfoDo {
	return r.refundInfoDo.WithContext(ctx)
}

func (r refundInfo) TableName() string { return r.refundInfoDo.TableName() }

func (r refundInfo) Alias() string { return r.refundInfoDo.Alias() }

func (r refundInfo) Columns(cols ...field.Expr) gen.Columns { return r.refundInfoDo.Columns(cols...) }

func (r *refundInfo) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *refundInfo) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 10)
	r.fieldMap["id"] = r.ID
	r.fieldMap["number"] = r.Number
	r.fieldMap["order_id"] = r.OrderID
	r.fieldMap["goods_id"] = r.GoodsID
	r.fieldMap["reason"] = r.Reason
	r.fieldMap["status"] = r.Status
	r.fieldMap["user_id"] = r.UserID
	r.fieldMap["created_at"] = r.CreatedAt
	r.fieldMap["updated_at"] = r.UpdatedAt
	r.fieldMap["deleted_at"] = r.DeletedAt
}

func (r refundInfo) clone(db *gorm.DB) refundInfo {
	r.refundInfoDo.ReplaceConnPool(db.Statement.ConnPool)
	return r
}

func (r refundInfo) replaceDB(db *gorm.DB) refundInfo {
	r.refundInfoDo.ReplaceDB(db)
	return r
}

type refundInfoDo struct{ gen.DO }

type IRefundInfoDo interface {
	gen.SubQuery
	Debug() IRefundInfoDo
	WithContext(ctx context.Context) IRefundInfoDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IRefundInfoDo
	WriteDB() IRefundInfoDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IRefundInfoDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IRefundInfoDo
	Not(conds ...gen.Condition) IRefundInfoDo
	Or(conds ...gen.Condition) IRefundInfoDo
	Select(conds ...field.Expr) IRefundInfoDo
	Where(conds ...gen.Condition) IRefundInfoDo
	Order(conds ...field.Expr) IRefundInfoDo
	Distinct(cols ...field.Expr) IRefundInfoDo
	Omit(cols ...field.Expr) IRefundInfoDo
	Join(table schema.Tabler, on ...field.Expr) IRefundInfoDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IRefundInfoDo
	RightJoin(table schema.Tabler, on ...field.Expr) IRefundInfoDo
	Group(cols ...field.Expr) IRefundInfoDo
	Having(conds ...gen.Condition) IRefundInfoDo
	Limit(limit int) IRefundInfoDo
	Offset(offset int) IRefundInfoDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IRefundInfoDo
	Unscoped() IRefundInfoDo
	Create(values ...*model.RefundInfo) error
	CreateInBatches(values []*model.RefundInfo, batchSize int) error
	Save(values ...*model.RefundInfo) error
	First() (*model.RefundInfo, error)
	Take() (*model.RefundInfo, error)
	Last() (*model.RefundInfo, error)
	Find() ([]*model.RefundInfo, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.RefundInfo, err error)
	FindInBatches(result *[]*model.RefundInfo, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.RefundInfo) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IRefundInfoDo
	Assign(attrs ...field.AssignExpr) IRefundInfoDo
	Joins(fields ...field.RelationField) IRefundInfoDo
	Preload(fields ...field.RelationField) IRefundInfoDo
	FirstOrInit() (*model.RefundInfo, error)
	FirstOrCreate() (*model.RefundInfo, error)
	FindByPage(offset int, limit int) (result []*model.RefundInfo, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IRefundInfoDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (r refundInfoDo) Debug() IRefundInfoDo {
	return r.withDO(r.DO.Debug())
}

func (r refundInfoDo) WithContext(ctx context.Context) IRefundInfoDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r refundInfoDo) ReadDB() IRefundInfoDo {
	return r.Clauses(dbresolver.Read)
}

func (r refundInfoDo) WriteDB() IRefundInfoDo {
	return r.Clauses(dbresolver.Write)
}

func (r refundInfoDo) Session(config *gorm.Session) IRefundInfoDo {
	return r.withDO(r.DO.Session(config))
}

func (r refundInfoDo) Clauses(conds ...clause.Expression) IRefundInfoDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r refundInfoDo) Returning(value interface{}, columns ...string) IRefundInfoDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r refundInfoDo) Not(conds ...gen.Condition) IRefundInfoDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r refundInfoDo) Or(conds ...gen.Condition) IRefundInfoDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r refundInfoDo) Select(conds ...field.Expr) IRefundInfoDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r refundInfoDo) Where(conds ...gen.Condition) IRefundInfoDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r refundInfoDo) Order(conds ...field.Expr) IRefundInfoDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r refundInfoDo) Distinct(cols ...field.Expr) IRefundInfoDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r refundInfoDo) Omit(cols ...field.Expr) IRefundInfoDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r refundInfoDo) Join(table schema.Tabler, on ...field.Expr) IRefundInfoDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r refundInfoDo) LeftJoin(table schema.Tabler, on ...field.Expr) IRefundInfoDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r refundInfoDo) RightJoin(table schema.Tabler, on ...field.Expr) IRefundInfoDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r refundInfoDo) Group(cols ...field.Expr) IRefundInfoDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r refundInfoDo) Having(conds ...gen.Condition) IRefundInfoDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r refundInfoDo) Limit(limit int) IRefundInfoDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r refundInfoDo) Offset(offset int) IRefundInfoDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r refundInfoDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IRefundInfoDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r refundInfoDo) Unscoped() IRefundInfoDo {
	return r.withDO(r.DO.Unscoped())
}

func (r refundInfoDo) Create(values ...*model.RefundInfo) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r refundInfoDo) CreateInBatches(values []*model.RefundInfo, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r refundInfoDo) Save(values ...*model.RefundInfo) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r refundInfoDo) First() (*model.RefundInfo, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.RefundInfo), nil
	}
}

func (r refundInfoDo) Take() (*model.RefundInfo, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.RefundInfo), nil
	}
}

func (r refundInfoDo) Last() (*model.RefundInfo, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.RefundInfo), nil
	}
}

func (r refundInfoDo) Find() ([]*model.RefundInfo, error) {
	result, err := r.DO.Find()
	return result.([]*model.RefundInfo), err
}

func (r refundInfoDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.RefundInfo, err error) {
	buf := make([]*model.RefundInfo, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r refundInfoDo) FindInBatches(result *[]*model.RefundInfo, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r refundInfoDo) Attrs(attrs ...field.AssignExpr) IRefundInfoDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r refundInfoDo) Assign(attrs ...field.AssignExpr) IRefundInfoDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r refundInfoDo) Joins(fields ...field.RelationField) IRefundInfoDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r refundInfoDo) Preload(fields ...field.RelationField) IRefundInfoDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r refundInfoDo) FirstOrInit() (*model.RefundInfo, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.RefundInfo), nil
	}
}

func (r refundInfoDo) FirstOrCreate() (*model.RefundInfo, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.RefundInfo), nil
	}
}

func (r refundInfoDo) FindByPage(offset int, limit int) (result []*model.RefundInfo, count int64, err error) {
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

func (r refundInfoDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r refundInfoDo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r refundInfoDo) Delete(models ...*model.RefundInfo) (result gen.ResultInfo, err error) {
	return r.DO.Delete(models)
}

func (r *refundInfoDo) withDO(do gen.Dao) *refundInfoDo {
	r.DO = *do.(*gen.DO)
	return r
}