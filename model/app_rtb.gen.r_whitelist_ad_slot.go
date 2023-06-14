package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _RWhitelistAdSlotMgr struct {
	*_BaseMgr
}

// RWhitelistAdSlotMgr open func
func RWhitelistAdSlotMgr(db *gorm.DB) *_RWhitelistAdSlotMgr {
	if db == nil {
		panic(fmt.Errorf("RWhitelistAdSlotMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_RWhitelistAdSlotMgr{_BaseMgr: &_BaseMgr{DB: db.Table("r_whitelist_ad_slot"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_RWhitelistAdSlotMgr) GetTableName() string {
	return "r_whitelist_ad_slot"
}

// Reset 重置gorm会话
func (obj *_RWhitelistAdSlotMgr) Reset() *_RWhitelistAdSlotMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_RWhitelistAdSlotMgr) Get() (result RWhitelistAdSlot, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RWhitelistAdSlot{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_RWhitelistAdSlotMgr) Gets() (results []*RWhitelistAdSlot, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RWhitelistAdSlot{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_RWhitelistAdSlotMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(RWhitelistAdSlot{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithWasID was_id获取
func (obj *_RWhitelistAdSlotMgr) WithWasID(wasID uint32) Option {
	return optionFunc(func(o *options) { o.query["was_id"] = wasID })
}

// WithWasBundle was_bundle获取
func (obj *_RWhitelistAdSlotMgr) WithWasBundle(wasBundle string) Option {
	return optionFunc(func(o *options) { o.query["was_bundle"] = wasBundle })
}

// WithWasWidth was_width获取
func (obj *_RWhitelistAdSlotMgr) WithWasWidth(wasWidth int) Option {
	return optionFunc(func(o *options) { o.query["was_width"] = wasWidth })
}

// WithWasHeight was_height获取
func (obj *_RWhitelistAdSlotMgr) WithWasHeight(wasHeight int) Option {
	return optionFunc(func(o *options) { o.query["was_height"] = wasHeight })
}

// WithWasPrice was_price获取
func (obj *_RWhitelistAdSlotMgr) WithWasPrice(wasPrice float32) Option {
	return optionFunc(func(o *options) { o.query["was_price"] = wasPrice })
}

// WithWasStatus was_status获取
func (obj *_RWhitelistAdSlotMgr) WithWasStatus(wasStatus int8) Option {
	return optionFunc(func(o *options) { o.query["was_status"] = wasStatus })
}

// GetByOption 功能选项模式获取
func (obj *_RWhitelistAdSlotMgr) GetByOption(opts ...Option) (result RWhitelistAdSlot, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(RWhitelistAdSlot{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_RWhitelistAdSlotMgr) GetByOptions(opts ...Option) (results []*RWhitelistAdSlot, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(RWhitelistAdSlot{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromWasID 通过was_id获取内容
func (obj *_RWhitelistAdSlotMgr) GetFromWasID(wasID uint32) (result RWhitelistAdSlot, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RWhitelistAdSlot{}).Where("`was_id` = ?", wasID).Find(&result).Error

	return
}

// GetBatchFromWasID 批量查找
func (obj *_RWhitelistAdSlotMgr) GetBatchFromWasID(wasIDs []uint32) (results []*RWhitelistAdSlot, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RWhitelistAdSlot{}).Where("`was_id` IN (?)", wasIDs).Find(&results).Error

	return
}

// GetFromWasBundle 通过was_bundle获取内容
func (obj *_RWhitelistAdSlotMgr) GetFromWasBundle(wasBundle string) (results []*RWhitelistAdSlot, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RWhitelistAdSlot{}).Where("`was_bundle` = ?", wasBundle).Find(&results).Error

	return
}

// GetBatchFromWasBundle 批量查找
func (obj *_RWhitelistAdSlotMgr) GetBatchFromWasBundle(wasBundles []string) (results []*RWhitelistAdSlot, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RWhitelistAdSlot{}).Where("`was_bundle` IN (?)", wasBundles).Find(&results).Error

	return
}

// GetFromWasWidth 通过was_width获取内容
func (obj *_RWhitelistAdSlotMgr) GetFromWasWidth(wasWidth int) (results []*RWhitelistAdSlot, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RWhitelistAdSlot{}).Where("`was_width` = ?", wasWidth).Find(&results).Error

	return
}

// GetBatchFromWasWidth 批量查找
func (obj *_RWhitelistAdSlotMgr) GetBatchFromWasWidth(wasWidths []int) (results []*RWhitelistAdSlot, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RWhitelistAdSlot{}).Where("`was_width` IN (?)", wasWidths).Find(&results).Error

	return
}

// GetFromWasHeight 通过was_height获取内容
func (obj *_RWhitelistAdSlotMgr) GetFromWasHeight(wasHeight int) (results []*RWhitelistAdSlot, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RWhitelistAdSlot{}).Where("`was_height` = ?", wasHeight).Find(&results).Error

	return
}

// GetBatchFromWasHeight 批量查找
func (obj *_RWhitelistAdSlotMgr) GetBatchFromWasHeight(wasHeights []int) (results []*RWhitelistAdSlot, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RWhitelistAdSlot{}).Where("`was_height` IN (?)", wasHeights).Find(&results).Error

	return
}

// GetFromWasPrice 通过was_price获取内容
func (obj *_RWhitelistAdSlotMgr) GetFromWasPrice(wasPrice float32) (results []*RWhitelistAdSlot, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RWhitelistAdSlot{}).Where("`was_price` = ?", wasPrice).Find(&results).Error

	return
}

// GetBatchFromWasPrice 批量查找
func (obj *_RWhitelistAdSlotMgr) GetBatchFromWasPrice(wasPrices []float32) (results []*RWhitelistAdSlot, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RWhitelistAdSlot{}).Where("`was_price` IN (?)", wasPrices).Find(&results).Error

	return
}

// GetFromWasStatus 通过was_status获取内容
func (obj *_RWhitelistAdSlotMgr) GetFromWasStatus(wasStatus int8) (results []*RWhitelistAdSlot, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RWhitelistAdSlot{}).Where("`was_status` = ?", wasStatus).Find(&results).Error

	return
}

// GetBatchFromWasStatus 批量查找
func (obj *_RWhitelistAdSlotMgr) GetBatchFromWasStatus(wasStatuss []int8) (results []*RWhitelistAdSlot, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RWhitelistAdSlot{}).Where("`was_status` IN (?)", wasStatuss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_RWhitelistAdSlotMgr) FetchByPrimaryKey(wasID uint32) (result RWhitelistAdSlot, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RWhitelistAdSlot{}).Where("`was_id` = ?", wasID).Find(&result).Error

	return
}
