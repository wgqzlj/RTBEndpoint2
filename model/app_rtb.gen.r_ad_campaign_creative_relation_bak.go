package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _RAdCampaignCreativeRelationBakMgr struct {
	*_BaseMgr
}

// RAdCampaignCreativeRelationBakMgr open func
func RAdCampaignCreativeRelationBakMgr(db *gorm.DB) *_RAdCampaignCreativeRelationBakMgr {
	if db == nil {
		panic(fmt.Errorf("RAdCampaignCreativeRelationBakMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_RAdCampaignCreativeRelationBakMgr{_BaseMgr: &_BaseMgr{DB: db.Table("r_ad_campaign_creative_relation_bak"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_RAdCampaignCreativeRelationBakMgr) GetTableName() string {
	return "r_ad_campaign_creative_relation_bak"
}

// Reset 重置gorm会话
func (obj *_RAdCampaignCreativeRelationBakMgr) Reset() *_RAdCampaignCreativeRelationBakMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_RAdCampaignCreativeRelationBakMgr) Get() (result RAdCampaignCreativeRelationBak, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaignCreativeRelationBak{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_RAdCampaignCreativeRelationBakMgr) Gets() (results []*RAdCampaignCreativeRelationBak, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaignCreativeRelationBak{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_RAdCampaignCreativeRelationBakMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(RAdCampaignCreativeRelationBak{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithAccrID accr_id获取
func (obj *_RAdCampaignCreativeRelationBakMgr) WithAccrID(accrID uint32) Option {
	return optionFunc(func(o *options) { o.query["accr_id"] = accrID })
}

// WithAccrAcID accr_ac_id获取
func (obj *_RAdCampaignCreativeRelationBakMgr) WithAccrAcID(accrAcID int) Option {
	return optionFunc(func(o *options) { o.query["accr_ac_id"] = accrAcID })
}

// WithAccrAcrID accr_acr_id获取
func (obj *_RAdCampaignCreativeRelationBakMgr) WithAccrAcrID(accrAcrID int) Option {
	return optionFunc(func(o *options) { o.query["accr_acr_id"] = accrAcrID })
}

// WithAccrAeID accr_ae_id获取
func (obj *_RAdCampaignCreativeRelationBakMgr) WithAccrAeID(accrAeID int) Option {
	return optionFunc(func(o *options) { o.query["accr_ae_id"] = accrAeID })
}

// WithAccrIstID accr_ist_id获取
func (obj *_RAdCampaignCreativeRelationBakMgr) WithAccrIstID(accrIstID int) Option {
	return optionFunc(func(o *options) { o.query["accr_ist_id"] = accrIstID })
}

// WithAccrStatus accr_status获取
func (obj *_RAdCampaignCreativeRelationBakMgr) WithAccrStatus(accrStatus int8) Option {
	return optionFunc(func(o *options) { o.query["accr_status"] = accrStatus })
}

// WithAccrCreateTime accr_create_time获取
func (obj *_RAdCampaignCreativeRelationBakMgr) WithAccrCreateTime(accrCreateTime int) Option {
	return optionFunc(func(o *options) { o.query["accr_create_time"] = accrCreateTime })
}

// WithAccrAtID accr_at_id获取
func (obj *_RAdCampaignCreativeRelationBakMgr) WithAccrAtID(accrAtID int) Option {
	return optionFunc(func(o *options) { o.query["accr_at_id"] = accrAtID })
}

// GetByOption 功能选项模式获取
func (obj *_RAdCampaignCreativeRelationBakMgr) GetByOption(opts ...Option) (result RAdCampaignCreativeRelationBak, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaignCreativeRelationBak{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_RAdCampaignCreativeRelationBakMgr) GetByOptions(opts ...Option) (results []*RAdCampaignCreativeRelationBak, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaignCreativeRelationBak{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromAccrID 通过accr_id获取内容
func (obj *_RAdCampaignCreativeRelationBakMgr) GetFromAccrID(accrID uint32) (result RAdCampaignCreativeRelationBak, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaignCreativeRelationBak{}).Where("`accr_id` = ?", accrID).Find(&result).Error

	return
}

// GetBatchFromAccrID 批量查找
func (obj *_RAdCampaignCreativeRelationBakMgr) GetBatchFromAccrID(accrIDs []uint32) (results []*RAdCampaignCreativeRelationBak, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaignCreativeRelationBak{}).Where("`accr_id` IN (?)", accrIDs).Find(&results).Error

	return
}

// GetFromAccrAcID 通过accr_ac_id获取内容
func (obj *_RAdCampaignCreativeRelationBakMgr) GetFromAccrAcID(accrAcID int) (results []*RAdCampaignCreativeRelationBak, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaignCreativeRelationBak{}).Where("`accr_ac_id` = ?", accrAcID).Find(&results).Error

	return
}

// GetBatchFromAccrAcID 批量查找
func (obj *_RAdCampaignCreativeRelationBakMgr) GetBatchFromAccrAcID(accrAcIDs []int) (results []*RAdCampaignCreativeRelationBak, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaignCreativeRelationBak{}).Where("`accr_ac_id` IN (?)", accrAcIDs).Find(&results).Error

	return
}

// GetFromAccrAcrID 通过accr_acr_id获取内容
func (obj *_RAdCampaignCreativeRelationBakMgr) GetFromAccrAcrID(accrAcrID int) (results []*RAdCampaignCreativeRelationBak, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaignCreativeRelationBak{}).Where("`accr_acr_id` = ?", accrAcrID).Find(&results).Error

	return
}

// GetBatchFromAccrAcrID 批量查找
func (obj *_RAdCampaignCreativeRelationBakMgr) GetBatchFromAccrAcrID(accrAcrIDs []int) (results []*RAdCampaignCreativeRelationBak, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaignCreativeRelationBak{}).Where("`accr_acr_id` IN (?)", accrAcrIDs).Find(&results).Error

	return
}

// GetFromAccrAeID 通过accr_ae_id获取内容
func (obj *_RAdCampaignCreativeRelationBakMgr) GetFromAccrAeID(accrAeID int) (results []*RAdCampaignCreativeRelationBak, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaignCreativeRelationBak{}).Where("`accr_ae_id` = ?", accrAeID).Find(&results).Error

	return
}

// GetBatchFromAccrAeID 批量查找
func (obj *_RAdCampaignCreativeRelationBakMgr) GetBatchFromAccrAeID(accrAeIDs []int) (results []*RAdCampaignCreativeRelationBak, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaignCreativeRelationBak{}).Where("`accr_ae_id` IN (?)", accrAeIDs).Find(&results).Error

	return
}

// GetFromAccrIstID 通过accr_ist_id获取内容
func (obj *_RAdCampaignCreativeRelationBakMgr) GetFromAccrIstID(accrIstID int) (results []*RAdCampaignCreativeRelationBak, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaignCreativeRelationBak{}).Where("`accr_ist_id` = ?", accrIstID).Find(&results).Error

	return
}

// GetBatchFromAccrIstID 批量查找
func (obj *_RAdCampaignCreativeRelationBakMgr) GetBatchFromAccrIstID(accrIstIDs []int) (results []*RAdCampaignCreativeRelationBak, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaignCreativeRelationBak{}).Where("`accr_ist_id` IN (?)", accrIstIDs).Find(&results).Error

	return
}

// GetFromAccrStatus 通过accr_status获取内容
func (obj *_RAdCampaignCreativeRelationBakMgr) GetFromAccrStatus(accrStatus int8) (results []*RAdCampaignCreativeRelationBak, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaignCreativeRelationBak{}).Where("`accr_status` = ?", accrStatus).Find(&results).Error

	return
}

// GetBatchFromAccrStatus 批量查找
func (obj *_RAdCampaignCreativeRelationBakMgr) GetBatchFromAccrStatus(accrStatuss []int8) (results []*RAdCampaignCreativeRelationBak, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaignCreativeRelationBak{}).Where("`accr_status` IN (?)", accrStatuss).Find(&results).Error

	return
}

// GetFromAccrCreateTime 通过accr_create_time获取内容
func (obj *_RAdCampaignCreativeRelationBakMgr) GetFromAccrCreateTime(accrCreateTime int) (results []*RAdCampaignCreativeRelationBak, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaignCreativeRelationBak{}).Where("`accr_create_time` = ?", accrCreateTime).Find(&results).Error

	return
}

// GetBatchFromAccrCreateTime 批量查找
func (obj *_RAdCampaignCreativeRelationBakMgr) GetBatchFromAccrCreateTime(accrCreateTimes []int) (results []*RAdCampaignCreativeRelationBak, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaignCreativeRelationBak{}).Where("`accr_create_time` IN (?)", accrCreateTimes).Find(&results).Error

	return
}

// GetFromAccrAtID 通过accr_at_id获取内容
func (obj *_RAdCampaignCreativeRelationBakMgr) GetFromAccrAtID(accrAtID int) (results []*RAdCampaignCreativeRelationBak, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaignCreativeRelationBak{}).Where("`accr_at_id` = ?", accrAtID).Find(&results).Error

	return
}

// GetBatchFromAccrAtID 批量查找
func (obj *_RAdCampaignCreativeRelationBakMgr) GetBatchFromAccrAtID(accrAtIDs []int) (results []*RAdCampaignCreativeRelationBak, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaignCreativeRelationBak{}).Where("`accr_at_id` IN (?)", accrAtIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_RAdCampaignCreativeRelationBakMgr) FetchByPrimaryKey(accrID uint32) (result RAdCampaignCreativeRelationBak, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaignCreativeRelationBak{}).Where("`accr_id` = ?", accrID).Find(&result).Error

	return
}
