package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _RSystemAdminMgr struct {
	*_BaseMgr
}

// RSystemAdminMgr open func
func RSystemAdminMgr(db *gorm.DB) *_RSystemAdminMgr {
	if db == nil {
		panic(fmt.Errorf("RSystemAdminMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_RSystemAdminMgr{_BaseMgr: &_BaseMgr{DB: db.Table("r_system_admin"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_RSystemAdminMgr) GetTableName() string {
	return "r_system_admin"
}

// Reset 重置gorm会话
func (obj *_RSystemAdminMgr) Reset() *_RSystemAdminMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_RSystemAdminMgr) Get() (result RSystemAdmin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSystemAdmin{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_RSystemAdminMgr) Gets() (results []*RSystemAdmin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSystemAdmin{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_RSystemAdminMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(RSystemAdmin{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSaID sa_id获取 系统管理员ID
func (obj *_RSystemAdminMgr) WithSaID(saID uint32) Option {
	return optionFunc(func(o *options) { o.query["sa_id"] = saID })
}

// WithSaUsername sa_username获取 系统管理员用户名
func (obj *_RSystemAdminMgr) WithSaUsername(saUsername string) Option {
	return optionFunc(func(o *options) { o.query["sa_username"] = saUsername })
}

// WithSaPassword sa_password获取 系统管理员密码
func (obj *_RSystemAdminMgr) WithSaPassword(saPassword string) Option {
	return optionFunc(func(o *options) { o.query["sa_password"] = saPassword })
}

// WithSaPasswordSalt sa_password_salt获取 密码盐
func (obj *_RSystemAdminMgr) WithSaPasswordSalt(saPasswordSalt string) Option {
	return optionFunc(func(o *options) { o.query["sa_password_salt"] = saPasswordSalt })
}

// WithSaLoginIP sa_login_ip获取 上一次登陆IP
func (obj *_RSystemAdminMgr) WithSaLoginIP(saLoginIP uint32) Option {
	return optionFunc(func(o *options) { o.query["sa_login_ip"] = saLoginIP })
}

// WithSaLoginTime sa_login_time获取 上一次登陆时间
func (obj *_RSystemAdminMgr) WithSaLoginTime(saLoginTime int) Option {
	return optionFunc(func(o *options) { o.query["sa_login_time"] = saLoginTime })
}

// WithSaStatus sa_status获取 系统管理员状态，1为正常，0为禁用
func (obj *_RSystemAdminMgr) WithSaStatus(saStatus int) Option {
	return optionFunc(func(o *options) { o.query["sa_status"] = saStatus })
}

// GetByOption 功能选项模式获取
func (obj *_RSystemAdminMgr) GetByOption(opts ...Option) (result RSystemAdmin, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(RSystemAdmin{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_RSystemAdminMgr) GetByOptions(opts ...Option) (results []*RSystemAdmin, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(RSystemAdmin{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromSaID 通过sa_id获取内容 系统管理员ID
func (obj *_RSystemAdminMgr) GetFromSaID(saID uint32) (result RSystemAdmin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSystemAdmin{}).Where("`sa_id` = ?", saID).Find(&result).Error

	return
}

// GetBatchFromSaID 批量查找 系统管理员ID
func (obj *_RSystemAdminMgr) GetBatchFromSaID(saIDs []uint32) (results []*RSystemAdmin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSystemAdmin{}).Where("`sa_id` IN (?)", saIDs).Find(&results).Error

	return
}

// GetFromSaUsername 通过sa_username获取内容 系统管理员用户名
func (obj *_RSystemAdminMgr) GetFromSaUsername(saUsername string) (result RSystemAdmin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSystemAdmin{}).Where("`sa_username` = ?", saUsername).Find(&result).Error

	return
}

// GetBatchFromSaUsername 批量查找 系统管理员用户名
func (obj *_RSystemAdminMgr) GetBatchFromSaUsername(saUsernames []string) (results []*RSystemAdmin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSystemAdmin{}).Where("`sa_username` IN (?)", saUsernames).Find(&results).Error

	return
}

// GetFromSaPassword 通过sa_password获取内容 系统管理员密码
func (obj *_RSystemAdminMgr) GetFromSaPassword(saPassword string) (results []*RSystemAdmin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSystemAdmin{}).Where("`sa_password` = ?", saPassword).Find(&results).Error

	return
}

// GetBatchFromSaPassword 批量查找 系统管理员密码
func (obj *_RSystemAdminMgr) GetBatchFromSaPassword(saPasswords []string) (results []*RSystemAdmin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSystemAdmin{}).Where("`sa_password` IN (?)", saPasswords).Find(&results).Error

	return
}

// GetFromSaPasswordSalt 通过sa_password_salt获取内容 密码盐
func (obj *_RSystemAdminMgr) GetFromSaPasswordSalt(saPasswordSalt string) (results []*RSystemAdmin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSystemAdmin{}).Where("`sa_password_salt` = ?", saPasswordSalt).Find(&results).Error

	return
}

// GetBatchFromSaPasswordSalt 批量查找 密码盐
func (obj *_RSystemAdminMgr) GetBatchFromSaPasswordSalt(saPasswordSalts []string) (results []*RSystemAdmin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSystemAdmin{}).Where("`sa_password_salt` IN (?)", saPasswordSalts).Find(&results).Error

	return
}

// GetFromSaLoginIP 通过sa_login_ip获取内容 上一次登陆IP
func (obj *_RSystemAdminMgr) GetFromSaLoginIP(saLoginIP uint32) (results []*RSystemAdmin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSystemAdmin{}).Where("`sa_login_ip` = ?", saLoginIP).Find(&results).Error

	return
}

// GetBatchFromSaLoginIP 批量查找 上一次登陆IP
func (obj *_RSystemAdminMgr) GetBatchFromSaLoginIP(saLoginIPs []uint32) (results []*RSystemAdmin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSystemAdmin{}).Where("`sa_login_ip` IN (?)", saLoginIPs).Find(&results).Error

	return
}

// GetFromSaLoginTime 通过sa_login_time获取内容 上一次登陆时间
func (obj *_RSystemAdminMgr) GetFromSaLoginTime(saLoginTime int) (results []*RSystemAdmin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSystemAdmin{}).Where("`sa_login_time` = ?", saLoginTime).Find(&results).Error

	return
}

// GetBatchFromSaLoginTime 批量查找 上一次登陆时间
func (obj *_RSystemAdminMgr) GetBatchFromSaLoginTime(saLoginTimes []int) (results []*RSystemAdmin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSystemAdmin{}).Where("`sa_login_time` IN (?)", saLoginTimes).Find(&results).Error

	return
}

// GetFromSaStatus 通过sa_status获取内容 系统管理员状态，1为正常，0为禁用
func (obj *_RSystemAdminMgr) GetFromSaStatus(saStatus int) (results []*RSystemAdmin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSystemAdmin{}).Where("`sa_status` = ?", saStatus).Find(&results).Error

	return
}

// GetBatchFromSaStatus 批量查找 系统管理员状态，1为正常，0为禁用
func (obj *_RSystemAdminMgr) GetBatchFromSaStatus(saStatuss []int) (results []*RSystemAdmin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSystemAdmin{}).Where("`sa_status` IN (?)", saStatuss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_RSystemAdminMgr) FetchByPrimaryKey(saID uint32) (result RSystemAdmin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSystemAdmin{}).Where("`sa_id` = ?", saID).Find(&result).Error

	return
}

// FetchUniqueByUniqueName primary or index 获取唯一内容
func (obj *_RSystemAdminMgr) FetchUniqueByUniqueName(saUsername string) (result RSystemAdmin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSystemAdmin{}).Where("`sa_username` = ?", saUsername).Find(&result).Error

	return
}

// FetchIndexBySaStatus  获取多个内容
func (obj *_RSystemAdminMgr) FetchIndexBySaStatus(saStatus int) (results []*RSystemAdmin, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RSystemAdmin{}).Where("`sa_status` = ?", saStatus).Find(&results).Error

	return
}
