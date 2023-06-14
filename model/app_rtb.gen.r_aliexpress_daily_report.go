package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _RAliexpressDailyReportMgr struct {
	*_BaseMgr
}

// RAliexpressDailyReportMgr open func
func RAliexpressDailyReportMgr(db *gorm.DB) *_RAliexpressDailyReportMgr {
	if db == nil {
		panic(fmt.Errorf("RAliexpressDailyReportMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_RAliexpressDailyReportMgr{_BaseMgr: &_BaseMgr{DB: db.Table("r_aliexpress_daily_report"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_RAliexpressDailyReportMgr) GetTableName() string {
	return "r_aliexpress_daily_report"
}

// Reset 重置gorm会话
func (obj *_RAliexpressDailyReportMgr) Reset() *_RAliexpressDailyReportMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_RAliexpressDailyReportMgr) Get() (result RAliexpressDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAliexpressDailyReport{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_RAliexpressDailyReportMgr) Gets() (results []*RAliexpressDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAliexpressDailyReport{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_RAliexpressDailyReportMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(RAliexpressDailyReport{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithAdrID adr_id获取
func (obj *_RAliexpressDailyReportMgr) WithAdrID(adrID uint32) Option {
	return optionFunc(func(o *options) { o.query["adr_id"] = adrID })
}

// WithAdrDate adr_date获取
func (obj *_RAliexpressDailyReportMgr) WithAdrDate(adrDate string) Option {
	return optionFunc(func(o *options) { o.query["adr_date"] = adrDate })
}

// WithAdrCampaignID adr_campaign_id获取
func (obj *_RAliexpressDailyReportMgr) WithAdrCampaignID(adrCampaignID string) Option {
	return optionFunc(func(o *options) { o.query["adr_campaign_id"] = adrCampaignID })
}

// WithAdrImpression adr_impression获取
func (obj *_RAliexpressDailyReportMgr) WithAdrImpression(adrImpression int) Option {
	return optionFunc(func(o *options) { o.query["adr_impression"] = adrImpression })
}

// WithAdrClick adr_click获取
func (obj *_RAliexpressDailyReportMgr) WithAdrClick(adrClick int) Option {
	return optionFunc(func(o *options) { o.query["adr_click"] = adrClick })
}

// WithAdrPostback adr_postback获取
func (obj *_RAliexpressDailyReportMgr) WithAdrPostback(adrPostback int) Option {
	return optionFunc(func(o *options) { o.query["adr_postback"] = adrPostback })
}

// WithAdrCost adr_cost获取
func (obj *_RAliexpressDailyReportMgr) WithAdrCost(adrCost float32) Option {
	return optionFunc(func(o *options) { o.query["adr_cost"] = adrCost })
}

// WithAdrTrueImpression adr_true_impression获取
func (obj *_RAliexpressDailyReportMgr) WithAdrTrueImpression(adrTrueImpression int) Option {
	return optionFunc(func(o *options) { o.query["adr_true_impression"] = adrTrueImpression })
}

// WithAdrTrueClick adr_true_click获取
func (obj *_RAliexpressDailyReportMgr) WithAdrTrueClick(adrTrueClick int) Option {
	return optionFunc(func(o *options) { o.query["adr_true_click"] = adrTrueClick })
}

// WithAdrTrueCost adr_true_cost获取
func (obj *_RAliexpressDailyReportMgr) WithAdrTrueCost(adrTrueCost float32) Option {
	return optionFunc(func(o *options) { o.query["adr_true_cost"] = adrTrueCost })
}

// WithAdrTruePostback adr_true_postback获取
func (obj *_RAliexpressDailyReportMgr) WithAdrTruePostback(adrTruePostback int) Option {
	return optionFunc(func(o *options) { o.query["adr_true_postback"] = adrTruePostback })
}

// WithAdrCreateTime adr_create_time获取
func (obj *_RAliexpressDailyReportMgr) WithAdrCreateTime(adrCreateTime int) Option {
	return optionFunc(func(o *options) { o.query["adr_create_time"] = adrCreateTime })
}

// WithAdrAeCost adr_ae_cost获取
func (obj *_RAliexpressDailyReportMgr) WithAdrAeCost(adrAeCost float32) Option {
	return optionFunc(func(o *options) { o.query["adr_ae_cost"] = adrAeCost })
}

// WithAdrAeUv adr_ae_uv获取
func (obj *_RAliexpressDailyReportMgr) WithAdrAeUv(adrAeUv int) Option {
	return optionFunc(func(o *options) { o.query["adr_ae_uv"] = adrAeUv })
}

// WithAdrAeMemberUv adr_ae_member_uv获取
func (obj *_RAliexpressDailyReportMgr) WithAdrAeMemberUv(adrAeMemberUv int) Option {
	return optionFunc(func(o *options) { o.query["adr_ae_member_uv"] = adrAeMemberUv })
}

// WithAdrAeCrtOrdByrCnt adr_ae_crt_ord_byr_cnt获取
func (obj *_RAliexpressDailyReportMgr) WithAdrAeCrtOrdByrCnt(adrAeCrtOrdByrCnt int) Option {
	return optionFunc(func(o *options) { o.query["adr_ae_crt_ord_byr_cnt"] = adrAeCrtOrdByrCnt })
}

// WithAdrAeCrtOrdCnt adr_ae_crt_ord_cnt获取
func (obj *_RAliexpressDailyReportMgr) WithAdrAeCrtOrdCnt(adrAeCrtOrdCnt int) Option {
	return optionFunc(func(o *options) { o.query["adr_ae_crt_ord_cnt"] = adrAeCrtOrdCnt })
}

// WithAdrAeOrdGmv adr_ae_ord_gmv获取
func (obj *_RAliexpressDailyReportMgr) WithAdrAeOrdGmv(adrAeOrdGmv float32) Option {
	return optionFunc(func(o *options) { o.query["adr_ae_ord_gmv"] = adrAeOrdGmv })
}

// WithAdrAeAffiCommission adr_ae_affi_commission获取
func (obj *_RAliexpressDailyReportMgr) WithAdrAeAffiCommission(adrAeAffiCommission float32) Option {
	return optionFunc(func(o *options) { o.query["adr_ae_affi_commission"] = adrAeAffiCommission })
}

// WithAdrAeP4pCost adr_ae_p4p_cost获取
func (obj *_RAliexpressDailyReportMgr) WithAdrAeP4pCost(adrAeP4pCost float32) Option {
	return optionFunc(func(o *options) { o.query["adr_ae_p4p_cost"] = adrAeP4pCost })
}

// WithAdrAeTotalIncome adr_ae_total_income获取
func (obj *_RAliexpressDailyReportMgr) WithAdrAeTotalIncome(adrAeTotalIncome float32) Option {
	return optionFunc(func(o *options) { o.query["adr_ae_total_income"] = adrAeTotalIncome })
}

// GetByOption 功能选项模式获取
func (obj *_RAliexpressDailyReportMgr) GetByOption(opts ...Option) (result RAliexpressDailyReport, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(RAliexpressDailyReport{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_RAliexpressDailyReportMgr) GetByOptions(opts ...Option) (results []*RAliexpressDailyReport, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(RAliexpressDailyReport{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromAdrID 通过adr_id获取内容
func (obj *_RAliexpressDailyReportMgr) GetFromAdrID(adrID uint32) (result RAliexpressDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAliexpressDailyReport{}).Where("`adr_id` = ?", adrID).Find(&result).Error

	return
}

// GetBatchFromAdrID 批量查找
func (obj *_RAliexpressDailyReportMgr) GetBatchFromAdrID(adrIDs []uint32) (results []*RAliexpressDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAliexpressDailyReport{}).Where("`adr_id` IN (?)", adrIDs).Find(&results).Error

	return
}

// GetFromAdrDate 通过adr_date获取内容
func (obj *_RAliexpressDailyReportMgr) GetFromAdrDate(adrDate string) (results []*RAliexpressDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAliexpressDailyReport{}).Where("`adr_date` = ?", adrDate).Find(&results).Error

	return
}

// GetBatchFromAdrDate 批量查找
func (obj *_RAliexpressDailyReportMgr) GetBatchFromAdrDate(adrDates []string) (results []*RAliexpressDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAliexpressDailyReport{}).Where("`adr_date` IN (?)", adrDates).Find(&results).Error

	return
}

// GetFromAdrCampaignID 通过adr_campaign_id获取内容
func (obj *_RAliexpressDailyReportMgr) GetFromAdrCampaignID(adrCampaignID string) (results []*RAliexpressDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAliexpressDailyReport{}).Where("`adr_campaign_id` = ?", adrCampaignID).Find(&results).Error

	return
}

// GetBatchFromAdrCampaignID 批量查找
func (obj *_RAliexpressDailyReportMgr) GetBatchFromAdrCampaignID(adrCampaignIDs []string) (results []*RAliexpressDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAliexpressDailyReport{}).Where("`adr_campaign_id` IN (?)", adrCampaignIDs).Find(&results).Error

	return
}

// GetFromAdrImpression 通过adr_impression获取内容
func (obj *_RAliexpressDailyReportMgr) GetFromAdrImpression(adrImpression int) (results []*RAliexpressDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAliexpressDailyReport{}).Where("`adr_impression` = ?", adrImpression).Find(&results).Error

	return
}

// GetBatchFromAdrImpression 批量查找
func (obj *_RAliexpressDailyReportMgr) GetBatchFromAdrImpression(adrImpressions []int) (results []*RAliexpressDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAliexpressDailyReport{}).Where("`adr_impression` IN (?)", adrImpressions).Find(&results).Error

	return
}

// GetFromAdrClick 通过adr_click获取内容
func (obj *_RAliexpressDailyReportMgr) GetFromAdrClick(adrClick int) (results []*RAliexpressDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAliexpressDailyReport{}).Where("`adr_click` = ?", adrClick).Find(&results).Error

	return
}

// GetBatchFromAdrClick 批量查找
func (obj *_RAliexpressDailyReportMgr) GetBatchFromAdrClick(adrClicks []int) (results []*RAliexpressDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAliexpressDailyReport{}).Where("`adr_click` IN (?)", adrClicks).Find(&results).Error

	return
}

// GetFromAdrPostback 通过adr_postback获取内容
func (obj *_RAliexpressDailyReportMgr) GetFromAdrPostback(adrPostback int) (results []*RAliexpressDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAliexpressDailyReport{}).Where("`adr_postback` = ?", adrPostback).Find(&results).Error

	return
}

// GetBatchFromAdrPostback 批量查找
func (obj *_RAliexpressDailyReportMgr) GetBatchFromAdrPostback(adrPostbacks []int) (results []*RAliexpressDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAliexpressDailyReport{}).Where("`adr_postback` IN (?)", adrPostbacks).Find(&results).Error

	return
}

// GetFromAdrCost 通过adr_cost获取内容
func (obj *_RAliexpressDailyReportMgr) GetFromAdrCost(adrCost float32) (results []*RAliexpressDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAliexpressDailyReport{}).Where("`adr_cost` = ?", adrCost).Find(&results).Error

	return
}

// GetBatchFromAdrCost 批量查找
func (obj *_RAliexpressDailyReportMgr) GetBatchFromAdrCost(adrCosts []float32) (results []*RAliexpressDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAliexpressDailyReport{}).Where("`adr_cost` IN (?)", adrCosts).Find(&results).Error

	return
}

// GetFromAdrTrueImpression 通过adr_true_impression获取内容
func (obj *_RAliexpressDailyReportMgr) GetFromAdrTrueImpression(adrTrueImpression int) (results []*RAliexpressDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAliexpressDailyReport{}).Where("`adr_true_impression` = ?", adrTrueImpression).Find(&results).Error

	return
}

// GetBatchFromAdrTrueImpression 批量查找
func (obj *_RAliexpressDailyReportMgr) GetBatchFromAdrTrueImpression(adrTrueImpressions []int) (results []*RAliexpressDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAliexpressDailyReport{}).Where("`adr_true_impression` IN (?)", adrTrueImpressions).Find(&results).Error

	return
}

// GetFromAdrTrueClick 通过adr_true_click获取内容
func (obj *_RAliexpressDailyReportMgr) GetFromAdrTrueClick(adrTrueClick int) (results []*RAliexpressDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAliexpressDailyReport{}).Where("`adr_true_click` = ?", adrTrueClick).Find(&results).Error

	return
}

// GetBatchFromAdrTrueClick 批量查找
func (obj *_RAliexpressDailyReportMgr) GetBatchFromAdrTrueClick(adrTrueClicks []int) (results []*RAliexpressDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAliexpressDailyReport{}).Where("`adr_true_click` IN (?)", adrTrueClicks).Find(&results).Error

	return
}

// GetFromAdrTrueCost 通过adr_true_cost获取内容
func (obj *_RAliexpressDailyReportMgr) GetFromAdrTrueCost(adrTrueCost float32) (results []*RAliexpressDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAliexpressDailyReport{}).Where("`adr_true_cost` = ?", adrTrueCost).Find(&results).Error

	return
}

// GetBatchFromAdrTrueCost 批量查找
func (obj *_RAliexpressDailyReportMgr) GetBatchFromAdrTrueCost(adrTrueCosts []float32) (results []*RAliexpressDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAliexpressDailyReport{}).Where("`adr_true_cost` IN (?)", adrTrueCosts).Find(&results).Error

	return
}

// GetFromAdrTruePostback 通过adr_true_postback获取内容
func (obj *_RAliexpressDailyReportMgr) GetFromAdrTruePostback(adrTruePostback int) (results []*RAliexpressDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAliexpressDailyReport{}).Where("`adr_true_postback` = ?", adrTruePostback).Find(&results).Error

	return
}

// GetBatchFromAdrTruePostback 批量查找
func (obj *_RAliexpressDailyReportMgr) GetBatchFromAdrTruePostback(adrTruePostbacks []int) (results []*RAliexpressDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAliexpressDailyReport{}).Where("`adr_true_postback` IN (?)", adrTruePostbacks).Find(&results).Error

	return
}

// GetFromAdrCreateTime 通过adr_create_time获取内容
func (obj *_RAliexpressDailyReportMgr) GetFromAdrCreateTime(adrCreateTime int) (results []*RAliexpressDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAliexpressDailyReport{}).Where("`adr_create_time` = ?", adrCreateTime).Find(&results).Error

	return
}

// GetBatchFromAdrCreateTime 批量查找
func (obj *_RAliexpressDailyReportMgr) GetBatchFromAdrCreateTime(adrCreateTimes []int) (results []*RAliexpressDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAliexpressDailyReport{}).Where("`adr_create_time` IN (?)", adrCreateTimes).Find(&results).Error

	return
}

// GetFromAdrAeCost 通过adr_ae_cost获取内容
func (obj *_RAliexpressDailyReportMgr) GetFromAdrAeCost(adrAeCost float32) (results []*RAliexpressDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAliexpressDailyReport{}).Where("`adr_ae_cost` = ?", adrAeCost).Find(&results).Error

	return
}

// GetBatchFromAdrAeCost 批量查找
func (obj *_RAliexpressDailyReportMgr) GetBatchFromAdrAeCost(adrAeCosts []float32) (results []*RAliexpressDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAliexpressDailyReport{}).Where("`adr_ae_cost` IN (?)", adrAeCosts).Find(&results).Error

	return
}

// GetFromAdrAeUv 通过adr_ae_uv获取内容
func (obj *_RAliexpressDailyReportMgr) GetFromAdrAeUv(adrAeUv int) (results []*RAliexpressDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAliexpressDailyReport{}).Where("`adr_ae_uv` = ?", adrAeUv).Find(&results).Error

	return
}

// GetBatchFromAdrAeUv 批量查找
func (obj *_RAliexpressDailyReportMgr) GetBatchFromAdrAeUv(adrAeUvs []int) (results []*RAliexpressDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAliexpressDailyReport{}).Where("`adr_ae_uv` IN (?)", adrAeUvs).Find(&results).Error

	return
}

// GetFromAdrAeMemberUv 通过adr_ae_member_uv获取内容
func (obj *_RAliexpressDailyReportMgr) GetFromAdrAeMemberUv(adrAeMemberUv int) (results []*RAliexpressDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAliexpressDailyReport{}).Where("`adr_ae_member_uv` = ?", adrAeMemberUv).Find(&results).Error

	return
}

// GetBatchFromAdrAeMemberUv 批量查找
func (obj *_RAliexpressDailyReportMgr) GetBatchFromAdrAeMemberUv(adrAeMemberUvs []int) (results []*RAliexpressDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAliexpressDailyReport{}).Where("`adr_ae_member_uv` IN (?)", adrAeMemberUvs).Find(&results).Error

	return
}

// GetFromAdrAeCrtOrdByrCnt 通过adr_ae_crt_ord_byr_cnt获取内容
func (obj *_RAliexpressDailyReportMgr) GetFromAdrAeCrtOrdByrCnt(adrAeCrtOrdByrCnt int) (results []*RAliexpressDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAliexpressDailyReport{}).Where("`adr_ae_crt_ord_byr_cnt` = ?", adrAeCrtOrdByrCnt).Find(&results).Error

	return
}

// GetBatchFromAdrAeCrtOrdByrCnt 批量查找
func (obj *_RAliexpressDailyReportMgr) GetBatchFromAdrAeCrtOrdByrCnt(adrAeCrtOrdByrCnts []int) (results []*RAliexpressDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAliexpressDailyReport{}).Where("`adr_ae_crt_ord_byr_cnt` IN (?)", adrAeCrtOrdByrCnts).Find(&results).Error

	return
}

// GetFromAdrAeCrtOrdCnt 通过adr_ae_crt_ord_cnt获取内容
func (obj *_RAliexpressDailyReportMgr) GetFromAdrAeCrtOrdCnt(adrAeCrtOrdCnt int) (results []*RAliexpressDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAliexpressDailyReport{}).Where("`adr_ae_crt_ord_cnt` = ?", adrAeCrtOrdCnt).Find(&results).Error

	return
}

// GetBatchFromAdrAeCrtOrdCnt 批量查找
func (obj *_RAliexpressDailyReportMgr) GetBatchFromAdrAeCrtOrdCnt(adrAeCrtOrdCnts []int) (results []*RAliexpressDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAliexpressDailyReport{}).Where("`adr_ae_crt_ord_cnt` IN (?)", adrAeCrtOrdCnts).Find(&results).Error

	return
}

// GetFromAdrAeOrdGmv 通过adr_ae_ord_gmv获取内容
func (obj *_RAliexpressDailyReportMgr) GetFromAdrAeOrdGmv(adrAeOrdGmv float32) (results []*RAliexpressDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAliexpressDailyReport{}).Where("`adr_ae_ord_gmv` = ?", adrAeOrdGmv).Find(&results).Error

	return
}

// GetBatchFromAdrAeOrdGmv 批量查找
func (obj *_RAliexpressDailyReportMgr) GetBatchFromAdrAeOrdGmv(adrAeOrdGmvs []float32) (results []*RAliexpressDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAliexpressDailyReport{}).Where("`adr_ae_ord_gmv` IN (?)", adrAeOrdGmvs).Find(&results).Error

	return
}

// GetFromAdrAeAffiCommission 通过adr_ae_affi_commission获取内容
func (obj *_RAliexpressDailyReportMgr) GetFromAdrAeAffiCommission(adrAeAffiCommission float32) (results []*RAliexpressDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAliexpressDailyReport{}).Where("`adr_ae_affi_commission` = ?", adrAeAffiCommission).Find(&results).Error

	return
}

// GetBatchFromAdrAeAffiCommission 批量查找
func (obj *_RAliexpressDailyReportMgr) GetBatchFromAdrAeAffiCommission(adrAeAffiCommissions []float32) (results []*RAliexpressDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAliexpressDailyReport{}).Where("`adr_ae_affi_commission` IN (?)", adrAeAffiCommissions).Find(&results).Error

	return
}

// GetFromAdrAeP4pCost 通过adr_ae_p4p_cost获取内容
func (obj *_RAliexpressDailyReportMgr) GetFromAdrAeP4pCost(adrAeP4pCost float32) (results []*RAliexpressDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAliexpressDailyReport{}).Where("`adr_ae_p4p_cost` = ?", adrAeP4pCost).Find(&results).Error

	return
}

// GetBatchFromAdrAeP4pCost 批量查找
func (obj *_RAliexpressDailyReportMgr) GetBatchFromAdrAeP4pCost(adrAeP4pCosts []float32) (results []*RAliexpressDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAliexpressDailyReport{}).Where("`adr_ae_p4p_cost` IN (?)", adrAeP4pCosts).Find(&results).Error

	return
}

// GetFromAdrAeTotalIncome 通过adr_ae_total_income获取内容
func (obj *_RAliexpressDailyReportMgr) GetFromAdrAeTotalIncome(adrAeTotalIncome float32) (results []*RAliexpressDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAliexpressDailyReport{}).Where("`adr_ae_total_income` = ?", adrAeTotalIncome).Find(&results).Error

	return
}

// GetBatchFromAdrAeTotalIncome 批量查找
func (obj *_RAliexpressDailyReportMgr) GetBatchFromAdrAeTotalIncome(adrAeTotalIncomes []float32) (results []*RAliexpressDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAliexpressDailyReport{}).Where("`adr_ae_total_income` IN (?)", adrAeTotalIncomes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_RAliexpressDailyReportMgr) FetchByPrimaryKey(adrID uint32) (result RAliexpressDailyReport, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAliexpressDailyReport{}).Where("`adr_id` = ?", adrID).Find(&result).Error

	return
}
