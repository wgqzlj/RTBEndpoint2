package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _RDpaFeedItemMgr struct {
	*_BaseMgr
}

// RDpaFeedItemMgr open func
func RDpaFeedItemMgr(db *gorm.DB) *_RDpaFeedItemMgr {
	if db == nil {
		panic(fmt.Errorf("RDpaFeedItemMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_RDpaFeedItemMgr{_BaseMgr: &_BaseMgr{DB: db.Table("r_dpa_feed_item"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_RDpaFeedItemMgr) GetTableName() string {
	return "r_dpa_feed_item"
}

// Reset 重置gorm会话
func (obj *_RDpaFeedItemMgr) Reset() *_RDpaFeedItemMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_RDpaFeedItemMgr) Get() (result RDpaFeedItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RDpaFeedItem{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_RDpaFeedItemMgr) Gets() (results []*RDpaFeedItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RDpaFeedItem{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_RDpaFeedItemMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(RDpaFeedItem{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithDfiID dfi_id获取
func (obj *_RDpaFeedItemMgr) WithDfiID(dfiID uint32) Option {
	return optionFunc(func(o *options) { o.query["dfi_id"] = dfiID })
}

// WithDfiSkuID dfi_sku_id获取
func (obj *_RDpaFeedItemMgr) WithDfiSkuID(dfiSkuID int64) Option {
	return optionFunc(func(o *options) { o.query["dfi_sku_id"] = dfiSkuID })
}

// WithDfiProductID dfi_product_id获取
func (obj *_RDpaFeedItemMgr) WithDfiProductID(dfiProductID int64) Option {
	return optionFunc(func(o *options) { o.query["dfi_product_id"] = dfiProductID })
}

// WithDfiProductNameCsvEsc dfi_product_name_csv_esc获取
func (obj *_RDpaFeedItemMgr) WithDfiProductNameCsvEsc(dfiProductNameCsvEsc string) Option {
	return optionFunc(func(o *options) { o.query["dfi_product_name_csv_esc"] = dfiProductNameCsvEsc })
}

// WithDfiBrandName dfi_brand_name获取
func (obj *_RDpaFeedItemMgr) WithDfiBrandName(dfiBrandName string) Option {
	return optionFunc(func(o *options) { o.query["dfi_brand_name"] = dfiBrandName })
}

// WithDfiRegionalCategory1Name dfi_regional_category1_name获取
func (obj *_RDpaFeedItemMgr) WithDfiRegionalCategory1Name(dfiRegionalCategory1Name string) Option {
	return optionFunc(func(o *options) { o.query["dfi_regional_category1_name"] = dfiRegionalCategory1Name })
}

// WithDfiRegionalCategory2Name dfi_regional_category2_name获取
func (obj *_RDpaFeedItemMgr) WithDfiRegionalCategory2Name(dfiRegionalCategory2Name string) Option {
	return optionFunc(func(o *options) { o.query["dfi_regional_category2_name"] = dfiRegionalCategory2Name })
}

// WithDfiRegionalCategory3Name dfi_regional_category3_name获取
func (obj *_RDpaFeedItemMgr) WithDfiRegionalCategory3Name(dfiRegionalCategory3Name string) Option {
	return optionFunc(func(o *options) { o.query["dfi_regional_category3_name"] = dfiRegionalCategory3Name })
}

// WithDfiProductURL dfi_product_url获取
func (obj *_RDpaFeedItemMgr) WithDfiProductURL(dfiProductURL string) Option {
	return optionFunc(func(o *options) { o.query["dfi_product_url"] = dfiProductURL })
}

// WithDfiURLMainImage dfi_url_main_image获取
func (obj *_RDpaFeedItemMgr) WithDfiURLMainImage(dfiURLMainImage string) Option {
	return optionFunc(func(o *options) { o.query["dfi_url_main_image"] = dfiURLMainImage })
}

// WithDfiProductMediumImg dfi_product_medium_img获取
func (obj *_RDpaFeedItemMgr) WithDfiProductMediumImg(dfiProductMediumImg string) Option {
	return optionFunc(func(o *options) { o.query["dfi_product_medium_img"] = dfiProductMediumImg })
}

// WithDfiProductBigImg dfi_product_big_img获取
func (obj *_RDpaFeedItemMgr) WithDfiProductBigImg(dfiProductBigImg string) Option {
	return optionFunc(func(o *options) { o.query["dfi_product_big_img"] = dfiProductBigImg })
}

// WithDfiImageURL2 dfi_image_url_2获取
func (obj *_RDpaFeedItemMgr) WithDfiImageURL2(dfiImageURL2 string) Option {
	return optionFunc(func(o *options) { o.query["dfi_image_url_2"] = dfiImageURL2 })
}

// WithDfiImageURL3 dfi_image_url_3获取
func (obj *_RDpaFeedItemMgr) WithDfiImageURL3(dfiImageURL3 string) Option {
	return optionFunc(func(o *options) { o.query["dfi_image_url_3"] = dfiImageURL3 })
}

// WithDfiImageURL4 dfi_image_url_4获取
func (obj *_RDpaFeedItemMgr) WithDfiImageURL4(dfiImageURL4 string) Option {
	return optionFunc(func(o *options) { o.query["dfi_image_url_4"] = dfiImageURL4 })
}

// WithDfiImageURL5 dfi_image_url_5获取
func (obj *_RDpaFeedItemMgr) WithDfiImageURL5(dfiImageURL5 string) Option {
	return optionFunc(func(o *options) { o.query["dfi_image_url_5"] = dfiImageURL5 })
}

// WithDfiDescription dfi_description获取
func (obj *_RDpaFeedItemMgr) WithDfiDescription(dfiDescription string) Option {
	return optionFunc(func(o *options) { o.query["dfi_description"] = dfiDescription })
}

// WithDfiVentureCategory1NameEn dfi_venture_category1_name_en获取
func (obj *_RDpaFeedItemMgr) WithDfiVentureCategory1NameEn(dfiVentureCategory1NameEn string) Option {
	return optionFunc(func(o *options) { o.query["dfi_venture_category1_name_en"] = dfiVentureCategory1NameEn })
}

// WithDfiVentureCategory2NameEn dfi_venture_category2_name_en获取
func (obj *_RDpaFeedItemMgr) WithDfiVentureCategory2NameEn(dfiVentureCategory2NameEn string) Option {
	return optionFunc(func(o *options) { o.query["dfi_venture_category2_name_en"] = dfiVentureCategory2NameEn })
}

// WithDfiVentureCategory3NameEn dfi_venture_category3_name_en获取
func (obj *_RDpaFeedItemMgr) WithDfiVentureCategory3NameEn(dfiVentureCategory3NameEn string) Option {
	return optionFunc(func(o *options) { o.query["dfi_venture_category3_name_en"] = dfiVentureCategory3NameEn })
}

// WithDfiPromotionTag dfi_promotion_tag获取
func (obj *_RDpaFeedItemMgr) WithDfiPromotionTag(dfiPromotionTag string) Option {
	return optionFunc(func(o *options) { o.query["dfi_promotion_tag"] = dfiPromotionTag })
}

// WithDfiPromotionPrice dfi_promotion_price获取
func (obj *_RDpaFeedItemMgr) WithDfiPromotionPrice(dfiPromotionPrice float32) Option {
	return optionFunc(func(o *options) { o.query["dfi_promotion_price"] = dfiPromotionPrice })
}

// WithDfiBrandLogo dfi_brand_logo获取
func (obj *_RDpaFeedItemMgr) WithDfiBrandLogo(dfiBrandLogo string) Option {
	return optionFunc(func(o *options) { o.query["dfi_brand_logo"] = dfiBrandLogo })
}

// WithDfiCurrentPrice dfi_current_price获取
func (obj *_RDpaFeedItemMgr) WithDfiCurrentPrice(dfiCurrentPrice float32) Option {
	return optionFunc(func(o *options) { o.query["dfi_current_price"] = dfiCurrentPrice })
}

// WithDfiPrice dfi_price获取
func (obj *_RDpaFeedItemMgr) WithDfiPrice(dfiPrice float32) Option {
	return optionFunc(func(o *options) { o.query["dfi_price"] = dfiPrice })
}

// WithDfiVenture dfi_venture获取
func (obj *_RDpaFeedItemMgr) WithDfiVenture(dfiVenture string) Option {
	return optionFunc(func(o *options) { o.query["dfi_venture"] = dfiVenture })
}

// WithDfiUpdateTime dfi_update_time获取
func (obj *_RDpaFeedItemMgr) WithDfiUpdateTime(dfiUpdateTime int) Option {
	return optionFunc(func(o *options) { o.query["dfi_update_time"] = dfiUpdateTime })
}

// WithDfiDiscountPercentage dfi_discount_percentage获取
func (obj *_RDpaFeedItemMgr) WithDfiDiscountPercentage(dfiDiscountPercentage string) Option {
	return optionFunc(func(o *options) { o.query["dfi_discount_percentage"] = dfiDiscountPercentage })
}

// WithDfiBenefitUsp dfi_benefit_usp获取
func (obj *_RDpaFeedItemMgr) WithDfiBenefitUsp(dfiBenefitUsp string) Option {
	return optionFunc(func(o *options) { o.query["dfi_benefit_usp"] = dfiBenefitUsp })
}

// GetByOption 功能选项模式获取
func (obj *_RDpaFeedItemMgr) GetByOption(opts ...Option) (result RDpaFeedItem, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(RDpaFeedItem{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_RDpaFeedItemMgr) GetByOptions(opts ...Option) (results []*RDpaFeedItem, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(RDpaFeedItem{}).Where(options.query).Find(&results).Error

	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromDfiID 通过dfi_id获取内容
func (obj *_RDpaFeedItemMgr) GetFromDfiID(dfiID uint32) (result RDpaFeedItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RDpaFeedItem{}).Where("`dfi_id` = ?", dfiID).Find(&result).Error

	return
}

// GetBatchFromDfiID 批量查找
func (obj *_RDpaFeedItemMgr) GetBatchFromDfiID(dfiIDs []uint32) (results []*RDpaFeedItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RDpaFeedItem{}).Where("`dfi_id` IN (?)", dfiIDs).Find(&results).Error

	return
}

// GetFromDfiSkuID 通过dfi_sku_id获取内容
func (obj *_RDpaFeedItemMgr) GetFromDfiSkuID(dfiSkuID int64) (results []*RDpaFeedItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RDpaFeedItem{}).Where("`dfi_sku_id` = ?", dfiSkuID).Find(&results).Error

	return
}

// GetBatchFromDfiSkuID 批量查找
func (obj *_RDpaFeedItemMgr) GetBatchFromDfiSkuID(dfiSkuIDs []int64) (results []*RDpaFeedItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RDpaFeedItem{}).Where("`dfi_sku_id` IN (?)", dfiSkuIDs).Find(&results).Error

	return
}

// GetFromDfiProductID 通过dfi_product_id获取内容
func (obj *_RDpaFeedItemMgr) GetFromDfiProductID(dfiProductID int64) (results []*RDpaFeedItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RDpaFeedItem{}).Where("`dfi_product_id` = ?", dfiProductID).Find(&results).Error

	return
}

// GetBatchFromDfiProductID 批量查找
func (obj *_RDpaFeedItemMgr) GetBatchFromDfiProductID(dfiProductIDs []int64) (results []*RDpaFeedItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RDpaFeedItem{}).Where("`dfi_product_id` IN (?)", dfiProductIDs).Find(&results).Error

	return
}

// GetFromDfiProductNameCsvEsc 通过dfi_product_name_csv_esc获取内容
func (obj *_RDpaFeedItemMgr) GetFromDfiProductNameCsvEsc(dfiProductNameCsvEsc string) (results []*RDpaFeedItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RDpaFeedItem{}).Where("`dfi_product_name_csv_esc` = ?", dfiProductNameCsvEsc).Find(&results).Error

	return
}

// GetBatchFromDfiProductNameCsvEsc 批量查找
func (obj *_RDpaFeedItemMgr) GetBatchFromDfiProductNameCsvEsc(dfiProductNameCsvEscs []string) (results []*RDpaFeedItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RDpaFeedItem{}).Where("`dfi_product_name_csv_esc` IN (?)", dfiProductNameCsvEscs).Find(&results).Error

	return
}

// GetFromDfiBrandName 通过dfi_brand_name获取内容
func (obj *_RDpaFeedItemMgr) GetFromDfiBrandName(dfiBrandName string) (results []*RDpaFeedItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RDpaFeedItem{}).Where("`dfi_brand_name` = ?", dfiBrandName).Find(&results).Error

	return
}

// GetBatchFromDfiBrandName 批量查找
func (obj *_RDpaFeedItemMgr) GetBatchFromDfiBrandName(dfiBrandNames []string) (results []*RDpaFeedItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RDpaFeedItem{}).Where("`dfi_brand_name` IN (?)", dfiBrandNames).Find(&results).Error

	return
}

// GetFromDfiRegionalCategory1Name 通过dfi_regional_category1_name获取内容
func (obj *_RDpaFeedItemMgr) GetFromDfiRegionalCategory1Name(dfiRegionalCategory1Name string) (results []*RDpaFeedItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RDpaFeedItem{}).Where("`dfi_regional_category1_name` = ?", dfiRegionalCategory1Name).Find(&results).Error

	return
}

// GetBatchFromDfiRegionalCategory1Name 批量查找
func (obj *_RDpaFeedItemMgr) GetBatchFromDfiRegionalCategory1Name(dfiRegionalCategory1Names []string) (results []*RDpaFeedItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RDpaFeedItem{}).Where("`dfi_regional_category1_name` IN (?)", dfiRegionalCategory1Names).Find(&results).Error

	return
}

// GetFromDfiRegionalCategory2Name 通过dfi_regional_category2_name获取内容
func (obj *_RDpaFeedItemMgr) GetFromDfiRegionalCategory2Name(dfiRegionalCategory2Name string) (results []*RDpaFeedItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RDpaFeedItem{}).Where("`dfi_regional_category2_name` = ?", dfiRegionalCategory2Name).Find(&results).Error

	return
}

// GetBatchFromDfiRegionalCategory2Name 批量查找
func (obj *_RDpaFeedItemMgr) GetBatchFromDfiRegionalCategory2Name(dfiRegionalCategory2Names []string) (results []*RDpaFeedItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RDpaFeedItem{}).Where("`dfi_regional_category2_name` IN (?)", dfiRegionalCategory2Names).Find(&results).Error

	return
}

// GetFromDfiRegionalCategory3Name 通过dfi_regional_category3_name获取内容
func (obj *_RDpaFeedItemMgr) GetFromDfiRegionalCategory3Name(dfiRegionalCategory3Name string) (results []*RDpaFeedItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RDpaFeedItem{}).Where("`dfi_regional_category3_name` = ?", dfiRegionalCategory3Name).Find(&results).Error

	return
}

// GetBatchFromDfiRegionalCategory3Name 批量查找
func (obj *_RDpaFeedItemMgr) GetBatchFromDfiRegionalCategory3Name(dfiRegionalCategory3Names []string) (results []*RDpaFeedItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RDpaFeedItem{}).Where("`dfi_regional_category3_name` IN (?)", dfiRegionalCategory3Names).Find(&results).Error

	return
}

// GetFromDfiProductURL 通过dfi_product_url获取内容
func (obj *_RDpaFeedItemMgr) GetFromDfiProductURL(dfiProductURL string) (results []*RDpaFeedItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RDpaFeedItem{}).Where("`dfi_product_url` = ?", dfiProductURL).Find(&results).Error

	return
}

// GetBatchFromDfiProductURL 批量查找
func (obj *_RDpaFeedItemMgr) GetBatchFromDfiProductURL(dfiProductURLs []string) (results []*RDpaFeedItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RDpaFeedItem{}).Where("`dfi_product_url` IN (?)", dfiProductURLs).Find(&results).Error

	return
}

// GetFromDfiURLMainImage 通过dfi_url_main_image获取内容
func (obj *_RDpaFeedItemMgr) GetFromDfiURLMainImage(dfiURLMainImage string) (results []*RDpaFeedItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RDpaFeedItem{}).Where("`dfi_url_main_image` = ?", dfiURLMainImage).Find(&results).Error

	return
}

// GetBatchFromDfiURLMainImage 批量查找
func (obj *_RDpaFeedItemMgr) GetBatchFromDfiURLMainImage(dfiURLMainImages []string) (results []*RDpaFeedItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RDpaFeedItem{}).Where("`dfi_url_main_image` IN (?)", dfiURLMainImages).Find(&results).Error

	return
}

// GetFromDfiProductMediumImg 通过dfi_product_medium_img获取内容
func (obj *_RDpaFeedItemMgr) GetFromDfiProductMediumImg(dfiProductMediumImg string) (results []*RDpaFeedItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RDpaFeedItem{}).Where("`dfi_product_medium_img` = ?", dfiProductMediumImg).Find(&results).Error

	return
}

// GetBatchFromDfiProductMediumImg 批量查找
func (obj *_RDpaFeedItemMgr) GetBatchFromDfiProductMediumImg(dfiProductMediumImgs []string) (results []*RDpaFeedItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RDpaFeedItem{}).Where("`dfi_product_medium_img` IN (?)", dfiProductMediumImgs).Find(&results).Error

	return
}

// GetFromDfiProductBigImg 通过dfi_product_big_img获取内容
func (obj *_RDpaFeedItemMgr) GetFromDfiProductBigImg(dfiProductBigImg string) (results []*RDpaFeedItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RDpaFeedItem{}).Where("`dfi_product_big_img` = ?", dfiProductBigImg).Find(&results).Error

	return
}

// GetBatchFromDfiProductBigImg 批量查找
func (obj *_RDpaFeedItemMgr) GetBatchFromDfiProductBigImg(dfiProductBigImgs []string) (results []*RDpaFeedItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RDpaFeedItem{}).Where("`dfi_product_big_img` IN (?)", dfiProductBigImgs).Find(&results).Error

	return
}

// GetFromDfiImageURL2 通过dfi_image_url_2获取内容
func (obj *_RDpaFeedItemMgr) GetFromDfiImageURL2(dfiImageURL2 string) (results []*RDpaFeedItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RDpaFeedItem{}).Where("`dfi_image_url_2` = ?", dfiImageURL2).Find(&results).Error

	return
}

// GetBatchFromDfiImageURL2 批量查找
func (obj *_RDpaFeedItemMgr) GetBatchFromDfiImageURL2(dfiImageURL2s []string) (results []*RDpaFeedItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RDpaFeedItem{}).Where("`dfi_image_url_2` IN (?)", dfiImageURL2s).Find(&results).Error

	return
}

// GetFromDfiImageURL3 通过dfi_image_url_3获取内容
func (obj *_RDpaFeedItemMgr) GetFromDfiImageURL3(dfiImageURL3 string) (results []*RDpaFeedItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RDpaFeedItem{}).Where("`dfi_image_url_3` = ?", dfiImageURL3).Find(&results).Error

	return
}

// GetBatchFromDfiImageURL3 批量查找
func (obj *_RDpaFeedItemMgr) GetBatchFromDfiImageURL3(dfiImageURL3s []string) (results []*RDpaFeedItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RDpaFeedItem{}).Where("`dfi_image_url_3` IN (?)", dfiImageURL3s).Find(&results).Error

	return
}

// GetFromDfiImageURL4 通过dfi_image_url_4获取内容
func (obj *_RDpaFeedItemMgr) GetFromDfiImageURL4(dfiImageURL4 string) (results []*RDpaFeedItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RDpaFeedItem{}).Where("`dfi_image_url_4` = ?", dfiImageURL4).Find(&results).Error

	return
}

// GetBatchFromDfiImageURL4 批量查找
func (obj *_RDpaFeedItemMgr) GetBatchFromDfiImageURL4(dfiImageURL4s []string) (results []*RDpaFeedItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RDpaFeedItem{}).Where("`dfi_image_url_4` IN (?)", dfiImageURL4s).Find(&results).Error

	return
}

// GetFromDfiImageURL5 通过dfi_image_url_5获取内容
func (obj *_RDpaFeedItemMgr) GetFromDfiImageURL5(dfiImageURL5 string) (results []*RDpaFeedItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RDpaFeedItem{}).Where("`dfi_image_url_5` = ?", dfiImageURL5).Find(&results).Error

	return
}

// GetBatchFromDfiImageURL5 批量查找
func (obj *_RDpaFeedItemMgr) GetBatchFromDfiImageURL5(dfiImageURL5s []string) (results []*RDpaFeedItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RDpaFeedItem{}).Where("`dfi_image_url_5` IN (?)", dfiImageURL5s).Find(&results).Error

	return
}

// GetFromDfiDescription 通过dfi_description获取内容
func (obj *_RDpaFeedItemMgr) GetFromDfiDescription(dfiDescription string) (results []*RDpaFeedItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RDpaFeedItem{}).Where("`dfi_description` = ?", dfiDescription).Find(&results).Error

	return
}

// GetBatchFromDfiDescription 批量查找
func (obj *_RDpaFeedItemMgr) GetBatchFromDfiDescription(dfiDescriptions []string) (results []*RDpaFeedItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RDpaFeedItem{}).Where("`dfi_description` IN (?)", dfiDescriptions).Find(&results).Error

	return
}

// GetFromDfiVentureCategory1NameEn 通过dfi_venture_category1_name_en获取内容
func (obj *_RDpaFeedItemMgr) GetFromDfiVentureCategory1NameEn(dfiVentureCategory1NameEn string) (results []*RDpaFeedItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RDpaFeedItem{}).Where("`dfi_venture_category1_name_en` = ?", dfiVentureCategory1NameEn).Find(&results).Error

	return
}

// GetBatchFromDfiVentureCategory1NameEn 批量查找
func (obj *_RDpaFeedItemMgr) GetBatchFromDfiVentureCategory1NameEn(dfiVentureCategory1NameEns []string) (results []*RDpaFeedItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RDpaFeedItem{}).Where("`dfi_venture_category1_name_en` IN (?)", dfiVentureCategory1NameEns).Find(&results).Error

	return
}

// GetFromDfiVentureCategory2NameEn 通过dfi_venture_category2_name_en获取内容
func (obj *_RDpaFeedItemMgr) GetFromDfiVentureCategory2NameEn(dfiVentureCategory2NameEn string) (results []*RDpaFeedItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RDpaFeedItem{}).Where("`dfi_venture_category2_name_en` = ?", dfiVentureCategory2NameEn).Find(&results).Error

	return
}

// GetBatchFromDfiVentureCategory2NameEn 批量查找
func (obj *_RDpaFeedItemMgr) GetBatchFromDfiVentureCategory2NameEn(dfiVentureCategory2NameEns []string) (results []*RDpaFeedItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RDpaFeedItem{}).Where("`dfi_venture_category2_name_en` IN (?)", dfiVentureCategory2NameEns).Find(&results).Error

	return
}

// GetFromDfiVentureCategory3NameEn 通过dfi_venture_category3_name_en获取内容
func (obj *_RDpaFeedItemMgr) GetFromDfiVentureCategory3NameEn(dfiVentureCategory3NameEn string) (results []*RDpaFeedItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RDpaFeedItem{}).Where("`dfi_venture_category3_name_en` = ?", dfiVentureCategory3NameEn).Find(&results).Error

	return
}

// GetBatchFromDfiVentureCategory3NameEn 批量查找
func (obj *_RDpaFeedItemMgr) GetBatchFromDfiVentureCategory3NameEn(dfiVentureCategory3NameEns []string) (results []*RDpaFeedItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RDpaFeedItem{}).Where("`dfi_venture_category3_name_en` IN (?)", dfiVentureCategory3NameEns).Find(&results).Error

	return
}

// GetFromDfiPromotionTag 通过dfi_promotion_tag获取内容
func (obj *_RDpaFeedItemMgr) GetFromDfiPromotionTag(dfiPromotionTag string) (results []*RDpaFeedItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RDpaFeedItem{}).Where("`dfi_promotion_tag` = ?", dfiPromotionTag).Find(&results).Error

	return
}

// GetBatchFromDfiPromotionTag 批量查找
func (obj *_RDpaFeedItemMgr) GetBatchFromDfiPromotionTag(dfiPromotionTags []string) (results []*RDpaFeedItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RDpaFeedItem{}).Where("`dfi_promotion_tag` IN (?)", dfiPromotionTags).Find(&results).Error

	return
}

// GetFromDfiPromotionPrice 通过dfi_promotion_price获取内容
func (obj *_RDpaFeedItemMgr) GetFromDfiPromotionPrice(dfiPromotionPrice float32) (results []*RDpaFeedItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RDpaFeedItem{}).Where("`dfi_promotion_price` = ?", dfiPromotionPrice).Find(&results).Error

	return
}

// GetBatchFromDfiPromotionPrice 批量查找
func (obj *_RDpaFeedItemMgr) GetBatchFromDfiPromotionPrice(dfiPromotionPrices []float32) (results []*RDpaFeedItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RDpaFeedItem{}).Where("`dfi_promotion_price` IN (?)", dfiPromotionPrices).Find(&results).Error

	return
}

// GetFromDfiBrandLogo 通过dfi_brand_logo获取内容
func (obj *_RDpaFeedItemMgr) GetFromDfiBrandLogo(dfiBrandLogo string) (results []*RDpaFeedItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RDpaFeedItem{}).Where("`dfi_brand_logo` = ?", dfiBrandLogo).Find(&results).Error

	return
}

// GetBatchFromDfiBrandLogo 批量查找
func (obj *_RDpaFeedItemMgr) GetBatchFromDfiBrandLogo(dfiBrandLogos []string) (results []*RDpaFeedItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RDpaFeedItem{}).Where("`dfi_brand_logo` IN (?)", dfiBrandLogos).Find(&results).Error

	return
}

// GetFromDfiCurrentPrice 通过dfi_current_price获取内容
func (obj *_RDpaFeedItemMgr) GetFromDfiCurrentPrice(dfiCurrentPrice float32) (results []*RDpaFeedItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RDpaFeedItem{}).Where("`dfi_current_price` = ?", dfiCurrentPrice).Find(&results).Error

	return
}

// GetBatchFromDfiCurrentPrice 批量查找
func (obj *_RDpaFeedItemMgr) GetBatchFromDfiCurrentPrice(dfiCurrentPrices []float32) (results []*RDpaFeedItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RDpaFeedItem{}).Where("`dfi_current_price` IN (?)", dfiCurrentPrices).Find(&results).Error

	return
}

// GetFromDfiPrice 通过dfi_price获取内容
func (obj *_RDpaFeedItemMgr) GetFromDfiPrice(dfiPrice float32) (results []*RDpaFeedItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RDpaFeedItem{}).Where("`dfi_price` = ?", dfiPrice).Find(&results).Error

	return
}

// GetBatchFromDfiPrice 批量查找
func (obj *_RDpaFeedItemMgr) GetBatchFromDfiPrice(dfiPrices []float32) (results []*RDpaFeedItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RDpaFeedItem{}).Where("`dfi_price` IN (?)", dfiPrices).Find(&results).Error

	return
}

// GetFromDfiVenture 通过dfi_venture获取内容
func (obj *_RDpaFeedItemMgr) GetFromDfiVenture(dfiVenture string) (results []*RDpaFeedItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RDpaFeedItem{}).Where("`dfi_venture` = ?", dfiVenture).Find(&results).Error

	return
}

// GetBatchFromDfiVenture 批量查找
func (obj *_RDpaFeedItemMgr) GetBatchFromDfiVenture(dfiVentures []string) (results []*RDpaFeedItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RDpaFeedItem{}).Where("`dfi_venture` IN (?)", dfiVentures).Find(&results).Error

	return
}

// GetFromDfiUpdateTime 通过dfi_update_time获取内容
func (obj *_RDpaFeedItemMgr) GetFromDfiUpdateTime(dfiUpdateTime int) (results []*RDpaFeedItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RDpaFeedItem{}).Where("`dfi_update_time` = ?", dfiUpdateTime).Find(&results).Error

	return
}

// GetBatchFromDfiUpdateTime 批量查找
func (obj *_RDpaFeedItemMgr) GetBatchFromDfiUpdateTime(dfiUpdateTimes []int) (results []*RDpaFeedItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RDpaFeedItem{}).Where("`dfi_update_time` IN (?)", dfiUpdateTimes).Find(&results).Error

	return
}

// GetFromDfiDiscountPercentage 通过dfi_discount_percentage获取内容
func (obj *_RDpaFeedItemMgr) GetFromDfiDiscountPercentage(dfiDiscountPercentage string) (results []*RDpaFeedItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RDpaFeedItem{}).Where("`dfi_discount_percentage` = ?", dfiDiscountPercentage).Find(&results).Error

	return
}

// GetBatchFromDfiDiscountPercentage 批量查找
func (obj *_RDpaFeedItemMgr) GetBatchFromDfiDiscountPercentage(dfiDiscountPercentages []string) (results []*RDpaFeedItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RDpaFeedItem{}).Where("`dfi_discount_percentage` IN (?)", dfiDiscountPercentages).Find(&results).Error

	return
}

// GetFromDfiBenefitUsp 通过dfi_benefit_usp获取内容
func (obj *_RDpaFeedItemMgr) GetFromDfiBenefitUsp(dfiBenefitUsp string) (results []*RDpaFeedItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RDpaFeedItem{}).Where("`dfi_benefit_usp` = ?", dfiBenefitUsp).Find(&results).Error

	return
}

// GetBatchFromDfiBenefitUsp 批量查找
func (obj *_RDpaFeedItemMgr) GetBatchFromDfiBenefitUsp(dfiBenefitUsps []string) (results []*RDpaFeedItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RDpaFeedItem{}).Where("`dfi_benefit_usp` IN (?)", dfiBenefitUsps).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_RDpaFeedItemMgr) FetchByPrimaryKey(dfiID uint32) (result RDpaFeedItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RDpaFeedItem{}).Where("`dfi_id` = ?", dfiID).Find(&result).Error

	return
}

// FetchIndexByDfiSkuID  获取多个内容
func (obj *_RDpaFeedItemMgr) FetchIndexByDfiSkuID(dfiSkuID int64) (results []*RDpaFeedItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RDpaFeedItem{}).Where("`dfi_sku_id` = ?", dfiSkuID).Find(&results).Error

	return
}

// FetchIndexByDfiProductID  获取多个内容
func (obj *_RDpaFeedItemMgr) FetchIndexByDfiProductID(dfiProductID int64) (results []*RDpaFeedItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RDpaFeedItem{}).Where("`dfi_product_id` = ?", dfiProductID).Find(&results).Error

	return
}

// FetchIndexByDfiVenture  获取多个内容
func (obj *_RDpaFeedItemMgr) FetchIndexByDfiVenture(dfiVenture string) (results []*RDpaFeedItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RDpaFeedItem{}).Where("`dfi_venture` = ?", dfiVenture).Find(&results).Error

	return
}

// FetchIndexByDfiUpdateTime  获取多个内容
func (obj *_RDpaFeedItemMgr) FetchIndexByDfiUpdateTime(dfiUpdateTime int) (results []*RDpaFeedItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(RDpaFeedItem{}).Where("`dfi_update_time` = ?", dfiUpdateTime).Find(&results).Error

	return
}
