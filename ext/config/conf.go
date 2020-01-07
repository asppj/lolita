package config

import (
	// "git.dustess.com/mk-base/es-driver/es"
	"t-mk-opentrace/ext/constant"
	// "git.dustess.com/mk-base/redis-driver/redis"
	"github.com/stevenroose/gonfig"
)

// MongoAuth mongodb认证信息结构
type MongoAuth struct {
	Username    string `id:"username" default:"foo"` // username
	Password    string `id:"password" default:"bar"` // password
	AuthSource  string `id:"source" default:"admin"` // auth source
	PasswordSet bool   `id:"on" default:"true"`      // is auth on
}

// GranType 授权类型
type GranType struct {
	AuthorizationCode string `id:"authorizationCode" default:"authorization_code"`
	ClientCredential  string `id:"clientCredential" default:"client_credential"`
	RefreshToken      string `id:"refreshToken" default:"refreshToken"`
}

// Config 配置结构
type Config struct {
	ConfigFile string `short:"c" default:"config.json"`

	Server struct {
		Host string `id:"host" default:"" desc:"listen addr"`                 // 监听地址，配置为空字符串表示监听0.0.0.0
		Port string `id:"port" default:"5000" desc:"http server listen port"` // 启动端口
	} `id:"server" desc:"server config"`

	RPC struct {
		Host string `id:"host" default:"" desc:"listen rpc addr"`              // rpc监听地址，配置为空字符串表示监听0.0.0.0
		Port string `id:"port" default:"50000" desc:"http server listen port"` // 启动端口
	} `id:"rpc" desc:"rpc config"`

	OpenTrace struct {
		Host string `id:"host" default:"localhost"`
		Port string `id:"port" default:"6831"`
	}
	Mode constant.RunTimeMode `id:"mode" default:"debug"` // 运行模式 debug|test|release
}

// Init 初始化
func (config *Config) Init() error {
	return gonfig.Load(config, gonfig.Conf{
		ConfigFileVariable:  constant.ConfigFileVariable, // enables passing --configfile myfile.conf
		FileDefaultFilename: constant.ConfigFileDefaultName,
		FileDecoder:         gonfig.DecoderJSON,
		EnvPrefix:           constant.ConfigEnvPrefix,
	})
}
