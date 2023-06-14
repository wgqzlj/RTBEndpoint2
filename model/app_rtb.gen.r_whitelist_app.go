package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _RWhitelistAppMgr struct {
	*_BaseMgr
}

// RWhitelistAppMgr open func
func RWhitelistAppMgr(db *gorm.DB) *_RWhitelistAppMgr {
	if db == nil {
		panic(fmt.Errorf("RWhitelistAppMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_RWhitelistAppMgr{_BaseMgr: &_BaseMgr{DB: db.Table("r_whitelist_app"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_RWhitelistAppMgr) GetTableName() string {
	return "r_whitelist_app"
}

// Reset 重置gorm会话
func (obj *_RWhitelistAppMgr) Reset() *_RWhitelistAppMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_RWhitelistAppMgr) Get() (result RWhitelistApp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RWhitelistApp{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_RWhitelistAppMgr) Gets() (results []*RWhitelistApp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RWhitelistApp{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_RWhitelistAppMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(RWhitelistApp{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithWaID wa_id获取
func (obj *_RWhitelistAppMgr) WithWaID(waID uint32) Option {
	return optionFunc(func(o *options) { o.query["wa_id"] = waID })
}

// WithWaName wa_name获取
func (obj *_RWhitelistAppMgr) WithWaName(waName string) Option {
	return optionFunc(func(o *options) { o.query["wa_name"] = waName })
}

// WithWaBundle wa_bundle获取
func (obj *_RWhitelistAppMgr) WithWaBundle(waBundle string) Option {
	return optionFunc(func(o *options) { o.query["wa_bundle"] = waBundle })
}

// WithWaCreatetime wa_createtime获取
func (obj *_RWhitelistAppMgr) WithWaCreatetime(waCreatetime int) Option {
	return optionFunc(func(o *options) { o.query["wa_createtime"] = waCreatetime })
}

// WithWaStatus wa_status获取
func (obj *_RWhitelistAppMgr) WithWaStatus(waStatus int8) Option {
	return optionFunc(func(o *options) { o.query["wa_status"] = waStatus })
}

// GetByOption 功能选项模式获取
func (obj *_RWhitelistAppMgr) GetByOption(opts ...Option) (result RWhitelistApp, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(RWhitelistApp{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_RWhitelistAppMgr) GetByOptions(opts ...Option) (results []*RWhitelistApp, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(RWhitelistApp{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromWaID 通过wa_id获取内容
func (obj *_RWhitelistAppMgr) GetFromWaID(waID uint32) (result RWhitelistApp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RWhitelistApp{}).Where("`wa_id` = ?", waID).Find(&result).Error

	return
}

// GetBatchFromWaID 批量查找
func (obj *_RWhitelistAppMgr) GetBatchFromWaID(waIDs []uint32) (results []*RWhitelistApp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RWhitelistApp{}).Where("`wa_id` IN (?)", waIDs).Find(&results).Error

	return
}

// GetFromWaName 通过wa_name获取内容
func (obj *_RWhitelistAppMgr) GetFromWaName(waName string) (results []*RWhitelistApp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RWhitelistApp{}).Where("`wa_name` = ?", waName).Find(&results).Error

	return
}

// GetBatchFromWaName 批量查找
func (obj *_RWhitelistAppMgr) GetBatchFromWaName(waNames []string) (results []*RWhitelistApp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RWhitelistApp{}).Where("`wa_name` IN (?)", waNames).Find(&results).Error

	return
}

// GetFromWaBundle 通过wa_bundle获取内容
func (obj *_RWhitelistAppMgr) GetFromWaBundle(waBundle string) (results []*RWhitelistApp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RWhitelistApp{}).Where("`wa_bundle` = ?", waBundle).Find(&results).Error

	return
}

// GetBatchFromWaBundle 批量查找
func (obj *_RWhitelistAppMgr) GetBatchFromWaBundle(waBundles []string) (results []*RWhitelistApp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RWhitelistApp{}).Where("`wa_bundle` IN (?)", waBundles).Find(&results).Error

	return
}

// GetFromWaCreatetime 通过wa_createtime获取内容
func (obj *_RWhitelistAppMgr) GetFromWaCreatetime(waCreatetime int) (results []*RWhitelistApp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RWhitelistApp{}).Where("`wa_createtime` = ?", waCreatetime).Find(&results).Error

	return
}

// GetBatchFromWaCreatetime 批量查找
func (obj *_RWhitelistAppMgr) GetBatchFromWaCreatetime(waCreatetimes []int) (results []*RWhitelistApp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RWhitelistApp{}).Where("`wa_createtime` IN (?)", waCreatetimes).Find(&results).Error

	return
}

// GetFromWaStatus 通过wa_status获取内容
func (obj *_RWhitelistAppMgr) GetFromWaStatus(waStatus int8) (results []*RWhitelistApp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RWhitelistApp{}).Where("`wa_status` = ?", waStatus).Find(&results).Error

	return
}

// GetBatchFromWaStatus 批量查找
func (obj *_RWhitelistAppMgr) GetBatchFromWaStatus(waStatuss []int8) (results []*RWhitelistApp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RWhitelistApp{}).Where("`wa_status` IN (?)", waStatuss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_RWhitelistAppMgr) FetchByPrimaryKey(waID uint32) (result RWhitelistApp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RWhitelistApp{}).Where("`wa_id` = ?", waID).Find(&result).Error

	return
}
