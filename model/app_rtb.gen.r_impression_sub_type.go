package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _RImpressionSubTypeMgr struct {
	*_BaseMgr
}

// RImpressionSubTypeMgr open func
func RImpressionSubTypeMgr(db *gorm.DB) *_RImpressionSubTypeMgr {
	if db == nil {
		panic(fmt.Errorf("RImpressionSubTypeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_RImpressionSubTypeMgr{_BaseMgr: &_BaseMgr{DB: db.Table("r_impression_sub_type"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_RImpressionSubTypeMgr) GetTableName() string {
	return "r_impression_sub_type"
}

// Reset 重置gorm会话
func (obj *_RImpressionSubTypeMgr) Reset() *_RImpressionSubTypeMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_RImpressionSubTypeMgr) Get() (result RImpressionSubType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RImpressionSubType{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_RImpressionSubTypeMgr) Gets() (results []*RImpressionSubType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RImpressionSubType{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_RImpressionSubTypeMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(RImpressionSubType{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithIstID ist_id获取
func (obj *_RImpressionSubTypeMgr) WithIstID(istID uint32) Option {
	return optionFunc(func(o *options) { o.query["ist_id"] = istID })
}

// WithIstCode ist_code获取
func (obj *_RImpressionSubTypeMgr) WithIstCode(istCode string) Option {
	return optionFunc(func(o *options) { o.query["ist_code"] = istCode })
}

// WithIstName ist_name获取
func (obj *_RImpressionSubTypeMgr) WithIstName(istName string) Option {
	return optionFunc(func(o *options) { o.query["ist_name"] = istName })
}

// WithIstItID ist_it_id获取
func (obj *_RImpressionSubTypeMgr) WithIstItID(istItID int) Option {
	return optionFunc(func(o *options) { o.query["ist_it_id"] = istItID })
}

// WithIstStatus ist_status获取
func (obj *_RImpressionSubTypeMgr) WithIstStatus(istStatus int8) Option {
	return optionFunc(func(o *options) { o.query["ist_status"] = istStatus })
}

// WithIstSuffix ist_suffix获取
func (obj *_RImpressionSubTypeMgr) WithIstSuffix(istSuffix string) Option {
	return optionFunc(func(o *options) { o.query["ist_suffix"] = istSuffix })
}

// GetByOption 功能选项模式获取
func (obj *_RImpressionSubTypeMgr) GetByOption(opts ...Option) (result RImpressionSubType, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(RImpressionSubType{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_RImpressionSubTypeMgr) GetByOptions(opts ...Option) (results []*RImpressionSubType, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(RImpressionSubType{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromIstID 通过ist_id获取内容
func (obj *_RImpressionSubTypeMgr) GetFromIstID(istID uint32) (result RImpressionSubType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RImpressionSubType{}).Where("`ist_id` = ?", istID).Find(&result).Error

	return
}

// GetBatchFromIstID 批量查找
func (obj *_RImpressionSubTypeMgr) GetBatchFromIstID(istIDs []uint32) (results []*RImpressionSubType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RImpressionSubType{}).Where("`ist_id` IN (?)", istIDs).Find(&results).Error

	return
}

// GetFromIstCode 通过ist_code获取内容
func (obj *_RImpressionSubTypeMgr) GetFromIstCode(istCode string) (results []*RImpressionSubType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RImpressionSubType{}).Where("`ist_code` = ?", istCode).Find(&results).Error

	return
}

// GetBatchFromIstCode 批量查找
func (obj *_RImpressionSubTypeMgr) GetBatchFromIstCode(istCodes []string) (results []*RImpressionSubType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RImpressionSubType{}).Where("`ist_code` IN (?)", istCodes).Find(&results).Error

	return
}

// GetFromIstName 通过ist_name获取内容
func (obj *_RImpressionSubTypeMgr) GetFromIstName(istName string) (results []*RImpressionSubType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RImpressionSubType{}).Where("`ist_name` = ?", istName).Find(&results).Error

	return
}

// GetBatchFromIstName 批量查找
func (obj *_RImpressionSubTypeMgr) GetBatchFromIstName(istNames []string) (results []*RImpressionSubType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RImpressionSubType{}).Where("`ist_name` IN (?)", istNames).Find(&results).Error

	return
}

// GetFromIstItID 通过ist_it_id获取内容
func (obj *_RImpressionSubTypeMgr) GetFromIstItID(istItID int) (results []*RImpressionSubType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RImpressionSubType{}).Where("`ist_it_id` = ?", istItID).Find(&results).Error

	return
}

// GetBatchFromIstItID 批量查找
func (obj *_RImpressionSubTypeMgr) GetBatchFromIstItID(istItIDs []int) (results []*RImpressionSubType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RImpressionSubType{}).Where("`ist_it_id` IN (?)", istItIDs).Find(&results).Error

	return
}

// GetFromIstStatus 通过ist_status获取内容
func (obj *_RImpressionSubTypeMgr) GetFromIstStatus(istStatus int8) (results []*RImpressionSubType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RImpressionSubType{}).Where("`ist_status` = ?", istStatus).Find(&results).Error

	return
}

// GetBatchFromIstStatus 批量查找
func (obj *_RImpressionSubTypeMgr) GetBatchFromIstStatus(istStatuss []int8) (results []*RImpressionSubType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RImpressionSubType{}).Where("`ist_status` IN (?)", istStatuss).Find(&results).Error

	return
}

// GetFromIstSuffix 通过ist_suffix获取内容
func (obj *_RImpressionSubTypeMgr) GetFromIstSuffix(istSuffix string) (results []*RImpressionSubType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RImpressionSubType{}).Where("`ist_suffix` = ?", istSuffix).Find(&results).Error

	return
}

// GetBatchFromIstSuffix 批量查找
func (obj *_RImpressionSubTypeMgr) GetBatchFromIstSuffix(istSuffixs []string) (results []*RImpressionSubType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RImpressionSubType{}).Where("`ist_suffix` IN (?)", istSuffixs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_RImpressionSubTypeMgr) FetchByPrimaryKey(istID uint32) (result RImpressionSubType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RImpressionSubType{}).Where("`ist_id` = ?", istID).Find(&result).Error

	return
}
