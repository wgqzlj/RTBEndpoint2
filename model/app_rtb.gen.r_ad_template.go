package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _RAdTemplateMgr struct {
	*_BaseMgr
}

// RAdTemplateMgr open func
func RAdTemplateMgr(db *gorm.DB) *_RAdTemplateMgr {
	if db == nil {
		panic(fmt.Errorf("RAdTemplateMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_RAdTemplateMgr{_BaseMgr: &_BaseMgr{DB: db.Table("r_ad_template"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_RAdTemplateMgr) GetTableName() string {
	return "r_ad_template"
}

// Reset 重置gorm会话
func (obj *_RAdTemplateMgr) Reset() *_RAdTemplateMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_RAdTemplateMgr) Get() (result RAdTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdTemplate{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_RAdTemplateMgr) Gets() (results []*RAdTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdTemplate{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_RAdTemplateMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(RAdTemplate{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithAtID at_id获取
func (obj *_RAdTemplateMgr) WithAtID(atID uint32) Option {
	return optionFunc(func(o *options) { o.query["at_id"] = atID })
}

// WithAtName at_name获取
func (obj *_RAdTemplateMgr) WithAtName(atName string) Option {
	return optionFunc(func(o *options) { o.query["at_name"] = atName })
}

// WithAtFilename at_filename获取
func (obj *_RAdTemplateMgr) WithAtFilename(atFilename string) Option {
	return optionFunc(func(o *options) { o.query["at_filename"] = atFilename })
}

// WithAtCode at_code获取
func (obj *_RAdTemplateMgr) WithAtCode(atCode string) Option {
	return optionFunc(func(o *options) { o.query["at_code"] = atCode })
}

// WithAtStatus at_status获取
func (obj *_RAdTemplateMgr) WithAtStatus(atStatus int8) Option {
	return optionFunc(func(o *options) { o.query["at_status"] = atStatus })
}

// WithAtContent at_content获取
func (obj *_RAdTemplateMgr) WithAtContent(atContent string) Option {
	return optionFunc(func(o *options) { o.query["at_content"] = atContent })
}

// WithAtItID at_it_id获取 1 banner 2 video
func (obj *_RAdTemplateMgr) WithAtItID(atItID int8) Option {
	return optionFunc(func(o *options) { o.query["at_it_id"] = atItID })
}

// WithAtRemark at_remark获取
func (obj *_RAdTemplateMgr) WithAtRemark(atRemark string) Option {
	return optionFunc(func(o *options) { o.query["at_remark"] = atRemark })
}

// WithAtCreateTime at_create_time获取
func (obj *_RAdTemplateMgr) WithAtCreateTime(atCreateTime int) Option {
	return optionFunc(func(o *options) { o.query["at_create_time"] = atCreateTime })
}

// GetByOption 功能选项模式获取
func (obj *_RAdTemplateMgr) GetByOption(opts ...Option) (result RAdTemplate, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(RAdTemplate{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_RAdTemplateMgr) GetByOptions(opts ...Option) (results []*RAdTemplate, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(RAdTemplate{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromAtID 通过at_id获取内容
func (obj *_RAdTemplateMgr) GetFromAtID(atID uint32) (result RAdTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdTemplate{}).Where("`at_id` = ?", atID).Find(&result).Error

	return
}

// GetBatchFromAtID 批量查找
func (obj *_RAdTemplateMgr) GetBatchFromAtID(atIDs []uint32) (results []*RAdTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdTemplate{}).Where("`at_id` IN (?)", atIDs).Find(&results).Error

	return
}

// GetFromAtName 通过at_name获取内容
func (obj *_RAdTemplateMgr) GetFromAtName(atName string) (results []*RAdTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdTemplate{}).Where("`at_name` = ?", atName).Find(&results).Error

	return
}

// GetBatchFromAtName 批量查找
func (obj *_RAdTemplateMgr) GetBatchFromAtName(atNames []string) (results []*RAdTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdTemplate{}).Where("`at_name` IN (?)", atNames).Find(&results).Error

	return
}

// GetFromAtFilename 通过at_filename获取内容
func (obj *_RAdTemplateMgr) GetFromAtFilename(atFilename string) (results []*RAdTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdTemplate{}).Where("`at_filename` = ?", atFilename).Find(&results).Error

	return
}

// GetBatchFromAtFilename 批量查找
func (obj *_RAdTemplateMgr) GetBatchFromAtFilename(atFilenames []string) (results []*RAdTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdTemplate{}).Where("`at_filename` IN (?)", atFilenames).Find(&results).Error

	return
}

// GetFromAtCode 通过at_code获取内容
func (obj *_RAdTemplateMgr) GetFromAtCode(atCode string) (results []*RAdTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdTemplate{}).Where("`at_code` = ?", atCode).Find(&results).Error

	return
}

// GetBatchFromAtCode 批量查找
func (obj *_RAdTemplateMgr) GetBatchFromAtCode(atCodes []string) (results []*RAdTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdTemplate{}).Where("`at_code` IN (?)", atCodes).Find(&results).Error

	return
}

// GetFromAtStatus 通过at_status获取内容
func (obj *_RAdTemplateMgr) GetFromAtStatus(atStatus int8) (results []*RAdTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdTemplate{}).Where("`at_status` = ?", atStatus).Find(&results).Error

	return
}

// GetBatchFromAtStatus 批量查找
func (obj *_RAdTemplateMgr) GetBatchFromAtStatus(atStatuss []int8) (results []*RAdTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdTemplate{}).Where("`at_status` IN (?)", atStatuss).Find(&results).Error

	return
}

// GetFromAtContent 通过at_content获取内容
func (obj *_RAdTemplateMgr) GetFromAtContent(atContent string) (results []*RAdTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdTemplate{}).Where("`at_content` = ?", atContent).Find(&results).Error

	return
}

// GetBatchFromAtContent 批量查找
func (obj *_RAdTemplateMgr) GetBatchFromAtContent(atContents []string) (results []*RAdTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdTemplate{}).Where("`at_content` IN (?)", atContents).Find(&results).Error

	return
}

// GetFromAtItID 通过at_it_id获取内容 1 banner 2 video
func (obj *_RAdTemplateMgr) GetFromAtItID(atItID int8) (results []*RAdTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdTemplate{}).Where("`at_it_id` = ?", atItID).Find(&results).Error

	return
}

// GetBatchFromAtItID 批量查找 1 banner 2 video
func (obj *_RAdTemplateMgr) GetBatchFromAtItID(atItIDs []int8) (results []*RAdTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdTemplate{}).Where("`at_it_id` IN (?)", atItIDs).Find(&results).Error

	return
}

// GetFromAtRemark 通过at_remark获取内容
func (obj *_RAdTemplateMgr) GetFromAtRemark(atRemark string) (results []*RAdTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdTemplate{}).Where("`at_remark` = ?", atRemark).Find(&results).Error

	return
}

// GetBatchFromAtRemark 批量查找
func (obj *_RAdTemplateMgr) GetBatchFromAtRemark(atRemarks []string) (results []*RAdTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdTemplate{}).Where("`at_remark` IN (?)", atRemarks).Find(&results).Error

	return
}

// GetFromAtCreateTime 通过at_create_time获取内容
func (obj *_RAdTemplateMgr) GetFromAtCreateTime(atCreateTime int) (results []*RAdTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdTemplate{}).Where("`at_create_time` = ?", atCreateTime).Find(&results).Error

	return
}

// GetBatchFromAtCreateTime 批量查找
func (obj *_RAdTemplateMgr) GetBatchFromAtCreateTime(atCreateTimes []int) (results []*RAdTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdTemplate{}).Where("`at_create_time` IN (?)", atCreateTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_RAdTemplateMgr) FetchByPrimaryKey(atID uint32) (result RAdTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RAdTemplate{}).Where("`at_id` = ?", atID).Find(&result).Error

	return
}
