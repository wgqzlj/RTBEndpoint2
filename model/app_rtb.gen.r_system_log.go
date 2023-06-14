package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _RSystemLogMgr struct {
	*_BaseMgr
}

// RSystemLogMgr open func
func RSystemLogMgr(db *gorm.DB) *_RSystemLogMgr {
	if db == nil {
		panic(fmt.Errorf("RSystemLogMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_RSystemLogMgr{_BaseMgr: &_BaseMgr{DB: db.Table("r_system_log"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_RSystemLogMgr) GetTableName() string {
	return "r_system_log"
}

// Reset 重置gorm会话
func (obj *_RSystemLogMgr) Reset() *_RSystemLogMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_RSystemLogMgr) Get() (result RSystemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSystemLog{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_RSystemLogMgr) Gets() (results []*RSystemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSystemLog{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_RSystemLogMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(RSystemLog{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithLrID lr_id获取
func (obj *_RSystemLogMgr) WithLrID(lrID uint32) Option {
	return optionFunc(func(o *options) { o.query["lr_id"] = lrID })
}

// WithLrUID lr_u_id获取
func (obj *_RSystemLogMgr) WithLrUID(lrUID int) Option {
	return optionFunc(func(o *options) { o.query["lr_u_id"] = lrUID })
}

// WithLrCode lr_code获取
func (obj *_RSystemLogMgr) WithLrCode(lrCode string) Option {
	return optionFunc(func(o *options) { o.query["lr_code"] = lrCode })
}

// WithLrMessage lr_message获取
func (obj *_RSystemLogMgr) WithLrMessage(lrMessage string) Option {
	return optionFunc(func(o *options) { o.query["lr_message"] = lrMessage })
}

// WithLrCreateTime lr_create_time获取
func (obj *_RSystemLogMgr) WithLrCreateTime(lrCreateTime int) Option {
	return optionFunc(func(o *options) { o.query["lr_create_time"] = lrCreateTime })
}

// GetByOption 功能选项模式获取
func (obj *_RSystemLogMgr) GetByOption(opts ...Option) (result RSystemLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(RSystemLog{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_RSystemLogMgr) GetByOptions(opts ...Option) (results []*RSystemLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(RSystemLog{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromLrID 通过lr_id获取内容
func (obj *_RSystemLogMgr) GetFromLrID(lrID uint32) (result RSystemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSystemLog{}).Where("`lr_id` = ?", lrID).Find(&result).Error

	return
}

// GetBatchFromLrID 批量查找
func (obj *_RSystemLogMgr) GetBatchFromLrID(lrIDs []uint32) (results []*RSystemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSystemLog{}).Where("`lr_id` IN (?)", lrIDs).Find(&results).Error

	return
}

// GetFromLrUID 通过lr_u_id获取内容
func (obj *_RSystemLogMgr) GetFromLrUID(lrUID int) (results []*RSystemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSystemLog{}).Where("`lr_u_id` = ?", lrUID).Find(&results).Error

	return
}

// GetBatchFromLrUID 批量查找
func (obj *_RSystemLogMgr) GetBatchFromLrUID(lrUIDs []int) (results []*RSystemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSystemLog{}).Where("`lr_u_id` IN (?)", lrUIDs).Find(&results).Error

	return
}

// GetFromLrCode 通过lr_code获取内容
func (obj *_RSystemLogMgr) GetFromLrCode(lrCode string) (results []*RSystemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSystemLog{}).Where("`lr_code` = ?", lrCode).Find(&results).Error

	return
}

// GetBatchFromLrCode 批量查找
func (obj *_RSystemLogMgr) GetBatchFromLrCode(lrCodes []string) (results []*RSystemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSystemLog{}).Where("`lr_code` IN (?)", lrCodes).Find(&results).Error

	return
}

// GetFromLrMessage 通过lr_message获取内容
func (obj *_RSystemLogMgr) GetFromLrMessage(lrMessage string) (results []*RSystemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSystemLog{}).Where("`lr_message` = ?", lrMessage).Find(&results).Error

	return
}

// GetBatchFromLrMessage 批量查找
func (obj *_RSystemLogMgr) GetBatchFromLrMessage(lrMessages []string) (results []*RSystemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSystemLog{}).Where("`lr_message` IN (?)", lrMessages).Find(&results).Error

	return
}

// GetFromLrCreateTime 通过lr_create_time获取内容
func (obj *_RSystemLogMgr) GetFromLrCreateTime(lrCreateTime int) (results []*RSystemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSystemLog{}).Where("`lr_create_time` = ?", lrCreateTime).Find(&results).Error

	return
}

// GetBatchFromLrCreateTime 批量查找
func (obj *_RSystemLogMgr) GetBatchFromLrCreateTime(lrCreateTimes []int) (results []*RSystemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSystemLog{}).Where("`lr_create_time` IN (?)", lrCreateTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_RSystemLogMgr) FetchByPrimaryKey(lrID uint32) (result RSystemLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSystemLog{}).Where("`lr_id` = ?", lrID).Find(&result).Error

	return
}
