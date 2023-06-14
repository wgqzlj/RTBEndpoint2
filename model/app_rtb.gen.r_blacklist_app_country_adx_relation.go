package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _RBlacklistAppCountryAdxRelationMgr struct {
	*_BaseMgr
}

// RBlacklistAppCountryAdxRelationMgr open func
func RBlacklistAppCountryAdxRelationMgr(db *gorm.DB) *_RBlacklistAppCountryAdxRelationMgr {
	if db == nil {
		panic(fmt.Errorf("RBlacklistAppCountryAdxRelationMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_RBlacklistAppCountryAdxRelationMgr{_BaseMgr: &_BaseMgr{DB: db.Table("r_blacklist_app_country_adx_relation"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_RBlacklistAppCountryAdxRelationMgr) GetTableName() string {
	return "r_blacklist_app_country_adx_relation"
}

// Reset 重置gorm会话
func (obj *_RBlacklistAppCountryAdxRelationMgr) Reset() *_RBlacklistAppCountryAdxRelationMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_RBlacklistAppCountryAdxRelationMgr) Get() (result RBlacklistAppCountryAdxRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistAppCountryAdxRelation{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_RBlacklistAppCountryAdxRelationMgr) Gets() (results []*RBlacklistAppCountryAdxRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistAppCountryAdxRelation{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_RBlacklistAppCountryAdxRelationMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(RBlacklistAppCountryAdxRelation{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithBacarID bacar_id获取
func (obj *_RBlacklistAppCountryAdxRelationMgr) WithBacarID(bacarID uint32) Option {
	return optionFunc(func(o *options) { o.query["bacar_id"] = bacarID })
}

// WithBacarBaID bacar_ba_id获取
func (obj *_RBlacklistAppCountryAdxRelationMgr) WithBacarBaID(bacarBaID int) Option {
	return optionFunc(func(o *options) { o.query["bacar_ba_id"] = bacarBaID })
}

// WithBacarCID bacar_c_id获取
func (obj *_RBlacklistAppCountryAdxRelationMgr) WithBacarCID(bacarCID int) Option {
	return optionFunc(func(o *options) { o.query["bacar_c_id"] = bacarCID })
}

// WithBacarAeID bacar_ae_id获取
func (obj *_RBlacklistAppCountryAdxRelationMgr) WithBacarAeID(bacarAeID int) Option {
	return optionFunc(func(o *options) { o.query["bacar_ae_id"] = bacarAeID })
}

// WithBacarStatus bacar_status获取
func (obj *_RBlacklistAppCountryAdxRelationMgr) WithBacarStatus(bacarStatus int8) Option {
	return optionFunc(func(o *options) { o.query["bacar_status"] = bacarStatus })
}

// WithBacarCreateTime bacar_create_time获取
func (obj *_RBlacklistAppCountryAdxRelationMgr) WithBacarCreateTime(bacarCreateTime int) Option {
	return optionFunc(func(o *options) { o.query["bacar_create_time"] = bacarCreateTime })
}

// GetByOption 功能选项模式获取
func (obj *_RBlacklistAppCountryAdxRelationMgr) GetByOption(opts ...Option) (result RBlacklistAppCountryAdxRelation, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistAppCountryAdxRelation{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_RBlacklistAppCountryAdxRelationMgr) GetByOptions(opts ...Option) (results []*RBlacklistAppCountryAdxRelation, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistAppCountryAdxRelation{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromBacarID 通过bacar_id获取内容
func (obj *_RBlacklistAppCountryAdxRelationMgr) GetFromBacarID(bacarID uint32) (result RBlacklistAppCountryAdxRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistAppCountryAdxRelation{}).Where("`bacar_id` = ?", bacarID).Find(&result).Error

	return
}

// GetBatchFromBacarID 批量查找
func (obj *_RBlacklistAppCountryAdxRelationMgr) GetBatchFromBacarID(bacarIDs []uint32) (results []*RBlacklistAppCountryAdxRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistAppCountryAdxRelation{}).Where("`bacar_id` IN (?)", bacarIDs).Find(&results).Error

	return
}

// GetFromBacarBaID 通过bacar_ba_id获取内容
func (obj *_RBlacklistAppCountryAdxRelationMgr) GetFromBacarBaID(bacarBaID int) (results []*RBlacklistAppCountryAdxRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistAppCountryAdxRelation{}).Where("`bacar_ba_id` = ?", bacarBaID).Find(&results).Error

	return
}

// GetBatchFromBacarBaID 批量查找
func (obj *_RBlacklistAppCountryAdxRelationMgr) GetBatchFromBacarBaID(bacarBaIDs []int) (results []*RBlacklistAppCountryAdxRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistAppCountryAdxRelation{}).Where("`bacar_ba_id` IN (?)", bacarBaIDs).Find(&results).Error

	return
}

// GetFromBacarCID 通过bacar_c_id获取内容
func (obj *_RBlacklistAppCountryAdxRelationMgr) GetFromBacarCID(bacarCID int) (results []*RBlacklistAppCountryAdxRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistAppCountryAdxRelation{}).Where("`bacar_c_id` = ?", bacarCID).Find(&results).Error

	return
}

// GetBatchFromBacarCID 批量查找
func (obj *_RBlacklistAppCountryAdxRelationMgr) GetBatchFromBacarCID(bacarCIDs []int) (results []*RBlacklistAppCountryAdxRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistAppCountryAdxRelation{}).Where("`bacar_c_id` IN (?)", bacarCIDs).Find(&results).Error

	return
}

// GetFromBacarAeID 通过bacar_ae_id获取内容
func (obj *_RBlacklistAppCountryAdxRelationMgr) GetFromBacarAeID(bacarAeID int) (results []*RBlacklistAppCountryAdxRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistAppCountryAdxRelation{}).Where("`bacar_ae_id` = ?", bacarAeID).Find(&results).Error

	return
}

// GetBatchFromBacarAeID 批量查找
func (obj *_RBlacklistAppCountryAdxRelationMgr) GetBatchFromBacarAeID(bacarAeIDs []int) (results []*RBlacklistAppCountryAdxRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistAppCountryAdxRelation{}).Where("`bacar_ae_id` IN (?)", bacarAeIDs).Find(&results).Error

	return
}

// GetFromBacarStatus 通过bacar_status获取内容
func (obj *_RBlacklistAppCountryAdxRelationMgr) GetFromBacarStatus(bacarStatus int8) (results []*RBlacklistAppCountryAdxRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistAppCountryAdxRelation{}).Where("`bacar_status` = ?", bacarStatus).Find(&results).Error

	return
}

// GetBatchFromBacarStatus 批量查找
func (obj *_RBlacklistAppCountryAdxRelationMgr) GetBatchFromBacarStatus(bacarStatuss []int8) (results []*RBlacklistAppCountryAdxRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistAppCountryAdxRelation{}).Where("`bacar_status` IN (?)", bacarStatuss).Find(&results).Error

	return
}

// GetFromBacarCreateTime 通过bacar_create_time获取内容
func (obj *_RBlacklistAppCountryAdxRelationMgr) GetFromBacarCreateTime(bacarCreateTime int) (results []*RBlacklistAppCountryAdxRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistAppCountryAdxRelation{}).Where("`bacar_create_time` = ?", bacarCreateTime).Find(&results).Error

	return
}

// GetBatchFromBacarCreateTime 批量查找
func (obj *_RBlacklistAppCountryAdxRelationMgr) GetBatchFromBacarCreateTime(bacarCreateTimes []int) (results []*RBlacklistAppCountryAdxRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistAppCountryAdxRelation{}).Where("`bacar_create_time` IN (?)", bacarCreateTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_RBlacklistAppCountryAdxRelationMgr) FetchByPrimaryKey(bacarID uint32) (result RBlacklistAppCountryAdxRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistAppCountryAdxRelation{}).Where("`bacar_id` = ?", bacarID).Find(&result).Error

	return
}

// FetchIndexByBacarBaID  获取多个内容
func (obj *_RBlacklistAppCountryAdxRelationMgr) FetchIndexByBacarBaID(bacarBaID int) (results []*RBlacklistAppCountryAdxRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistAppCountryAdxRelation{}).Where("`bacar_ba_id` = ?", bacarBaID).Find(&results).Error

	return
}

// FetchIndexByBacarCID  获取多个内容
func (obj *_RBlacklistAppCountryAdxRelationMgr) FetchIndexByBacarCID(bacarCID int) (results []*RBlacklistAppCountryAdxRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistAppCountryAdxRelation{}).Where("`bacar_c_id` = ?", bacarCID).Find(&results).Error

	return
}

// FetchIndexByBacarAeID  获取多个内容
func (obj *_RBlacklistAppCountryAdxRelationMgr) FetchIndexByBacarAeID(bacarAeID int) (results []*RBlacklistAppCountryAdxRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistAppCountryAdxRelation{}).Where("`bacar_ae_id` = ?", bacarAeID).Find(&results).Error

	return
}

// FetchIndexByBacarStatus  获取多个内容
func (obj *_RBlacklistAppCountryAdxRelationMgr) FetchIndexByBacarStatus(bacarStatus int8) (results []*RBlacklistAppCountryAdxRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistAppCountryAdxRelation{}).Where("`bacar_status` = ?", bacarStatus).Find(&results).Error

	return
}

// FetchIndexByBacarCreateTime  获取多个内容
func (obj *_RBlacklistAppCountryAdxRelationMgr) FetchIndexByBacarCreateTime(bacarCreateTime int) (results []*RBlacklistAppCountryAdxRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistAppCountryAdxRelation{}).Where("`bacar_create_time` = ?", bacarCreateTime).Find(&results).Error

	return
}
