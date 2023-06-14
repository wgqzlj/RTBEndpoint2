package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _RAdExchangeMgr struct {
	*_BaseMgr
}

// RAdExchangeMgr open func
func RAdExchangeMgr(db *gorm.DB) *_RAdExchangeMgr {
	if db == nil {
		panic(fmt.Errorf("RAdExchangeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_RAdExchangeMgr{_BaseMgr: &_BaseMgr{DB: db.Table("r_ad_exchange"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_RAdExchangeMgr) GetTableName() string {
	return "r_ad_exchange"
}

// Reset 重置gorm会话
func (obj *_RAdExchangeMgr) Reset() *_RAdExchangeMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_RAdExchangeMgr) Get() (result RAdExchange, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdExchange{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_RAdExchangeMgr) Gets() (results []*RAdExchange, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdExchange{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_RAdExchangeMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(RAdExchange{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithAeID ae_id获取
func (obj *_RAdExchangeMgr) WithAeID(aeID uint32) Option {
	return optionFunc(func(o *options) { o.query["ae_id"] = aeID })
}

// WithAeName ae_name获取
func (obj *_RAdExchangeMgr) WithAeName(aeName string) Option {
	return optionFunc(func(o *options) { o.query["ae_name"] = aeName })
}

// WithAeCode ae_code获取
func (obj *_RAdExchangeMgr) WithAeCode(aeCode string) Option {
	return optionFunc(func(o *options) { o.query["ae_code"] = aeCode })
}

// WithAeRemark ae_remark获取
func (obj *_RAdExchangeMgr) WithAeRemark(aeRemark string) Option {
	return optionFunc(func(o *options) { o.query["ae_remark"] = aeRemark })
}

// WithAeStatus ae_status获取
func (obj *_RAdExchangeMgr) WithAeStatus(aeStatus int8) Option {
	return optionFunc(func(o *options) { o.query["ae_status"] = aeStatus })
}

// WithAeSeq ae_seq获取
func (obj *_RAdExchangeMgr) WithAeSeq(aeSeq int) Option {
	return optionFunc(func(o *options) { o.query["ae_seq"] = aeSeq })
}

// GetByOption 功能选项模式获取
func (obj *_RAdExchangeMgr) GetByOption(opts ...Option) (result RAdExchange, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(RAdExchange{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_RAdExchangeMgr) GetByOptions(opts ...Option) (results []*RAdExchange, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(RAdExchange{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromAeID 通过ae_id获取内容
func (obj *_RAdExchangeMgr) GetFromAeID(aeID uint32) (result RAdExchange, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdExchange{}).Where("`ae_id` = ?", aeID).Find(&result).Error

	return
}

// GetBatchFromAeID 批量查找
func (obj *_RAdExchangeMgr) GetBatchFromAeID(aeIDs []uint32) (results []*RAdExchange, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdExchange{}).Where("`ae_id` IN (?)", aeIDs).Find(&results).Error

	return
}

// GetFromAeName 通过ae_name获取内容
func (obj *_RAdExchangeMgr) GetFromAeName(aeName string) (result RAdExchange, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdExchange{}).Where("`ae_name` = ?", aeName).Find(&result).Error

	return
}

// GetBatchFromAeName 批量查找
func (obj *_RAdExchangeMgr) GetBatchFromAeName(aeNames []string) (results []*RAdExchange, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdExchange{}).Where("`ae_name` IN (?)", aeNames).Find(&results).Error

	return
}

// GetFromAeCode 通过ae_code获取内容
func (obj *_RAdExchangeMgr) GetFromAeCode(aeCode string) (result RAdExchange, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdExchange{}).Where("`ae_code` = ?", aeCode).Find(&result).Error

	return
}

// GetBatchFromAeCode 批量查找
func (obj *_RAdExchangeMgr) GetBatchFromAeCode(aeCodes []string) (results []*RAdExchange, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdExchange{}).Where("`ae_code` IN (?)", aeCodes).Find(&results).Error

	return
}

// GetFromAeRemark 通过ae_remark获取内容
func (obj *_RAdExchangeMgr) GetFromAeRemark(aeRemark string) (results []*RAdExchange, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdExchange{}).Where("`ae_remark` = ?", aeRemark).Find(&results).Error

	return
}

// GetBatchFromAeRemark 批量查找
func (obj *_RAdExchangeMgr) GetBatchFromAeRemark(aeRemarks []string) (results []*RAdExchange, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdExchange{}).Where("`ae_remark` IN (?)", aeRemarks).Find(&results).Error

	return
}

// GetFromAeStatus 通过ae_status获取内容
func (obj *_RAdExchangeMgr) GetFromAeStatus(aeStatus int8) (results []*RAdExchange, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdExchange{}).Where("`ae_status` = ?", aeStatus).Find(&results).Error

	return
}

// GetBatchFromAeStatus 批量查找
func (obj *_RAdExchangeMgr) GetBatchFromAeStatus(aeStatuss []int8) (results []*RAdExchange, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdExchange{}).Where("`ae_status` IN (?)", aeStatuss).Find(&results).Error

	return
}

// GetFromAeSeq 通过ae_seq获取内容
func (obj *_RAdExchangeMgr) GetFromAeSeq(aeSeq int) (results []*RAdExchange, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdExchange{}).Where("`ae_seq` = ?", aeSeq).Find(&results).Error

	return
}

// GetBatchFromAeSeq 批量查找
func (obj *_RAdExchangeMgr) GetBatchFromAeSeq(aeSeqs []int) (results []*RAdExchange, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdExchange{}).Where("`ae_seq` IN (?)", aeSeqs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_RAdExchangeMgr) FetchByPrimaryKey(aeID uint32) (result RAdExchange, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdExchange{}).Where("`ae_id` = ?", aeID).Find(&result).Error

	return
}

// FetchUniqueByAName primary or index 获取唯一内容
func (obj *_RAdExchangeMgr) FetchUniqueByAName(aeName string) (result RAdExchange, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdExchange{}).Where("`ae_name` = ?", aeName).Find(&result).Error

	return
}

// FetchUniqueByACode primary or index 获取唯一内容
func (obj *_RAdExchangeMgr) FetchUniqueByACode(aeCode string) (result RAdExchange, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdExchange{}).Where("`ae_code` = ?", aeCode).Find(&result).Error

	return
}
