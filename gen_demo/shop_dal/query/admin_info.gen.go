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

func newAdminInfo(db *gorm.DB, opts ...gen.DOOption) adminInfo {
	_adminInfo := adminInfo{}

	_adminInfo.adminInfoDo.UseDB(db, opts...)
	_adminInfo.adminInfoDo.UseModel(&model.AdminInfo{})

	tableName := _adminInfo.adminInfoDo.TableName()
	_adminInfo.ALL = field.NewAsterisk(tableName)
	_adminInfo.ID = field.NewInt32(tableName, "id")
	_adminInfo.Name = field.NewString(tableName, "name")
	_adminInfo.Password = field.NewString(tableName, "password")
	_adminInfo.RoleIds = field.NewString(tableName, "role_ids")
	_adminInfo.CreatedAt = field.NewTime(tableName, "created_at")
	_adminInfo.UpdatedAt = field.NewTime(tableName, "updated_at")
	_adminInfo.UserSalt = field.NewString(tableName, "user_salt")
	_adminInfo.IsAdmin = field.NewBool(tableName, "is_admin")

	_adminInfo.fillFieldMap()

	return _adminInfo
}

type adminInfo struct {
	adminInfoDo adminInfoDo

	ALL       field.Asterisk
	ID        field.Int32
	Name      field.String // 用户名
	Password  field.String // 密码
	RoleIds   field.String // 角色ids
	CreatedAt field.Time
	UpdatedAt field.Time
	UserSalt  field.String // 加密盐
	IsAdmin   field.Bool   // 是否超级管理员

	fieldMap map[string]field.Expr
}

func (a adminInfo) Table(newTableName string) *adminInfo {
	a.adminInfoDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a adminInfo) As(alias string) *adminInfo {
	a.adminInfoDo.DO = *(a.adminInfoDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *adminInfo) updateTableName(table string) *adminInfo {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt32(table, "id")
	a.Name = field.NewString(table, "name")
	a.Password = field.NewString(table, "password")
	a.RoleIds = field.NewString(table, "role_ids")
	a.CreatedAt = field.NewTime(table, "created_at")
	a.UpdatedAt = field.NewTime(table, "updated_at")
	a.UserSalt = field.NewString(table, "user_salt")
	a.IsAdmin = field.NewBool(table, "is_admin")

	a.fillFieldMap()

	return a
}

func (a *adminInfo) WithContext(ctx context.Context) IAdminInfoDo {
	return a.adminInfoDo.WithContext(ctx)
}

func (a adminInfo) TableName() string { return a.adminInfoDo.TableName() }

func (a adminInfo) Alias() string { return a.adminInfoDo.Alias() }

func (a adminInfo) Columns(cols ...field.Expr) gen.Columns { return a.adminInfoDo.Columns(cols...) }

func (a *adminInfo) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *adminInfo) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 8)
	a.fieldMap["id"] = a.ID
	a.fieldMap["name"] = a.Name
	a.fieldMap["password"] = a.Password
	a.fieldMap["role_ids"] = a.RoleIds
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["updated_at"] = a.UpdatedAt
	a.fieldMap["user_salt"] = a.UserSalt
	a.fieldMap["is_admin"] = a.IsAdmin
}

func (a adminInfo) clone(db *gorm.DB) adminInfo {
	a.adminInfoDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a adminInfo) replaceDB(db *gorm.DB) adminInfo {
	a.adminInfoDo.ReplaceDB(db)
	return a
}

type adminInfoDo struct{ gen.DO }

type IAdminInfoDo interface {
	gen.SubQuery
	Debug() IAdminInfoDo
	WithContext(ctx context.Context) IAdminInfoDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAdminInfoDo
	WriteDB() IAdminInfoDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAdminInfoDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAdminInfoDo
	Not(conds ...gen.Condition) IAdminInfoDo
	Or(conds ...gen.Condition) IAdminInfoDo
	Select(conds ...field.Expr) IAdminInfoDo
	Where(conds ...gen.Condition) IAdminInfoDo
	Order(conds ...field.Expr) IAdminInfoDo
	Distinct(cols ...field.Expr) IAdminInfoDo
	Omit(cols ...field.Expr) IAdminInfoDo
	Join(table schema.Tabler, on ...field.Expr) IAdminInfoDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAdminInfoDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAdminInfoDo
	Group(cols ...field.Expr) IAdminInfoDo
	Having(conds ...gen.Condition) IAdminInfoDo
	Limit(limit int) IAdminInfoDo
	Offset(offset int) IAdminInfoDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAdminInfoDo
	Unscoped() IAdminInfoDo
	Create(values ...*model.AdminInfo) error
	CreateInBatches(values []*model.AdminInfo, batchSize int) error
	Save(values ...*model.AdminInfo) error
	First() (*model.AdminInfo, error)
	Take() (*model.AdminInfo, error)
	Last() (*model.AdminInfo, error)
	Find() ([]*model.AdminInfo, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AdminInfo, err error)
	FindInBatches(result *[]*model.AdminInfo, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.AdminInfo) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAdminInfoDo
	Assign(attrs ...field.AssignExpr) IAdminInfoDo
	Joins(fields ...field.RelationField) IAdminInfoDo
	Preload(fields ...field.RelationField) IAdminInfoDo
	FirstOrInit() (*model.AdminInfo, error)
	FirstOrCreate() (*model.AdminInfo, error)
	FindByPage(offset int, limit int) (result []*model.AdminInfo, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAdminInfoDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a adminInfoDo) Debug() IAdminInfoDo {
	return a.withDO(a.DO.Debug())
}

func (a adminInfoDo) WithContext(ctx context.Context) IAdminInfoDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a adminInfoDo) ReadDB() IAdminInfoDo {
	return a.Clauses(dbresolver.Read)
}

func (a adminInfoDo) WriteDB() IAdminInfoDo {
	return a.Clauses(dbresolver.Write)
}

func (a adminInfoDo) Session(config *gorm.Session) IAdminInfoDo {
	return a.withDO(a.DO.Session(config))
}

func (a adminInfoDo) Clauses(conds ...clause.Expression) IAdminInfoDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a adminInfoDo) Returning(value interface{}, columns ...string) IAdminInfoDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a adminInfoDo) Not(conds ...gen.Condition) IAdminInfoDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a adminInfoDo) Or(conds ...gen.Condition) IAdminInfoDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a adminInfoDo) Select(conds ...field.Expr) IAdminInfoDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a adminInfoDo) Where(conds ...gen.Condition) IAdminInfoDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a adminInfoDo) Order(conds ...field.Expr) IAdminInfoDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a adminInfoDo) Distinct(cols ...field.Expr) IAdminInfoDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a adminInfoDo) Omit(cols ...field.Expr) IAdminInfoDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a adminInfoDo) Join(table schema.Tabler, on ...field.Expr) IAdminInfoDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a adminInfoDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAdminInfoDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a adminInfoDo) RightJoin(table schema.Tabler, on ...field.Expr) IAdminInfoDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a adminInfoDo) Group(cols ...field.Expr) IAdminInfoDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a adminInfoDo) Having(conds ...gen.Condition) IAdminInfoDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a adminInfoDo) Limit(limit int) IAdminInfoDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a adminInfoDo) Offset(offset int) IAdminInfoDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a adminInfoDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAdminInfoDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a adminInfoDo) Unscoped() IAdminInfoDo {
	return a.withDO(a.DO.Unscoped())
}

func (a adminInfoDo) Create(values ...*model.AdminInfo) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a adminInfoDo) CreateInBatches(values []*model.AdminInfo, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a adminInfoDo) Save(values ...*model.AdminInfo) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a adminInfoDo) First() (*model.AdminInfo, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.AdminInfo), nil
	}
}

func (a adminInfoDo) Take() (*model.AdminInfo, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.AdminInfo), nil
	}
}

func (a adminInfoDo) Last() (*model.AdminInfo, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.AdminInfo), nil
	}
}

func (a adminInfoDo) Find() ([]*model.AdminInfo, error) {
	result, err := a.DO.Find()
	return result.([]*model.AdminInfo), err
}

func (a adminInfoDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AdminInfo, err error) {
	buf := make([]*model.AdminInfo, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a adminInfoDo) FindInBatches(result *[]*model.AdminInfo, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a adminInfoDo) Attrs(attrs ...field.AssignExpr) IAdminInfoDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a adminInfoDo) Assign(attrs ...field.AssignExpr) IAdminInfoDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a adminInfoDo) Joins(fields ...field.RelationField) IAdminInfoDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a adminInfoDo) Preload(fields ...field.RelationField) IAdminInfoDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a adminInfoDo) FirstOrInit() (*model.AdminInfo, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.AdminInfo), nil
	}
}

func (a adminInfoDo) FirstOrCreate() (*model.AdminInfo, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.AdminInfo), nil
	}
}

func (a adminInfoDo) FindByPage(offset int, limit int) (result []*model.AdminInfo, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a adminInfoDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a adminInfoDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a adminInfoDo) Delete(models ...*model.AdminInfo) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *adminInfoDo) withDO(do gen.Dao) *adminInfoDo {
	a.DO = *do.(*gen.DO)
	return a
}
