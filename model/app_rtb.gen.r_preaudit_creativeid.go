package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _RPreauditCreativeidMgr struct {
	*_BaseMgr
}

// RPreauditCreativeidMgr open func
func RPreauditCreativeidMgr(db *gorm.DB) *_RPreauditCreativeidMgr {
	if db == nil {
		panic(fmt.Errorf("RPreauditCreativeidMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_RPreauditCreativeidMgr{_BaseMgr: &_BaseMgr{DB: db.Table("r_preaudit_creativeid"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_RPreauditCreativeidMgr) GetTableName() string {
	return "r_preaudit_creativeid"
}

// Reset 重置gorm会话
func (obj *_RPreauditCreativeidMgr) Reset() *_RPreauditCreativeidMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_RPreauditCreativeidMgr) Get() (result RPreauditCreativeid, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RPreauditCreativeid{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_RPreauditCreativeidMgr) Gets() (results []*RPreauditCreativeid, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RPreauditCreativeid{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_RPreauditCreativeidMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(RPreauditCreativeid{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithPcID pc_id获取
func (obj *_RPreauditCreativeidMgr) WithPcID(pcID uint32) Option {
	return optionFunc(func(o *options) { o.query["pc_id"] = pcID })
}

// WithPcCode pc_code获取
func (obj *_RPreauditCreativeidMgr) WithPcCode(pcCode string) Option {
	return optionFunc(func(o *options) { o.query["pc_code"] = pcCode })
}

// WithPcAeID pc_ae_id获取
func (obj *_RPreauditCreativeidMgr) WithPcAeID(pcAeID int) Option {
	return optionFunc(func(o *options) { o.query["pc_ae_id"] = pcAeID })
}

// WithPcCID pc_c_id获取
func (obj *_RPreauditCreativeidMgr) WithPcCID(pcCID int) Option {
	return optionFunc(func(o *options) { o.query["pc_c_id"] = pcCID })
}

// WithPcCreateTime pc_create_time获取
func (obj *_RPreauditCreativeidMgr) WithPcCreateTime(pcCreateTime uint32) Option {
	return optionFunc(func(o *options) { o.query["pc_create_time"] = pcCreateTime })
}

// WithPcPreAuditTime pc_pre_audit_time获取
func (obj *_RPreauditCreativeidMgr) WithPcPreAuditTime(pcPreAuditTime uint32) Option {
	return optionFunc(func(o *options) { o.query["pc_pre_audit_time"] = pcPreAuditTime })
}

// WithPcStartEnableTime pc_start_enable_time获取
func (obj *_RPreauditCreativeidMgr) WithPcStartEnableTime(pcStartEnableTime uint32) Option {
	return optionFunc(func(o *options) { o.query["pc_start_enable_time"] = pcStartEnableTime })
}

// WithPcStatus pc_status获取 1 正常 0删除 2待提交 3 被封禁
func (obj *_RPreauditCreativeidMgr) WithPcStatus(pcStatus int8) Option {
	return optionFunc(func(o *options) { o.query["pc_status"] = pcStatus })
}

// WithPcAcID pc_ac_id获取
func (obj *_RPreauditCreativeidMgr) WithPcAcID(pcAcID int) Option {
	return optionFunc(func(o *options) { o.query["pc_ac_id"] = pcAcID })
}

// WithPcSubmitCount pc_submit_count获取
func (obj *_RPreauditCreativeidMgr) WithPcSubmitCount(pcSubmitCount int) Option {
	return optionFunc(func(o *options) { o.query["pc_submit_count"] = pcSubmitCount })
}

// GetByOption 功能选项模式获取
func (obj *_RPreauditCreativeidMgr) GetByOption(opts ...Option) (result RPreauditCreativeid, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(RPreauditCreativeid{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_RPreauditCreativeidMgr) GetByOptions(opts ...Option) (results []*RPreauditCreativeid, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(RPreauditCreativeid{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromPcID 通过pc_id获取内容
func (obj *_RPreauditCreativeidMgr) GetFromPcID(pcID uint32) (result RPreauditCreativeid, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RPreauditCreativeid{}).Where("`pc_id` = ?", pcID).Find(&result).Error

	return
}

// GetBatchFromPcID 批量查找
func (obj *_RPreauditCreativeidMgr) GetBatchFromPcID(pcIDs []uint32) (results []*RPreauditCreativeid, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RPreauditCreativeid{}).Where("`pc_id` IN (?)", pcIDs).Find(&results).Error

	return
}

// GetFromPcCode 通过pc_code获取内容
func (obj *_RPreauditCreativeidMgr) GetFromPcCode(pcCode string) (results []*RPreauditCreativeid, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RPreauditCreativeid{}).Where("`pc_code` = ?", pcCode).Find(&results).Error

	return
}

// GetBatchFromPcCode 批量查找
func (obj *_RPreauditCreativeidMgr) GetBatchFromPcCode(pcCodes []string) (results []*RPreauditCreativeid, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RPreauditCreativeid{}).Where("`pc_code` IN (?)", pcCodes).Find(&results).Error

	return
}

// GetFromPcAeID 通过pc_ae_id获取内容
func (obj *_RPreauditCreativeidMgr) GetFromPcAeID(pcAeID int) (results []*RPreauditCreativeid, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RPreauditCreativeid{}).Where("`pc_ae_id` = ?", pcAeID).Find(&results).Error

	return
}

// GetBatchFromPcAeID 批量查找
func (obj *_RPreauditCreativeidMgr) GetBatchFromPcAeID(pcAeIDs []int) (results []*RPreauditCreativeid, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RPreauditCreativeid{}).Where("`pc_ae_id` IN (?)", pcAeIDs).Find(&results).Error

	return
}

// GetFromPcCID 通过pc_c_id获取内容
func (obj *_RPreauditCreativeidMgr) GetFromPcCID(pcCID int) (results []*RPreauditCreativeid, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RPreauditCreativeid{}).Where("`pc_c_id` = ?", pcCID).Find(&results).Error

	return
}

// GetBatchFromPcCID 批量查找
func (obj *_RPreauditCreativeidMgr) GetBatchFromPcCID(pcCIDs []int) (results []*RPreauditCreativeid, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RPreauditCreativeid{}).Where("`pc_c_id` IN (?)", pcCIDs).Find(&results).Error

	return
}

// GetFromPcCreateTime 通过pc_create_time获取内容
func (obj *_RPreauditCreativeidMgr) GetFromPcCreateTime(pcCreateTime uint32) (results []*RPreauditCreativeid, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RPreauditCreativeid{}).Where("`pc_create_time` = ?", pcCreateTime).Find(&results).Error

	return
}

// GetBatchFromPcCreateTime 批量查找
func (obj *_RPreauditCreativeidMgr) GetBatchFromPcCreateTime(pcCreateTimes []uint32) (results []*RPreauditCreativeid, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RPreauditCreativeid{}).Where("`pc_create_time` IN (?)", pcCreateTimes).Find(&results).Error

	return
}

// GetFromPcPreAuditTime 通过pc_pre_audit_time获取内容
func (obj *_RPreauditCreativeidMgr) GetFromPcPreAuditTime(pcPreAuditTime uint32) (results []*RPreauditCreativeid, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RPreauditCreativeid{}).Where("`pc_pre_audit_time` = ?", pcPreAuditTime).Find(&results).Error

	return
}

// GetBatchFromPcPreAuditTime 批量查找
func (obj *_RPreauditCreativeidMgr) GetBatchFromPcPreAuditTime(pcPreAuditTimes []uint32) (results []*RPreauditCreativeid, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RPreauditCreativeid{}).Where("`pc_pre_audit_time` IN (?)", pcPreAuditTimes).Find(&results).Error

	return
}

// GetFromPcStartEnableTime 通过pc_start_enable_time获取内容
func (obj *_RPreauditCreativeidMgr) GetFromPcStartEnableTime(pcStartEnableTime uint32) (results []*RPreauditCreativeid, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RPreauditCreativeid{}).Where("`pc_start_enable_time` = ?", pcStartEnableTime).Find(&results).Error

	return
}

// GetBatchFromPcStartEnableTime 批量查找
func (obj *_RPreauditCreativeidMgr) GetBatchFromPcStartEnableTime(pcStartEnableTimes []uint32) (results []*RPreauditCreativeid, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RPreauditCreativeid{}).Where("`pc_start_enable_time` IN (?)", pcStartEnableTimes).Find(&results).Error

	return
}

// GetFromPcStatus 通过pc_status获取内容 1 正常 0删除 2待提交 3 被封禁
func (obj *_RPreauditCreativeidMgr) GetFromPcStatus(pcStatus int8) (results []*RPreauditCreativeid, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RPreauditCreativeid{}).Where("`pc_status` = ?", pcStatus).Find(&results).Error

	return
}

// GetBatchFromPcStatus 批量查找 1 正常 0删除 2待提交 3 被封禁
func (obj *_RPreauditCreativeidMgr) GetBatchFromPcStatus(pcStatuss []int8) (results []*RPreauditCreativeid, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RPreauditCreativeid{}).Where("`pc_status` IN (?)", pcStatuss).Find(&results).Error

	return
}

// GetFromPcAcID 通过pc_ac_id获取内容
func (obj *_RPreauditCreativeidMgr) GetFromPcAcID(pcAcID int) (results []*RPreauditCreativeid, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RPreauditCreativeid{}).Where("`pc_ac_id` = ?", pcAcID).Find(&results).Error

	return
}

// GetBatchFromPcAcID 批量查找
func (obj *_RPreauditCreativeidMgr) GetBatchFromPcAcID(pcAcIDs []int) (results []*RPreauditCreativeid, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RPreauditCreativeid{}).Where("`pc_ac_id` IN (?)", pcAcIDs).Find(&results).Error

	return
}

// GetFromPcSubmitCount 通过pc_submit_count获取内容
func (obj *_RPreauditCreativeidMgr) GetFromPcSubmitCount(pcSubmitCount int) (results []*RPreauditCreativeid, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RPreauditCreativeid{}).Where("`pc_submit_count` = ?", pcSubmitCount).Find(&results).Error

	return
}

// GetBatchFromPcSubmitCount 批量查找
func (obj *_RPreauditCreativeidMgr) GetBatchFromPcSubmitCount(pcSubmitCounts []int) (results []*RPreauditCreativeid, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RPreauditCreativeid{}).Where("`pc_submit_count` IN (?)", pcSubmitCounts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_RPreauditCreativeidMgr) FetchByPrimaryKey(pcID uint32) (result RPreauditCreativeid, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RPreauditCreativeid{}).Where("`pc_id` = ?", pcID).Find(&result).Error

	return
}
