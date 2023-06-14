package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _RAdCreativeCollectionMgr struct {
	*_BaseMgr
}

// RAdCreativeCollectionMgr open func
func RAdCreativeCollectionMgr(db *gorm.DB) *_RAdCreativeCollectionMgr {
	if db == nil {
		panic(fmt.Errorf("RAdCreativeCollectionMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_RAdCreativeCollectionMgr{_BaseMgr: &_BaseMgr{DB: db.Table("r_ad_creative_collection"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_RAdCreativeCollectionMgr) GetTableName() string {
	return "r_ad_creative_collection"
}

// Reset 重置gorm会话
func (obj *_RAdCreativeCollectionMgr) Reset() *_RAdCreativeCollectionMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_RAdCreativeCollectionMgr) Get() (result RAdCreativeCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreativeCollection{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_RAdCreativeCollectionMgr) Gets() (results []*RAdCreativeCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreativeCollection{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_RAdCreativeCollectionMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(RAdCreativeCollection{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithAcrcID acrc_id获取
func (obj *_RAdCreativeCollectionMgr) WithAcrcID(acrcID uint32) Option {
	return optionFunc(func(o *options) { o.query["acrc_id"] = acrcID })
}

// WithAcrcName acrc_name获取
func (obj *_RAdCreativeCollectionMgr) WithAcrcName(acrcName string) Option {
	return optionFunc(func(o *options) { o.query["acrc_name"] = acrcName })
}

// WithAcrcCreateTime acrc_create_time获取
func (obj *_RAdCreativeCollectionMgr) WithAcrcCreateTime(acrcCreateTime uint32) Option {
	return optionFunc(func(o *options) { o.query["acrc_create_time"] = acrcCreateTime })
}

// WithAcrcStatus acrc_status获取
func (obj *_RAdCreativeCollectionMgr) WithAcrcStatus(acrcStatus int8) Option {
	return optionFunc(func(o *options) { o.query["acrc_status"] = acrcStatus })
}

// WithAcrcCID acrc_c_id获取
func (obj *_RAdCreativeCollectionMgr) WithAcrcCID(acrcCID int) Option {
	return optionFunc(func(o *options) { o.query["acrc_c_id"] = acrcCID })
}

// WithAcrcOID acrc_o_id获取
func (obj *_RAdCreativeCollectionMgr) WithAcrcOID(acrcOID int) Option {
	return optionFunc(func(o *options) { o.query["acrc_o_id"] = acrcOID })
}

// GetByOption 功能选项模式获取
func (obj *_RAdCreativeCollectionMgr) GetByOption(opts ...Option) (result RAdCreativeCollection, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(RAdCreativeCollection{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_RAdCreativeCollectionMgr) GetByOptions(opts ...Option) (results []*RAdCreativeCollection, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(RAdCreativeCollection{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromAcrcID 通过acrc_id获取内容
func (obj *_RAdCreativeCollectionMgr) GetFromAcrcID(acrcID uint32) (result RAdCreativeCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreativeCollection{}).Where("`acrc_id` = ?", acrcID).Find(&result).Error

	return
}

// GetBatchFromAcrcID 批量查找
func (obj *_RAdCreativeCollectionMgr) GetBatchFromAcrcID(acrcIDs []uint32) (results []*RAdCreativeCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreativeCollection{}).Where("`acrc_id` IN (?)", acrcIDs).Find(&results).Error

	return
}

// GetFromAcrcName 通过acrc_name获取内容
func (obj *_RAdCreativeCollectionMgr) GetFromAcrcName(acrcName string) (results []*RAdCreativeCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreativeCollection{}).Where("`acrc_name` = ?", acrcName).Find(&results).Error

	return
}

// GetBatchFromAcrcName 批量查找
func (obj *_RAdCreativeCollectionMgr) GetBatchFromAcrcName(acrcNames []string) (results []*RAdCreativeCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreativeCollection{}).Where("`acrc_name` IN (?)", acrcNames).Find(&results).Error

	return
}

// GetFromAcrcCreateTime 通过acrc_create_time获取内容
func (obj *_RAdCreativeCollectionMgr) GetFromAcrcCreateTime(acrcCreateTime uint32) (results []*RAdCreativeCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreativeCollection{}).Where("`acrc_create_time` = ?", acrcCreateTime).Find(&results).Error

	return
}

// GetBatchFromAcrcCreateTime 批量查找
func (obj *_RAdCreativeCollectionMgr) GetBatchFromAcrcCreateTime(acrcCreateTimes []uint32) (results []*RAdCreativeCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreativeCollection{}).Where("`acrc_create_time` IN (?)", acrcCreateTimes).Find(&results).Error

	return
}

// GetFromAcrcStatus 通过acrc_status获取内容
func (obj *_RAdCreativeCollectionMgr) GetFromAcrcStatus(acrcStatus int8) (results []*RAdCreativeCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreativeCollection{}).Where("`acrc_status` = ?", acrcStatus).Find(&results).Error

	return
}

// GetBatchFromAcrcStatus 批量查找
func (obj *_RAdCreativeCollectionMgr) GetBatchFromAcrcStatus(acrcStatuss []int8) (results []*RAdCreativeCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreativeCollection{}).Where("`acrc_status` IN (?)", acrcStatuss).Find(&results).Error

	return
}

// GetFromAcrcCID 通过acrc_c_id获取内容
func (obj *_RAdCreativeCollectionMgr) GetFromAcrcCID(acrcCID int) (results []*RAdCreativeCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreativeCollection{}).Where("`acrc_c_id` = ?", acrcCID).Find(&results).Error

	return
}

// GetBatchFromAcrcCID 批量查找
func (obj *_RAdCreativeCollectionMgr) GetBatchFromAcrcCID(acrcCIDs []int) (results []*RAdCreativeCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreativeCollection{}).Where("`acrc_c_id` IN (?)", acrcCIDs).Find(&results).Error

	return
}

// GetFromAcrcOID 通过acrc_o_id获取内容
func (obj *_RAdCreativeCollectionMgr) GetFromAcrcOID(acrcOID int) (results []*RAdCreativeCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreativeCollection{}).Where("`acrc_o_id` = ?", acrcOID).Find(&results).Error

	return
}

// GetBatchFromAcrcOID 批量查找
func (obj *_RAdCreativeCollectionMgr) GetBatchFromAcrcOID(acrcOIDs []int) (results []*RAdCreativeCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreativeCollection{}).Where("`acrc_o_id` IN (?)", acrcOIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_RAdCreativeCollectionMgr) FetchByPrimaryKey(acrcID uint32) (result RAdCreativeCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreativeCollection{}).Where("`acrc_id` = ?", acrcID).Find(&result).Error

	return
}

// FetchIndexByAcrcCID  获取多个内容
func (obj *_RAdCreativeCollectionMgr) FetchIndexByAcrcCID(acrcCID int) (results []*RAdCreativeCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreativeCollection{}).Where("`acrc_c_id` = ?", acrcCID).Find(&results).Error

	return
}

// FetchIndexByAcrcOID  获取多个内容
func (obj *_RAdCreativeCollectionMgr) FetchIndexByAcrcOID(acrcOID int) (results []*RAdCreativeCollection, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdCreativeCollection{}).Where("`acrc_o_id` = ?", acrcOID).Find(&results).Error

	return
}
