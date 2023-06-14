package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _RAliexpressCampaignMgr struct {
	*_BaseMgr
}

// RAliexpressCampaignMgr open func
func RAliexpressCampaignMgr(db *gorm.DB) *_RAliexpressCampaignMgr {
	if db == nil {
		panic(fmt.Errorf("RAliexpressCampaignMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_RAliexpressCampaignMgr{_BaseMgr: &_BaseMgr{DB: db.Table("r_aliexpress_campaign"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_RAliexpressCampaignMgr) GetTableName() string {
	return "r_aliexpress_campaign"
}

// Reset 重置gorm会话
func (obj *_RAliexpressCampaignMgr) Reset() *_RAliexpressCampaignMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_RAliexpressCampaignMgr) Get() (result RAliexpressCampaign, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAliexpressCampaign{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_RAliexpressCampaignMgr) Gets() (results []*RAliexpressCampaign, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAliexpressCampaign{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_RAliexpressCampaignMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(RAliexpressCampaign{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithAcID ac_id获取
func (obj *_RAliexpressCampaignMgr) WithAcID(acID uint32) Option {
	return optionFunc(func(o *options) { o.query["ac_id"] = acID })
}

// WithAcCampaignID ac_campaign_id获取
func (obj *_RAliexpressCampaignMgr) WithAcCampaignID(acCampaignID string) Option {
	return optionFunc(func(o *options) { o.query["ac_campaign_id"] = acCampaignID })
}

// WithAcCampaignName ac_campaign_name获取
func (obj *_RAliexpressCampaignMgr) WithAcCampaignName(acCampaignName string) Option {
	return optionFunc(func(o *options) { o.query["ac_campaign_name"] = acCampaignName })
}

// WithAcBusinessType ac_business_type获取
func (obj *_RAliexpressCampaignMgr) WithAcBusinessType(acBusinessType string) Option {
	return optionFunc(func(o *options) { o.query["ac_business_type"] = acBusinessType })
}

// WithAcMaterialType ac_material_type获取
func (obj *_RAliexpressCampaignMgr) WithAcMaterialType(acMaterialType string) Option {
	return optionFunc(func(o *options) { o.query["ac_material_type"] = acMaterialType })
}

// WithAcCountry ac_country获取
func (obj *_RAliexpressCampaignMgr) WithAcCountry(acCountry string) Option {
	return optionFunc(func(o *options) { o.query["ac_country"] = acCountry })
}

// WithAcTargetAudiences ac_target_audiences获取
func (obj *_RAliexpressCampaignMgr) WithAcTargetAudiences(acTargetAudiences string) Option {
	return optionFunc(func(o *options) { o.query["ac_target_audiences"] = acTargetAudiences })
}

// WithAcCreateTime ac_create_time获取
func (obj *_RAliexpressCampaignMgr) WithAcCreateTime(acCreateTime int) Option {
	return optionFunc(func(o *options) { o.query["ac_create_time"] = acCreateTime })
}

// WithAcOnline ac_online获取
func (obj *_RAliexpressCampaignMgr) WithAcOnline(acOnline int8) Option {
	return optionFunc(func(o *options) { o.query["ac_online"] = acOnline })
}

// WithAcLandingPageURL ac_landing_page_url获取
func (obj *_RAliexpressCampaignMgr) WithAcLandingPageURL(acLandingPageURL string) Option {
	return optionFunc(func(o *options) { o.query["ac_landing_page_url"] = acLandingPageURL })
}

// GetByOption 功能选项模式获取
func (obj *_RAliexpressCampaignMgr) GetByOption(opts ...Option) (result RAliexpressCampaign, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(RAliexpressCampaign{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_RAliexpressCampaignMgr) GetByOptions(opts ...Option) (results []*RAliexpressCampaign, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(RAliexpressCampaign{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromAcID 通过ac_id获取内容
func (obj *_RAliexpressCampaignMgr) GetFromAcID(acID uint32) (result RAliexpressCampaign, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAliexpressCampaign{}).Where("`ac_id` = ?", acID).Find(&result).Error

	return
}

// GetBatchFromAcID 批量查找
func (obj *_RAliexpressCampaignMgr) GetBatchFromAcID(acIDs []uint32) (results []*RAliexpressCampaign, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAliexpressCampaign{}).Where("`ac_id` IN (?)", acIDs).Find(&results).Error

	return
}

// GetFromAcCampaignID 通过ac_campaign_id获取内容
func (obj *_RAliexpressCampaignMgr) GetFromAcCampaignID(acCampaignID string) (results []*RAliexpressCampaign, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAliexpressCampaign{}).Where("`ac_campaign_id` = ?", acCampaignID).Find(&results).Error

	return
}

// GetBatchFromAcCampaignID 批量查找
func (obj *_RAliexpressCampaignMgr) GetBatchFromAcCampaignID(acCampaignIDs []string) (results []*RAliexpressCampaign, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAliexpressCampaign{}).Where("`ac_campaign_id` IN (?)", acCampaignIDs).Find(&results).Error

	return
}

// GetFromAcCampaignName 通过ac_campaign_name获取内容
func (obj *_RAliexpressCampaignMgr) GetFromAcCampaignName(acCampaignName string) (results []*RAliexpressCampaign, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAliexpressCampaign{}).Where("`ac_campaign_name` = ?", acCampaignName).Find(&results).Error

	return
}

// GetBatchFromAcCampaignName 批量查找
func (obj *_RAliexpressCampaignMgr) GetBatchFromAcCampaignName(acCampaignNames []string) (results []*RAliexpressCampaign, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAliexpressCampaign{}).Where("`ac_campaign_name` IN (?)", acCampaignNames).Find(&results).Error

	return
}

// GetFromAcBusinessType 通过ac_business_type获取内容
func (obj *_RAliexpressCampaignMgr) GetFromAcBusinessType(acBusinessType string) (results []*RAliexpressCampaign, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAliexpressCampaign{}).Where("`ac_business_type` = ?", acBusinessType).Find(&results).Error

	return
}

// GetBatchFromAcBusinessType 批量查找
func (obj *_RAliexpressCampaignMgr) GetBatchFromAcBusinessType(acBusinessTypes []string) (results []*RAliexpressCampaign, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAliexpressCampaign{}).Where("`ac_business_type` IN (?)", acBusinessTypes).Find(&results).Error

	return
}

// GetFromAcMaterialType 通过ac_material_type获取内容
func (obj *_RAliexpressCampaignMgr) GetFromAcMaterialType(acMaterialType string) (results []*RAliexpressCampaign, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAliexpressCampaign{}).Where("`ac_material_type` = ?", acMaterialType).Find(&results).Error

	return
}

// GetBatchFromAcMaterialType 批量查找
func (obj *_RAliexpressCampaignMgr) GetBatchFromAcMaterialType(acMaterialTypes []string) (results []*RAliexpressCampaign, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAliexpressCampaign{}).Where("`ac_material_type` IN (?)", acMaterialTypes).Find(&results).Error

	return
}

// GetFromAcCountry 通过ac_country获取内容
func (obj *_RAliexpressCampaignMgr) GetFromAcCountry(acCountry string) (results []*RAliexpressCampaign, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAliexpressCampaign{}).Where("`ac_country` = ?", acCountry).Find(&results).Error

	return
}

// GetBatchFromAcCountry 批量查找
func (obj *_RAliexpressCampaignMgr) GetBatchFromAcCountry(acCountrys []string) (results []*RAliexpressCampaign, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAliexpressCampaign{}).Where("`ac_country` IN (?)", acCountrys).Find(&results).Error

	return
}

// GetFromAcTargetAudiences 通过ac_target_audiences获取内容
func (obj *_RAliexpressCampaignMgr) GetFromAcTargetAudiences(acTargetAudiences string) (results []*RAliexpressCampaign, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAliexpressCampaign{}).Where("`ac_target_audiences` = ?", acTargetAudiences).Find(&results).Error

	return
}

// GetBatchFromAcTargetAudiences 批量查找
func (obj *_RAliexpressCampaignMgr) GetBatchFromAcTargetAudiences(acTargetAudiencess []string) (results []*RAliexpressCampaign, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAliexpressCampaign{}).Where("`ac_target_audiences` IN (?)", acTargetAudiencess).Find(&results).Error

	return
}

// GetFromAcCreateTime 通过ac_create_time获取内容
func (obj *_RAliexpressCampaignMgr) GetFromAcCreateTime(acCreateTime int) (results []*RAliexpressCampaign, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAliexpressCampaign{}).Where("`ac_create_time` = ?", acCreateTime).Find(&results).Error

	return
}

// GetBatchFromAcCreateTime 批量查找
func (obj *_RAliexpressCampaignMgr) GetBatchFromAcCreateTime(acCreateTimes []int) (results []*RAliexpressCampaign, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAliexpressCampaign{}).Where("`ac_create_time` IN (?)", acCreateTimes).Find(&results).Error

	return
}

// GetFromAcOnline 通过ac_online获取内容
func (obj *_RAliexpressCampaignMgr) GetFromAcOnline(acOnline int8) (results []*RAliexpressCampaign, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAliexpressCampaign{}).Where("`ac_online` = ?", acOnline).Find(&results).Error

	return
}

// GetBatchFromAcOnline 批量查找
func (obj *_RAliexpressCampaignMgr) GetBatchFromAcOnline(acOnlines []int8) (results []*RAliexpressCampaign, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAliexpressCampaign{}).Where("`ac_online` IN (?)", acOnlines).Find(&results).Error

	return
}

// GetFromAcLandingPageURL 通过ac_landing_page_url获取内容
func (obj *_RAliexpressCampaignMgr) GetFromAcLandingPageURL(acLandingPageURL string) (results []*RAliexpressCampaign, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAliexpressCampaign{}).Where("`ac_landing_page_url` = ?", acLandingPageURL).Find(&results).Error

	return
}

// GetBatchFromAcLandingPageURL 批量查找
func (obj *_RAliexpressCampaignMgr) GetBatchFromAcLandingPageURL(acLandingPageURLs []string) (results []*RAliexpressCampaign, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAliexpressCampaign{}).Where("`ac_landing_page_url` IN (?)", acLandingPageURLs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_RAliexpressCampaignMgr) FetchByPrimaryKey(acID uint32) (result RAliexpressCampaign, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAliexpressCampaign{}).Where("`ac_id` = ?", acID).Find(&result).Error

	return
}
