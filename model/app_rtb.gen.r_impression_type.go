package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _RImpressionTypeMgr struct {
	*_BaseMgr
}

// RImpressionTypeMgr open func
func RImpressionTypeMgr(db *gorm.DB) *_RImpressionTypeMgr {
	if db == nil {
		panic(fmt.Errorf("RImpressionTypeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_RImpressionTypeMgr{_BaseMgr: &_BaseMgr{DB: db.Table("r_impression_type"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_RImpressionTypeMgr) GetTableName() string {
	return "r_impression_type"
}

// Reset 重置gorm会话
func (obj *_RImpressionTypeMgr) Reset() *_RImpressionTypeMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_RImpressionTypeMgr) Get() (result RImpressionType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RImpressionType{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_RImpressionTypeMgr) Gets() (results []*RImpressionType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RImpressionType{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_RImpressionTypeMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(RImpressionType{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithItID it_id获取
func (obj *_RImpressionTypeMgr) WithItID(itID int) Option {
	return optionFunc(func(o *options) { o.query["it_id"] = itID })
}

// WithItCode it_code获取
func (obj *_RImpressionTypeMgr) WithItCode(itCode string) Option {
	return optionFunc(func(o *options) { o.query["it_code"] = itCode })
}

// WithItName it_name获取
func (obj *_RImpressionTypeMgr) WithItName(itName string) Option {
	return optionFunc(func(o *options) { o.query["it_name"] = itName })
}

// WithItStatus it_status获取
func (obj *_RImpressionTypeMgr) WithItStatus(itStatus int8) Option {
	return optionFunc(func(o *options) { o.query["it_status"] = itStatus })
}

// GetByOption 功能选项模式获取
func (obj *_RImpressionTypeMgr) GetByOption(opts ...Option) (result RImpressionType, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(RImpressionType{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_RImpressionTypeMgr) GetByOptions(opts ...Option) (results []*RImpressionType, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(RImpressionType{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromItID 通过it_id获取内容
func (obj *_RImpressionTypeMgr) GetFromItID(itID int) (result RImpressionType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RImpressionType{}).Where("`it_id` = ?", itID).Find(&result).Error

	return
}

// GetBatchFromItID 批量查找
func (obj *_RImpressionTypeMgr) GetBatchFromItID(itIDs []int) (results []*RImpressionType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RImpressionType{}).Where("`it_id` IN (?)", itIDs).Find(&results).Error

	return
}

// GetFromItCode 通过it_code获取内容
func (obj *_RImpressionTypeMgr) GetFromItCode(itCode string) (results []*RImpressionType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RImpressionType{}).Where("`it_code` = ?", itCode).Find(&results).Error

	return
}

// GetBatchFromItCode 批量查找
func (obj *_RImpressionTypeMgr) GetBatchFromItCode(itCodes []string) (results []*RImpressionType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RImpressionType{}).Where("`it_code` IN (?)", itCodes).Find(&results).Error

	return
}

// GetFromItName 通过it_name获取内容
func (obj *_RImpressionTypeMgr) GetFromItName(itName string) (results []*RImpressionType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RImpressionType{}).Where("`it_name` = ?", itName).Find(&results).Error

	return
}

// GetBatchFromItName 批量查找
func (obj *_RImpressionTypeMgr) GetBatchFromItName(itNames []string) (results []*RImpressionType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RImpressionType{}).Where("`it_name` IN (?)", itNames).Find(&results).Error

	return
}

// GetFromItStatus 通过it_status获取内容
func (obj *_RImpressionTypeMgr) GetFromItStatus(itStatus int8) (results []*RImpressionType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RImpressionType{}).Where("`it_status` = ?", itStatus).Find(&results).Error

	return
}

// GetBatchFromItStatus 批量查找
func (obj *_RImpressionTypeMgr) GetBatchFromItStatus(itStatuss []int8) (results []*RImpressionType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RImpressionType{}).Where("`it_status` IN (?)", itStatuss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_RImpressionTypeMgr) FetchByPrimaryKey(itID int) (result RImpressionType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RImpressionType{}).Where("`it_id` = ?", itID).Find(&result).Error

	return
}
