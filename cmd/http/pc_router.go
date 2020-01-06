package http

import (
	"t-mk-opentrace/interval/material/article"
	"t-mk-opentrace/interval/statics"

	"github.com/gin-gonic/gin"
)

// registerPCRouter 注册路由
// /v1/pc
func registerPCRouter(router *gin.RouterGroup) {
	router.GET("article", article.ShowArticle)
	router.GET("files", statics.ServerFile)
	router.GET("stream", statics.Stream)
}
