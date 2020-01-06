package mongo

// Auth mongodb认证信息结构
type Auth struct {
	Username    string `id:"username" default:"foo"` // username
	Password    string `id:"password" default:"bar"` // password
	AuthSource  string `id:"source" default:"admin"` // auth source
	PasswordSet bool   `id:"on" default:"true"`      // is auth on
}

// Config 配置
type Config struct {
	ClientName ClientName
	Addr       string
	Auth       Auth
}
