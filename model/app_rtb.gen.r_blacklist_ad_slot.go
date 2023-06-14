package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _RBlacklistAdSlotMgr struct {
	*_BaseMgr
}

// RBlacklistAdSlotMgr open func
func RBlacklistAdSlotMgr(db *gorm.DB) *_RBlacklistAdSlotMgr {
	if db == nil {
		panic(fmt.Errorf("RBlacklistAdSlotMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_RBlacklistAdSlotMgr{_BaseMgr: &_BaseMgr{DB: db.Table("r_blacklist_ad_slot"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_RBlacklistAdSlotMgr) GetTableName() string {
	return "r_blacklist_ad_slot"
}

// Reset 重置gorm会话
func (obj *_RBlacklistAdSlotMgr) Reset() *_RBlacklistAdSlotMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_RBlacklistAdSlotMgr) Get() (result RBlacklistAdSlot, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistAdSlot{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_RBlacklistAdSlotMgr) Gets() (results []*RBlacklistAdSlot, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistAdSlot{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_RBlacklistAdSlotMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(RBlacklistAdSlot{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithBasID bas_id获取
func (obj *_RBlacklistAdSlotMgr) WithBasID(basID uint32) Option {
	return optionFunc(func(o *options) { o.query["bas_id"] = basID })
}

// WithBasBundle bas_bundle获取
func (obj *_RBlacklistAdSlotMgr) WithBasBundle(basBundle string) Option {
	return optionFunc(func(o *options) { o.query["bas_bundle"] = basBundle })
}

// WithBasWidth bas_width获取
func (obj *_RBlacklistAdSlotMgr) WithBasWidth(basWidth int) Option {
	return optionFunc(func(o *options) { o.query["bas_width"] = basWidth })
}

// WithBasHeight bas_height获取
func (obj *_RBlacklistAdSlotMgr) WithBasHeight(basHeight int) Option {
	return optionFunc(func(o *options) { o.query["bas_height"] = basHeight })
}

// WithBasStatus bas_status获取
func (obj *_RBlacklistAdSlotMgr) WithBasStatus(basStatus int8) Option {
	return optionFunc(func(o *options) { o.query["bas_status"] = basStatus })
}

// WithBasIstID bas_ist_id获取
func (obj *_RBlacklistAdSlotMgr) WithBasIstID(basIstID int) Option {
	return optionFunc(func(o *options) { o.query["bas_ist_id"] = basIstID })
}

// GetByOption 功能选项模式获取
func (obj *_RBlacklistAdSlotMgr) GetByOption(opts ...Option) (result RBlacklistAdSlot, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistAdSlot{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_RBlacklistAdSlotMgr) GetByOptions(opts ...Option) (results []*RBlacklistAdSlot, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistAdSlot{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromBasID 通过bas_id获取内容
func (obj *_RBlacklistAdSlotMgr) GetFromBasID(basID uint32) (result RBlacklistAdSlot, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistAdSlot{}).Where("`bas_id` = ?", basID).Find(&result).Error

	return
}

// GetBatchFromBasID 批量查找
func (obj *_RBlacklistAdSlotMgr) GetBatchFromBasID(basIDs []uint32) (results []*RBlacklistAdSlot, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistAdSlot{}).Where("`bas_id` IN (?)", basIDs).Find(&results).Error

	return
}

// GetFromBasBundle 通过bas_bundle获取内容
func (obj *_RBlacklistAdSlotMgr) GetFromBasBundle(basBundle string) (results []*RBlacklistAdSlot, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistAdSlot{}).Where("`bas_bundle` = ?", basBundle).Find(&results).Error

	return
}

// GetBatchFromBasBundle 批量查找
func (obj *_RBlacklistAdSlotMgr) GetBatchFromBasBundle(basBundles []string) (results []*RBlacklistAdSlot, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistAdSlot{}).Where("`bas_bundle` IN (?)", basBundles).Find(&results).Error

	return
}

// GetFromBasWidth 通过bas_width获取内容
func (obj *_RBlacklistAdSlotMgr) GetFromBasWidth(basWidth int) (results []*RBlacklistAdSlot, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistAdSlot{}).Where("`bas_width` = ?", basWidth).Find(&results).Error

	return
}

// GetBatchFromBasWidth 批量查找
func (obj *_RBlacklistAdSlotMgr) GetBatchFromBasWidth(basWidths []int) (results []*RBlacklistAdSlot, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistAdSlot{}).Where("`bas_width` IN (?)", basWidths).Find(&results).Error

	return
}

// GetFromBasHeight 通过bas_height获取内容
func (obj *_RBlacklistAdSlotMgr) GetFromBasHeight(basHeight int) (results []*RBlacklistAdSlot, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistAdSlot{}).Where("`bas_height` = ?", basHeight).Find(&results).Error

	return
}

// GetBatchFromBasHeight 批量查找
func (obj *_RBlacklistAdSlotMgr) GetBatchFromBasHeight(basHeights []int) (results []*RBlacklistAdSlot, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistAdSlot{}).Where("`bas_height` IN (?)", basHeights).Find(&results).Error

	return
}

// GetFromBasStatus 通过bas_status获取内容
func (obj *_RBlacklistAdSlotMgr) GetFromBasStatus(basStatus int8) (results []*RBlacklistAdSlot, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistAdSlot{}).Where("`bas_status` = ?", basStatus).Find(&results).Error

	return
}

// GetBatchFromBasStatus 批量查找
func (obj *_RBlacklistAdSlotMgr) GetBatchFromBasStatus(basStatuss []int8) (results []*RBlacklistAdSlot, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistAdSlot{}).Where("`bas_status` IN (?)", basStatuss).Find(&results).Error

	return
}

// GetFromBasIstID 通过bas_ist_id获取内容
func (obj *_RBlacklistAdSlotMgr) GetFromBasIstID(basIstID int) (results []*RBlacklistAdSlot, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistAdSlot{}).Where("`bas_ist_id` = ?", basIstID).Find(&results).Error

	return
}

// GetBatchFromBasIstID 批量查找
func (obj *_RBlacklistAdSlotMgr) GetBatchFromBasIstID(basIstIDs []int) (results []*RBlacklistAdSlot, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistAdSlot{}).Where("`bas_ist_id` IN (?)", basIstIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_RBlacklistAdSlotMgr) FetchByPrimaryKey(basID uint32) (result RBlacklistAdSlot, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistAdSlot{}).Where("`bas_id` = ?", basID).Find(&result).Error

	return
}

// FetchIndexByBasBundle  获取多个内容
func (obj *_RBlacklistAdSlotMgr) FetchIndexByBasBundle(basBundle string) (results []*RBlacklistAdSlot, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistAdSlot{}).Where("`bas_bundle` = ?", basBundle).Find(&results).Error

	return
}

// FetchIndexByBasWidth  获取多个内容
func (obj *_RBlacklistAdSlotMgr) FetchIndexByBasWidth(basWidth int) (results []*RBlacklistAdSlot, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistAdSlot{}).Where("`bas_width` = ?", basWidth).Find(&results).Error

	return
}

// FetchIndexByBasHeight  获取多个内容
func (obj *_RBlacklistAdSlotMgr) FetchIndexByBasHeight(basHeight int) (results []*RBlacklistAdSlot, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistAdSlot{}).Where("`bas_height` = ?", basHeight).Find(&results).Error

	return
}

// FetchIndexByBasStatus  获取多个内容
func (obj *_RBlacklistAdSlotMgr) FetchIndexByBasStatus(basStatus int8) (results []*RBlacklistAdSlot, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RBlacklistAdSlot{}).Where("`bas_status` = ?", basStatus).Find(&results).Error

	return
}
