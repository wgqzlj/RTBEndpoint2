package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _RAdCampaignCreativeRelationMgr struct {
	*_BaseMgr
}

// RAdCampaignCreativeRelationMgr open func
func RAdCampaignCreativeRelationMgr(db *gorm.DB) *_RAdCampaignCreativeRelationMgr {
	if db == nil {
		panic(fmt.Errorf("RAdCampaignCreativeRelationMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_RAdCampaignCreativeRelationMgr{_BaseMgr: &_BaseMgr{DB: db.Table("r_ad_campaign_creative_relation"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_RAdCampaignCreativeRelationMgr) GetTableName() string {
	return "r_ad_campaign_creative_relation"
}

// Reset 重置gorm会话
func (obj *_RAdCampaignCreativeRelationMgr) Reset() *_RAdCampaignCreativeRelationMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_RAdCampaignCreativeRelationMgr) Get() (result RAdCampaignCreativeRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaignCreativeRelation{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_RAdCampaignCreativeRelationMgr) Gets() (results []*RAdCampaignCreativeRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaignCreativeRelation{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_RAdCampaignCreativeRelationMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(RAdCampaignCreativeRelation{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithAccrID accr_id获取
func (obj *_RAdCampaignCreativeRelationMgr) WithAccrID(accrID uint32) Option {
	return optionFunc(func(o *options) { o.query["accr_id"] = accrID })
}

// WithAccrAcID accr_ac_id获取
func (obj *_RAdCampaignCreativeRelationMgr) WithAccrAcID(accrAcID int) Option {
	return optionFunc(func(o *options) { o.query["accr_ac_id"] = accrAcID })
}

// WithAccrAcrID accr_acr_id获取
func (obj *_RAdCampaignCreativeRelationMgr) WithAccrAcrID(accrAcrID int) Option {
	return optionFunc(func(o *options) { o.query["accr_acr_id"] = accrAcrID })
}

// WithAccrAeID accr_ae_id获取
func (obj *_RAdCampaignCreativeRelationMgr) WithAccrAeID(accrAeID int) Option {
	return optionFunc(func(o *options) { o.query["accr_ae_id"] = accrAeID })
}

// WithAccrIstID accr_ist_id获取
func (obj *_RAdCampaignCreativeRelationMgr) WithAccrIstID(accrIstID int) Option {
	return optionFunc(func(o *options) { o.query["accr_ist_id"] = accrIstID })
}

// WithAccrStatus accr_status获取
func (obj *_RAdCampaignCreativeRelationMgr) WithAccrStatus(accrStatus int8) Option {
	return optionFunc(func(o *options) { o.query["accr_status"] = accrStatus })
}

// WithAccrCreateTime accr_create_time获取
func (obj *_RAdCampaignCreativeRelationMgr) WithAccrCreateTime(accrCreateTime int) Option {
	return optionFunc(func(o *options) { o.query["accr_create_time"] = accrCreateTime })
}

// WithAccrAtID accr_at_id获取
func (obj *_RAdCampaignCreativeRelationMgr) WithAccrAtID(accrAtID int) Option {
	return optionFunc(func(o *options) { o.query["accr_at_id"] = accrAtID })
}

// GetByOption 功能选项模式获取
func (obj *_RAdCampaignCreativeRelationMgr) GetByOption(opts ...Option) (result RAdCampaignCreativeRelation, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaignCreativeRelation{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_RAdCampaignCreativeRelationMgr) GetByOptions(opts ...Option) (results []*RAdCampaignCreativeRelation, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaignCreativeRelation{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromAccrID 通过accr_id获取内容
func (obj *_RAdCampaignCreativeRelationMgr) GetFromAccrID(accrID uint32) (result RAdCampaignCreativeRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaignCreativeRelation{}).Where("`accr_id` = ?", accrID).Find(&result).Error

	return
}

// GetBatchFromAccrID 批量查找
func (obj *_RAdCampaignCreativeRelationMgr) GetBatchFromAccrID(accrIDs []uint32) (results []*RAdCampaignCreativeRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaignCreativeRelation{}).Where("`accr_id` IN (?)", accrIDs).Find(&results).Error

	return
}

// GetFromAccrAcID 通过accr_ac_id获取内容
func (obj *_RAdCampaignCreativeRelationMgr) GetFromAccrAcID(accrAcID int) (results []*RAdCampaignCreativeRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaignCreativeRelation{}).Where("`accr_ac_id` = ?", accrAcID).Find(&results).Error

	return
}

// GetBatchFromAccrAcID 批量查找
func (obj *_RAdCampaignCreativeRelationMgr) GetBatchFromAccrAcID(accrAcIDs []int) (results []*RAdCampaignCreativeRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaignCreativeRelation{}).Where("`accr_ac_id` IN (?)", accrAcIDs).Find(&results).Error

	return
}

// GetFromAccrAcrID 通过accr_acr_id获取内容
func (obj *_RAdCampaignCreativeRelationMgr) GetFromAccrAcrID(accrAcrID int) (results []*RAdCampaignCreativeRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaignCreativeRelation{}).Where("`accr_acr_id` = ?", accrAcrID).Find(&results).Error

	return
}

// GetBatchFromAccrAcrID 批量查找
func (obj *_RAdCampaignCreativeRelationMgr) GetBatchFromAccrAcrID(accrAcrIDs []int) (results []*RAdCampaignCreativeRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaignCreativeRelation{}).Where("`accr_acr_id` IN (?)", accrAcrIDs).Find(&results).Error

	return
}

// GetFromAccrAeID 通过accr_ae_id获取内容
func (obj *_RAdCampaignCreativeRelationMgr) GetFromAccrAeID(accrAeID int) (results []*RAdCampaignCreativeRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaignCreativeRelation{}).Where("`accr_ae_id` = ?", accrAeID).Find(&results).Error

	return
}

// GetBatchFromAccrAeID 批量查找
func (obj *_RAdCampaignCreativeRelationMgr) GetBatchFromAccrAeID(accrAeIDs []int) (results []*RAdCampaignCreativeRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaignCreativeRelation{}).Where("`accr_ae_id` IN (?)", accrAeIDs).Find(&results).Error

	return
}

// GetFromAccrIstID 通过accr_ist_id获取内容
func (obj *_RAdCampaignCreativeRelationMgr) GetFromAccrIstID(accrIstID int) (results []*RAdCampaignCreativeRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaignCreativeRelation{}).Where("`accr_ist_id` = ?", accrIstID).Find(&results).Error

	return
}

// GetBatchFromAccrIstID 批量查找
func (obj *_RAdCampaignCreativeRelationMgr) GetBatchFromAccrIstID(accrIstIDs []int) (results []*RAdCampaignCreativeRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaignCreativeRelation{}).Where("`accr_ist_id` IN (?)", accrIstIDs).Find(&results).Error

	return
}

// GetFromAccrStatus 通过accr_status获取内容
func (obj *_RAdCampaignCreativeRelationMgr) GetFromAccrStatus(accrStatus int8) (results []*RAdCampaignCreativeRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaignCreativeRelation{}).Where("`accr_status` = ?", accrStatus).Find(&results).Error

	return
}

// GetBatchFromAccrStatus 批量查找
func (obj *_RAdCampaignCreativeRelationMgr) GetBatchFromAccrStatus(accrStatuss []int8) (results []*RAdCampaignCreativeRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaignCreativeRelation{}).Where("`accr_status` IN (?)", accrStatuss).Find(&results).Error

	return
}

// GetFromAccrCreateTime 通过accr_create_time获取内容
func (obj *_RAdCampaignCreativeRelationMgr) GetFromAccrCreateTime(accrCreateTime int) (results []*RAdCampaignCreativeRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaignCreativeRelation{}).Where("`accr_create_time` = ?", accrCreateTime).Find(&results).Error

	return
}

// GetBatchFromAccrCreateTime 批量查找
func (obj *_RAdCampaignCreativeRelationMgr) GetBatchFromAccrCreateTime(accrCreateTimes []int) (results []*RAdCampaignCreativeRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaignCreativeRelation{}).Where("`accr_create_time` IN (?)", accrCreateTimes).Find(&results).Error

	return
}

// GetFromAccrAtID 通过accr_at_id获取内容
func (obj *_RAdCampaignCreativeRelationMgr) GetFromAccrAtID(accrAtID int) (results []*RAdCampaignCreativeRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaignCreativeRelation{}).Where("`accr_at_id` = ?", accrAtID).Find(&results).Error

	return
}

// GetBatchFromAccrAtID 批量查找
func (obj *_RAdCampaignCreativeRelationMgr) GetBatchFromAccrAtID(accrAtIDs []int) (results []*RAdCampaignCreativeRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaignCreativeRelation{}).Where("`accr_at_id` IN (?)", accrAtIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_RAdCampaignCreativeRelationMgr) FetchByPrimaryKey(accrID uint32) (result RAdCampaignCreativeRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaignCreativeRelation{}).Where("`accr_id` = ?", accrID).Find(&result).Error

	return
}
