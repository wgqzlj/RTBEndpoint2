package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _RSmartPriceAdSlotMgr struct {
	*_BaseMgr
}

// RSmartPriceAdSlotMgr open func
func RSmartPriceAdSlotMgr(db *gorm.DB) *_RSmartPriceAdSlotMgr {
	if db == nil {
		panic(fmt.Errorf("RSmartPriceAdSlotMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_RSmartPriceAdSlotMgr{_BaseMgr: &_BaseMgr{DB: db.Table("r_smart_price_ad_slot"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_RSmartPriceAdSlotMgr) GetTableName() string {
	return "r_smart_price_ad_slot"
}

// Reset 重置gorm会话
func (obj *_RSmartPriceAdSlotMgr) Reset() *_RSmartPriceAdSlotMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_RSmartPriceAdSlotMgr) Get() (result RSmartPriceAdSlot, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAdSlot{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_RSmartPriceAdSlotMgr) Gets() (results []*RSmartPriceAdSlot, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAdSlot{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_RSmartPriceAdSlotMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(RSmartPriceAdSlot{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSpasID spas_id获取
func (obj *_RSmartPriceAdSlotMgr) WithSpasID(spasID uint32) Option {
	return optionFunc(func(o *options) { o.query["spas_id"] = spasID })
}

// WithSpasBundle spas_bundle获取
func (obj *_RSmartPriceAdSlotMgr) WithSpasBundle(spasBundle string) Option {
	return optionFunc(func(o *options) { o.query["spas_bundle"] = spasBundle })
}

// WithSpasWidth spas_width获取
func (obj *_RSmartPriceAdSlotMgr) WithSpasWidth(spasWidth int) Option {
	return optionFunc(func(o *options) { o.query["spas_width"] = spasWidth })
}

// WithSpasHeight spas_height获取
func (obj *_RSmartPriceAdSlotMgr) WithSpasHeight(spasHeight int) Option {
	return optionFunc(func(o *options) { o.query["spas_height"] = spasHeight })
}

// WithSpasCreateTime spas_create_time获取
func (obj *_RSmartPriceAdSlotMgr) WithSpasCreateTime(spasCreateTime int) Option {
	return optionFunc(func(o *options) { o.query["spas_create_time"] = spasCreateTime })
}

// WithSpasStatus spas_status获取
func (obj *_RSmartPriceAdSlotMgr) WithSpasStatus(spasStatus int8) Option {
	return optionFunc(func(o *options) { o.query["spas_status"] = spasStatus })
}

// WithSpasIstID spas_ist_id获取
func (obj *_RSmartPriceAdSlotMgr) WithSpasIstID(spasIstID int) Option {
	return optionFunc(func(o *options) { o.query["spas_ist_id"] = spasIstID })
}

// GetByOption 功能选项模式获取
func (obj *_RSmartPriceAdSlotMgr) GetByOption(opts ...Option) (result RSmartPriceAdSlot, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAdSlot{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_RSmartPriceAdSlotMgr) GetByOptions(opts ...Option) (results []*RSmartPriceAdSlot, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAdSlot{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromSpasID 通过spas_id获取内容
func (obj *_RSmartPriceAdSlotMgr) GetFromSpasID(spasID uint32) (result RSmartPriceAdSlot, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAdSlot{}).Where("`spas_id` = ?", spasID).Find(&result).Error

	return
}

// GetBatchFromSpasID 批量查找
func (obj *_RSmartPriceAdSlotMgr) GetBatchFromSpasID(spasIDs []uint32) (results []*RSmartPriceAdSlot, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAdSlot{}).Where("`spas_id` IN (?)", spasIDs).Find(&results).Error

	return
}

// GetFromSpasBundle 通过spas_bundle获取内容
func (obj *_RSmartPriceAdSlotMgr) GetFromSpasBundle(spasBundle string) (results []*RSmartPriceAdSlot, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAdSlot{}).Where("`spas_bundle` = ?", spasBundle).Find(&results).Error

	return
}

// GetBatchFromSpasBundle 批量查找
func (obj *_RSmartPriceAdSlotMgr) GetBatchFromSpasBundle(spasBundles []string) (results []*RSmartPriceAdSlot, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAdSlot{}).Where("`spas_bundle` IN (?)", spasBundles).Find(&results).Error

	return
}

// GetFromSpasWidth 通过spas_width获取内容
func (obj *_RSmartPriceAdSlotMgr) GetFromSpasWidth(spasWidth int) (results []*RSmartPriceAdSlot, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAdSlot{}).Where("`spas_width` = ?", spasWidth).Find(&results).Error

	return
}

// GetBatchFromSpasWidth 批量查找
func (obj *_RSmartPriceAdSlotMgr) GetBatchFromSpasWidth(spasWidths []int) (results []*RSmartPriceAdSlot, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAdSlot{}).Where("`spas_width` IN (?)", spasWidths).Find(&results).Error

	return
}

// GetFromSpasHeight 通过spas_height获取内容
func (obj *_RSmartPriceAdSlotMgr) GetFromSpasHeight(spasHeight int) (results []*RSmartPriceAdSlot, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAdSlot{}).Where("`spas_height` = ?", spasHeight).Find(&results).Error

	return
}

// GetBatchFromSpasHeight 批量查找
func (obj *_RSmartPriceAdSlotMgr) GetBatchFromSpasHeight(spasHeights []int) (results []*RSmartPriceAdSlot, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAdSlot{}).Where("`spas_height` IN (?)", spasHeights).Find(&results).Error

	return
}

// GetFromSpasCreateTime 通过spas_create_time获取内容
func (obj *_RSmartPriceAdSlotMgr) GetFromSpasCreateTime(spasCreateTime int) (results []*RSmartPriceAdSlot, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAdSlot{}).Where("`spas_create_time` = ?", spasCreateTime).Find(&results).Error

	return
}

// GetBatchFromSpasCreateTime 批量查找
func (obj *_RSmartPriceAdSlotMgr) GetBatchFromSpasCreateTime(spasCreateTimes []int) (results []*RSmartPriceAdSlot, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAdSlot{}).Where("`spas_create_time` IN (?)", spasCreateTimes).Find(&results).Error

	return
}

// GetFromSpasStatus 通过spas_status获取内容
func (obj *_RSmartPriceAdSlotMgr) GetFromSpasStatus(spasStatus int8) (results []*RSmartPriceAdSlot, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAdSlot{}).Where("`spas_status` = ?", spasStatus).Find(&results).Error

	return
}

// GetBatchFromSpasStatus 批量查找
func (obj *_RSmartPriceAdSlotMgr) GetBatchFromSpasStatus(spasStatuss []int8) (results []*RSmartPriceAdSlot, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAdSlot{}).Where("`spas_status` IN (?)", spasStatuss).Find(&results).Error

	return
}

// GetFromSpasIstID 通过spas_ist_id获取内容
func (obj *_RSmartPriceAdSlotMgr) GetFromSpasIstID(spasIstID int) (results []*RSmartPriceAdSlot, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAdSlot{}).Where("`spas_ist_id` = ?", spasIstID).Find(&results).Error

	return
}

// GetBatchFromSpasIstID 批量查找
func (obj *_RSmartPriceAdSlotMgr) GetBatchFromSpasIstID(spasIstIDs []int) (results []*RSmartPriceAdSlot, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAdSlot{}).Where("`spas_ist_id` IN (?)", spasIstIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_RSmartPriceAdSlotMgr) FetchByPrimaryKey(spasID uint32) (result RSmartPriceAdSlot, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSmartPriceAdSlot{}).Where("`spas_id` = ?", spasID).Find(&result).Error

	return
}
