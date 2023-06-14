package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _RAdCampaignTypeMgr struct {
	*_BaseMgr
}

// RAdCampaignTypeMgr open func
func RAdCampaignTypeMgr(db *gorm.DB) *_RAdCampaignTypeMgr {
	if db == nil {
		panic(fmt.Errorf("RAdCampaignTypeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_RAdCampaignTypeMgr{_BaseMgr: &_BaseMgr{DB: db.Table("r_ad_campaign_type"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_RAdCampaignTypeMgr) GetTableName() string {
	return "r_ad_campaign_type"
}

// Reset 重置gorm会话
func (obj *_RAdCampaignTypeMgr) Reset() *_RAdCampaignTypeMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_RAdCampaignTypeMgr) Get() (result RAdCampaignType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaignType{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_RAdCampaignTypeMgr) Gets() (results []*RAdCampaignType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaignType{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_RAdCampaignTypeMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(RAdCampaignType{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithActID act_id获取
func (obj *_RAdCampaignTypeMgr) WithActID(actID uint32) Option {
	return optionFunc(func(o *options) { o.query["act_id"] = actID })
}

// WithActCode act_code获取
func (obj *_RAdCampaignTypeMgr) WithActCode(actCode string) Option {
	return optionFunc(func(o *options) { o.query["act_code"] = actCode })
}

// WithActName act_name获取
func (obj *_RAdCampaignTypeMgr) WithActName(actName string) Option {
	return optionFunc(func(o *options) { o.query["act_name"] = actName })
}

// WithActStatus act_status获取
func (obj *_RAdCampaignTypeMgr) WithActStatus(actStatus int8) Option {
	return optionFunc(func(o *options) { o.query["act_status"] = actStatus })
}

// GetByOption 功能选项模式获取
func (obj *_RAdCampaignTypeMgr) GetByOption(opts ...Option) (result RAdCampaignType, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaignType{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_RAdCampaignTypeMgr) GetByOptions(opts ...Option) (results []*RAdCampaignType, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaignType{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromActID 通过act_id获取内容
func (obj *_RAdCampaignTypeMgr) GetFromActID(actID uint32) (result RAdCampaignType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaignType{}).Where("`act_id` = ?", actID).Find(&result).Error

	return
}

// GetBatchFromActID 批量查找
func (obj *_RAdCampaignTypeMgr) GetBatchFromActID(actIDs []uint32) (results []*RAdCampaignType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaignType{}).Where("`act_id` IN (?)", actIDs).Find(&results).Error

	return
}

// GetFromActCode 通过act_code获取内容
func (obj *_RAdCampaignTypeMgr) GetFromActCode(actCode string) (results []*RAdCampaignType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaignType{}).Where("`act_code` = ?", actCode).Find(&results).Error

	return
}

// GetBatchFromActCode 批量查找
func (obj *_RAdCampaignTypeMgr) GetBatchFromActCode(actCodes []string) (results []*RAdCampaignType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaignType{}).Where("`act_code` IN (?)", actCodes).Find(&results).Error

	return
}

// GetFromActName 通过act_name获取内容
func (obj *_RAdCampaignTypeMgr) GetFromActName(actName string) (results []*RAdCampaignType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaignType{}).Where("`act_name` = ?", actName).Find(&results).Error

	return
}

// GetBatchFromActName 批量查找
func (obj *_RAdCampaignTypeMgr) GetBatchFromActName(actNames []string) (results []*RAdCampaignType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaignType{}).Where("`act_name` IN (?)", actNames).Find(&results).Error

	return
}

// GetFromActStatus 通过act_status获取内容
func (obj *_RAdCampaignTypeMgr) GetFromActStatus(actStatus int8) (results []*RAdCampaignType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaignType{}).Where("`act_status` = ?", actStatus).Find(&results).Error

	return
}

// GetBatchFromActStatus 批量查找
func (obj *_RAdCampaignTypeMgr) GetBatchFromActStatus(actStatuss []int8) (results []*RAdCampaignType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaignType{}).Where("`act_status` IN (?)", actStatuss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_RAdCampaignTypeMgr) FetchByPrimaryKey(actID uint32) (result RAdCampaignType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaignType{}).Where("`act_id` = ?", actID).Find(&result).Error

	return
}
