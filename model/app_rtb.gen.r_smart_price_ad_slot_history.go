package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _RSmartPriceAdSlotHistoryMgr struct {
	*_BaseMgr
}

// RSmartPriceAdSlotHistoryMgr open func
func RSmartPriceAdSlotHistoryMgr(db *gorm.DB) *_RSmartPriceAdSlotHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("RSmartPriceAdSlotHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_RSmartPriceAdSlotHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("r_smart_price_ad_slot_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_RSmartPriceAdSlotHistoryMgr) GetTableName() string {
	return "r_smart_price_ad_slot_history"
}

// Reset 重置gorm会话
func (obj *_RSmartPriceAdSlotHistoryMgr) Reset() *_RSmartPriceAdSlotHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_RSmartPriceAdSlotHistoryMgr) Get() (result RSmartPriceAdSlotHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAdSlotHistory{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_RSmartPriceAdSlotHistoryMgr) Gets() (results []*RSmartPriceAdSlotHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAdSlotHistory{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_RSmartPriceAdSlotHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(RSmartPriceAdSlotHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSpashID spash_id获取
func (obj *_RSmartPriceAdSlotHistoryMgr) WithSpashID(spashID uint32) Option {
	return optionFunc(func(o *options) { o.query["spash_id"] = spashID })
}

// WithSpashSpasID spash_spas_id获取
func (obj *_RSmartPriceAdSlotHistoryMgr) WithSpashSpasID(spashSpasID int) Option {
	return optionFunc(func(o *options) { o.query["spash_spas_id"] = spashSpasID })
}

// WithSpashAeID spash_ae_id获取
func (obj *_RSmartPriceAdSlotHistoryMgr) WithSpashAeID(spashAeID int) Option {
	return optionFunc(func(o *options) { o.query["spash_ae_id"] = spashAeID })
}

// WithSpashCID spash_c_id获取
func (obj *_RSmartPriceAdSlotHistoryMgr) WithSpashCID(spashCID int) Option {
	return optionFunc(func(o *options) { o.query["spash_c_id"] = spashCID })
}

// WithSpashTime spash_time获取
func (obj *_RSmartPriceAdSlotHistoryMgr) WithSpashTime(spashTime int) Option {
	return optionFunc(func(o *options) { o.query["spash_time"] = spashTime })
}

// WithSpashBidCount spash_bid_count获取
func (obj *_RSmartPriceAdSlotHistoryMgr) WithSpashBidCount(spashBidCount int) Option {
	return optionFunc(func(o *options) { o.query["spash_bid_count"] = spashBidCount })
}

// WithSpashImpCount spash_imp_count获取
func (obj *_RSmartPriceAdSlotHistoryMgr) WithSpashImpCount(spashImpCount int) Option {
	return optionFunc(func(o *options) { o.query["spash_imp_count"] = spashImpCount })
}

// WithSpashClickCount spash_click_count获取
func (obj *_RSmartPriceAdSlotHistoryMgr) WithSpashClickCount(spashClickCount int) Option {
	return optionFunc(func(o *options) { o.query["spash_click_count"] = spashClickCount })
}

// WithSpashPostbackCount spash_postback_count获取
func (obj *_RSmartPriceAdSlotHistoryMgr) WithSpashPostbackCount(spashPostbackCount int) Option {
	return optionFunc(func(o *options) { o.query["spash_postback_count"] = spashPostbackCount })
}

// WithSpashCost spash_cost获取
func (obj *_RSmartPriceAdSlotHistoryMgr) WithSpashCost(spashCost float32) Option {
	return optionFunc(func(o *options) { o.query["spash_cost"] = spashCost })
}

// WithSpashRoi spash_roi获取
func (obj *_RSmartPriceAdSlotHistoryMgr) WithSpashRoi(spashRoi float32) Option {
	return optionFunc(func(o *options) { o.query["spash_roi"] = spashRoi })
}

// WithSpashEcpm spash_ecpm获取
func (obj *_RSmartPriceAdSlotHistoryMgr) WithSpashEcpm(spashEcpm float32) Option {
	return optionFunc(func(o *options) { o.query["spash_ecpm"] = spashEcpm })
}

// GetByOption 功能选项模式获取
func (obj *_RSmartPriceAdSlotHistoryMgr) GetByOption(opts ...Option) (result RSmartPriceAdSlotHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAdSlotHistory{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_RSmartPriceAdSlotHistoryMgr) GetByOptions(opts ...Option) (results []*RSmartPriceAdSlotHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAdSlotHistory{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromSpashID 通过spash_id获取内容
func (obj *_RSmartPriceAdSlotHistoryMgr) GetFromSpashID(spashID uint32) (result RSmartPriceAdSlotHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAdSlotHistory{}).Where("`spash_id` = ?", spashID).Find(&result).Error

	return
}

// GetBatchFromSpashID 批量查找
func (obj *_RSmartPriceAdSlotHistoryMgr) GetBatchFromSpashID(spashIDs []uint32) (results []*RSmartPriceAdSlotHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAdSlotHistory{}).Where("`spash_id` IN (?)", spashIDs).Find(&results).Error

	return
}

// GetFromSpashSpasID 通过spash_spas_id获取内容
func (obj *_RSmartPriceAdSlotHistoryMgr) GetFromSpashSpasID(spashSpasID int) (results []*RSmartPriceAdSlotHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAdSlotHistory{}).Where("`spash_spas_id` = ?", spashSpasID).Find(&results).Error

	return
}

// GetBatchFromSpashSpasID 批量查找
func (obj *_RSmartPriceAdSlotHistoryMgr) GetBatchFromSpashSpasID(spashSpasIDs []int) (results []*RSmartPriceAdSlotHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAdSlotHistory{}).Where("`spash_spas_id` IN (?)", spashSpasIDs).Find(&results).Error

	return
}

// GetFromSpashAeID 通过spash_ae_id获取内容
func (obj *_RSmartPriceAdSlotHistoryMgr) GetFromSpashAeID(spashAeID int) (results []*RSmartPriceAdSlotHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAdSlotHistory{}).Where("`spash_ae_id` = ?", spashAeID).Find(&results).Error

	return
}

// GetBatchFromSpashAeID 批量查找
func (obj *_RSmartPriceAdSlotHistoryMgr) GetBatchFromSpashAeID(spashAeIDs []int) (results []*RSmartPriceAdSlotHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAdSlotHistory{}).Where("`spash_ae_id` IN (?)", spashAeIDs).Find(&results).Error

	return
}

// GetFromSpashCID 通过spash_c_id获取内容
func (obj *_RSmartPriceAdSlotHistoryMgr) GetFromSpashCID(spashCID int) (results []*RSmartPriceAdSlotHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAdSlotHistory{}).Where("`spash_c_id` = ?", spashCID).Find(&results).Error

	return
}

// GetBatchFromSpashCID 批量查找
func (obj *_RSmartPriceAdSlotHistoryMgr) GetBatchFromSpashCID(spashCIDs []int) (results []*RSmartPriceAdSlotHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAdSlotHistory{}).Where("`spash_c_id` IN (?)", spashCIDs).Find(&results).Error

	return
}

// GetFromSpashTime 通过spash_time获取内容
func (obj *_RSmartPriceAdSlotHistoryMgr) GetFromSpashTime(spashTime int) (results []*RSmartPriceAdSlotHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAdSlotHistory{}).Where("`spash_time` = ?", spashTime).Find(&results).Error

	return
}

// GetBatchFromSpashTime 批量查找
func (obj *_RSmartPriceAdSlotHistoryMgr) GetBatchFromSpashTime(spashTimes []int) (results []*RSmartPriceAdSlotHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAdSlotHistory{}).Where("`spash_time` IN (?)", spashTimes).Find(&results).Error

	return
}

// GetFromSpashBidCount 通过spash_bid_count获取内容
func (obj *_RSmartPriceAdSlotHistoryMgr) GetFromSpashBidCount(spashBidCount int) (results []*RSmartPriceAdSlotHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAdSlotHistory{}).Where("`spash_bid_count` = ?", spashBidCount).Find(&results).Error

	return
}

// GetBatchFromSpashBidCount 批量查找
func (obj *_RSmartPriceAdSlotHistoryMgr) GetBatchFromSpashBidCount(spashBidCounts []int) (results []*RSmartPriceAdSlotHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAdSlotHistory{}).Where("`spash_bid_count` IN (?)", spashBidCounts).Find(&results).Error

	return
}

// GetFromSpashImpCount 通过spash_imp_count获取内容
func (obj *_RSmartPriceAdSlotHistoryMgr) GetFromSpashImpCount(spashImpCount int) (results []*RSmartPriceAdSlotHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAdSlotHistory{}).Where("`spash_imp_count` = ?", spashImpCount).Find(&results).Error

	return
}

// GetBatchFromSpashImpCount 批量查找
func (obj *_RSmartPriceAdSlotHistoryMgr) GetBatchFromSpashImpCount(spashImpCounts []int) (results []*RSmartPriceAdSlotHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAdSlotHistory{}).Where("`spash_imp_count` IN (?)", spashImpCounts).Find(&results).Error

	return
}

// GetFromSpashClickCount 通过spash_click_count获取内容
func (obj *_RSmartPriceAdSlotHistoryMgr) GetFromSpashClickCount(spashClickCount int) (results []*RSmartPriceAdSlotHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAdSlotHistory{}).Where("`spash_click_count` = ?", spashClickCount).Find(&results).Error

	return
}

// GetBatchFromSpashClickCount 批量查找
func (obj *_RSmartPriceAdSlotHistoryMgr) GetBatchFromSpashClickCount(spashClickCounts []int) (results []*RSmartPriceAdSlotHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAdSlotHistory{}).Where("`spash_click_count` IN (?)", spashClickCounts).Find(&results).Error

	return
}

// GetFromSpashPostbackCount 通过spash_postback_count获取内容
func (obj *_RSmartPriceAdSlotHistoryMgr) GetFromSpashPostbackCount(spashPostbackCount int) (results []*RSmartPriceAdSlotHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAdSlotHistory{}).Where("`spash_postback_count` = ?", spashPostbackCount).Find(&results).Error

	return
}

// GetBatchFromSpashPostbackCount 批量查找
func (obj *_RSmartPriceAdSlotHistoryMgr) GetBatchFromSpashPostbackCount(spashPostbackCounts []int) (results []*RSmartPriceAdSlotHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAdSlotHistory{}).Where("`spash_postback_count` IN (?)", spashPostbackCounts).Find(&results).Error

	return
}

// GetFromSpashCost 通过spash_cost获取内容
func (obj *_RSmartPriceAdSlotHistoryMgr) GetFromSpashCost(spashCost float32) (results []*RSmartPriceAdSlotHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAdSlotHistory{}).Where("`spash_cost` = ?", spashCost).Find(&results).Error

	return
}

// GetBatchFromSpashCost 批量查找
func (obj *_RSmartPriceAdSlotHistoryMgr) GetBatchFromSpashCost(spashCosts []float32) (results []*RSmartPriceAdSlotHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAdSlotHistory{}).Where("`spash_cost` IN (?)", spashCosts).Find(&results).Error

	return
}

// GetFromSpashRoi 通过spash_roi获取内容
func (obj *_RSmartPriceAdSlotHistoryMgr) GetFromSpashRoi(spashRoi float32) (results []*RSmartPriceAdSlotHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAdSlotHistory{}).Where("`spash_roi` = ?", spashRoi).Find(&results).Error

	return
}

// GetBatchFromSpashRoi 批量查找
func (obj *_RSmartPriceAdSlotHistoryMgr) GetBatchFromSpashRoi(spashRois []float32) (results []*RSmartPriceAdSlotHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAdSlotHistory{}).Where("`spash_roi` IN (?)", spashRois).Find(&results).Error

	return
}

// GetFromSpashEcpm 通过spash_ecpm获取内容
func (obj *_RSmartPriceAdSlotHistoryMgr) GetFromSpashEcpm(spashEcpm float32) (results []*RSmartPriceAdSlotHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAdSlotHistory{}).Where("`spash_ecpm` = ?", spashEcpm).Find(&results).Error

	return
}

// GetBatchFromSpashEcpm 批量查找
func (obj *_RSmartPriceAdSlotHistoryMgr) GetBatchFromSpashEcpm(spashEcpms []float32) (results []*RSmartPriceAdSlotHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAdSlotHistory{}).Where("`spash_ecpm` IN (?)", spashEcpms).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_RSmartPriceAdSlotHistoryMgr) FetchByPrimaryKey(spashID uint32) (result RSmartPriceAdSlotHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAdSlotHistory{}).Where("`spash_id` = ?", spashID).Find(&result).Error

	return
}
