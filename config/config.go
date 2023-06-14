/*
 * @Descripttion:
 * @version:
 * @Author: WGQ
 * @Date: 2021-10-27 16:08:40
 * @LastEditors: WGQ
 * @LastEditTime: 2021-11-12 11:12:02
 */
package config

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Main                  MainConfig  `yaml:"main"`
	RedisCluster          RedisConfig `yaml:"redisCluster"`
	CPSRedisCluster       RedisConfig `yaml:"cpsRedisCluster"`
	ConsoleServerEndpoint string      `yaml:"consoleServerEndpoint"`
	StatsServerHost       string      `yaml:"statsServerHost"`
	LandingPageServer     string      `yaml:"landingPageServer"`
	MySQLDSN              string      `yaml:"mysqlDSN"`
}

type MainConfig struct {
	Listen string `yaml:"listen"`
	Dev    bool   `yaml:"dev"`
}

type RedisConfig struct {
	Conn     string `yaml:"conn"`
	Password string `yaml:"password"`
}

var (
	_config *Config
)

func IsDev() bool {
	return Read().Main.Dev
}

func Read() *Config {

	if _config == nil {
		configFileName := "config.dist.yml"
		if os.Getenv("DEV") == "true" || os.Getenv("DEV") == "1" {
			configFileName = "config.dev.yml"
		}

		if os.Getenv("REGION") == "US" {
			configFileName = "config.us.dist.yml"
		}

		if os.Getenv("REGION") == "USEAST" {
			configFileName = "config.useast.dist.yml"
		}

		buf, err := ioutil.ReadFile("config/" + configFileName)
		if err != nil {
			// todo: 验证是否会因为文件读取的原因导致错误推出？
			panic(err)
		}

		// Unmarshal yml
		var config Config
		err = yaml.Unmarshal(buf, &config)
		if err != nil {
			panic(err)
		}
		_config = &config
		return &config
	} else {
		return _config
	}

}
