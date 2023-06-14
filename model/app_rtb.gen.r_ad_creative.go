package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _RAdCreativeMgr struct {
	*_BaseMgr
}

// RAdCreativeMgr open func
func RAdCreativeMgr(db *gorm.DB) *_RAdCreativeMgr {
	if db == nil {
		panic(fmt.Errorf("RAdCreativeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_RAdCreativeMgr{_BaseMgr: &_BaseMgr{DB: db.Table("r_ad_creative"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_RAdCreativeMgr) GetTableName() string {
	return "r_ad_creative"
}

// Reset 重置gorm会话
func (obj *_RAdCreativeMgr) Reset() *_RAdCreativeMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_RAdCreativeMgr) Get() (result RAdCreative, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreative{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_RAdCreativeMgr) Gets() (results []*RAdCreative, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreative{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_RAdCreativeMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(RAdCreative{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithAcrID acr_id获取
func (obj *_RAdCreativeMgr) WithAcrID(acrID uint32) Option {
	return optionFunc(func(o *options) { o.query["acr_id"] = acrID })
}

// WithAcrName acr_name获取
func (obj *_RAdCreativeMgr) WithAcrName(acrName string) Option {
	return optionFunc(func(o *options) { o.query["acr_name"] = acrName })
}

// WithAcrCID acr_c_id获取 国家ID
func (obj *_RAdCreativeMgr) WithAcrCID(acrCID int) Option {
	return optionFunc(func(o *options) { o.query["acr_c_id"] = acrCID })
}

// WithAcrOID acr_o_id获取
func (obj *_RAdCreativeMgr) WithAcrOID(acrOID int) Option {
	return optionFunc(func(o *options) { o.query["acr_o_id"] = acrOID })
}

// WithAcrURL acr_url获取
func (obj *_RAdCreativeMgr) WithAcrURL(acrURL string) Option {
	return optionFunc(func(o *options) { o.query["acr_url"] = acrURL })
}

// WithAcrCloakURL acr_cloak_url获取
func (obj *_RAdCreativeMgr) WithAcrCloakURL(acrCloakURL string) Option {
	return optionFunc(func(o *options) { o.query["acr_cloak_url"] = acrCloakURL })
}

// WithAcrWidth acr_width获取
func (obj *_RAdCreativeMgr) WithAcrWidth(acrWidth float32) Option {
	return optionFunc(func(o *options) { o.query["acr_width"] = acrWidth })
}

// WithAcrHeight acr_height获取
func (obj *_RAdCreativeMgr) WithAcrHeight(acrHeight float32) Option {
	return optionFunc(func(o *options) { o.query["acr_height"] = acrHeight })
}

// WithAcrStatus acr_status获取
func (obj *_RAdCreativeMgr) WithAcrStatus(acrStatus int8) Option {
	return optionFunc(func(o *options) { o.query["acr_status"] = acrStatus })
}

// WithAcrItID acr_it_id获取 1 banner 2 video
func (obj *_RAdCreativeMgr) WithAcrItID(acrItID int) Option {
	return optionFunc(func(o *options) { o.query["acr_it_id"] = acrItID })
}

// WithAcrExt acr_ext获取
func (obj *_RAdCreativeMgr) WithAcrExt(acrExt string) Option {
	return optionFunc(func(o *options) { o.query["acr_ext"] = acrExt })
}

// WithAcrRemark acr_remark获取
func (obj *_RAdCreativeMgr) WithAcrRemark(acrRemark string) Option {
	return optionFunc(func(o *options) { o.query["acr_remark"] = acrRemark })
}

// GetByOption 功能选项模式获取
func (obj *_RAdCreativeMgr) GetByOption(opts ...Option) (result RAdCreative, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(RAdCreative{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_RAdCreativeMgr) GetByOptions(opts ...Option) (results []*RAdCreative, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(RAdCreative{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromAcrID 通过acr_id获取内容
func (obj *_RAdCreativeMgr) GetFromAcrID(acrID uint32) (result RAdCreative, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreative{}).Where("`acr_id` = ?", acrID).Find(&result).Error

	return
}

// GetBatchFromAcrID 批量查找
func (obj *_RAdCreativeMgr) GetBatchFromAcrID(acrIDs []uint32) (results []*RAdCreative, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreative{}).Where("`acr_id` IN (?)", acrIDs).Find(&results).Error

	return
}

// GetFromAcrName 通过acr_name获取内容
func (obj *_RAdCreativeMgr) GetFromAcrName(acrName string) (results []*RAdCreative, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreative{}).Where("`acr_name` = ?", acrName).Find(&results).Error

	return
}

// GetBatchFromAcrName 批量查找
func (obj *_RAdCreativeMgr) GetBatchFromAcrName(acrNames []string) (results []*RAdCreative, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreative{}).Where("`acr_name` IN (?)", acrNames).Find(&results).Error

	return
}

// GetFromAcrCID 通过acr_c_id获取内容 国家ID
func (obj *_RAdCreativeMgr) GetFromAcrCID(acrCID int) (results []*RAdCreative, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreative{}).Where("`acr_c_id` = ?", acrCID).Find(&results).Error

	return
}

// GetBatchFromAcrCID 批量查找 国家ID
func (obj *_RAdCreativeMgr) GetBatchFromAcrCID(acrCIDs []int) (results []*RAdCreative, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreative{}).Where("`acr_c_id` IN (?)", acrCIDs).Find(&results).Error

	return
}

// GetFromAcrOID 通过acr_o_id获取内容
func (obj *_RAdCreativeMgr) GetFromAcrOID(acrOID int) (results []*RAdCreative, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreative{}).Where("`acr_o_id` = ?", acrOID).Find(&results).Error

	return
}

// GetBatchFromAcrOID 批量查找
func (obj *_RAdCreativeMgr) GetBatchFromAcrOID(acrOIDs []int) (results []*RAdCreative, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreative{}).Where("`acr_o_id` IN (?)", acrOIDs).Find(&results).Error

	return
}

// GetFromAcrURL 通过acr_url获取内容
func (obj *_RAdCreativeMgr) GetFromAcrURL(acrURL string) (results []*RAdCreative, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreative{}).Where("`acr_url` = ?", acrURL).Find(&results).Error

	return
}

// GetBatchFromAcrURL 批量查找
func (obj *_RAdCreativeMgr) GetBatchFromAcrURL(acrURLs []string) (results []*RAdCreative, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreative{}).Where("`acr_url` IN (?)", acrURLs).Find(&results).Error

	return
}

// GetFromAcrCloakURL 通过acr_cloak_url获取内容
func (obj *_RAdCreativeMgr) GetFromAcrCloakURL(acrCloakURL string) (results []*RAdCreative, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreative{}).Where("`acr_cloak_url` = ?", acrCloakURL).Find(&results).Error

	return
}

// GetBatchFromAcrCloakURL 批量查找
func (obj *_RAdCreativeMgr) GetBatchFromAcrCloakURL(acrCloakURLs []string) (results []*RAdCreative, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreative{}).Where("`acr_cloak_url` IN (?)", acrCloakURLs).Find(&results).Error

	return
}

// GetFromAcrWidth 通过acr_width获取内容
func (obj *_RAdCreativeMgr) GetFromAcrWidth(acrWidth float32) (results []*RAdCreative, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreative{}).Where("`acr_width` = ?", acrWidth).Find(&results).Error

	return
}

// GetBatchFromAcrWidth 批量查找
func (obj *_RAdCreativeMgr) GetBatchFromAcrWidth(acrWidths []float32) (results []*RAdCreative, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreative{}).Where("`acr_width` IN (?)", acrWidths).Find(&results).Error

	return
}

// GetFromAcrHeight 通过acr_height获取内容
func (obj *_RAdCreativeMgr) GetFromAcrHeight(acrHeight float32) (results []*RAdCreative, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreative{}).Where("`acr_height` = ?", acrHeight).Find(&results).Error

	return
}

// GetBatchFromAcrHeight 批量查找
func (obj *_RAdCreativeMgr) GetBatchFromAcrHeight(acrHeights []float32) (results []*RAdCreative, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreative{}).Where("`acr_height` IN (?)", acrHeights).Find(&results).Error

	return
}

// GetFromAcrStatus 通过acr_status获取内容
func (obj *_RAdCreativeMgr) GetFromAcrStatus(acrStatus int8) (results []*RAdCreative, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreative{}).Where("`acr_status` = ?", acrStatus).Find(&results).Error

	return
}

// GetBatchFromAcrStatus 批量查找
func (obj *_RAdCreativeMgr) GetBatchFromAcrStatus(acrStatuss []int8) (results []*RAdCreative, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreative{}).Where("`acr_status` IN (?)", acrStatuss).Find(&results).Error

	return
}

// GetFromAcrItID 通过acr_it_id获取内容 1 banner 2 video
func (obj *_RAdCreativeMgr) GetFromAcrItID(acrItID int) (results []*RAdCreative, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreative{}).Where("`acr_it_id` = ?", acrItID).Find(&results).Error

	return
}

// GetBatchFromAcrItID 批量查找 1 banner 2 video
func (obj *_RAdCreativeMgr) GetBatchFromAcrItID(acrItIDs []int) (results []*RAdCreative, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreative{}).Where("`acr_it_id` IN (?)", acrItIDs).Find(&results).Error

	return
}

// GetFromAcrExt 通过acr_ext获取内容
func (obj *_RAdCreativeMgr) GetFromAcrExt(acrExt string) (results []*RAdCreative, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreative{}).Where("`acr_ext` = ?", acrExt).Find(&results).Error

	return
}

// GetBatchFromAcrExt 批量查找
func (obj *_RAdCreativeMgr) GetBatchFromAcrExt(acrExts []string) (results []*RAdCreative, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreative{}).Where("`acr_ext` IN (?)", acrExts).Find(&results).Error

	return
}

// GetFromAcrRemark 通过acr_remark获取内容
func (obj *_RAdCreativeMgr) GetFromAcrRemark(acrRemark string) (results []*RAdCreative, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreative{}).Where("`acr_remark` = ?", acrRemark).Find(&results).Error

	return
}

// GetBatchFromAcrRemark 批量查找
func (obj *_RAdCreativeMgr) GetBatchFromAcrRemark(acrRemarks []string) (results []*RAdCreative, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreative{}).Where("`acr_remark` IN (?)", acrRemarks).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_RAdCreativeMgr) FetchByPrimaryKey(acrID uint32) (result RAdCreative, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreative{}).Where("`acr_id` = ?", acrID).Find(&result).Error

	return
}
