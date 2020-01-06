package constant

// 应用常量定义
const (
	OSSTTL   = 3600
	OSSTOKEN = "marketing_oss_token"

	PlatForm = "mk-pc"
)

// RunTimeMode 运行环境
type RunTimeMode string

const (
	// DebugMode 调试模式
	DebugMode RunTimeMode = "debug"

	// ReleaseMode 发布模式
	ReleaseMode RunTimeMode = "release"

	// TestMode 测试模式
	TestMode RunTimeMode = "test"
)

// 配置文件相关配置
const (
	ConfigFileVariable    = "configfile"  // 配置文件变量
	ConfigEnvPrefix       = "MK_BIZ_"     // ConfigEnvPrefix 配置文件，环境变量前缀
	ConfigFileDefaultName = "config.json" // 默认配置文件名
)
