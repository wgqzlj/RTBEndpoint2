package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _RAdCampaignMgr struct {
	*_BaseMgr
}

// RAdCampaignMgr open func
func RAdCampaignMgr(db *gorm.DB) *_RAdCampaignMgr {
	if db == nil {
		panic(fmt.Errorf("RAdCampaignMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_RAdCampaignMgr{_BaseMgr: &_BaseMgr{DB: db.Table("r_ad_campaign"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_RAdCampaignMgr) GetTableName() string {
	return "r_ad_campaign"
}

// Reset 重置gorm会话
func (obj *_RAdCampaignMgr) Reset() *_RAdCampaignMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_RAdCampaignMgr) Get() (result RAdCampaign, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaign{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_RAdCampaignMgr) Gets() (results []*RAdCampaign, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaign{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_RAdCampaignMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(RAdCampaign{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithAcID ac_id获取
func (obj *_RAdCampaignMgr) WithAcID(acID uint64) Option {
	return optionFunc(func(o *options) { o.query["ac_id"] = acID })
}

// WithAcCode ac_code获取
func (obj *_RAdCampaignMgr) WithAcCode(acCode string) Option {
	return optionFunc(func(o *options) { o.query["ac_code"] = acCode })
}

// WithAcCID ac_c_id获取
func (obj *_RAdCampaignMgr) WithAcCID(acCID int) Option {
	return optionFunc(func(o *options) { o.query["ac_c_id"] = acCID })
}

// WithAcOID ac_o_id获取
func (obj *_RAdCampaignMgr) WithAcOID(acOID int) Option {
	return optionFunc(func(o *options) { o.query["ac_o_id"] = acOID })
}

// WithAcName ac_name获取
func (obj *_RAdCampaignMgr) WithAcName(acName string) Option {
	return optionFunc(func(o *options) { o.query["ac_name"] = acName })
}

// WithAcDeeplink ac_deeplink获取
func (obj *_RAdCampaignMgr) WithAcDeeplink(acDeeplink string) Option {
	return optionFunc(func(o *options) { o.query["ac_deeplink"] = acDeeplink })
}

// WithAcClickURL ac_click_url获取
func (obj *_RAdCampaignMgr) WithAcClickURL(acClickURL string) Option {
	return optionFunc(func(o *options) { o.query["ac_click_url"] = acClickURL })
}

// WithAcImpressionTrackingURL ac_impression_tracking_url获取
func (obj *_RAdCampaignMgr) WithAcImpressionTrackingURL(acImpressionTrackingURL string) Option {
	return optionFunc(func(o *options) { o.query["ac_impression_tracking_url"] = acImpressionTrackingURL })
}

// WithAcActID ac_act_id获取
func (obj *_RAdCampaignMgr) WithAcActID(acActID int) Option {
	return optionFunc(func(o *options) { o.query["ac_act_id"] = acActID })
}

// WithAcClickTrackingURL ac_click_tracking_url获取
func (obj *_RAdCampaignMgr) WithAcClickTrackingURL(acClickTrackingURL string) Option {
	return optionFunc(func(o *options) { o.query["ac_click_tracking_url"] = acClickTrackingURL })
}

// WithAcMemberID ac_member_id获取
func (obj *_RAdCampaignMgr) WithAcMemberID(acMemberID string) Option {
	return optionFunc(func(o *options) { o.query["ac_member_id"] = acMemberID })
}

// WithAcStatus ac_status获取
func (obj *_RAdCampaignMgr) WithAcStatus(acStatus int8) Option {
	return optionFunc(func(o *options) { o.query["ac_status"] = acStatus })
}

// WithAcValidStartTime ac_valid_start_time获取
func (obj *_RAdCampaignMgr) WithAcValidStartTime(acValidStartTime int) Option {
	return optionFunc(func(o *options) { o.query["ac_valid_start_time"] = acValidStartTime })
}

// WithAcValidEndTime ac_valid_end_time获取
func (obj *_RAdCampaignMgr) WithAcValidEndTime(acValidEndTime int) Option {
	return optionFunc(func(o *options) { o.query["ac_valid_end_time"] = acValidEndTime })
}

// WithAcCreateTime ac_create_time获取
func (obj *_RAdCampaignMgr) WithAcCreateTime(acCreateTime int) Option {
	return optionFunc(func(o *options) { o.query["ac_create_time"] = acCreateTime })
}

// WithAcPriority ac_priority获取
func (obj *_RAdCampaignMgr) WithAcPriority(acPriority int) Option {
	return optionFunc(func(o *options) { o.query["ac_priority"] = acPriority })
}

// WithAcRemark ac_remark获取
func (obj *_RAdCampaignMgr) WithAcRemark(acRemark string) Option {
	return optionFunc(func(o *options) { o.query["ac_remark"] = acRemark })
}

// WithAcAcrCount ac_acr_count获取
func (obj *_RAdCampaignMgr) WithAcAcrCount(acAcrCount int) Option {
	return optionFunc(func(o *options) { o.query["ac_acr_count"] = acAcrCount })
}

// WithAcAcrcID ac_acrc_id获取 关联的r_ad_creative_collection id
func (obj *_RAdCampaignMgr) WithAcAcrcID(acAcrcID int) Option {
	return optionFunc(func(o *options) { o.query["ac_acrc_id"] = acAcrcID })
}

// GetByOption 功能选项模式获取
func (obj *_RAdCampaignMgr) GetByOption(opts ...Option) (result RAdCampaign, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaign{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_RAdCampaignMgr) GetByOptions(opts ...Option) (results []*RAdCampaign, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaign{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromAcID 通过ac_id获取内容
func (obj *_RAdCampaignMgr) GetFromAcID(acID uint64) (result RAdCampaign, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaign{}).Where("`ac_id` = ?", acID).Find(&result).Error

	return
}

// GetBatchFromAcID 批量查找
func (obj *_RAdCampaignMgr) GetBatchFromAcID(acIDs []uint64) (results []*RAdCampaign, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaign{}).Where("`ac_id` IN (?)", acIDs).Find(&results).Error

	return
}

// GetFromAcCode 通过ac_code获取内容
func (obj *_RAdCampaignMgr) GetFromAcCode(acCode string) (results []*RAdCampaign, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaign{}).Where("`ac_code` = ?", acCode).Find(&results).Error

	return
}

// GetBatchFromAcCode 批量查找
func (obj *_RAdCampaignMgr) GetBatchFromAcCode(acCodes []string) (results []*RAdCampaign, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaign{}).Where("`ac_code` IN (?)", acCodes).Find(&results).Error

	return
}

// GetFromAcCID 通过ac_c_id获取内容
func (obj *_RAdCampaignMgr) GetFromAcCID(acCID int) (results []*RAdCampaign, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaign{}).Where("`ac_c_id` = ?", acCID).Find(&results).Error

	return
}

// GetBatchFromAcCID 批量查找
func (obj *_RAdCampaignMgr) GetBatchFromAcCID(acCIDs []int) (results []*RAdCampaign, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaign{}).Where("`ac_c_id` IN (?)", acCIDs).Find(&results).Error

	return
}

// GetFromAcOID 通过ac_o_id获取内容
func (obj *_RAdCampaignMgr) GetFromAcOID(acOID int) (results []*RAdCampaign, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaign{}).Where("`ac_o_id` = ?", acOID).Find(&results).Error

	return
}

// GetBatchFromAcOID 批量查找
func (obj *_RAdCampaignMgr) GetBatchFromAcOID(acOIDs []int) (results []*RAdCampaign, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaign{}).Where("`ac_o_id` IN (?)", acOIDs).Find(&results).Error

	return
}

// GetFromAcName 通过ac_name获取内容
func (obj *_RAdCampaignMgr) GetFromAcName(acName string) (results []*RAdCampaign, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaign{}).Where("`ac_name` = ?", acName).Find(&results).Error

	return
}

// GetBatchFromAcName 批量查找
func (obj *_RAdCampaignMgr) GetBatchFromAcName(acNames []string) (results []*RAdCampaign, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaign{}).Where("`ac_name` IN (?)", acNames).Find(&results).Error

	return
}

// GetFromAcDeeplink 通过ac_deeplink获取内容
func (obj *_RAdCampaignMgr) GetFromAcDeeplink(acDeeplink string) (results []*RAdCampaign, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaign{}).Where("`ac_deeplink` = ?", acDeeplink).Find(&results).Error

	return
}

// GetBatchFromAcDeeplink 批量查找
func (obj *_RAdCampaignMgr) GetBatchFromAcDeeplink(acDeeplinks []string) (results []*RAdCampaign, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaign{}).Where("`ac_deeplink` IN (?)", acDeeplinks).Find(&results).Error

	return
}

// GetFromAcClickURL 通过ac_click_url获取内容
func (obj *_RAdCampaignMgr) GetFromAcClickURL(acClickURL string) (results []*RAdCampaign, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaign{}).Where("`ac_click_url` = ?", acClickURL).Find(&results).Error

	return
}

// GetBatchFromAcClickURL 批量查找
func (obj *_RAdCampaignMgr) GetBatchFromAcClickURL(acClickURLs []string) (results []*RAdCampaign, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaign{}).Where("`ac_click_url` IN (?)", acClickURLs).Find(&results).Error

	return
}

// GetFromAcImpressionTrackingURL 通过ac_impression_tracking_url获取内容
func (obj *_RAdCampaignMgr) GetFromAcImpressionTrackingURL(acImpressionTrackingURL string) (results []*RAdCampaign, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaign{}).Where("`ac_impression_tracking_url` = ?", acImpressionTrackingURL).Find(&results).Error

	return
}

// GetBatchFromAcImpressionTrackingURL 批量查找
func (obj *_RAdCampaignMgr) GetBatchFromAcImpressionTrackingURL(acImpressionTrackingURLs []string) (results []*RAdCampaign, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaign{}).Where("`ac_impression_tracking_url` IN (?)", acImpressionTrackingURLs).Find(&results).Error

	return
}

// GetFromAcActID 通过ac_act_id获取内容
func (obj *_RAdCampaignMgr) GetFromAcActID(acActID int) (results []*RAdCampaign, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaign{}).Where("`ac_act_id` = ?", acActID).Find(&results).Error

	return
}

// GetBatchFromAcActID 批量查找
func (obj *_RAdCampaignMgr) GetBatchFromAcActID(acActIDs []int) (results []*RAdCampaign, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaign{}).Where("`ac_act_id` IN (?)", acActIDs).Find(&results).Error

	return
}

// GetFromAcClickTrackingURL 通过ac_click_tracking_url获取内容
func (obj *_RAdCampaignMgr) GetFromAcClickTrackingURL(acClickTrackingURL string) (results []*RAdCampaign, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaign{}).Where("`ac_click_tracking_url` = ?", acClickTrackingURL).Find(&results).Error

	return
}

// GetBatchFromAcClickTrackingURL 批量查找
func (obj *_RAdCampaignMgr) GetBatchFromAcClickTrackingURL(acClickTrackingURLs []string) (results []*RAdCampaign, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaign{}).Where("`ac_click_tracking_url` IN (?)", acClickTrackingURLs).Find(&results).Error

	return
}

// GetFromAcMemberID 通过ac_member_id获取内容
func (obj *_RAdCampaignMgr) GetFromAcMemberID(acMemberID string) (results []*RAdCampaign, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaign{}).Where("`ac_member_id` = ?", acMemberID).Find(&results).Error

	return
}

// GetBatchFromAcMemberID 批量查找
func (obj *_RAdCampaignMgr) GetBatchFromAcMemberID(acMemberIDs []string) (results []*RAdCampaign, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaign{}).Where("`ac_member_id` IN (?)", acMemberIDs).Find(&results).Error

	return
}

// GetFromAcStatus 通过ac_status获取内容
func (obj *_RAdCampaignMgr) GetFromAcStatus(acStatus int8) (results []*RAdCampaign, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaign{}).Where("`ac_status` = ?", acStatus).Find(&results).Error

	return
}

// GetBatchFromAcStatus 批量查找
func (obj *_RAdCampaignMgr) GetBatchFromAcStatus(acStatuss []int8) (results []*RAdCampaign, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaign{}).Where("`ac_status` IN (?)", acStatuss).Find(&results).Error

	return
}

// GetFromAcValidStartTime 通过ac_valid_start_time获取内容
func (obj *_RAdCampaignMgr) GetFromAcValidStartTime(acValidStartTime int) (results []*RAdCampaign, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaign{}).Where("`ac_valid_start_time` = ?", acValidStartTime).Find(&results).Error

	return
}

// GetBatchFromAcValidStartTime 批量查找
func (obj *_RAdCampaignMgr) GetBatchFromAcValidStartTime(acValidStartTimes []int) (results []*RAdCampaign, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaign{}).Where("`ac_valid_start_time` IN (?)", acValidStartTimes).Find(&results).Error

	return
}

// GetFromAcValidEndTime 通过ac_valid_end_time获取内容
func (obj *_RAdCampaignMgr) GetFromAcValidEndTime(acValidEndTime int) (results []*RAdCampaign, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaign{}).Where("`ac_valid_end_time` = ?", acValidEndTime).Find(&results).Error

	return
}

// GetBatchFromAcValidEndTime 批量查找
func (obj *_RAdCampaignMgr) GetBatchFromAcValidEndTime(acValidEndTimes []int) (results []*RAdCampaign, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaign{}).Where("`ac_valid_end_time` IN (?)", acValidEndTimes).Find(&results).Error

	return
}

// GetFromAcCreateTime 通过ac_create_time获取内容
func (obj *_RAdCampaignMgr) GetFromAcCreateTime(acCreateTime int) (results []*RAdCampaign, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaign{}).Where("`ac_create_time` = ?", acCreateTime).Find(&results).Error

	return
}

// GetBatchFromAcCreateTime 批量查找
func (obj *_RAdCampaignMgr) GetBatchFromAcCreateTime(acCreateTimes []int) (results []*RAdCampaign, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaign{}).Where("`ac_create_time` IN (?)", acCreateTimes).Find(&results).Error

	return
}

// GetFromAcPriority 通过ac_priority获取内容
func (obj *_RAdCampaignMgr) GetFromAcPriority(acPriority int) (results []*RAdCampaign, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaign{}).Where("`ac_priority` = ?", acPriority).Find(&results).Error

	return
}

// GetBatchFromAcPriority 批量查找
func (obj *_RAdCampaignMgr) GetBatchFromAcPriority(acPrioritys []int) (results []*RAdCampaign, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaign{}).Where("`ac_priority` IN (?)", acPrioritys).Find(&results).Error

	return
}

// GetFromAcRemark 通过ac_remark获取内容
func (obj *_RAdCampaignMgr) GetFromAcRemark(acRemark string) (results []*RAdCampaign, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaign{}).Where("`ac_remark` = ?", acRemark).Find(&results).Error

	return
}

// GetBatchFromAcRemark 批量查找
func (obj *_RAdCampaignMgr) GetBatchFromAcRemark(acRemarks []string) (results []*RAdCampaign, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaign{}).Where("`ac_remark` IN (?)", acRemarks).Find(&results).Error

	return
}

// GetFromAcAcrCount 通过ac_acr_count获取内容
func (obj *_RAdCampaignMgr) GetFromAcAcrCount(acAcrCount int) (results []*RAdCampaign, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaign{}).Where("`ac_acr_count` = ?", acAcrCount).Find(&results).Error

	return
}

// GetBatchFromAcAcrCount 批量查找
func (obj *_RAdCampaignMgr) GetBatchFromAcAcrCount(acAcrCounts []int) (results []*RAdCampaign, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaign{}).Where("`ac_acr_count` IN (?)", acAcrCounts).Find(&results).Error

	return
}

// GetFromAcAcrcID 通过ac_acrc_id获取内容 关联的r_ad_creative_collection id
func (obj *_RAdCampaignMgr) GetFromAcAcrcID(acAcrcID int) (results []*RAdCampaign, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaign{}).Where("`ac_acrc_id` = ?", acAcrcID).Find(&results).Error

	return
}

// GetBatchFromAcAcrcID 批量查找 关联的r_ad_creative_collection id
func (obj *_RAdCampaignMgr) GetBatchFromAcAcrcID(acAcrcIDs []int) (results []*RAdCampaign, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaign{}).Where("`ac_acrc_id` IN (?)", acAcrcIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_RAdCampaignMgr) FetchByPrimaryKey(acID uint64) (result RAdCampaign, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaign{}).Where("`ac_id` = ?", acID).Find(&result).Error

	return
}

// FetchIndexByAcCID  获取多个内容
func (obj *_RAdCampaignMgr) FetchIndexByAcCID(acCID int) (results []*RAdCampaign, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaign{}).Where("`ac_c_id` = ?", acCID).Find(&results).Error

	return
}

// FetchIndexByAcOID  获取多个内容
func (obj *_RAdCampaignMgr) FetchIndexByAcOID(acOID int) (results []*RAdCampaign, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaign{}).Where("`ac_o_id` = ?", acOID).Find(&results).Error

	return
}

// FetchIndexByAcActID  获取多个内容
func (obj *_RAdCampaignMgr) FetchIndexByAcActID(acActID int) (results []*RAdCampaign, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaign{}).Where("`ac_act_id` = ?", acActID).Find(&results).Error

	return
}

// FetchIndexByAcAcrcID  获取多个内容
func (obj *_RAdCampaignMgr) FetchIndexByAcAcrcID(acAcrcID int) (results []*RAdCampaign, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCampaign{}).Where("`ac_acrc_id` = ?", acAcrcID).Find(&results).Error

	return
}
