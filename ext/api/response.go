package api

import (
	"net/http"

	"github.com/asppj/lolita/ext/errors"
	"github.com/asppj/lolita/ext/log-driver/log"

	"github.com/gin-gonic/gin"
)

// SendParamParseError 返回参数解析失败错误
func SendParamParseError(ctx *gin.Context, err *errors.Error) {
	res := failedRes()
	res.Msg = ReqDataValError
	if err != nil {
		if code := err.Code(); code != 0 {
			res.Code = code
		}
	} else {
		log.Error("SendParamParseError err is nil")
	}
	ctx.JSON(http.StatusOK, res)

}

// SendInternalServerError 服务器内部错误
func SendInternalServerError(ctx *gin.Context, err *errors.Error) {
	res := failedRes()
	res.Msg = InternalServerError
	if err != nil {
		if code := err.Code(); code != 0 {
			res.Code = code
		}
	} else {
		log.Error("SendInternalServerError err is nil")
	}
	ctx.JSON(http.StatusOK, res)
}

// SendNotFound 服务器内部错误
func SendNotFound(ctx *gin.Context, err *errors.Error) {
	res := failedRes()
	res.Msg = NotFound
	if err != nil {
		if code := err.Code(); code != 0 {
			res.Code = code
		}
	} else {
		log.Error("SendNotFound err is nil")
	}
	ctx.JSON(http.StatusOK, res)
}

// SendUnauthorized 身份认证失败
func SendUnauthorized(ctx *gin.Context, err *errors.Error) {
	res := failedRes()
	res.Msg = Unauthorized
	if err != nil {
		if code := err.Code(); code != 0 {
			res.Code = code
		}
	} else {
		log.Error("SendUnauthorized err is nil")
	}
	ctx.JSON(http.StatusOK, res)
}

// SendRefuse 拒绝请求
func SendRefuse(ctx *gin.Context, err *errors.Error) {
	res := failedRes()
	res.Msg = Refuse
	if err != nil {
		if code := err.Code(); code != 0 {
			res.Code = code
		}
	} else {
		log.Error("SendRefuse err is nil")
	}
	ctx.JSON(http.StatusOK, res)
}

// SendOK 成功
func SendOK(ctx *gin.Context) {
	res := successRes()
	ctx.JSON(http.StatusOK, res)
}

// Send 返回结果
func Send(ctx *gin.Context, resData interface{}, err *errors.Error) {
	var res *Response
	if err != nil {
		res = failedRes()
		res.Code = err.Code()
		res.Msg = err.Msg()
	} else {
		res = successRes()
		res.Data = resData
	}
	ctx.JSON(http.StatusOK, res)
}
