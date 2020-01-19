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

// 服务器定义错误码，统一处理
const (
	// StatusUnCustomize  未定义
	StatusUnCustomize = 50000
	// StatusOK StatusOK
	StatusOK = 500200 // RFC 7231, 6.3.1
	// StatusBadRequest  StatusBadRequest
	StatusBadRequest = 500400 // RFC 7231, 6.5.1
	// StatusUnauthorized  StatusUnauthorized
	StatusUnauthorized = 500401 // RFC 7235, 3.1
	// StatusForbidden StatusForbidden
	StatusForbidden = 500403 // RFC 7231, 6.5.3
	// StatusNotFound StatusNotFound
	StatusNotFound = 500404 // RFC 7231, 6.5.4
	// StatusNotAcceptable StatusNotAcceptable
	StatusNotAcceptable = 500406 // RFC 7231, 6.5.6
	// StatusRequestTimeout StatusRequestTimeout
	StatusRequestTimeout = 500408 // RFC 7231, 6.5.7
	// StatusInternalServerError StatusInternalServerError
	StatusInternalServerError = 500500 // RFC 7231, 6.6.1
	// StatusServiceUnavailable StatusServiceUnavailable
	StatusServiceUnavailable = 500503 // RFC 7231, 6.6.4
)

const (
	// 客户端自定义，同一个错误码每个接口含义不同

	// StatusCustomizeNotFound 接口级别错误码
	StatusCustomizeNotFound = 400401
	// StatusCustomizeRefused 接口级别错误码
	StatusCustomizeRefused = 400402
)
