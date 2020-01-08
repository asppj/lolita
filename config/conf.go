package config

import (
	ginConfig "github.com/asppj/t-go-opentrace/ext/config"

	ginConstant "github.com/asppj/t-go-opentrace/ext/constant"

	"github.com/stevenroose/gonfig"
)

// Config 配置结构
type Config struct {
	ConfigFile string
	ginConfig.Config
}

var config *Config

// Init 初始化配置
func Init() error {
	config = &Config{}
	if err := config.Init(); err != nil {
		return err
	}
	return gonfig.Load(config, gonfig.Conf{
		ConfigFileVariable:  ginConstant.ConfigFileVariable, // enables passing --configfile myfile.conf
		FileDefaultFilename: ginConstant.ConfigFileDefaultName,
		FileDecoder:         gonfig.DecoderJSON,
		EnvPrefix:           ginConstant.ConfigEnvPrefix,
	})
}

// Get 获取配置
func Get() *Config {
	return config
}
