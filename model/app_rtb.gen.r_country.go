package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _RCountryMgr struct {
	*_BaseMgr
}

// RCountryMgr open func
func RCountryMgr(db *gorm.DB) *_RCountryMgr {
	if db == nil {
		panic(fmt.Errorf("RCountryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_RCountryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("r_country"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_RCountryMgr) GetTableName() string {
	return "r_country"
}

// Reset 重置gorm会话
func (obj *_RCountryMgr) Reset() *_RCountryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_RCountryMgr) Get() (result RCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RCountry{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_RCountryMgr) Gets() (results []*RCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RCountry{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_RCountryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(RCountry{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithCID c_id获取
func (obj *_RCountryMgr) WithCID(cID uint32) Option {
	return optionFunc(func(o *options) { o.query["c_id"] = cID })
}

// WithCName c_name获取
func (obj *_RCountryMgr) WithCName(cName string) Option {
	return optionFunc(func(o *options) { o.query["c_name"] = cName })
}

// WithCCode3 c_code3获取
func (obj *_RCountryMgr) WithCCode3(cCode3 string) Option {
	return optionFunc(func(o *options) { o.query["c_code3"] = cCode3 })
}

// WithCCode2 c_code2获取
func (obj *_RCountryMgr) WithCCode2(cCode2 string) Option {
	return optionFunc(func(o *options) { o.query["c_code2"] = cCode2 })
}

// WithCTimezoneUtc c_timezone_utc获取
func (obj *_RCountryMgr) WithCTimezoneUtc(cTimezoneUtc int) Option {
	return optionFunc(func(o *options) { o.query["c_timezone_utc"] = cTimezoneUtc })
}

// WithCStatus c_status获取
func (obj *_RCountryMgr) WithCStatus(cStatus int8) Option {
	return optionFunc(func(o *options) { o.query["c_status"] = cStatus })
}

// GetByOption 功能选项模式获取
func (obj *_RCountryMgr) GetByOption(opts ...Option) (result RCountry, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(RCountry{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_RCountryMgr) GetByOptions(opts ...Option) (results []*RCountry, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(RCountry{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromCID 通过c_id获取内容
func (obj *_RCountryMgr) GetFromCID(cID uint32) (result RCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RCountry{}).Where("`c_id` = ?", cID).Find(&result).Error

	return
}

// GetBatchFromCID 批量查找
func (obj *_RCountryMgr) GetBatchFromCID(cIDs []uint32) (results []*RCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RCountry{}).Where("`c_id` IN (?)", cIDs).Find(&results).Error

	return
}

// GetFromCName 通过c_name获取内容
func (obj *_RCountryMgr) GetFromCName(cName string) (results []*RCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RCountry{}).Where("`c_name` = ?", cName).Find(&results).Error

	return
}

// GetBatchFromCName 批量查找
func (obj *_RCountryMgr) GetBatchFromCName(cNames []string) (results []*RCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RCountry{}).Where("`c_name` IN (?)", cNames).Find(&results).Error

	return
}

// GetFromCCode3 通过c_code3获取内容
func (obj *_RCountryMgr) GetFromCCode3(cCode3 string) (result RCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RCountry{}).Where("`c_code3` = ?", cCode3).Find(&result).Error

	return
}

// GetBatchFromCCode3 批量查找
func (obj *_RCountryMgr) GetBatchFromCCode3(cCode3s []string) (results []*RCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RCountry{}).Where("`c_code3` IN (?)", cCode3s).Find(&results).Error

	return
}

// GetFromCCode2 通过c_code2获取内容
func (obj *_RCountryMgr) GetFromCCode2(cCode2 string) (result RCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RCountry{}).Where("`c_code2` = ?", cCode2).Find(&result).Error

	return
}

// GetBatchFromCCode2 批量查找
func (obj *_RCountryMgr) GetBatchFromCCode2(cCode2s []string) (results []*RCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RCountry{}).Where("`c_code2` IN (?)", cCode2s).Find(&results).Error

	return
}

// GetFromCTimezoneUtc 通过c_timezone_utc获取内容
func (obj *_RCountryMgr) GetFromCTimezoneUtc(cTimezoneUtc int) (results []*RCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RCountry{}).Where("`c_timezone_utc` = ?", cTimezoneUtc).Find(&results).Error

	return
}

// GetBatchFromCTimezoneUtc 批量查找
func (obj *_RCountryMgr) GetBatchFromCTimezoneUtc(cTimezoneUtcs []int) (results []*RCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RCountry{}).Where("`c_timezone_utc` IN (?)", cTimezoneUtcs).Find(&results).Error

	return
}

// GetFromCStatus 通过c_status获取内容
func (obj *_RCountryMgr) GetFromCStatus(cStatus int8) (results []*RCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RCountry{}).Where("`c_status` = ?", cStatus).Find(&results).Error

	return
}

// GetBatchFromCStatus 批量查找
func (obj *_RCountryMgr) GetBatchFromCStatus(cStatuss []int8) (results []*RCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RCountry{}).Where("`c_status` IN (?)", cStatuss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_RCountryMgr) FetchByPrimaryKey(cID uint32) (result RCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RCountry{}).Where("`c_id` = ?", cID).Find(&result).Error

	return
}

// FetchUniqueByCCode3 primary or index 获取唯一内容
func (obj *_RCountryMgr) FetchUniqueByCCode3(cCode3 string) (result RCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RCountry{}).Where("`c_code3` = ?", cCode3).Find(&result).Error

	return
}

// FetchUniqueByCCode2 primary or index 获取唯一内容
func (obj *_RCountryMgr) FetchUniqueByCCode2(cCode2 string) (result RCountry, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RCountry{}).Where("`c_code2` = ?", cCode2).Find(&result).Error

	return
}
