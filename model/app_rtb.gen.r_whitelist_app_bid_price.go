package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _RWhitelistAppBidPriceMgr struct {
	*_BaseMgr
}

// RWhitelistAppBidPriceMgr open func
func RWhitelistAppBidPriceMgr(db *gorm.DB) *_RWhitelistAppBidPriceMgr {
	if db == nil {
		panic(fmt.Errorf("RWhitelistAppBidPriceMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_RWhitelistAppBidPriceMgr{_BaseMgr: &_BaseMgr{DB: db.Table("r_whitelist_app_bid_price"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_RWhitelistAppBidPriceMgr) GetTableName() string {
	return "r_whitelist_app_bid_price"
}

// Reset 重置gorm会话
func (obj *_RWhitelistAppBidPriceMgr) Reset() *_RWhitelistAppBidPriceMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_RWhitelistAppBidPriceMgr) Get() (result RWhitelistAppBidPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RWhitelistAppBidPrice{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_RWhitelistAppBidPriceMgr) Gets() (results []*RWhitelistAppBidPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RWhitelistAppBidPrice{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_RWhitelistAppBidPriceMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(RWhitelistAppBidPrice{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithWabpID wabp_id获取
func (obj *_RWhitelistAppBidPriceMgr) WithWabpID(wabpID uint32) Option {
	return optionFunc(func(o *options) { o.query["wabp_id"] = wabpID })
}

// GetByOption 功能选项模式获取
func (obj *_RWhitelistAppBidPriceMgr) GetByOption(opts ...Option) (result RWhitelistAppBidPrice, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(RWhitelistAppBidPrice{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_RWhitelistAppBidPriceMgr) GetByOptions(opts ...Option) (results []*RWhitelistAppBidPrice, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(RWhitelistAppBidPrice{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromWabpID 通过wabp_id获取内容
func (obj *_RWhitelistAppBidPriceMgr) GetFromWabpID(wabpID uint32) (result RWhitelistAppBidPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RWhitelistAppBidPrice{}).Where("`wabp_id` = ?", wabpID).Find(&result).Error

	return
}

// GetBatchFromWabpID 批量查找
func (obj *_RWhitelistAppBidPriceMgr) GetBatchFromWabpID(wabpIDs []uint32) (results []*RWhitelistAppBidPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RWhitelistAppBidPrice{}).Where("`wabp_id` IN (?)", wabpIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_RWhitelistAppBidPriceMgr) FetchByPrimaryKey(wabpID uint32) (result RWhitelistAppBidPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RWhitelistAppBidPrice{}).Where("`wabp_id` = ?", wabpID).Find(&result).Error

	return
}
