package plan

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SearchPlan test open trace
func SearchPlan(ctx *gin.Context) {
	ctx.String(http.StatusOK, "router oks")
}
