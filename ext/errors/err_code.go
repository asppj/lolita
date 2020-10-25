package errors

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

	// StatusCustomize401 接口级别错误码
	StatusCustomize401 = 400401
	// StatusCustomize402 接口级别错误码
	StatusCustomize402 = 400402
	// StatusCustomizeRefused 接口级别错误码
	StatusCustomizeRefused = 400403
	// StatusCustomizeNotFound 接口级别错误码
	StatusCustomizeNotFound = 400404
	// StatusCustomize405 接口级别错误码
	StatusCustomize405 = 400405
	// StatusCustomize406 接口级别错误码
	StatusCustomize406 = 400406
)
