package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _RBlacklistAppMgr struct {
	*_BaseMgr
}

// RBlacklistAppMgr open func
func RBlacklistAppMgr(db *gorm.DB) *_RBlacklistAppMgr {
	if db == nil {
		panic(fmt.Errorf("RBlacklistAppMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_RBlacklistAppMgr{_BaseMgr: &_BaseMgr{DB: db.Table("r_blacklist_app"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_RBlacklistAppMgr) GetTableName() string {
	return "r_blacklist_app"
}

// Reset 重置gorm会话
func (obj *_RBlacklistAppMgr) Reset() *_RBlacklistAppMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_RBlacklistAppMgr) Get() (result RBlacklistApp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistApp{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_RBlacklistAppMgr) Gets() (results []*RBlacklistApp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistApp{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_RBlacklistAppMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(RBlacklistApp{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithBaID ba_id获取
func (obj *_RBlacklistAppMgr) WithBaID(baID uint32) Option {
	return optionFunc(func(o *options) { o.query["ba_id"] = baID })
}

// WithBaName ba_name获取
func (obj *_RBlacklistAppMgr) WithBaName(baName string) Option {
	return optionFunc(func(o *options) { o.query["ba_name"] = baName })
}

// WithBaBundle ba_bundle获取
func (obj *_RBlacklistAppMgr) WithBaBundle(baBundle string) Option {
	return optionFunc(func(o *options) { o.query["ba_bundle"] = baBundle })
}

// WithBaCreatetime ba_createtime获取
func (obj *_RBlacklistAppMgr) WithBaCreatetime(baCreatetime int) Option {
	return optionFunc(func(o *options) { o.query["ba_createtime"] = baCreatetime })
}

// WithBaStatus ba_status获取
func (obj *_RBlacklistAppMgr) WithBaStatus(baStatus int8) Option {
	return optionFunc(func(o *options) { o.query["ba_status"] = baStatus })
}

// WithBaItID ba_it_id获取
func (obj *_RBlacklistAppMgr) WithBaItID(baItID int) Option {
	return optionFunc(func(o *options) { o.query["ba_it_id"] = baItID })
}

// GetByOption 功能选项模式获取
func (obj *_RBlacklistAppMgr) GetByOption(opts ...Option) (result RBlacklistApp, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistApp{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_RBlacklistAppMgr) GetByOptions(opts ...Option) (results []*RBlacklistApp, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistApp{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromBaID 通过ba_id获取内容
func (obj *_RBlacklistAppMgr) GetFromBaID(baID uint32) (result RBlacklistApp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistApp{}).Where("`ba_id` = ?", baID).Find(&result).Error

	return
}

// GetBatchFromBaID 批量查找
func (obj *_RBlacklistAppMgr) GetBatchFromBaID(baIDs []uint32) (results []*RBlacklistApp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistApp{}).Where("`ba_id` IN (?)", baIDs).Find(&results).Error

	return
}

// GetFromBaName 通过ba_name获取内容
func (obj *_RBlacklistAppMgr) GetFromBaName(baName string) (results []*RBlacklistApp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistApp{}).Where("`ba_name` = ?", baName).Find(&results).Error

	return
}

// GetBatchFromBaName 批量查找
func (obj *_RBlacklistAppMgr) GetBatchFromBaName(baNames []string) (results []*RBlacklistApp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistApp{}).Where("`ba_name` IN (?)", baNames).Find(&results).Error

	return
}

// GetFromBaBundle 通过ba_bundle获取内容
func (obj *_RBlacklistAppMgr) GetFromBaBundle(baBundle string) (results []*RBlacklistApp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistApp{}).Where("`ba_bundle` = ?", baBundle).Find(&results).Error

	return
}

// GetBatchFromBaBundle 批量查找
func (obj *_RBlacklistAppMgr) GetBatchFromBaBundle(baBundles []string) (results []*RBlacklistApp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistApp{}).Where("`ba_bundle` IN (?)", baBundles).Find(&results).Error

	return
}

// GetFromBaCreatetime 通过ba_createtime获取内容
func (obj *_RBlacklistAppMgr) GetFromBaCreatetime(baCreatetime int) (results []*RBlacklistApp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistApp{}).Where("`ba_createtime` = ?", baCreatetime).Find(&results).Error

	return
}

// GetBatchFromBaCreatetime 批量查找
func (obj *_RBlacklistAppMgr) GetBatchFromBaCreatetime(baCreatetimes []int) (results []*RBlacklistApp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistApp{}).Where("`ba_createtime` IN (?)", baCreatetimes).Find(&results).Error

	return
}

// GetFromBaStatus 通过ba_status获取内容
func (obj *_RBlacklistAppMgr) GetFromBaStatus(baStatus int8) (results []*RBlacklistApp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistApp{}).Where("`ba_status` = ?", baStatus).Find(&results).Error

	return
}

// GetBatchFromBaStatus 批量查找
func (obj *_RBlacklistAppMgr) GetBatchFromBaStatus(baStatuss []int8) (results []*RBlacklistApp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistApp{}).Where("`ba_status` IN (?)", baStatuss).Find(&results).Error

	return
}

// GetFromBaItID 通过ba_it_id获取内容
func (obj *_RBlacklistAppMgr) GetFromBaItID(baItID int) (results []*RBlacklistApp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistApp{}).Where("`ba_it_id` = ?", baItID).Find(&results).Error

	return
}

// GetBatchFromBaItID 批量查找
func (obj *_RBlacklistAppMgr) GetBatchFromBaItID(baItIDs []int) (results []*RBlacklistApp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistApp{}).Where("`ba_it_id` IN (?)", baItIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_RBlacklistAppMgr) FetchByPrimaryKey(baID uint32) (result RBlacklistApp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistApp{}).Where("`ba_id` = ?", baID).Find(&result).Error

	return
}

// FetchIndexByBaStatus  获取多个内容
func (obj *_RBlacklistAppMgr) FetchIndexByBaStatus(baStatus int8) (results []*RBlacklistApp, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistApp{}).Where("`ba_status` = ?", baStatus).Find(&results).Error

	return
}
