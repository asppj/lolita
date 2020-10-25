package api

const (
	// SUCCESS 成功
	SUCCESS = "ok" // 成功
	// ReqDataValError 请求数据校验失败
	ReqDataValError = "参数解析失败" // 请求数据校验失败
	// InternalServerError 服务器内部错误
	InternalServerError = "服务器内部错误"
	// Unauthorized session认证失败
	Unauthorized = "认证失败" // session认证失败
	// NotFound 未找到相关资源
	NotFound = "未找到相关资源" // 未找到相关资源
	// Expired 页面过期
	Expired = "页面过期" // 页面过期
	// Refuse 正在发布拒绝执行
	Refuse      = "请求被拒绝，建议刷新页面" // 正在发布拒绝执行
	unCustomize = "未定义"
)
