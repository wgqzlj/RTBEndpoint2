package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _RBlacklistAdSlotCountryAdxRelationMgr struct {
	*_BaseMgr
}

// RBlacklistAdSlotCountryAdxRelationMgr open func
func RBlacklistAdSlotCountryAdxRelationMgr(db *gorm.DB) *_RBlacklistAdSlotCountryAdxRelationMgr {
	if db == nil {
		panic(fmt.Errorf("RBlacklistAdSlotCountryAdxRelationMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_RBlacklistAdSlotCountryAdxRelationMgr{_BaseMgr: &_BaseMgr{DB: db.Table("r_blacklist_ad_slot_country_adx_relation"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_RBlacklistAdSlotCountryAdxRelationMgr) GetTableName() string {
	return "r_blacklist_ad_slot_country_adx_relation"
}

// Reset 重置gorm会话
func (obj *_RBlacklistAdSlotCountryAdxRelationMgr) Reset() *_RBlacklistAdSlotCountryAdxRelationMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_RBlacklistAdSlotCountryAdxRelationMgr) Get() (result RBlacklistAdSlotCountryAdxRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistAdSlotCountryAdxRelation{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_RBlacklistAdSlotCountryAdxRelationMgr) Gets() (results []*RBlacklistAdSlotCountryAdxRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistAdSlotCountryAdxRelation{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_RBlacklistAdSlotCountryAdxRelationMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(RBlacklistAdSlotCountryAdxRelation{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithBascarID bascar_id获取
func (obj *_RBlacklistAdSlotCountryAdxRelationMgr) WithBascarID(bascarID uint32) Option {
	return optionFunc(func(o *options) { o.query["bascar_id"] = bascarID })
}

// WithBascarBasID bascar_bas_id获取
func (obj *_RBlacklistAdSlotCountryAdxRelationMgr) WithBascarBasID(bascarBasID int) Option {
	return optionFunc(func(o *options) { o.query["bascar_bas_id"] = bascarBasID })
}

// WithBascarCID bascar_c_id获取
func (obj *_RBlacklistAdSlotCountryAdxRelationMgr) WithBascarCID(bascarCID int) Option {
	return optionFunc(func(o *options) { o.query["bascar_c_id"] = bascarCID })
}

// WithBascarAeID bascar_ae_id获取
func (obj *_RBlacklistAdSlotCountryAdxRelationMgr) WithBascarAeID(bascarAeID int) Option {
	return optionFunc(func(o *options) { o.query["bascar_ae_id"] = bascarAeID })
}

// WithBascarStatus bascar_status获取
func (obj *_RBlacklistAdSlotCountryAdxRelationMgr) WithBascarStatus(bascarStatus int8) Option {
	return optionFunc(func(o *options) { o.query["bascar_status"] = bascarStatus })
}

// WithBascarCreateTime bascar_create_time获取
func (obj *_RBlacklistAdSlotCountryAdxRelationMgr) WithBascarCreateTime(bascarCreateTime int) Option {
	return optionFunc(func(o *options) { o.query["bascar_create_time"] = bascarCreateTime })
}

// GetByOption 功能选项模式获取
func (obj *_RBlacklistAdSlotCountryAdxRelationMgr) GetByOption(opts ...Option) (result RBlacklistAdSlotCountryAdxRelation, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistAdSlotCountryAdxRelation{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_RBlacklistAdSlotCountryAdxRelationMgr) GetByOptions(opts ...Option) (results []*RBlacklistAdSlotCountryAdxRelation, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistAdSlotCountryAdxRelation{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromBascarID 通过bascar_id获取内容
func (obj *_RBlacklistAdSlotCountryAdxRelationMgr) GetFromBascarID(bascarID uint32) (result RBlacklistAdSlotCountryAdxRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistAdSlotCountryAdxRelation{}).Where("`bascar_id` = ?", bascarID).Find(&result).Error

	return
}

// GetBatchFromBascarID 批量查找
func (obj *_RBlacklistAdSlotCountryAdxRelationMgr) GetBatchFromBascarID(bascarIDs []uint32) (results []*RBlacklistAdSlotCountryAdxRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistAdSlotCountryAdxRelation{}).Where("`bascar_id` IN (?)", bascarIDs).Find(&results).Error

	return
}

// GetFromBascarBasID 通过bascar_bas_id获取内容
func (obj *_RBlacklistAdSlotCountryAdxRelationMgr) GetFromBascarBasID(bascarBasID int) (results []*RBlacklistAdSlotCountryAdxRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistAdSlotCountryAdxRelation{}).Where("`bascar_bas_id` = ?", bascarBasID).Find(&results).Error

	return
}

// GetBatchFromBascarBasID 批量查找
func (obj *_RBlacklistAdSlotCountryAdxRelationMgr) GetBatchFromBascarBasID(bascarBasIDs []int) (results []*RBlacklistAdSlotCountryAdxRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistAdSlotCountryAdxRelation{}).Where("`bascar_bas_id` IN (?)", bascarBasIDs).Find(&results).Error

	return
}

// GetFromBascarCID 通过bascar_c_id获取内容
func (obj *_RBlacklistAdSlotCountryAdxRelationMgr) GetFromBascarCID(bascarCID int) (results []*RBlacklistAdSlotCountryAdxRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistAdSlotCountryAdxRelation{}).Where("`bascar_c_id` = ?", bascarCID).Find(&results).Error

	return
}

// GetBatchFromBascarCID 批量查找
func (obj *_RBlacklistAdSlotCountryAdxRelationMgr) GetBatchFromBascarCID(bascarCIDs []int) (results []*RBlacklistAdSlotCountryAdxRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistAdSlotCountryAdxRelation{}).Where("`bascar_c_id` IN (?)", bascarCIDs).Find(&results).Error

	return
}

// GetFromBascarAeID 通过bascar_ae_id获取内容
func (obj *_RBlacklistAdSlotCountryAdxRelationMgr) GetFromBascarAeID(bascarAeID int) (results []*RBlacklistAdSlotCountryAdxRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistAdSlotCountryAdxRelation{}).Where("`bascar_ae_id` = ?", bascarAeID).Find(&results).Error

	return
}

// GetBatchFromBascarAeID 批量查找
func (obj *_RBlacklistAdSlotCountryAdxRelationMgr) GetBatchFromBascarAeID(bascarAeIDs []int) (results []*RBlacklistAdSlotCountryAdxRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistAdSlotCountryAdxRelation{}).Where("`bascar_ae_id` IN (?)", bascarAeIDs).Find(&results).Error

	return
}

// GetFromBascarStatus 通过bascar_status获取内容
func (obj *_RBlacklistAdSlotCountryAdxRelationMgr) GetFromBascarStatus(bascarStatus int8) (results []*RBlacklistAdSlotCountryAdxRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistAdSlotCountryAdxRelation{}).Where("`bascar_status` = ?", bascarStatus).Find(&results).Error

	return
}

// GetBatchFromBascarStatus 批量查找
func (obj *_RBlacklistAdSlotCountryAdxRelationMgr) GetBatchFromBascarStatus(bascarStatuss []int8) (results []*RBlacklistAdSlotCountryAdxRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistAdSlotCountryAdxRelation{}).Where("`bascar_status` IN (?)", bascarStatuss).Find(&results).Error

	return
}

// GetFromBascarCreateTime 通过bascar_create_time获取内容
func (obj *_RBlacklistAdSlotCountryAdxRelationMgr) GetFromBascarCreateTime(bascarCreateTime int) (results []*RBlacklistAdSlotCountryAdxRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistAdSlotCountryAdxRelation{}).Where("`bascar_create_time` = ?", bascarCreateTime).Find(&results).Error

	return
}

// GetBatchFromBascarCreateTime 批量查找
func (obj *_RBlacklistAdSlotCountryAdxRelationMgr) GetBatchFromBascarCreateTime(bascarCreateTimes []int) (results []*RBlacklistAdSlotCountryAdxRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistAdSlotCountryAdxRelation{}).Where("`bascar_create_time` IN (?)", bascarCreateTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_RBlacklistAdSlotCountryAdxRelationMgr) FetchByPrimaryKey(bascarID uint32) (result RBlacklistAdSlotCountryAdxRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistAdSlotCountryAdxRelation{}).Where("`bascar_id` = ?", bascarID).Find(&result).Error

	return
}

// FetchIndexByBascarBasID  获取多个内容
func (obj *_RBlacklistAdSlotCountryAdxRelationMgr) FetchIndexByBascarBasID(bascarBasID int) (results []*RBlacklistAdSlotCountryAdxRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistAdSlotCountryAdxRelation{}).Where("`bascar_bas_id` = ?", bascarBasID).Find(&results).Error

	return
}

// FetchIndexByBascarCID  获取多个内容
func (obj *_RBlacklistAdSlotCountryAdxRelationMgr) FetchIndexByBascarCID(bascarCID int) (results []*RBlacklistAdSlotCountryAdxRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistAdSlotCountryAdxRelation{}).Where("`bascar_c_id` = ?", bascarCID).Find(&results).Error

	return
}

// FetchIndexByBascarAeID  获取多个内容
func (obj *_RBlacklistAdSlotCountryAdxRelationMgr) FetchIndexByBascarAeID(bascarAeID int) (results []*RBlacklistAdSlotCountryAdxRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistAdSlotCountryAdxRelation{}).Where("`bascar_ae_id` = ?", bascarAeID).Find(&results).Error

	return
}
