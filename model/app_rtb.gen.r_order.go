package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _ROrderMgr struct {
	*_BaseMgr
}

// ROrderMgr open func
func ROrderMgr(db *gorm.DB) *_ROrderMgr {
	if db == nil {
		panic(fmt.Errorf("ROrderMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ROrderMgr{_BaseMgr: &_BaseMgr{DB: db.Table("r_order"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ROrderMgr) GetTableName() string {
	return "r_order"
}

// Reset 重置gorm会话
func (obj *_ROrderMgr) Reset() *_ROrderMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_ROrderMgr) Get() (result ROrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ROrder{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ROrderMgr) Gets() (results []*ROrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ROrder{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_ROrderMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(ROrder{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithOID o_id获取
func (obj *_ROrderMgr) WithOID(oID uint32) Option {
	return optionFunc(func(o *options) { o.query["o_id"] = oID })
}

// WithOName o_name获取
func (obj *_ROrderMgr) WithOName(oName string) Option {
	return optionFunc(func(o *options) { o.query["o_name"] = oName })
}

// WithOCode o_code获取
func (obj *_ROrderMgr) WithOCode(oCode string) Option {
	return optionFunc(func(o *options) { o.query["o_code"] = oCode })
}

// WithOStatus o_status获取
func (obj *_ROrderMgr) WithOStatus(oStatus int8) Option {
	return optionFunc(func(o *options) { o.query["o_status"] = oStatus })
}

// WithOCreateTime o_create_time获取
func (obj *_ROrderMgr) WithOCreateTime(oCreateTime int) Option {
	return optionFunc(func(o *options) { o.query["o_create_time"] = oCreateTime })
}

// GetByOption 功能选项模式获取
func (obj *_ROrderMgr) GetByOption(opts ...Option) (result ROrder, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(ROrder{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_ROrderMgr) GetByOptions(opts ...Option) (results []*ROrder, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(ROrder{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromOID 通过o_id获取内容
func (obj *_ROrderMgr) GetFromOID(oID uint32) (result ROrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ROrder{}).Where("`o_id` = ?", oID).Find(&result).Error

	return
}

// GetBatchFromOID 批量查找
func (obj *_ROrderMgr) GetBatchFromOID(oIDs []uint32) (results []*ROrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ROrder{}).Where("`o_id` IN (?)", oIDs).Find(&results).Error

	return
}

// GetFromOName 通过o_name获取内容
func (obj *_ROrderMgr) GetFromOName(oName string) (results []*ROrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ROrder{}).Where("`o_name` = ?", oName).Find(&results).Error

	return
}

// GetBatchFromOName 批量查找
func (obj *_ROrderMgr) GetBatchFromOName(oNames []string) (results []*ROrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ROrder{}).Where("`o_name` IN (?)", oNames).Find(&results).Error

	return
}

// GetFromOCode 通过o_code获取内容
func (obj *_ROrderMgr) GetFromOCode(oCode string) (results []*ROrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ROrder{}).Where("`o_code` = ?", oCode).Find(&results).Error

	return
}

// GetBatchFromOCode 批量查找
func (obj *_ROrderMgr) GetBatchFromOCode(oCodes []string) (results []*ROrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ROrder{}).Where("`o_code` IN (?)", oCodes).Find(&results).Error

	return
}

// GetFromOStatus 通过o_status获取内容
func (obj *_ROrderMgr) GetFromOStatus(oStatus int8) (results []*ROrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ROrder{}).Where("`o_status` = ?", oStatus).Find(&results).Error

	return
}

// GetBatchFromOStatus 批量查找
func (obj *_ROrderMgr) GetBatchFromOStatus(oStatuss []int8) (results []*ROrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ROrder{}).Where("`o_status` IN (?)", oStatuss).Find(&results).Error

	return
}

// GetFromOCreateTime 通过o_create_time获取内容
func (obj *_ROrderMgr) GetFromOCreateTime(oCreateTime int) (results []*ROrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ROrder{}).Where("`o_create_time` = ?", oCreateTime).Find(&results).Error

	return
}

// GetBatchFromOCreateTime 批量查找
func (obj *_ROrderMgr) GetBatchFromOCreateTime(oCreateTimes []int) (results []*ROrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ROrder{}).Where("`o_create_time` IN (?)", oCreateTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_ROrderMgr) FetchByPrimaryKey(oID uint32) (result ROrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ROrder{}).Where("`o_id` = ?", oID).Find(&result).Error

	return
}
