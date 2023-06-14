package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _RSmartPriceRuleMgr struct {
	*_BaseMgr
}

// RSmartPriceRuleMgr open func
func RSmartPriceRuleMgr(db *gorm.DB) *_RSmartPriceRuleMgr {
	if db == nil {
		panic(fmt.Errorf("RSmartPriceRuleMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_RSmartPriceRuleMgr{_BaseMgr: &_BaseMgr{DB: db.Table("r_smart_price_rule"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_RSmartPriceRuleMgr) GetTableName() string {
	return "r_smart_price_rule"
}

// Reset 重置gorm会话
func (obj *_RSmartPriceRuleMgr) Reset() *_RSmartPriceRuleMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_RSmartPriceRuleMgr) Get() (result RSmartPriceRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceRule{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_RSmartPriceRuleMgr) Gets() (results []*RSmartPriceRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceRule{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_RSmartPriceRuleMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(RSmartPriceRule{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSmrID smr_id获取
func (obj *_RSmartPriceRuleMgr) WithSmrID(smrID uint32) Option {
	return optionFunc(func(o *options) { o.query["smr_id"] = smrID })
}

// WithSmrCID smr_c_id获取
func (obj *_RSmartPriceRuleMgr) WithSmrCID(smrCID int) Option {
	return optionFunc(func(o *options) { o.query["smr_c_id"] = smrCID })
}

// WithSmrIstID smr_ist_id获取
func (obj *_RSmartPriceRuleMgr) WithSmrIstID(smrIstID int) Option {
	return optionFunc(func(o *options) { o.query["smr_ist_id"] = smrIstID })
}

// WithSmrAnchorPrice smr_anchor_price获取
func (obj *_RSmartPriceRuleMgr) WithSmrAnchorPrice(smrAnchorPrice float32) Option {
	return optionFunc(func(o *options) { o.query["smr_anchor_price"] = smrAnchorPrice })
}

// GetByOption 功能选项模式获取
func (obj *_RSmartPriceRuleMgr) GetByOption(opts ...Option) (result RSmartPriceRule, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceRule{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_RSmartPriceRuleMgr) GetByOptions(opts ...Option) (results []*RSmartPriceRule, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceRule{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromSmrID 通过smr_id获取内容
func (obj *_RSmartPriceRuleMgr) GetFromSmrID(smrID uint32) (result RSmartPriceRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceRule{}).Where("`smr_id` = ?", smrID).Find(&result).Error

	return
}

// GetBatchFromSmrID 批量查找
func (obj *_RSmartPriceRuleMgr) GetBatchFromSmrID(smrIDs []uint32) (results []*RSmartPriceRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceRule{}).Where("`smr_id` IN (?)", smrIDs).Find(&results).Error

	return
}

// GetFromSmrCID 通过smr_c_id获取内容
func (obj *_RSmartPriceRuleMgr) GetFromSmrCID(smrCID int) (results []*RSmartPriceRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceRule{}).Where("`smr_c_id` = ?", smrCID).Find(&results).Error

	return
}

// GetBatchFromSmrCID 批量查找
func (obj *_RSmartPriceRuleMgr) GetBatchFromSmrCID(smrCIDs []int) (results []*RSmartPriceRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceRule{}).Where("`smr_c_id` IN (?)", smrCIDs).Find(&results).Error

	return
}

// GetFromSmrIstID 通过smr_ist_id获取内容
func (obj *_RSmartPriceRuleMgr) GetFromSmrIstID(smrIstID int) (results []*RSmartPriceRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceRule{}).Where("`smr_ist_id` = ?", smrIstID).Find(&results).Error

	return
}

// GetBatchFromSmrIstID 批量查找
func (obj *_RSmartPriceRuleMgr) GetBatchFromSmrIstID(smrIstIDs []int) (results []*RSmartPriceRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceRule{}).Where("`smr_ist_id` IN (?)", smrIstIDs).Find(&results).Error

	return
}

// GetFromSmrAnchorPrice 通过smr_anchor_price获取内容
func (obj *_RSmartPriceRuleMgr) GetFromSmrAnchorPrice(smrAnchorPrice float32) (results []*RSmartPriceRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceRule{}).Where("`smr_anchor_price` = ?", smrAnchorPrice).Find(&results).Error

	return
}

// GetBatchFromSmrAnchorPrice 批量查找
func (obj *_RSmartPriceRuleMgr) GetBatchFromSmrAnchorPrice(smrAnchorPrices []float32) (results []*RSmartPriceRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceRule{}).Where("`smr_anchor_price` IN (?)", smrAnchorPrices).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_RSmartPriceRuleMgr) FetchByPrimaryKey(smrID uint32) (result RSmartPriceRule, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceRule{}).Where("`smr_id` = ?", smrID).Find(&result).Error

	return
}
