package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _RSmartPriceAppMgr struct {
	*_BaseMgr
}

// RSmartPriceAppMgr open func
func RSmartPriceAppMgr(db *gorm.DB) *_RSmartPriceAppMgr {
	if db == nil {
		panic(fmt.Errorf("RSmartPriceAppMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_RSmartPriceAppMgr{_BaseMgr: &_BaseMgr{DB: db.Table("r_smart_price_app"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_RSmartPriceAppMgr) GetTableName() string {
	return "r_smart_price_app"
}

// Reset 重置gorm会话
func (obj *_RSmartPriceAppMgr) Reset() *_RSmartPriceAppMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_RSmartPriceAppMgr) Get() (result RSmartPriceApp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceApp{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_RSmartPriceAppMgr) Gets() (results []*RSmartPriceApp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceApp{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_RSmartPriceAppMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(RSmartPriceApp{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSpaID spa_id获取
func (obj *_RSmartPriceAppMgr) WithSpaID(spaID uint32) Option {
	return optionFunc(func(o *options) { o.query["spa_id"] = spaID })
}

// WithSpaBundle spa_bundle获取
func (obj *_RSmartPriceAppMgr) WithSpaBundle(spaBundle string) Option {
	return optionFunc(func(o *options) { o.query["spa_bundle"] = spaBundle })
}

// WithSpaStatus spa_status获取
func (obj *_RSmartPriceAppMgr) WithSpaStatus(spaStatus int8) Option {
	return optionFunc(func(o *options) { o.query["spa_status"] = spaStatus })
}

// WithSpaCreateTime spa_create_time获取
func (obj *_RSmartPriceAppMgr) WithSpaCreateTime(spaCreateTime int) Option {
	return optionFunc(func(o *options) { o.query["spa_create_time"] = spaCreateTime })
}

// GetByOption 功能选项模式获取
func (obj *_RSmartPriceAppMgr) GetByOption(opts ...Option) (result RSmartPriceApp, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceApp{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_RSmartPriceAppMgr) GetByOptions(opts ...Option) (results []*RSmartPriceApp, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceApp{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromSpaID 通过spa_id获取内容
func (obj *_RSmartPriceAppMgr) GetFromSpaID(spaID uint32) (result RSmartPriceApp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceApp{}).Where("`spa_id` = ?", spaID).Find(&result).Error

	return
}

// GetBatchFromSpaID 批量查找
func (obj *_RSmartPriceAppMgr) GetBatchFromSpaID(spaIDs []uint32) (results []*RSmartPriceApp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceApp{}).Where("`spa_id` IN (?)", spaIDs).Find(&results).Error

	return
}

// GetFromSpaBundle 通过spa_bundle获取内容
func (obj *_RSmartPriceAppMgr) GetFromSpaBundle(spaBundle string) (results []*RSmartPriceApp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceApp{}).Where("`spa_bundle` = ?", spaBundle).Find(&results).Error

	return
}

// GetBatchFromSpaBundle 批量查找
func (obj *_RSmartPriceAppMgr) GetBatchFromSpaBundle(spaBundles []string) (results []*RSmartPriceApp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceApp{}).Where("`spa_bundle` IN (?)", spaBundles).Find(&results).Error

	return
}

// GetFromSpaStatus 通过spa_status获取内容
func (obj *_RSmartPriceAppMgr) GetFromSpaStatus(spaStatus int8) (results []*RSmartPriceApp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceApp{}).Where("`spa_status` = ?", spaStatus).Find(&results).Error

	return
}

// GetBatchFromSpaStatus 批量查找
func (obj *_RSmartPriceAppMgr) GetBatchFromSpaStatus(spaStatuss []int8) (results []*RSmartPriceApp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceApp{}).Where("`spa_status` IN (?)", spaStatuss).Find(&results).Error

	return
}

// GetFromSpaCreateTime 通过spa_create_time获取内容
func (obj *_RSmartPriceAppMgr) GetFromSpaCreateTime(spaCreateTime int) (results []*RSmartPriceApp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceApp{}).Where("`spa_create_time` = ?", spaCreateTime).Find(&results).Error

	return
}

// GetBatchFromSpaCreateTime 批量查找
func (obj *_RSmartPriceAppMgr) GetBatchFromSpaCreateTime(spaCreateTimes []int) (results []*RSmartPriceApp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceApp{}).Where("`spa_create_time` IN (?)", spaCreateTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_RSmartPriceAppMgr) FetchByPrimaryKey(spaID uint32) (result RSmartPriceApp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceApp{}).Where("`spa_id` = ?", spaID).Find(&result).Error

	return
}
