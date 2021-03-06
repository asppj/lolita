package http

import (
	"github.com/asppj/lolita/cmd/http/middleware"
	"github.com/asppj/lolita/interval/plan"

	"github.com/gin-gonic/gin"
)

const prefix = "/v1/"
const pc = "pc"
const mobile = "m"
const api = "api"

// RegisterRouter 注册路由
func RegisterRouter(router *gin.Engine) error {
	// 注册端点
	middleware.RegisterEndpoint(router, "")
	apiRouter := router.Group(prefix + api)
	registerAPIRouter(apiRouter)
	mRouter := router.Group(prefix + mobile)
	registerMobileRouter(mRouter)
	pcRouter := router.Group(prefix + pc)
	registerPCRouter(pcRouter)
	router.GET("", plan.SearchPlan)
	router.GET("/user", plan.GetUser)
	return nil
}
