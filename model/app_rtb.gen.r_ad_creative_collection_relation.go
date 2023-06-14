package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _RAdCreativeCollectionRelationMgr struct {
	*_BaseMgr
}

// RAdCreativeCollectionRelationMgr open func
func RAdCreativeCollectionRelationMgr(db *gorm.DB) *_RAdCreativeCollectionRelationMgr {
	if db == nil {
		panic(fmt.Errorf("RAdCreativeCollectionRelationMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_RAdCreativeCollectionRelationMgr{_BaseMgr: &_BaseMgr{DB: db.Table("r_ad_creative_collection_relation"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_RAdCreativeCollectionRelationMgr) GetTableName() string {
	return "r_ad_creative_collection_relation"
}

// Reset 重置gorm会话
func (obj *_RAdCreativeCollectionRelationMgr) Reset() *_RAdCreativeCollectionRelationMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_RAdCreativeCollectionRelationMgr) Get() (result RAdCreativeCollectionRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreativeCollectionRelation{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_RAdCreativeCollectionRelationMgr) Gets() (results []*RAdCreativeCollectionRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreativeCollectionRelation{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_RAdCreativeCollectionRelationMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(RAdCreativeCollectionRelation{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithAcrcrID acrcr_id获取
func (obj *_RAdCreativeCollectionRelationMgr) WithAcrcrID(acrcrID uint32) Option {
	return optionFunc(func(o *options) { o.query["acrcr_id"] = acrcrID })
}

// WithAcrcrAcrcID acrcr_acrc_id获取
func (obj *_RAdCreativeCollectionRelationMgr) WithAcrcrAcrcID(acrcrAcrcID int) Option {
	return optionFunc(func(o *options) { o.query["acrcr_acrc_id"] = acrcrAcrcID })
}

// WithAcrcrAcrID acrcr_acr_id获取
func (obj *_RAdCreativeCollectionRelationMgr) WithAcrcrAcrID(acrcrAcrID int) Option {
	return optionFunc(func(o *options) { o.query["acrcr_acr_id"] = acrcrAcrID })
}

// WithAcrcrAeID acrcr_ae_id获取
func (obj *_RAdCreativeCollectionRelationMgr) WithAcrcrAeID(acrcrAeID int) Option {
	return optionFunc(func(o *options) { o.query["acrcr_ae_id"] = acrcrAeID })
}

// WithAcrcrIstID acrcr_ist_id获取
func (obj *_RAdCreativeCollectionRelationMgr) WithAcrcrIstID(acrcrIstID int) Option {
	return optionFunc(func(o *options) { o.query["acrcr_ist_id"] = acrcrIstID })
}

// WithAcrcrStatus acrcr_status获取
func (obj *_RAdCreativeCollectionRelationMgr) WithAcrcrStatus(acrcrStatus int8) Option {
	return optionFunc(func(o *options) { o.query["acrcr_status"] = acrcrStatus })
}

// WithAcrcrCreateTime acrcr_create_time获取
func (obj *_RAdCreativeCollectionRelationMgr) WithAcrcrCreateTime(acrcrCreateTime int) Option {
	return optionFunc(func(o *options) { o.query["acrcr_create_time"] = acrcrCreateTime })
}

// WithAcrcrAtID acrcr_at_id获取
func (obj *_RAdCreativeCollectionRelationMgr) WithAcrcrAtID(acrcrAtID int) Option {
	return optionFunc(func(o *options) { o.query["acrcr_at_id"] = acrcrAtID })
}

// GetByOption 功能选项模式获取
func (obj *_RAdCreativeCollectionRelationMgr) GetByOption(opts ...Option) (result RAdCreativeCollectionRelation, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(RAdCreativeCollectionRelation{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_RAdCreativeCollectionRelationMgr) GetByOptions(opts ...Option) (results []*RAdCreativeCollectionRelation, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(RAdCreativeCollectionRelation{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromAcrcrID 通过acrcr_id获取内容
func (obj *_RAdCreativeCollectionRelationMgr) GetFromAcrcrID(acrcrID uint32) (result RAdCreativeCollectionRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreativeCollectionRelation{}).Where("`acrcr_id` = ?", acrcrID).Find(&result).Error

	return
}

// GetBatchFromAcrcrID 批量查找
func (obj *_RAdCreativeCollectionRelationMgr) GetBatchFromAcrcrID(acrcrIDs []uint32) (results []*RAdCreativeCollectionRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreativeCollectionRelation{}).Where("`acrcr_id` IN (?)", acrcrIDs).Find(&results).Error

	return
}

// GetFromAcrcrAcrcID 通过acrcr_acrc_id获取内容
func (obj *_RAdCreativeCollectionRelationMgr) GetFromAcrcrAcrcID(acrcrAcrcID int) (results []*RAdCreativeCollectionRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreativeCollectionRelation{}).Where("`acrcr_acrc_id` = ?", acrcrAcrcID).Find(&results).Error

	return
}

// GetBatchFromAcrcrAcrcID 批量查找
func (obj *_RAdCreativeCollectionRelationMgr) GetBatchFromAcrcrAcrcID(acrcrAcrcIDs []int) (results []*RAdCreativeCollectionRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreativeCollectionRelation{}).Where("`acrcr_acrc_id` IN (?)", acrcrAcrcIDs).Find(&results).Error

	return
}

// GetFromAcrcrAcrID 通过acrcr_acr_id获取内容
func (obj *_RAdCreativeCollectionRelationMgr) GetFromAcrcrAcrID(acrcrAcrID int) (results []*RAdCreativeCollectionRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreativeCollectionRelation{}).Where("`acrcr_acr_id` = ?", acrcrAcrID).Find(&results).Error

	return
}

// GetBatchFromAcrcrAcrID 批量查找
func (obj *_RAdCreativeCollectionRelationMgr) GetBatchFromAcrcrAcrID(acrcrAcrIDs []int) (results []*RAdCreativeCollectionRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreativeCollectionRelation{}).Where("`acrcr_acr_id` IN (?)", acrcrAcrIDs).Find(&results).Error

	return
}

// GetFromAcrcrAeID 通过acrcr_ae_id获取内容
func (obj *_RAdCreativeCollectionRelationMgr) GetFromAcrcrAeID(acrcrAeID int) (results []*RAdCreativeCollectionRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreativeCollectionRelation{}).Where("`acrcr_ae_id` = ?", acrcrAeID).Find(&results).Error

	return
}

// GetBatchFromAcrcrAeID 批量查找
func (obj *_RAdCreativeCollectionRelationMgr) GetBatchFromAcrcrAeID(acrcrAeIDs []int) (results []*RAdCreativeCollectionRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreativeCollectionRelation{}).Where("`acrcr_ae_id` IN (?)", acrcrAeIDs).Find(&results).Error

	return
}

// GetFromAcrcrIstID 通过acrcr_ist_id获取内容
func (obj *_RAdCreativeCollectionRelationMgr) GetFromAcrcrIstID(acrcrIstID int) (results []*RAdCreativeCollectionRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreativeCollectionRelation{}).Where("`acrcr_ist_id` = ?", acrcrIstID).Find(&results).Error

	return
}

// GetBatchFromAcrcrIstID 批量查找
func (obj *_RAdCreativeCollectionRelationMgr) GetBatchFromAcrcrIstID(acrcrIstIDs []int) (results []*RAdCreativeCollectionRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreativeCollectionRelation{}).Where("`acrcr_ist_id` IN (?)", acrcrIstIDs).Find(&results).Error

	return
}

// GetFromAcrcrStatus 通过acrcr_status获取内容
func (obj *_RAdCreativeCollectionRelationMgr) GetFromAcrcrStatus(acrcrStatus int8) (results []*RAdCreativeCollectionRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreativeCollectionRelation{}).Where("`acrcr_status` = ?", acrcrStatus).Find(&results).Error

	return
}

// GetBatchFromAcrcrStatus 批量查找
func (obj *_RAdCreativeCollectionRelationMgr) GetBatchFromAcrcrStatus(acrcrStatuss []int8) (results []*RAdCreativeCollectionRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreativeCollectionRelation{}).Where("`acrcr_status` IN (?)", acrcrStatuss).Find(&results).Error

	return
}

// GetFromAcrcrCreateTime 通过acrcr_create_time获取内容
func (obj *_RAdCreativeCollectionRelationMgr) GetFromAcrcrCreateTime(acrcrCreateTime int) (results []*RAdCreativeCollectionRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreativeCollectionRelation{}).Where("`acrcr_create_time` = ?", acrcrCreateTime).Find(&results).Error

	return
}

// GetBatchFromAcrcrCreateTime 批量查找
func (obj *_RAdCreativeCollectionRelationMgr) GetBatchFromAcrcrCreateTime(acrcrCreateTimes []int) (results []*RAdCreativeCollectionRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreativeCollectionRelation{}).Where("`acrcr_create_time` IN (?)", acrcrCreateTimes).Find(&results).Error

	return
}

// GetFromAcrcrAtID 通过acrcr_at_id获取内容
func (obj *_RAdCreativeCollectionRelationMgr) GetFromAcrcrAtID(acrcrAtID int) (results []*RAdCreativeCollectionRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreativeCollectionRelation{}).Where("`acrcr_at_id` = ?", acrcrAtID).Find(&results).Error

	return
}

// GetBatchFromAcrcrAtID 批量查找
func (obj *_RAdCreativeCollectionRelationMgr) GetBatchFromAcrcrAtID(acrcrAtIDs []int) (results []*RAdCreativeCollectionRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreativeCollectionRelation{}).Where("`acrcr_at_id` IN (?)", acrcrAtIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_RAdCreativeCollectionRelationMgr) FetchByPrimaryKey(acrcrID uint32) (result RAdCreativeCollectionRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreativeCollectionRelation{}).Where("`acrcr_id` = ?", acrcrID).Find(&result).Error

	return
}

// FetchIndexByAcrcrAcrcID  获取多个内容
func (obj *_RAdCreativeCollectionRelationMgr) FetchIndexByAcrcrAcrcID(acrcrAcrcID int) (results []*RAdCreativeCollectionRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreativeCollectionRelation{}).Where("`acrcr_acrc_id` = ?", acrcrAcrcID).Find(&results).Error

	return
}

// FetchIndexByAcrcrAcrID  获取多个内容
func (obj *_RAdCreativeCollectionRelationMgr) FetchIndexByAcrcrAcrID(acrcrAcrID int) (results []*RAdCreativeCollectionRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreativeCollectionRelation{}).Where("`acrcr_acr_id` = ?", acrcrAcrID).Find(&results).Error

	return
}

// FetchIndexByAcrcrAeID  获取多个内容
func (obj *_RAdCreativeCollectionRelationMgr) FetchIndexByAcrcrAeID(acrcrAeID int) (results []*RAdCreativeCollectionRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreativeCollectionRelation{}).Where("`acrcr_ae_id` = ?", acrcrAeID).Find(&results).Error

	return
}

// FetchIndexByAcrcrIstID  获取多个内容
func (obj *_RAdCreativeCollectionRelationMgr) FetchIndexByAcrcrIstID(acrcrIstID int) (results []*RAdCreativeCollectionRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreativeCollectionRelation{}).Where("`acrcr_ist_id` = ?", acrcrIstID).Find(&results).Error

	return
}

// FetchIndexByAcrcrStatus  获取多个内容
func (obj *_RAdCreativeCollectionRelationMgr) FetchIndexByAcrcrStatus(acrcrStatus int8) (results []*RAdCreativeCollectionRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreativeCollectionRelation{}).Where("`acrcr_status` = ?", acrcrStatus).Find(&results).Error

	return
}

// FetchIndexByAcrcrAtID  获取多个内容
func (obj *_RAdCreativeCollectionRelationMgr) FetchIndexByAcrcrAtID(acrcrAtID int) (results []*RAdCreativeCollectionRelation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreativeCollectionRelation{}).Where("`acrcr_at_id` = ?", acrcrAtID).Find(&results).Error

	return
}
