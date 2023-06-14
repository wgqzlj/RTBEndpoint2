package model

// RAdCampaign [...]
type RAdCampaign struct {
	AcID                    uint64 `gorm:"primaryKey;column:ac_id;type:bigint(20) unsigned;not null"`
	AcCode                  string `gorm:"column:ac_code;type:varchar(100);not null;default:''"`
	AcCID                   int    `gorm:"index:ac_c_id;column:ac_c_id;type:int(11);default:0"`
	AcOID                   int    `gorm:"index:ac_o_id;column:ac_o_id;type:int(11);default:0"`
	AcName                  string `gorm:"column:ac_name;type:varchar(200);default:''"`
	AcDeeplink              string `gorm:"column:ac_deeplink;type:varchar(1000);default:''"`
	AcClickURL              string `gorm:"column:ac_click_url;type:varchar(1000);default:''"`
	AcImpressionTrackingURL string `gorm:"column:ac_impression_tracking_url;type:varchar(1000);default:''"`
	AcActID                 int    `gorm:"index:ac_act_id;column:ac_act_id;type:int(11);default:0"`
	AcClickTrackingURL      string `gorm:"column:ac_click_tracking_url;type:varchar(1000);default:''"`
	AcMemberID              string `gorm:"column:ac_member_id;type:varchar(200);default:''"`
	AcStatus                int8   `gorm:"column:ac_status;type:tinyint(4);default:1"`
	AcValidStartTime        int    `gorm:"column:ac_valid_start_time;type:int(11);default:0"`
	AcValidEndTime          int    `gorm:"column:ac_valid_end_time;type:int(11);default:0"`
	AcCreateTime            int    `gorm:"column:ac_create_time;type:int(11);default:0"`
	AcPriority              int    `gorm:"column:ac_priority;type:int(11);default:0"`
	AcRemark                string `gorm:"column:ac_remark;type:varchar(2000);default:''"`
	AcAcrCount              int    `gorm:"column:ac_acr_count;type:int(11);not null;default:0"`
	AcAcrcID                int    `gorm:"index:ac_acrc_id;column:ac_acrc_id;type:int(11);default:0"` // 关联的r_ad_creative_collection id
}

// RAdCampaignColumns get sql column name.获取数据库列名
var RAdCampaignColumns = struct {
	AcID                    string
	AcCode                  string
	AcCID                   string
	AcOID                   string
	AcName                  string
	AcDeeplink              string
	AcClickURL              string
	AcImpressionTrackingURL string
	AcActID                 string
	AcClickTrackingURL      string
	AcMemberID              string
	AcStatus                string
	AcValidStartTime        string
	AcValidEndTime          string
	AcCreateTime            string
	AcPriority              string
	AcRemark                string
	AcAcrCount              string
	AcAcrcID                string
}{
	AcID:                    "ac_id",
	AcCode:                  "ac_code",
	AcCID:                   "ac_c_id",
	AcOID:                   "ac_o_id",
	AcName:                  "ac_name",
	AcDeeplink:              "ac_deeplink",
	AcClickURL:              "ac_click_url",
	AcImpressionTrackingURL: "ac_impression_tracking_url",
	AcActID:                 "ac_act_id",
	AcClickTrackingURL:      "ac_click_tracking_url",
	AcMemberID:              "ac_member_id",
	AcStatus:                "ac_status",
	AcValidStartTime:        "ac_valid_start_time",
	AcValidEndTime:          "ac_valid_end_time",
	AcCreateTime:            "ac_create_time",
	AcPriority:              "ac_priority",
	AcRemark:                "ac_remark",
	AcAcrCount:              "ac_acr_count",
	AcAcrcID:                "ac_acrc_id",
}

// RAdCampaignCreativeRelation [...]
type RAdCampaignCreativeRelation struct {
	AccrID         uint32 `gorm:"primaryKey;column:accr_id;type:int(10) unsigned;not null"`
	AccrAcID       int    `gorm:"column:accr_ac_id;type:int(11);default:0"`
	AccrAcrID      int    `gorm:"column:accr_acr_id;type:int(11);default:0"`
	AccrAeID       int    `gorm:"column:accr_ae_id;type:int(11);default:0"`
	AccrIstID      int    `gorm:"column:accr_ist_id;type:int(11);default:0"`
	AccrStatus     int8   `gorm:"column:accr_status;type:tinyint(4);default:1"`
	AccrCreateTime int    `gorm:"column:accr_create_time;type:int(11);default:0"`
	AccrAtID       int    `gorm:"column:accr_at_id;type:int(11);default:0"`
}

// RAdCampaignCreativeRelationColumns get sql column name.获取数据库列名
var RAdCampaignCreativeRelationColumns = struct {
	AccrID         string
	AccrAcID       string
	AccrAcrID      string
	AccrAeID       string
	AccrIstID      string
	AccrStatus     string
	AccrCreateTime string
	AccrAtID       string
}{
	AccrID:         "accr_id",
	AccrAcID:       "accr_ac_id",
	AccrAcrID:      "accr_acr_id",
	AccrAeID:       "accr_ae_id",
	AccrIstID:      "accr_ist_id",
	AccrStatus:     "accr_status",
	AccrCreateTime: "accr_create_time",
	AccrAtID:       "accr_at_id",
}

// RAdCampaignCreativeRelationBak [...]
type RAdCampaignCreativeRelationBak struct {
	AccrID         uint32 `gorm:"primaryKey;column:accr_id;type:int(10) unsigned;not null"`
	AccrAcID       int    `gorm:"column:accr_ac_id;type:int(11);default:0"`
	AccrAcrID      int    `gorm:"column:accr_acr_id;type:int(11);default:0"`
	AccrAeID       int    `gorm:"column:accr_ae_id;type:int(11);default:0"`
	AccrIstID      int    `gorm:"column:accr_ist_id;type:int(11);default:0"`
	AccrStatus     int8   `gorm:"column:accr_status;type:tinyint(4);default:1"`
	AccrCreateTime int    `gorm:"column:accr_create_time;type:int(11);default:0"`
	AccrAtID       int    `gorm:"column:accr_at_id;type:int(11);default:0"`
}

// RAdCampaignCreativeRelationBakColumns get sql column name.获取数据库列名
var RAdCampaignCreativeRelationBakColumns = struct {
	AccrID         string
	AccrAcID       string
	AccrAcrID      string
	AccrAeID       string
	AccrIstID      string
	AccrStatus     string
	AccrCreateTime string
	AccrAtID       string
}{
	AccrID:         "accr_id",
	AccrAcID:       "accr_ac_id",
	AccrAcrID:      "accr_acr_id",
	AccrAeID:       "accr_ae_id",
	AccrIstID:      "accr_ist_id",
	AccrStatus:     "accr_status",
	AccrCreateTime: "accr_create_time",
	AccrAtID:       "accr_at_id",
}

// RAdCampaignCreativeRelationBak2 [...]
type RAdCampaignCreativeRelationBak2 struct {
	AccrID         uint32 `gorm:"primaryKey;column:accr_id;type:int(10) unsigned;not null"`
	AccrAcID       int    `gorm:"column:accr_ac_id;type:int(11);default:0"`
	AccrAcrID      int    `gorm:"column:accr_acr_id;type:int(11);default:0"`
	AccrAeID       int    `gorm:"column:accr_ae_id;type:int(11);default:0"`
	AccrIstID      int    `gorm:"column:accr_ist_id;type:int(11);default:0"`
	AccrStatus     int8   `gorm:"column:accr_status;type:tinyint(4);default:1"`
	AccrCreateTime int    `gorm:"column:accr_create_time;type:int(11);default:0"`
	AccrAtID       int    `gorm:"column:accr_at_id;type:int(11);default:0"`
}

// RAdCampaignCreativeRelationBak2Columns get sql column name.获取数据库列名
var RAdCampaignCreativeRelationBak2Columns = struct {
	AccrID         string
	AccrAcID       string
	AccrAcrID      string
	AccrAeID       string
	AccrIstID      string
	AccrStatus     string
	AccrCreateTime string
	AccrAtID       string
}{
	AccrID:         "accr_id",
	AccrAcID:       "accr_ac_id",
	AccrAcrID:      "accr_acr_id",
	AccrAeID:       "accr_ae_id",
	AccrIstID:      "accr_ist_id",
	AccrStatus:     "accr_status",
	AccrCreateTime: "accr_create_time",
	AccrAtID:       "accr_at_id",
}

// RAdCampaignType [...]
type RAdCampaignType struct {
	ActID     uint32 `gorm:"primaryKey;column:act_id;type:int(10) unsigned;not null"`
	ActCode   string `gorm:"column:act_code;type:varchar(200);default:''"`
	ActName   string `gorm:"column:act_name;type:varchar(200);default:''"`
	ActStatus int8   `gorm:"column:act_status;type:tinyint(4);not null;default:1"`
}

// RAdCampaignTypeColumns get sql column name.获取数据库列名
var RAdCampaignTypeColumns = struct {
	ActID     string
	ActCode   string
	ActName   string
	ActStatus string
}{
	ActID:     "act_id",
	ActCode:   "act_code",
	ActName:   "act_name",
	ActStatus: "act_status",
}

// RAdCreative [...]
type RAdCreative struct {
	AcrID       uint32  `gorm:"primaryKey;column:acr_id;type:int(10) unsigned;not null"`
	AcrName     string  `gorm:"column:acr_name;type:varchar(200);default:''"`
	AcrCID      int     `gorm:"column:acr_c_id;type:int(11);default:0"` // 国家ID
	AcrOID      int     `gorm:"column:acr_o_id;type:int(11);default:0"`
	AcrURL      string  `gorm:"column:acr_url;type:varchar(200);default:''"`
	AcrCloakURL string  `gorm:"column:acr_cloak_url;type:varchar(2000);not null;default:''"`
	AcrWidth    float32 `gorm:"column:acr_width;type:float;default:0"`
	AcrHeight   float32 `gorm:"column:acr_height;type:float;default:0"`
	AcrStatus   int8    `gorm:"column:acr_status;type:tinyint(4);default:1"`
	AcrItID     int     `gorm:"column:acr_it_id;type:int(11);default:0"` // 1 banner 2 video
	AcrExt      string  `gorm:"column:acr_ext;type:varchar(2000);default:''"`
	AcrRemark   string  `gorm:"column:acr_remark;type:varchar(200);default:''"`
}

// RAdCreativeColumns get sql column name.获取数据库列名
var RAdCreativeColumns = struct {
	AcrID       string
	AcrName     string
	AcrCID      string
	AcrOID      string
	AcrURL      string
	AcrCloakURL string
	AcrWidth    string
	AcrHeight   string
	AcrStatus   string
	AcrItID     string
	AcrExt      string
	AcrRemark   string
}{
	AcrID:       "acr_id",
	AcrName:     "acr_name",
	AcrCID:      "acr_c_id",
	AcrOID:      "acr_o_id",
	AcrURL:      "acr_url",
	AcrCloakURL: "acr_cloak_url",
	AcrWidth:    "acr_width",
	AcrHeight:   "acr_height",
	AcrStatus:   "acr_status",
	AcrItID:     "acr_it_id",
	AcrExt:      "acr_ext",
	AcrRemark:   "acr_remark",
}

// RAdCreativeCollection [...]
type RAdCreativeCollection struct {
	AcrcID         uint32 `gorm:"primaryKey;column:acrc_id;type:int(11) unsigned;not null"`
	AcrcName       string `gorm:"column:acrc_name;type:varchar(200);not null"`
	AcrcCreateTime uint32 `gorm:"column:acrc_create_time;type:int(11) unsigned;default:0"`
	AcrcStatus     int8   `gorm:"column:acrc_status;type:tinyint(4);default:1"`
	AcrcCID        int    `gorm:"index:acrc_c_id;column:acrc_c_id;type:int(11);default:0"`
	AcrcOID        int    `gorm:"index:acrc_o_id;column:acrc_o_id;type:int(11);default:0"`
}

// RAdCreativeCollectionColumns get sql column name.获取数据库列名
var RAdCreativeCollectionColumns = struct {
	AcrcID         string
	AcrcName       string
	AcrcCreateTime string
	AcrcStatus     string
	AcrcCID        string
	AcrcOID        string
}{
	AcrcID:         "acrc_id",
	AcrcName:       "acrc_name",
	AcrcCreateTime: "acrc_create_time",
	AcrcStatus:     "acrc_status",
	AcrcCID:        "acrc_c_id",
	AcrcOID:        "acrc_o_id",
}

// RAdCreativeCollectionRelation [...]
type RAdCreativeCollectionRelation struct {
	AcrcrID         uint32 `gorm:"primaryKey;column:acrcr_id;type:int(10) unsigned;not null"`
	AcrcrAcrcID     int    `gorm:"index:acrcr_acrc_id;column:acrcr_acrc_id;type:int(11);default:0"`
	AcrcrAcrID      int    `gorm:"index:acrcr_acr_id;column:acrcr_acr_id;type:int(11);default:0"`
	AcrcrAeID       int    `gorm:"index:acrcr_ae_id;column:acrcr_ae_id;type:int(11);default:0"`
	AcrcrIstID      int    `gorm:"index:acrcr_ist_id;column:acrcr_ist_id;type:int(11);default:0"`
	AcrcrStatus     int8   `gorm:"index:acrcr_status;column:acrcr_status;type:tinyint(4);default:1"`
	AcrcrCreateTime int    `gorm:"column:acrcr_create_time;type:int(11);default:0"`
	AcrcrAtID       int    `gorm:"index:acrcr_at_id;column:acrcr_at_id;type:int(11);default:0"`
}

// RAdCreativeCollectionRelationColumns get sql column name.获取数据库列名
var RAdCreativeCollectionRelationColumns = struct {
	AcrcrID         string
	AcrcrAcrcID     string
	AcrcrAcrID      string
	AcrcrAeID       string
	AcrcrIstID      string
	AcrcrStatus     string
	AcrcrCreateTime string
	AcrcrAtID       string
}{
	AcrcrID:         "acrcr_id",
	AcrcrAcrcID:     "acrcr_acrc_id",
	AcrcrAcrID:      "acrcr_acr_id",
	AcrcrAeID:       "acrcr_ae_id",
	AcrcrIstID:      "acrcr_ist_id",
	AcrcrStatus:     "acrcr_status",
	AcrcrCreateTime: "acrcr_create_time",
	AcrcrAtID:       "acrcr_at_id",
}

// RAdCreativeVideo [...]
type RAdCreativeVideo struct {
	AccvID                     uint32 `gorm:"primaryKey;column:accv_id;type:int(10) unsigned;not null"`
	AccvAcrID                  int    `gorm:"column:accv_acr_id;type:int(11);default:0"`
	AccvTitle                  string `gorm:"column:accv_title;type:varchar(200);default:''"`
	AccvDescription            string `gorm:"column:accv_description;type:varchar(200);default:''"`
	AccvCta                    string `gorm:"column:accv_cta;type:varchar(200);default:''"`
	AccvDuration               string `gorm:"column:accv_duration;type:varchar(200);default:''"`
	AccvIcon                   string `gorm:"column:accv_icon;type:varchar(200);default:''"`
	AccvCampanionImageURL      string `gorm:"column:accv_campanion_image_url;type:varchar(200);default:''"`
	AccvCampanionImageCloakURL string `gorm:"column:accv_campanion_image_cloak_url;type:varchar(2000);not null;default:''"`
}

// RAdCreativeVideoColumns get sql column name.获取数据库列名
var RAdCreativeVideoColumns = struct {
	AccvID                     string
	AccvAcrID                  string
	AccvTitle                  string
	AccvDescription            string
	AccvCta                    string
	AccvDuration               string
	AccvIcon                   string
	AccvCampanionImageURL      string
	AccvCampanionImageCloakURL string
}{
	AccvID:                     "accv_id",
	AccvAcrID:                  "accv_acr_id",
	AccvTitle:                  "accv_title",
	AccvDescription:            "accv_description",
	AccvCta:                    "accv_cta",
	AccvDuration:               "accv_duration",
	AccvIcon:                   "accv_icon",
	AccvCampanionImageURL:      "accv_campanion_image_url",
	AccvCampanionImageCloakURL: "accv_campanion_image_cloak_url",
}

// RAdExchange [...]
type RAdExchange struct {
	AeID     uint32 `gorm:"primaryKey;column:ae_id;type:int(10) unsigned;not null"`
	AeName   string `gorm:"unique;column:ae_name;type:varchar(200);default:''"`
	AeCode   string `gorm:"unique;column:ae_code;type:varchar(200);default:''"`
	AeRemark string `gorm:"column:ae_remark;type:varchar(200);default:''"`
	AeStatus int8   `gorm:"column:ae_status;type:tinyint(4);default:1"`
	AeSeq    int    `gorm:"column:ae_seq;type:int(11);default:0"`
}

// RAdExchangeColumns get sql column name.获取数据库列名
var RAdExchangeColumns = struct {
	AeID     string
	AeName   string
	AeCode   string
	AeRemark string
	AeStatus string
	AeSeq    string
}{
	AeID:     "ae_id",
	AeName:   "ae_name",
	AeCode:   "ae_code",
	AeRemark: "ae_remark",
	AeStatus: "ae_status",
	AeSeq:    "ae_seq",
}

// RAdTemplate [...]
type RAdTemplate struct {
	AtID         uint32 `gorm:"primaryKey;column:at_id;type:int(10) unsigned;not null"`
	AtName       string `gorm:"column:at_name;type:varchar(200);default:''"`
	AtFilename   string `gorm:"column:at_filename;type:varchar(200);default:''"`
	AtCode       string `gorm:"column:at_code;type:varchar(200);default:''"`
	AtStatus     int8   `gorm:"column:at_status;type:tinyint(4);default:1"`
	AtContent    string `gorm:"column:at_content;type:longtext"`
	AtItID       int8   `gorm:"column:at_it_id;type:tinyint(4);default:1"` // 1 banner 2 video
	AtRemark     string `gorm:"column:at_remark;type:varchar(200);default:''"`
	AtCreateTime int    `gorm:"column:at_create_time;type:int(11);default:0"`
}

// RAdTemplateColumns get sql column name.获取数据库列名
var RAdTemplateColumns = struct {
	AtID         string
	AtName       string
	AtFilename   string
	AtCode       string
	AtStatus     string
	AtContent    string
	AtItID       string
	AtRemark     string
	AtCreateTime string
}{
	AtID:         "at_id",
	AtName:       "at_name",
	AtFilename:   "at_filename",
	AtCode:       "at_code",
	AtStatus:     "at_status",
	AtContent:    "at_content",
	AtItID:       "at_it_id",
	AtRemark:     "at_remark",
	AtCreateTime: "at_create_time",
}

// RAliexpressCampaign [...]
type RAliexpressCampaign struct {
	AcID              uint32 `gorm:"primaryKey;column:ac_id;type:int(11) unsigned;not null"`
	AcCampaignID      string `gorm:"column:ac_campaign_id;type:varchar(100);default:''"`
	AcCampaignName    string `gorm:"column:ac_campaign_name;type:varchar(100);default:''"`
	AcBusinessType    string `gorm:"column:ac_business_type;type:varchar(100);default:''"`
	AcMaterialType    string `gorm:"column:ac_material_type;type:varchar(100);default:''"`
	AcCountry         string `gorm:"column:ac_country;type:varchar(100);default:''"`
	AcTargetAudiences string `gorm:"column:ac_target_audiences;type:varchar(100);default:''"`
	AcCreateTime      int    `gorm:"column:ac_create_time;type:int(11);default:0"`
	AcOnline          int8   `gorm:"column:ac_online;type:tinyint(4);default:1"`
	AcLandingPageURL  string `gorm:"column:ac_landing_page_url;type:varchar(500);default:''"`
}

// RAliexpressCampaignColumns get sql column name.获取数据库列名
var RAliexpressCampaignColumns = struct {
	AcID              string
	AcCampaignID      string
	AcCampaignName    string
	AcBusinessType    string
	AcMaterialType    string
	AcCountry         string
	AcTargetAudiences string
	AcCreateTime      string
	AcOnline          string
	AcLandingPageURL  string
}{
	AcID:              "ac_id",
	AcCampaignID:      "ac_campaign_id",
	AcCampaignName:    "ac_campaign_name",
	AcBusinessType:    "ac_business_type",
	AcMaterialType:    "ac_material_type",
	AcCountry:         "ac_country",
	AcTargetAudiences: "ac_target_audiences",
	AcCreateTime:      "ac_create_time",
	AcOnline:          "ac_online",
	AcLandingPageURL:  "ac_landing_page_url",
}

// RAliexpressDailyReport [...]
type RAliexpressDailyReport struct {
	AdrID               uint32  `gorm:"primaryKey;column:adr_id;type:int(11) unsigned;not null"`
	AdrDate             string  `gorm:"column:adr_date;type:varchar(20);default:''"`
	AdrCampaignID       string  `gorm:"column:adr_campaign_id;type:varchar(200);default:''"`
	AdrImpression       int     `gorm:"column:adr_impression;type:int(11);default:0"`
	AdrClick            int     `gorm:"column:adr_click;type:int(11);default:0"`
	AdrPostback         int     `gorm:"column:adr_postback;type:int(11);default:0"`
	AdrCost             float32 `gorm:"column:adr_cost;type:float;default:0"`
	AdrTrueImpression   int     `gorm:"column:adr_true_impression;type:int(11);default:0"`
	AdrTrueClick        int     `gorm:"column:adr_true_click;type:int(11);default:0"`
	AdrTrueCost         float32 `gorm:"column:adr_true_cost;type:float;default:0"`
	AdrTruePostback     int     `gorm:"column:adr_true_postback;type:int(11);default:0"`
	AdrCreateTime       int     `gorm:"column:adr_create_time;type:int(11);default:0"`
	AdrAeCost           float32 `gorm:"column:adr_ae_cost;type:float;default:0"`
	AdrAeUv             int     `gorm:"column:adr_ae_uv;type:int(11);default:0"`
	AdrAeMemberUv       int     `gorm:"column:adr_ae_member_uv;type:int(11);default:0"`
	AdrAeCrtOrdByrCnt   int     `gorm:"column:adr_ae_crt_ord_byr_cnt;type:int(11);default:0"`
	AdrAeCrtOrdCnt      int     `gorm:"column:adr_ae_crt_ord_cnt;type:int(11);default:0"`
	AdrAeOrdGmv         float32 `gorm:"column:adr_ae_ord_gmv;type:float;default:0"`
	AdrAeAffiCommission float32 `gorm:"column:adr_ae_affi_commission;type:float;default:0"`
	AdrAeP4pCost        float32 `gorm:"column:adr_ae_p4p_cost;type:float;default:0"`
	AdrAeTotalIncome    float32 `gorm:"column:adr_ae_total_income;type:float;default:0"`
}

// RAliexpressDailyReportColumns get sql column name.获取数据库列名
var RAliexpressDailyReportColumns = struct {
	AdrID               string
	AdrDate             string
	AdrCampaignID       string
	AdrImpression       string
	AdrClick            string
	AdrPostback         string
	AdrCost             string
	AdrTrueImpression   string
	AdrTrueClick        string
	AdrTrueCost         string
	AdrTruePostback     string
	AdrCreateTime       string
	AdrAeCost           string
	AdrAeUv             string
	AdrAeMemberUv       string
	AdrAeCrtOrdByrCnt   string
	AdrAeCrtOrdCnt      string
	AdrAeOrdGmv         string
	AdrAeAffiCommission string
	AdrAeP4pCost        string
	AdrAeTotalIncome    string
}{
	AdrID:               "adr_id",
	AdrDate:             "adr_date",
	AdrCampaignID:       "adr_campaign_id",
	AdrImpression:       "adr_impression",
	AdrClick:            "adr_click",
	AdrPostback:         "adr_postback",
	AdrCost:             "adr_cost",
	AdrTrueImpression:   "adr_true_impression",
	AdrTrueClick:        "adr_true_click",
	AdrTrueCost:         "adr_true_cost",
	AdrTruePostback:     "adr_true_postback",
	AdrCreateTime:       "adr_create_time",
	AdrAeCost:           "adr_ae_cost",
	AdrAeUv:             "adr_ae_uv",
	AdrAeMemberUv:       "adr_ae_member_uv",
	AdrAeCrtOrdByrCnt:   "adr_ae_crt_ord_byr_cnt",
	AdrAeCrtOrdCnt:      "adr_ae_crt_ord_cnt",
	AdrAeOrdGmv:         "adr_ae_ord_gmv",
	AdrAeAffiCommission: "adr_ae_affi_commission",
	AdrAeP4pCost:        "adr_ae_p4p_cost",
	AdrAeTotalIncome:    "adr_ae_total_income",
}

// RBlacklistAdSlot [...]
type RBlacklistAdSlot struct {
	BasID     uint32 `gorm:"primaryKey;column:bas_id;type:int(10) unsigned;not null"`
	BasBundle string `gorm:"index:bas_bundle;column:bas_bundle;type:varchar(300);default:''"`
	BasWidth  int    `gorm:"index:bas_width;column:bas_width;type:int(11);default:0"`
	BasHeight int    `gorm:"index:bas_height;column:bas_height;type:int(11);default:0"`
	BasStatus int8   `gorm:"index:bas_status;column:bas_status;type:tinyint(4);default:1"`
	BasIstID  int    `gorm:"column:bas_ist_id;type:int(11);default:1"`
}

// RBlacklistAdSlotColumns get sql column name.获取数据库列名
var RBlacklistAdSlotColumns = struct {
	BasID     string
	BasBundle string
	BasWidth  string
	BasHeight string
	BasStatus string
	BasIstID  string
}{
	BasID:     "bas_id",
	BasBundle: "bas_bundle",
	BasWidth:  "bas_width",
	BasHeight: "bas_height",
	BasStatus: "bas_status",
	BasIstID:  "bas_ist_id",
}

// RBlacklistAdSlotCountryAdxRelation [...]
type RBlacklistAdSlotCountryAdxRelation struct {
	BascarID         uint32 `gorm:"primaryKey;column:bascar_id;type:int(10) unsigned;not null"`
	BascarBasID      int    `gorm:"index:bascar_bas_id;column:bascar_bas_id;type:int(11);default:0"`
	BascarCID        int    `gorm:"index:bascar_c_id;column:bascar_c_id;type:int(11);default:0"`
	BascarAeID       int    `gorm:"index:bascar_ae_id;column:bascar_ae_id;type:int(11);default:0"`
	BascarStatus     int8   `gorm:"column:bascar_status;type:tinyint(4);default:1"`
	BascarCreateTime int    `gorm:"column:bascar_create_time;type:int(11);default:0"`
}

// RBlacklistAdSlotCountryAdxRelationColumns get sql column name.获取数据库列名
var RBlacklistAdSlotCountryAdxRelationColumns = struct {
	BascarID         string
	BascarBasID      string
	BascarCID        string
	BascarAeID       string
	BascarStatus     string
	BascarCreateTime string
}{
	BascarID:         "bascar_id",
	BascarBasID:      "bascar_bas_id",
	BascarCID:        "bascar_c_id",
	BascarAeID:       "bascar_ae_id",
	BascarStatus:     "bascar_status",
	BascarCreateTime: "bascar_create_time",
}

// RBlacklistApp [...]
type RBlacklistApp struct {
	BaID         uint32 `gorm:"primaryKey;column:ba_id;type:int(10) unsigned;not null"`
	BaName       string `gorm:"column:ba_name;type:varchar(200);default:''"`
	BaBundle     string `gorm:"column:ba_bundle;type:varchar(300);default:''"`
	BaCreatetime int    `gorm:"column:ba_createtime;type:int(11);default:0"`
	BaStatus     int8   `gorm:"index:ba_status;column:ba_status;type:tinyint(4);default:1"`
	BaItID       int    `gorm:"column:ba_it_id;type:int(11);not null;default:0"`
}

// RBlacklistAppColumns get sql column name.获取数据库列名
var RBlacklistAppColumns = struct {
	BaID         string
	BaName       string
	BaBundle     string
	BaCreatetime string
	BaStatus     string
	BaItID       string
}{
	BaID:         "ba_id",
	BaName:       "ba_name",
	BaBundle:     "ba_bundle",
	BaCreatetime: "ba_createtime",
	BaStatus:     "ba_status",
	BaItID:       "ba_it_id",
}

// RBlacklistAppCountryAdxRelation [...]
type RBlacklistAppCountryAdxRelation struct {
	BacarID         uint32 `gorm:"primaryKey;column:bacar_id;type:int(10) unsigned;not null"`
	BacarBaID       int    `gorm:"index:bacar_ba_id;column:bacar_ba_id;type:int(11);default:0"`
	BacarCID        int    `gorm:"index:bacar_c_id;column:bacar_c_id;type:int(11);default:0"`
	BacarAeID       int    `gorm:"index:bacar_ae_id;column:bacar_ae_id;type:int(11);default:0"`
	BacarStatus     int8   `gorm:"index:bacar_status;column:bacar_status;type:tinyint(4);default:1"`
	BacarCreateTime int    `gorm:"index:bacar_create_time;column:bacar_create_time;type:int(11);default:0"`
}

// RBlacklistAppCountryAdxRelationColumns get sql column name.获取数据库列名
var RBlacklistAppCountryAdxRelationColumns = struct {
	BacarID         string
	BacarBaID       string
	BacarCID        string
	BacarAeID       string
	BacarStatus     string
	BacarCreateTime string
}{
	BacarID:         "bacar_id",
	BacarBaID:       "bacar_ba_id",
	BacarCID:        "bacar_c_id",
	BacarAeID:       "bacar_ae_id",
	BacarStatus:     "bacar_status",
	BacarCreateTime: "bacar_create_time",
}

// RCountry [...]
type RCountry struct {
	CID          uint32 `gorm:"primaryKey;column:c_id;type:int(10) unsigned;not null"`
	CName        string `gorm:"column:c_name;type:varchar(200);default:''"`
	CCode3       string `gorm:"unique;column:c_code3;type:varchar(10);default:''"`
	CCode2       string `gorm:"unique;column:c_code2;type:varchar(10);default:''"`
	CTimezoneUtc int    `gorm:"column:c_timezone_utc;type:int(11);default:0"`
	CStatus      int8   `gorm:"column:c_status;type:tinyint(4);default:1"`
}

// RCountryColumns get sql column name.获取数据库列名
var RCountryColumns = struct {
	CID          string
	CName        string
	CCode3       string
	CCode2       string
	CTimezoneUtc string
	CStatus      string
}{
	CID:          "c_id",
	CName:        "c_name",
	CCode3:       "c_code3",
	CCode2:       "c_code2",
	CTimezoneUtc: "c_timezone_utc",
	CStatus:      "c_status",
}

// RDpaFeedItem [...]
type RDpaFeedItem struct {
	DfiID                     uint32  `gorm:"primaryKey;column:dfi_id;type:int(10) unsigned;not null"`
	DfiSkuID                  int64   `gorm:"index:dfi_sku_id;column:dfi_sku_id;type:bigint(20);default:0"`
	DfiProductID              int64   `gorm:"index:dfi_product_id;column:dfi_product_id;type:bigint(20);default:0"`
	DfiProductNameCsvEsc      string  `gorm:"column:dfi_product_name_csv_esc;type:varchar(1000);default:''"`
	DfiBrandName              string  `gorm:"column:dfi_brand_name;type:varchar(200);default:''"`
	DfiRegionalCategory1Name  string  `gorm:"column:dfi_regional_category1_name;type:varchar(200);default:''"`
	DfiRegionalCategory2Name  string  `gorm:"column:dfi_regional_category2_name;type:varchar(200);default:''"`
	DfiRegionalCategory3Name  string  `gorm:"column:dfi_regional_category3_name;type:varchar(200);default:''"`
	DfiProductURL             string  `gorm:"column:dfi_product_url;type:varchar(1000);default:''"`
	DfiURLMainImage           string  `gorm:"column:dfi_url_main_image;type:varchar(200);default:''"`
	DfiProductMediumImg       string  `gorm:"column:dfi_product_medium_img;type:varchar(200);default:''"`
	DfiProductBigImg          string  `gorm:"column:dfi_product_big_img;type:varchar(200);default:''"`
	DfiImageURL2              string  `gorm:"column:dfi_image_url_2;type:varchar(200);default:''"`
	DfiImageURL3              string  `gorm:"column:dfi_image_url_3;type:varchar(200);default:''"`
	DfiImageURL4              string  `gorm:"column:dfi_image_url_4;type:varchar(200);default:''"`
	DfiImageURL5              string  `gorm:"column:dfi_image_url_5;type:varchar(200);default:''"`
	DfiDescription            string  `gorm:"column:dfi_description;type:varchar(200);default:''"`
	DfiVentureCategory1NameEn string  `gorm:"column:dfi_venture_category1_name_en;type:varchar(11);default:''"`
	DfiVentureCategory2NameEn string  `gorm:"column:dfi_venture_category2_name_en;type:varchar(200);default:''"`
	DfiVentureCategory3NameEn string  `gorm:"column:dfi_venture_category3_name_en;type:varchar(200);default:''"`
	DfiPromotionTag           string  `gorm:"column:dfi_promotion_tag;type:varchar(200);default:''"`
	DfiPromotionPrice         float32 `gorm:"column:dfi_promotion_price;type:float;default:0"`
	DfiBrandLogo              string  `gorm:"column:dfi_brand_logo;type:varchar(11);default:''"`
	DfiCurrentPrice           float32 `gorm:"column:dfi_current_price;type:float;default:0"`
	DfiPrice                  float32 `gorm:"column:dfi_price;type:float;default:0"`
	DfiVenture                string  `gorm:"index:dfi_venture;column:dfi_venture;type:varchar(20);default:''"`
	DfiUpdateTime             int     `gorm:"index:dfi_update_time;column:dfi_update_time;type:int(11);not null;default:0"`
	DfiDiscountPercentage     string  `gorm:"column:dfi_discount_percentage;type:varchar(50);not null;default:''"`
	DfiBenefitUsp             string  `gorm:"column:dfi_benefit_usp;type:varchar(200);not null;default:''"`
}

// RDpaFeedItemColumns get sql column name.获取数据库列名
var RDpaFeedItemColumns = struct {
	DfiID                     string
	DfiSkuID                  string
	DfiProductID              string
	DfiProductNameCsvEsc      string
	DfiBrandName              string
	DfiRegionalCategory1Name  string
	DfiRegionalCategory2Name  string
	DfiRegionalCategory3Name  string
	DfiProductURL             string
	DfiURLMainImage           string
	DfiProductMediumImg       string
	DfiProductBigImg          string
	DfiImageURL2              string
	DfiImageURL3              string
	DfiImageURL4              string
	DfiImageURL5              string
	DfiDescription            string
	DfiVentureCategory1NameEn string
	DfiVentureCategory2NameEn string
	DfiVentureCategory3NameEn string
	DfiPromotionTag           string
	DfiPromotionPrice         string
	DfiBrandLogo              string
	DfiCurrentPrice           string
	DfiPrice                  string
	DfiVenture                string
	DfiUpdateTime             string
	DfiDiscountPercentage     string
	DfiBenefitUsp             string
}{
	DfiID:                     "dfi_id",
	DfiSkuID:                  "dfi_sku_id",
	DfiProductID:              "dfi_product_id",
	DfiProductNameCsvEsc:      "dfi_product_name_csv_esc",
	DfiBrandName:              "dfi_brand_name",
	DfiRegionalCategory1Name:  "dfi_regional_category1_name",
	DfiRegionalCategory2Name:  "dfi_regional_category2_name",
	DfiRegionalCategory3Name:  "dfi_regional_category3_name",
	DfiProductURL:             "dfi_product_url",
	DfiURLMainImage:           "dfi_url_main_image",
	DfiProductMediumImg:       "dfi_product_medium_img",
	DfiProductBigImg:          "dfi_product_big_img",
	DfiImageURL2:              "dfi_image_url_2",
	DfiImageURL3:              "dfi_image_url_3",
	DfiImageURL4:              "dfi_image_url_4",
	DfiImageURL5:              "dfi_image_url_5",
	DfiDescription:            "dfi_description",
	DfiVentureCategory1NameEn: "dfi_venture_category1_name_en",
	DfiVentureCategory2NameEn: "dfi_venture_category2_name_en",
	DfiVentureCategory3NameEn: "dfi_venture_category3_name_en",
	DfiPromotionTag:           "dfi_promotion_tag",
	DfiPromotionPrice:         "dfi_promotion_price",
	DfiBrandLogo:              "dfi_brand_logo",
	DfiCurrentPrice:           "dfi_current_price",
	DfiPrice:                  "dfi_price",
	DfiVenture:                "dfi_venture",
	DfiUpdateTime:             "dfi_update_time",
	DfiDiscountPercentage:     "dfi_discount_percentage",
	DfiBenefitUsp:             "dfi_benefit_usp",
}

// RImpressionSubType [...]
type RImpressionSubType struct {
	IstID     uint32 `gorm:"primaryKey;column:ist_id;type:int(10) unsigned;not null"`
	IstCode   string `gorm:"column:ist_code;type:varchar(50);default:''"`
	IstName   string `gorm:"column:ist_name;type:varchar(50);default:''"`
	IstItID   int    `gorm:"column:ist_it_id;type:int(11);default:0"`
	IstStatus int8   `gorm:"column:ist_status;type:tinyint(4);default:1"`
	IstSuffix string `gorm:"column:ist_suffix;type:varchar(50);not null;default:''"`
}

// RImpressionSubTypeColumns get sql column name.获取数据库列名
var RImpressionSubTypeColumns = struct {
	IstID     string
	IstCode   string
	IstName   string
	IstItID   string
	IstStatus string
	IstSuffix string
}{
	IstID:     "ist_id",
	IstCode:   "ist_code",
	IstName:   "ist_name",
	IstItID:   "ist_it_id",
	IstStatus: "ist_status",
	IstSuffix: "ist_suffix",
}

// RImpressionType [...]
type RImpressionType struct {
	ItID     int    `gorm:"primaryKey;column:it_id;type:int(11);not null"`
	ItCode   string `gorm:"column:it_code;type:varchar(255)"`
	ItName   string `gorm:"column:it_name;type:varchar(255)"`
	ItStatus int8   `gorm:"column:it_status;type:tinyint(4);default:1"`
}

// RImpressionTypeColumns get sql column name.获取数据库列名
var RImpressionTypeColumns = struct {
	ItID     string
	ItCode   string
	ItName   string
	ItStatus string
}{
	ItID:     "it_id",
	ItCode:   "it_code",
	ItName:   "it_name",
	ItStatus: "it_status",
}

// ROrder [...]
type ROrder struct {
	OID         uint32 `gorm:"primaryKey;column:o_id;type:int(10) unsigned;not null"`
	OName       string `gorm:"column:o_name;type:varchar(200);default:''"`
	OCode       string `gorm:"column:o_code;type:varchar(200);default:''"`
	OStatus     int8   `gorm:"column:o_status;type:tinyint(4);default:1"`
	OCreateTime int    `gorm:"column:o_create_time;type:int(11);default:0"`
}

// ROrderColumns get sql column name.获取数据库列名
var ROrderColumns = struct {
	OID         string
	OName       string
	OCode       string
	OStatus     string
	OCreateTime string
}{
	OID:         "o_id",
	OName:       "o_name",
	OCode:       "o_code",
	OStatus:     "o_status",
	OCreateTime: "o_create_time",
}

// RPreauditCreativeid [...]
type RPreauditCreativeid struct {
	PcID              uint32 `gorm:"primaryKey;column:pc_id;type:int(10) unsigned;not null"`
	PcCode            string `gorm:"column:pc_code;type:varchar(200);default:''"`
	PcAeID            int    `gorm:"column:pc_ae_id;type:int(11);default:0"`
	PcCID             int    `gorm:"column:pc_c_id;type:int(11);default:0"`
	PcCreateTime      uint32 `gorm:"column:pc_create_time;type:int(10) unsigned;default:0"`
	PcPreAuditTime    uint32 `gorm:"column:pc_pre_audit_time;type:int(10) unsigned;default:0"`
	PcStartEnableTime uint32 `gorm:"column:pc_start_enable_time;type:int(10) unsigned;default:0"`
	PcStatus          int8   `gorm:"column:pc_status;type:tinyint(4);default:1"` // 1 正常 0删除 2待提交 3 被封禁
	PcAcID            int    `gorm:"column:pc_ac_id;type:int(11);default:0"`
	PcSubmitCount     int    `gorm:"column:pc_submit_count;type:int(11);not null;default:0"`
}

// RPreauditCreativeidColumns get sql column name.获取数据库列名
var RPreauditCreativeidColumns = struct {
	PcID              string
	PcCode            string
	PcAeID            string
	PcCID             string
	PcCreateTime      string
	PcPreAuditTime    string
	PcStartEnableTime string
	PcStatus          string
	PcAcID            string
	PcSubmitCount     string
}{
	PcID:              "pc_id",
	PcCode:            "pc_code",
	PcAeID:            "pc_ae_id",
	PcCID:             "pc_c_id",
	PcCreateTime:      "pc_create_time",
	PcPreAuditTime:    "pc_pre_audit_time",
	PcStartEnableTime: "pc_start_enable_time",
	PcStatus:          "pc_status",
	PcAcID:            "pc_ac_id",
	PcSubmitCount:     "pc_submit_count",
}

// RSmartPriceAdSlot [...]
type RSmartPriceAdSlot struct {
	SpasID         uint32 `gorm:"primaryKey;column:spas_id;type:int(10) unsigned;not null"`
	SpasBundle     string `gorm:"column:spas_bundle;type:varchar(200);default:''"`
	SpasWidth      int    `gorm:"column:spas_width;type:int(11);default:0"`
	SpasHeight     int    `gorm:"column:spas_height;type:int(11);default:0"`
	SpasCreateTime int    `gorm:"column:spas_create_time;type:int(11);default:0"`
	SpasStatus     int8   `gorm:"column:spas_status;type:tinyint(4);default:1"`
	SpasIstID      int    `gorm:"column:spas_ist_id;type:int(11);default:0"`
}

// RSmartPriceAdSlotColumns get sql column name.获取数据库列名
var RSmartPriceAdSlotColumns = struct {
	SpasID         string
	SpasBundle     string
	SpasWidth      string
	SpasHeight     string
	SpasCreateTime string
	SpasStatus     string
	SpasIstID      string
}{
	SpasID:         "spas_id",
	SpasBundle:     "spas_bundle",
	SpasWidth:      "spas_width",
	SpasHeight:     "spas_height",
	SpasCreateTime: "spas_create_time",
	SpasStatus:     "spas_status",
	SpasIstID:      "spas_ist_id",
}

// RSmartPriceAdSlotHistory [...]
type RSmartPriceAdSlotHistory struct {
	SpashID            uint32  `gorm:"primaryKey;column:spash_id;type:int(10) unsigned;not null"`
	SpashSpasID        int     `gorm:"column:spash_spas_id;type:int(11);default:0"`
	SpashAeID          int     `gorm:"column:spash_ae_id;type:int(11);default:0"`
	SpashCID           int     `gorm:"column:spash_c_id;type:int(11);default:0"`
	SpashTime          int     `gorm:"column:spash_time;type:int(11);default:0"`
	SpashBidCount      int     `gorm:"column:spash_bid_count;type:int(11);default:0"`
	SpashImpCount      int     `gorm:"column:spash_imp_count;type:int(11);default:0"`
	SpashClickCount    int     `gorm:"column:spash_click_count;type:int(11);default:0"`
	SpashPostbackCount int     `gorm:"column:spash_postback_count;type:int(11);default:0"`
	SpashCost          float32 `gorm:"column:spash_cost;type:float;default:0"`
	SpashRoi           float32 `gorm:"column:spash_roi;type:float;default:0"`
	SpashEcpm          float32 `gorm:"column:spash_ecpm;type:float;default:0"`
}

// RSmartPriceAdSlotHistoryColumns get sql column name.获取数据库列名
var RSmartPriceAdSlotHistoryColumns = struct {
	SpashID            string
	SpashSpasID        string
	SpashAeID          string
	SpashCID           string
	SpashTime          string
	SpashBidCount      string
	SpashImpCount      string
	SpashClickCount    string
	SpashPostbackCount string
	SpashCost          string
	SpashRoi           string
	SpashEcpm          string
}{
	SpashID:            "spash_id",
	SpashSpasID:        "spash_spas_id",
	SpashAeID:          "spash_ae_id",
	SpashCID:           "spash_c_id",
	SpashTime:          "spash_time",
	SpashBidCount:      "spash_bid_count",
	SpashImpCount:      "spash_imp_count",
	SpashClickCount:    "spash_click_count",
	SpashPostbackCount: "spash_postback_count",
	SpashCost:          "spash_cost",
	SpashRoi:           "spash_roi",
	SpashEcpm:          "spash_ecpm",
}

// RSmartPriceApp [...]
type RSmartPriceApp struct {
	SpaID         uint32 `gorm:"primaryKey;column:spa_id;type:int(10) unsigned;not null"`
	SpaBundle     string `gorm:"column:spa_bundle;type:varchar(200);default:''"`
	SpaStatus     int8   `gorm:"column:spa_status;type:tinyint(4);default:1"`
	SpaCreateTime int    `gorm:"column:spa_create_time;type:int(11);default:0"`
}

// RSmartPriceAppColumns get sql column name.获取数据库列名
var RSmartPriceAppColumns = struct {
	SpaID         string
	SpaBundle     string
	SpaStatus     string
	SpaCreateTime string
}{
	SpaID:         "spa_id",
	SpaBundle:     "spa_bundle",
	SpaStatus:     "spa_status",
	SpaCreateTime: "spa_create_time",
}

// RSmartPriceAppHistory [...]
type RSmartPriceAppHistory struct {
	SpahID            uint32  `gorm:"primaryKey;column:spah_id;type:int(10) unsigned;not null"`
	SpahSpaID         int     `gorm:"column:spah_spa_id;type:int(11);default:0"`
	SpahAeID          int     `gorm:"column:spah_ae_id;type:int(11);default:0"`
	SpahCID           int     `gorm:"column:spah_c_id;type:int(11);default:0"`
	SpahItID          int     `gorm:"column:spah_it_id;type:int(11);default:0"`
	SpahTime          int     `gorm:"column:spah_time;type:int(11);default:0"`
	SpahBidCount      int     `gorm:"column:spah_bid_count;type:int(11);default:0"`
	SpahImpCount      int     `gorm:"column:spah_imp_count;type:int(11);default:0"`
	SpahClickCount    int     `gorm:"column:spah_click_count;type:int(11);default:0"`
	SpahPostbackCount int     `gorm:"column:spah_postback_count;type:int(11);default:0"`
	SpahCost          float32 `gorm:"column:spah_cost;type:float;default:0"`
	SpahRoi           float32 `gorm:"column:spah_roi;type:float;default:0"`
	SpahEcpm          float32 `gorm:"column:spah_ecpm;type:float;default:0"`
}

// RSmartPriceAppHistoryColumns get sql column name.获取数据库列名
var RSmartPriceAppHistoryColumns = struct {
	SpahID            string
	SpahSpaID         string
	SpahAeID          string
	SpahCID           string
	SpahItID          string
	SpahTime          string
	SpahBidCount      string
	SpahImpCount      string
	SpahClickCount    string
	SpahPostbackCount string
	SpahCost          string
	SpahRoi           string
	SpahEcpm          string
}{
	SpahID:            "spah_id",
	SpahSpaID:         "spah_spa_id",
	SpahAeID:          "spah_ae_id",
	SpahCID:           "spah_c_id",
	SpahItID:          "spah_it_id",
	SpahTime:          "spah_time",
	SpahBidCount:      "spah_bid_count",
	SpahImpCount:      "spah_imp_count",
	SpahClickCount:    "spah_click_count",
	SpahPostbackCount: "spah_postback_count",
	SpahCost:          "spah_cost",
	SpahRoi:           "spah_roi",
	SpahEcpm:          "spah_ecpm",
}

// RSmartPriceRule [...]
type RSmartPriceRule struct {
	SmrID          uint32  `gorm:"primaryKey;column:smr_id;type:int(10) unsigned;not null"`
	SmrCID         int     `gorm:"column:smr_c_id;type:int(11);default:0"`
	SmrIstID       int     `gorm:"column:smr_ist_id;type:int(11);default:0"`
	SmrAnchorPrice float32 `gorm:"column:smr_anchor_price;type:float;default:0"`
}

// RSmartPriceRuleColumns get sql column name.获取数据库列名
var RSmartPriceRuleColumns = struct {
	SmrID          string
	SmrCID         string
	SmrIstID       string
	SmrAnchorPrice string
}{
	SmrID:          "smr_id",
	SmrCID:         "smr_c_id",
	SmrIstID:       "smr_ist_id",
	SmrAnchorPrice: "smr_anchor_price",
}

// RSystemAdmin 后台用户信息表
type RSystemAdmin struct {
	SaID           uint32 `gorm:"primaryKey;column:sa_id;type:int(10) unsigned;not null"`  // 系统管理员ID
	SaUsername     string `gorm:"unique;column:sa_username;type:varchar(100);default:''"`  // 系统管理员用户名
	SaPassword     string `gorm:"column:sa_password;type:varchar(100);default:''"`         // 系统管理员密码
	SaPasswordSalt string `gorm:"column:sa_password_salt;type:varchar(200);default:''"`    // 密码盐
	SaLoginIP      uint32 `gorm:"column:sa_login_ip;type:int(10) unsigned;default:0"`      // 上一次登陆IP
	SaLoginTime    int    `gorm:"column:sa_login_time;type:int(11);default:0"`             // 上一次登陆时间
	SaStatus       int    `gorm:"index:sa_status;column:sa_status;type:int(11);default:1"` // 系统管理员状态，1为正常，0为禁用
}

// RSystemAdminColumns get sql column name.获取数据库列名
var RSystemAdminColumns = struct {
	SaID           string
	SaUsername     string
	SaPassword     string
	SaPasswordSalt string
	SaLoginIP      string
	SaLoginTime    string
	SaStatus       string
}{
	SaID:           "sa_id",
	SaUsername:     "sa_username",
	SaPassword:     "sa_password",
	SaPasswordSalt: "sa_password_salt",
	SaLoginIP:      "sa_login_ip",
	SaLoginTime:    "sa_login_time",
	SaStatus:       "sa_status",
}

// RSystemLog 包含日志清洗，系统操作等日志
type RSystemLog struct {
	LrID         uint32 `gorm:"primaryKey;column:lr_id;type:int(10) unsigned;not null"`
	LrUID        int    `gorm:"column:lr_u_id;type:int(11);default:0"`
	LrCode       string `gorm:"column:lr_code;type:varchar(200);default:''"`
	LrMessage    string `gorm:"column:lr_message;type:varchar(200);default:''"`
	LrCreateTime int    `gorm:"column:lr_create_time;type:int(11);default:0"`
}

// RSystemLogColumns get sql column name.获取数据库列名
var RSystemLogColumns = struct {
	LrID         string
	LrUID        string
	LrCode       string
	LrMessage    string
	LrCreateTime string
}{
	LrID:         "lr_id",
	LrUID:        "lr_u_id",
	LrCode:       "lr_code",
	LrMessage:    "lr_message",
	LrCreateTime: "lr_create_time",
}

// RWhitelistAdSlot [...]
type RWhitelistAdSlot struct {
	WasID     uint32  `gorm:"primaryKey;column:was_id;type:int(10) unsigned;not null"`
	WasBundle string  `gorm:"column:was_bundle;type:varchar(50);default:''"`
	WasWidth  int     `gorm:"column:was_width;type:int(11);default:0"`
	WasHeight int     `gorm:"column:was_height;type:int(11);default:0"`
	WasPrice  float32 `gorm:"column:was_price;type:float;default:0"`
	WasStatus int8    `gorm:"column:was_status;type:tinyint(4);default:1"`
}

// RWhitelistAdSlotColumns get sql column name.获取数据库列名
var RWhitelistAdSlotColumns = struct {
	WasID     string
	WasBundle string
	WasWidth  string
	WasHeight string
	WasPrice  string
	WasStatus string
}{
	WasID:     "was_id",
	WasBundle: "was_bundle",
	WasWidth:  "was_width",
	WasHeight: "was_height",
	WasPrice:  "was_price",
	WasStatus: "was_status",
}

// RWhitelistApp [...]
type RWhitelistApp struct {
	WaID         uint32 `gorm:"primaryKey;column:wa_id;type:int(10) unsigned;not null"`
	WaName       string `gorm:"column:wa_name;type:varchar(200);default:''"`
	WaBundle     string `gorm:"column:wa_bundle;type:varchar(200);default:''"`
	WaCreatetime int    `gorm:"column:wa_createtime;type:int(11);default:0"`
	WaStatus     int8   `gorm:"column:wa_status;type:tinyint(4);default:1"`
}

// RWhitelistAppColumns get sql column name.获取数据库列名
var RWhitelistAppColumns = struct {
	WaID         string
	WaName       string
	WaBundle     string
	WaCreatetime string
	WaStatus     string
}{
	WaID:         "wa_id",
	WaName:       "wa_name",
	WaBundle:     "wa_bundle",
	WaCreatetime: "wa_createtime",
	WaStatus:     "wa_status",
}

// RWhitelistAppBidPrice [...]
type RWhitelistAppBidPrice struct {
	WabpID uint32 `gorm:"primaryKey;column:wabp_id;type:int(10) unsigned;not null"`
}

// RWhitelistAppBidPriceColumns get sql column name.获取数据库列名
var RWhitelistAppBidPriceColumns = struct {
	WabpID string
}{
	WabpID: "wabp_id",
}
