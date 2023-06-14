package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"rtbep/bid"
	"rtbep/rtb/bid/aliexpress"
	"rtbep/rtb/common/pricerule"
	"rtbep/rtb/common/utils"
	"rtbep/rtb/config"
	"strconv"
	"time"

	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func OpenRTBHeaderMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("x-openrtb-version", "2.5")
		c.Next()
	}
}

func initLogurs() {
	//  初始化 Logus 日志
	logPath := ""
	if config.IsDev() {
		logPath = "./rtb.log"
	} else {
		logPath = "/var/log/rtb.log"
	}
	f, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	logrus.SetFormatter(&logrus.JSONFormatter{})

	if config.IsDev() {
		wrt := io.MultiWriter(os.Stdout, f)
		logrus.SetOutput(wrt)
	} else {
		wrt := io.Writer(f)
		logrus.SetOutput(wrt)
	}
}

func main() {
	initLogurs()
	// 刷新 banner adslot, instl bundle, video bundle, rewardedvideo 动态出价
	go pricerule.RefreshBidPriceData()
	go pricerule.RefreshAliExpressBidPriceData()
	go pricerule.RefreshAliExpressLongTrailData()

	// 生产环境用 new
	// 是否用New？
	r := gin.Default()

	// !!! 这两个必须一起使用
	// r := gin.New()
	// r.Use(gin.Recovery())
	// // gin.SetMode(gin.ReleaseMode)

	configData := config.Read()
	r.Use(OpenRTBHeaderMiddleWare())
	r.Use(gzip.Gzip(gzip.DefaultCompression)) // 开启gzip压缩

	if configData.Main.Dev {
		pprof.Register(r) // 性能调试
	}

	r.POST("/v1/bid", bid.BidHandler)
	r.POST("/v1/bid/:adxid", bid.BidHandler)
	r.POST("/unity", bid.BidHandler)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"now": time.Now().Format("2006-01-02 15:04:05"),
		})
	})

	r.GET("/v1/bid", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Request by Post method, please:)",
		})
	})
	r.GET("/v1/bid/:adxid", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Request by Post method, please:)",
		})
	})

	r.GET("/lzd/go", func(c *gin.Context) {
		url := c.Query("url")
		fmt.Println(url)
		c.Redirect(302, url)
	})

	r.GET("/timezone", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"UTC_HOUR": time.Now().UTC().Format("15"),
			"SGP":      utils.GetHourByCountry("SGP"),
			"MYS":      utils.GetHourByCountry("MYS"),
			"VNM":      utils.GetHourByCountry("VNM"),
			"PHL":      utils.GetHourByCountry("PHL"),
			"IDN":      utils.GetHourByCountry("IDN"),
			"THA":      utils.GetHourByCountry("THA"),
		})
	})

	r.GET("/bidprice", func(c *gin.Context) {
		module := c.Query("module")
		switch module {
		case "banner":
			c.JSON(200, gin.H{
				"data": pricerule.BannerAdSlotBidRecords,
			})

		case "bannercampaign":
			c.JSON(200, gin.H{
				"data": pricerule.BannerAdSlotCampaignBidRecords,
			})
		case "instl":
			c.JSON(200, gin.H{
				"data": pricerule.InstlBidRecords,
			})
		case "instlcampaign":
			c.JSON(200, gin.H{
				"data": pricerule.InstlCampaignBidRecords,
			})
		case "video":
			c.JSON(200, gin.H{
				"data": pricerule.VideoBidRecords,
			})
		case "videocampaign":
			c.JSON(200, gin.H{
				"data": pricerule.VideoCampaignBidRecords,
			})
		case "rewardedvideo":
			c.JSON(200, gin.H{
				"data": pricerule.RewardedVideoBidRecords,
			})
		case "rewardedvideocampaign":
			c.JSON(200, gin.H{
				"data": pricerule.RewardedVideoCampaignBidRecords,
			})
		case "":
			c.JSON(200, gin.H{
				"data": "module query required",
			})
		}
	})

	r.GET("/aliexpress/bidprice", func(c *gin.Context) {
		module := c.Query("module")
		switch module {
		case "banner":
			c.JSON(200, gin.H{
				"data": pricerule.AliExpressBannerAdSlotBidRecords,
			})

		case "bannercampaign":
			c.JSON(200, gin.H{
				"data": pricerule.AliExpressBannerAdSlotCampaignBidRecords,
			})
		case "instl":
			c.JSON(200, gin.H{
				"data": pricerule.AliExpressInstlBidRecords,
			})
		case "instlcampaign":
			c.JSON(200, gin.H{
				"data": pricerule.AliExpressInstlCampaignBidRecords,
			})
		case "video":
			c.JSON(200, gin.H{
				"data": pricerule.AliExpressVideoBidRecords,
			})
		case "videocampaign":
			c.JSON(200, gin.H{
				"data": pricerule.AliExpressVideoCampaignBidRecords,
			})
		case "rewardedvideo":
			c.JSON(200, gin.H{
				"data": pricerule.AliExpressRewardedVideoBidRecords,
			})
		case "rewardedvideocampaign":
			c.JSON(200, gin.H{
				"data": pricerule.AliExpressRewardedVideoCampaignBidRecords,
			})
		case "":
			c.JSON(200, gin.H{
				"data": "module query required",
			})
		}
	})

	r.GET("/cpistack", utils.CpiLength)
	r.GET("/cpiana", utils.CpcAnalytics)

	r.GET("/aliexpress/clearcache", aliexpress.ClearCache)

	// r.GET("/lzdreport", bid.LazadaReport)

	r.GET("/dpaitem", func(c *gin.Context) {
		// 生成各个国家的Google Template CSS
		itemId := c.DefaultQuery("itemid", "")
		itemIdInt, _ := strconv.Atoi(itemId)
		dpaItem, _ := utils.GetDPAItem([]int64{int64(itemIdInt)})
		c.JSON(200, gin.H{
			"data": dpaItem,
		})
	})

	r.GET("/pricerefreshtime", func(c *gin.Context) {
		// 生成各个国家的Google Template CSS
		c.JSON(200, gin.H{
			"data": pricerule.PriceUpdateTime(),
		})
	})

	r.GET("/aliexpress/pricerefreshtime", func(c *gin.Context) {
		// 生成各个国家的Google Template CSS
		c.JSON(200, gin.H{
			"data": pricerule.AliExpressPriceUpdateTime(),
		})
	})

	r.GET("/aliexpress/longtrailrefreshtime", func(c *gin.Context) {
		// 生成各个国家的Google Template CSS
		c.JSON(200, gin.H{
			"data": pricerule.AliExpressLongTrailUpdateTime(),
		})
	})

	r.GET("/countryqps", func(c *gin.Context) {
		utils.PrintISStats(c)
	})

	r.GET("/ae/devices", func(c *gin.Context) {
		utils.PrintDeviceId(c)
	})

	r.GET("/lzd/cps/device", func(c *gin.Context) {
		deviceId := c.Query("deviceid")
		country := c.Query("country")
		rt, err := utils.CPSUserTag(deviceId, country)
		if err != nil {
			c.JSON(200, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(200, gin.H{
				"data": rt,
			})
		}
	})

	r.GET("/aeqps", func(c *gin.Context) {
		utils.PrintAEQPSStats(c)
	})

	r.GET("/operaqps", func(c *gin.Context) {
		utils.PrintOperaStats(c)
	})

	r.GET("/braddj/count", bid.SummerDDJStats)
	r.GET("/braddj/stats", bid.SummerDDJStats)
	r.GET("/braddj/length", bid.SummerDDJStats)
	r.GET("/braddj/clicks", bid.SummerDDJClicks)

	r.GET("/lzdddj/count", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"count": utils.GetLazadaDdjDeviceCount(),
		})
	})
	r.GET("/lzdddj/targetcampaign", utils.PrintLazadaRTATargetCampaign)
	r.GET("/lzdddj/length", utils.LZDDDJListLength)

	r.GET("/lzdddj/clicks", func(c *gin.Context) {
		utils.PrintLazadaDDJClicks(c)
	})

	r.GET("/lzdddj/hits", func(c *gin.Context) {
		utils.PrintLazadaCPIHITCounts(c)
	})

	r.GET("/lzdddj/notinwhitelistmodel", func(c *gin.Context) {
		utils.PrintLazadaCPINotInWhiteListCounts(c)
	})

	r.GET("/lzdddj/inwhitelistmodel", func(c *gin.Context) {
		utils.PrintLazadaCPIInWhiteListCounts(c)
	})

	r.GET("/region", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"region": os.Getenv("REGION"),
		})
	})

	r.Run(configData.Main.Listen)
}
