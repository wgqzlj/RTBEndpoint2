package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _RAdCreativeVideoMgr struct {
	*_BaseMgr
}

// RAdCreativeVideoMgr open func
func RAdCreativeVideoMgr(db *gorm.DB) *_RAdCreativeVideoMgr {
	if db == nil {
		panic(fmt.Errorf("RAdCreativeVideoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_RAdCreativeVideoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("r_ad_creative_video"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_RAdCreativeVideoMgr) GetTableName() string {
	return "r_ad_creative_video"
}

// Reset 重置gorm会话
func (obj *_RAdCreativeVideoMgr) Reset() *_RAdCreativeVideoMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_RAdCreativeVideoMgr) Get() (result RAdCreativeVideo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreativeVideo{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_RAdCreativeVideoMgr) Gets() (results []*RAdCreativeVideo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreativeVideo{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_RAdCreativeVideoMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(RAdCreativeVideo{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithAccvID accv_id获取
func (obj *_RAdCreativeVideoMgr) WithAccvID(accvID uint32) Option {
	return optionFunc(func(o *options) { o.query["accv_id"] = accvID })
}

// WithAccvAcrID accv_acr_id获取
func (obj *_RAdCreativeVideoMgr) WithAccvAcrID(accvAcrID int) Option {
	return optionFunc(func(o *options) { o.query["accv_acr_id"] = accvAcrID })
}

// WithAccvTitle accv_title获取
func (obj *_RAdCreativeVideoMgr) WithAccvTitle(accvTitle string) Option {
	return optionFunc(func(o *options) { o.query["accv_title"] = accvTitle })
}

// WithAccvDescription accv_description获取
func (obj *_RAdCreativeVideoMgr) WithAccvDescription(accvDescription string) Option {
	return optionFunc(func(o *options) { o.query["accv_description"] = accvDescription })
}

// WithAccvCta accv_cta获取
func (obj *_RAdCreativeVideoMgr) WithAccvCta(accvCta string) Option {
	return optionFunc(func(o *options) { o.query["accv_cta"] = accvCta })
}

// WithAccvDuration accv_duration获取
func (obj *_RAdCreativeVideoMgr) WithAccvDuration(accvDuration string) Option {
	return optionFunc(func(o *options) { o.query["accv_duration"] = accvDuration })
}

// WithAccvIcon accv_icon获取
func (obj *_RAdCreativeVideoMgr) WithAccvIcon(accvIcon string) Option {
	return optionFunc(func(o *options) { o.query["accv_icon"] = accvIcon })
}

// WithAccvCampanionImageURL accv_campanion_image_url获取
func (obj *_RAdCreativeVideoMgr) WithAccvCampanionImageURL(accvCampanionImageURL string) Option {
	return optionFunc(func(o *options) { o.query["accv_campanion_image_url"] = accvCampanionImageURL })
}

// WithAccvCampanionImageCloakURL accv_campanion_image_cloak_url获取
func (obj *_RAdCreativeVideoMgr) WithAccvCampanionImageCloakURL(accvCampanionImageCloakURL string) Option {
	return optionFunc(func(o *options) { o.query["accv_campanion_image_cloak_url"] = accvCampanionImageCloakURL })
}

// GetByOption 功能选项模式获取
func (obj *_RAdCreativeVideoMgr) GetByOption(opts ...Option) (result RAdCreativeVideo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(RAdCreativeVideo{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_RAdCreativeVideoMgr) GetByOptions(opts ...Option) (results []*RAdCreativeVideo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(RAdCreativeVideo{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromAccvID 通过accv_id获取内容
func (obj *_RAdCreativeVideoMgr) GetFromAccvID(accvID uint32) (result RAdCreativeVideo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreativeVideo{}).Where("`accv_id` = ?", accvID).Find(&result).Error

	return
}

// GetBatchFromAccvID 批量查找
func (obj *_RAdCreativeVideoMgr) GetBatchFromAccvID(accvIDs []uint32) (results []*RAdCreativeVideo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreativeVideo{}).Where("`accv_id` IN (?)", accvIDs).Find(&results).Error

	return
}

// GetFromAccvAcrID 通过accv_acr_id获取内容
func (obj *_RAdCreativeVideoMgr) GetFromAccvAcrID(accvAcrID int) (results []*RAdCreativeVideo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreativeVideo{}).Where("`accv_acr_id` = ?", accvAcrID).Find(&results).Error

	return
}

// GetBatchFromAccvAcrID 批量查找
func (obj *_RAdCreativeVideoMgr) GetBatchFromAccvAcrID(accvAcrIDs []int) (results []*RAdCreativeVideo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreativeVideo{}).Where("`accv_acr_id` IN (?)", accvAcrIDs).Find(&results).Error

	return
}

// GetFromAccvTitle 通过accv_title获取内容
func (obj *_RAdCreativeVideoMgr) GetFromAccvTitle(accvTitle string) (results []*RAdCreativeVideo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreativeVideo{}).Where("`accv_title` = ?", accvTitle).Find(&results).Error

	return
}

// GetBatchFromAccvTitle 批量查找
func (obj *_RAdCreativeVideoMgr) GetBatchFromAccvTitle(accvTitles []string) (results []*RAdCreativeVideo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreativeVideo{}).Where("`accv_title` IN (?)", accvTitles).Find(&results).Error

	return
}

// GetFromAccvDescription 通过accv_description获取内容
func (obj *_RAdCreativeVideoMgr) GetFromAccvDescription(accvDescription string) (results []*RAdCreativeVideo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreativeVideo{}).Where("`accv_description` = ?", accvDescription).Find(&results).Error

	return
}

// GetBatchFromAccvDescription 批量查找
func (obj *_RAdCreativeVideoMgr) GetBatchFromAccvDescription(accvDescriptions []string) (results []*RAdCreativeVideo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreativeVideo{}).Where("`accv_description` IN (?)", accvDescriptions).Find(&results).Error

	return
}

// GetFromAccvCta 通过accv_cta获取内容
func (obj *_RAdCreativeVideoMgr) GetFromAccvCta(accvCta string) (results []*RAdCreativeVideo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreativeVideo{}).Where("`accv_cta` = ?", accvCta).Find(&results).Error

	return
}

// GetBatchFromAccvCta 批量查找
func (obj *_RAdCreativeVideoMgr) GetBatchFromAccvCta(accvCtas []string) (results []*RAdCreativeVideo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreativeVideo{}).Where("`accv_cta` IN (?)", accvCtas).Find(&results).Error

	return
}

// GetFromAccvDuration 通过accv_duration获取内容
func (obj *_RAdCreativeVideoMgr) GetFromAccvDuration(accvDuration string) (results []*RAdCreativeVideo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreativeVideo{}).Where("`accv_duration` = ?", accvDuration).Find(&results).Error

	return
}

// GetBatchFromAccvDuration 批量查找
func (obj *_RAdCreativeVideoMgr) GetBatchFromAccvDuration(accvDurations []string) (results []*RAdCreativeVideo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreativeVideo{}).Where("`accv_duration` IN (?)", accvDurations).Find(&results).Error

	return
}

// GetFromAccvIcon 通过accv_icon获取内容
func (obj *_RAdCreativeVideoMgr) GetFromAccvIcon(accvIcon string) (results []*RAdCreativeVideo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreativeVideo{}).Where("`accv_icon` = ?", accvIcon).Find(&results).Error

	return
}

// GetBatchFromAccvIcon 批量查找
func (obj *_RAdCreativeVideoMgr) GetBatchFromAccvIcon(accvIcons []string) (results []*RAdCreativeVideo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreativeVideo{}).Where("`accv_icon` IN (?)", accvIcons).Find(&results).Error

	return
}

// GetFromAccvCampanionImageURL 通过accv_campanion_image_url获取内容
func (obj *_RAdCreativeVideoMgr) GetFromAccvCampanionImageURL(accvCampanionImageURL string) (results []*RAdCreativeVideo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreativeVideo{}).Where("`accv_campanion_image_url` = ?", accvCampanionImageURL).Find(&results).Error

	return
}

// GetBatchFromAccvCampanionImageURL 批量查找
func (obj *_RAdCreativeVideoMgr) GetBatchFromAccvCampanionImageURL(accvCampanionImageURLs []string) (results []*RAdCreativeVideo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreativeVideo{}).Where("`accv_campanion_image_url` IN (?)", accvCampanionImageURLs).Find(&results).Error

	return
}

// GetFromAccvCampanionImageCloakURL 通过accv_campanion_image_cloak_url获取内容
func (obj *_RAdCreativeVideoMgr) GetFromAccvCampanionImageCloakURL(accvCampanionImageCloakURL string) (results []*RAdCreativeVideo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreativeVideo{}).Where("`accv_campanion_image_cloak_url` = ?", accvCampanionImageCloakURL).Find(&results).Error

	return
}

// GetBatchFromAccvCampanionImageCloakURL 批量查找
func (obj *_RAdCreativeVideoMgr) GetBatchFromAccvCampanionImageCloakURL(accvCampanionImageCloakURLs []string) (results []*RAdCreativeVideo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreativeVideo{}).Where("`accv_campanion_image_cloak_url` IN (?)", accvCampanionImageCloakURLs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_RAdCreativeVideoMgr) FetchByPrimaryKey(accvID uint32) (result RAdCreativeVideo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreativeVideo{}).Where("`accv_id` = ?", accvID).Find(&result).Error

	return
}
