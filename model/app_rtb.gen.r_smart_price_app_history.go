package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _RSmartPriceAppHistoryMgr struct {
	*_BaseMgr
}

// RSmartPriceAppHistoryMgr open func
func RSmartPriceAppHistoryMgr(db *gorm.DB) *_RSmartPriceAppHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("RSmartPriceAppHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_RSmartPriceAppHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("r_smart_price_app_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_RSmartPriceAppHistoryMgr) GetTableName() string {
	return "r_smart_price_app_history"
}

// Reset 重置gorm会话
func (obj *_RSmartPriceAppHistoryMgr) Reset() *_RSmartPriceAppHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_RSmartPriceAppHistoryMgr) Get() (result RSmartPriceAppHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAppHistory{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_RSmartPriceAppHistoryMgr) Gets() (results []*RSmartPriceAppHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAppHistory{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_RSmartPriceAppHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(RSmartPriceAppHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSpahID spah_id获取
func (obj *_RSmartPriceAppHistoryMgr) WithSpahID(spahID uint32) Option {
	return optionFunc(func(o *options) { o.query["spah_id"] = spahID })
}

// WithSpahSpaID spah_spa_id获取
func (obj *_RSmartPriceAppHistoryMgr) WithSpahSpaID(spahSpaID int) Option {
	return optionFunc(func(o *options) { o.query["spah_spa_id"] = spahSpaID })
}

// WithSpahAeID spah_ae_id获取
func (obj *_RSmartPriceAppHistoryMgr) WithSpahAeID(spahAeID int) Option {
	return optionFunc(func(o *options) { o.query["spah_ae_id"] = spahAeID })
}

// WithSpahCID spah_c_id获取
func (obj *_RSmartPriceAppHistoryMgr) WithSpahCID(spahCID int) Option {
	return optionFunc(func(o *options) { o.query["spah_c_id"] = spahCID })
}

// WithSpahItID spah_it_id获取
func (obj *_RSmartPriceAppHistoryMgr) WithSpahItID(spahItID int) Option {
	return optionFunc(func(o *options) { o.query["spah_it_id"] = spahItID })
}

// WithSpahTime spah_time获取
func (obj *_RSmartPriceAppHistoryMgr) WithSpahTime(spahTime int) Option {
	return optionFunc(func(o *options) { o.query["spah_time"] = spahTime })
}

// WithSpahBidCount spah_bid_count获取
func (obj *_RSmartPriceAppHistoryMgr) WithSpahBidCount(spahBidCount int) Option {
	return optionFunc(func(o *options) { o.query["spah_bid_count"] = spahBidCount })
}

// WithSpahImpCount spah_imp_count获取
func (obj *_RSmartPriceAppHistoryMgr) WithSpahImpCount(spahImpCount int) Option {
	return optionFunc(func(o *options) { o.query["spah_imp_count"] = spahImpCount })
}

// WithSpahClickCount spah_click_count获取
func (obj *_RSmartPriceAppHistoryMgr) WithSpahClickCount(spahClickCount int) Option {
	return optionFunc(func(o *options) { o.query["spah_click_count"] = spahClickCount })
}

// WithSpahPostbackCount spah_postback_count获取
func (obj *_RSmartPriceAppHistoryMgr) WithSpahPostbackCount(spahPostbackCount int) Option {
	return optionFunc(func(o *options) { o.query["spah_postback_count"] = spahPostbackCount })
}

// WithSpahCost spah_cost获取
func (obj *_RSmartPriceAppHistoryMgr) WithSpahCost(spahCost float32) Option {
	return optionFunc(func(o *options) { o.query["spah_cost"] = spahCost })
}

// WithSpahRoi spah_roi获取
func (obj *_RSmartPriceAppHistoryMgr) WithSpahRoi(spahRoi float32) Option {
	return optionFunc(func(o *options) { o.query["spah_roi"] = spahRoi })
}

// WithSpahEcpm spah_ecpm获取
func (obj *_RSmartPriceAppHistoryMgr) WithSpahEcpm(spahEcpm float32) Option {
	return optionFunc(func(o *options) { o.query["spah_ecpm"] = spahEcpm })
}

// GetByOption 功能选项模式获取
func (obj *_RSmartPriceAppHistoryMgr) GetByOption(opts ...Option) (result RSmartPriceAppHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAppHistory{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_RSmartPriceAppHistoryMgr) GetByOptions(opts ...Option) (results []*RSmartPriceAppHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAppHistory{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromSpahID 通过spah_id获取内容
func (obj *_RSmartPriceAppHistoryMgr) GetFromSpahID(spahID uint32) (result RSmartPriceAppHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAppHistory{}).Where("`spah_id` = ?", spahID).Find(&result).Error

	return
}

// GetBatchFromSpahID 批量查找
func (obj *_RSmartPriceAppHistoryMgr) GetBatchFromSpahID(spahIDs []uint32) (results []*RSmartPriceAppHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAppHistory{}).Where("`spah_id` IN (?)", spahIDs).Find(&results).Error

	return
}

// GetFromSpahSpaID 通过spah_spa_id获取内容
func (obj *_RSmartPriceAppHistoryMgr) GetFromSpahSpaID(spahSpaID int) (results []*RSmartPriceAppHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAppHistory{}).Where("`spah_spa_id` = ?", spahSpaID).Find(&results).Error

	return
}

// GetBatchFromSpahSpaID 批量查找
func (obj *_RSmartPriceAppHistoryMgr) GetBatchFromSpahSpaID(spahSpaIDs []int) (results []*RSmartPriceAppHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAppHistory{}).Where("`spah_spa_id` IN (?)", spahSpaIDs).Find(&results).Error

	return
}

// GetFromSpahAeID 通过spah_ae_id获取内容
func (obj *_RSmartPriceAppHistoryMgr) GetFromSpahAeID(spahAeID int) (results []*RSmartPriceAppHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAppHistory{}).Where("`spah_ae_id` = ?", spahAeID).Find(&results).Error

	return
}

// GetBatchFromSpahAeID 批量查找
func (obj *_RSmartPriceAppHistoryMgr) GetBatchFromSpahAeID(spahAeIDs []int) (results []*RSmartPriceAppHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAppHistory{}).Where("`spah_ae_id` IN (?)", spahAeIDs).Find(&results).Error

	return
}

// GetFromSpahCID 通过spah_c_id获取内容
func (obj *_RSmartPriceAppHistoryMgr) GetFromSpahCID(spahCID int) (results []*RSmartPriceAppHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAppHistory{}).Where("`spah_c_id` = ?", spahCID).Find(&results).Error

	return
}

// GetBatchFromSpahCID 批量查找
func (obj *_RSmartPriceAppHistoryMgr) GetBatchFromSpahCID(spahCIDs []int) (results []*RSmartPriceAppHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAppHistory{}).Where("`spah_c_id` IN (?)", spahCIDs).Find(&results).Error

	return
}

// GetFromSpahItID 通过spah_it_id获取内容
func (obj *_RSmartPriceAppHistoryMgr) GetFromSpahItID(spahItID int) (results []*RSmartPriceAppHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAppHistory{}).Where("`spah_it_id` = ?", spahItID).Find(&results).Error

	return
}

// GetBatchFromSpahItID 批量查找
func (obj *_RSmartPriceAppHistoryMgr) GetBatchFromSpahItID(spahItIDs []int) (results []*RSmartPriceAppHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAppHistory{}).Where("`spah_it_id` IN (?)", spahItIDs).Find(&results).Error

	return
}

// GetFromSpahTime 通过spah_time获取内容
func (obj *_RSmartPriceAppHistoryMgr) GetFromSpahTime(spahTime int) (results []*RSmartPriceAppHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAppHistory{}).Where("`spah_time` = ?", spahTime).Find(&results).Error

	return
}

// GetBatchFromSpahTime 批量查找
func (obj *_RSmartPriceAppHistoryMgr) GetBatchFromSpahTime(spahTimes []int) (results []*RSmartPriceAppHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAppHistory{}).Where("`spah_time` IN (?)", spahTimes).Find(&results).Error

	return
}

// GetFromSpahBidCount 通过spah_bid_count获取内容
func (obj *_RSmartPriceAppHistoryMgr) GetFromSpahBidCount(spahBidCount int) (results []*RSmartPriceAppHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAppHistory{}).Where("`spah_bid_count` = ?", spahBidCount).Find(&results).Error

	return
}

// GetBatchFromSpahBidCount 批量查找
func (obj *_RSmartPriceAppHistoryMgr) GetBatchFromSpahBidCount(spahBidCounts []int) (results []*RSmartPriceAppHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAppHistory{}).Where("`spah_bid_count` IN (?)", spahBidCounts).Find(&results).Error

	return
}

// GetFromSpahImpCount 通过spah_imp_count获取内容
func (obj *_RSmartPriceAppHistoryMgr) GetFromSpahImpCount(spahImpCount int) (results []*RSmartPriceAppHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAppHistory{}).Where("`spah_imp_count` = ?", spahImpCount).Find(&results).Error

	return
}

// GetBatchFromSpahImpCount 批量查找
func (obj *_RSmartPriceAppHistoryMgr) GetBatchFromSpahImpCount(spahImpCounts []int) (results []*RSmartPriceAppHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAppHistory{}).Where("`spah_imp_count` IN (?)", spahImpCounts).Find(&results).Error

	return
}

// GetFromSpahClickCount 通过spah_click_count获取内容
func (obj *_RSmartPriceAppHistoryMgr) GetFromSpahClickCount(spahClickCount int) (results []*RSmartPriceAppHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAppHistory{}).Where("`spah_click_count` = ?", spahClickCount).Find(&results).Error

	return
}

// GetBatchFromSpahClickCount 批量查找
func (obj *_RSmartPriceAppHistoryMgr) GetBatchFromSpahClickCount(spahClickCounts []int) (results []*RSmartPriceAppHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAppHistory{}).Where("`spah_click_count` IN (?)", spahClickCounts).Find(&results).Error

	return
}

// GetFromSpahPostbackCount 通过spah_postback_count获取内容
func (obj *_RSmartPriceAppHistoryMgr) GetFromSpahPostbackCount(spahPostbackCount int) (results []*RSmartPriceAppHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAppHistory{}).Where("`spah_postback_count` = ?", spahPostbackCount).Find(&results).Error

	return
}

// GetBatchFromSpahPostbackCount 批量查找
func (obj *_RSmartPriceAppHistoryMgr) GetBatchFromSpahPostbackCount(spahPostbackCounts []int) (results []*RSmartPriceAppHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAppHistory{}).Where("`spah_postback_count` IN (?)", spahPostbackCounts).Find(&results).Error

	return
}

// GetFromSpahCost 通过spah_cost获取内容
func (obj *_RSmartPriceAppHistoryMgr) GetFromSpahCost(spahCost float32) (results []*RSmartPriceAppHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAppHistory{}).Where("`spah_cost` = ?", spahCost).Find(&results).Error

	return
}

// GetBatchFromSpahCost 批量查找
func (obj *_RSmartPriceAppHistoryMgr) GetBatchFromSpahCost(spahCosts []float32) (results []*RSmartPriceAppHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAppHistory{}).Where("`spah_cost` IN (?)", spahCosts).Find(&results).Error

	return
}

// GetFromSpahRoi 通过spah_roi获取内容
func (obj *_RSmartPriceAppHistoryMgr) GetFromSpahRoi(spahRoi float32) (results []*RSmartPriceAppHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAppHistory{}).Where("`spah_roi` = ?", spahRoi).Find(&results).Error

	return
}

// GetBatchFromSpahRoi 批量查找
func (obj *_RSmartPriceAppHistoryMgr) GetBatchFromSpahRoi(spahRois []float32) (results []*RSmartPriceAppHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAppHistory{}).Where("`spah_roi` IN (?)", spahRois).Find(&results).Error

	return
}

// GetFromSpahEcpm 通过spah_ecpm获取内容
func (obj *_RSmartPriceAppHistoryMgr) GetFromSpahEcpm(spahEcpm float32) (results []*RSmartPriceAppHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAppHistory{}).Where("`spah_ecpm` = ?", spahEcpm).Find(&results).Error

	return
}

// GetBatchFromSpahEcpm 批量查找
func (obj *_RSmartPriceAppHistoryMgr) GetBatchFromSpahEcpm(spahEcpms []float32) (results []*RSmartPriceAppHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAppHistory{}).Where("`spah_ecpm` IN (?)", spahEcpms).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_RSmartPriceAppHistoryMgr) FetchByPrimaryKey(spahID uint32) (result RSmartPriceAppHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAppHistory{}).Where("`spah_id` = ?", spahID).Find(&result).Error

	return
}
