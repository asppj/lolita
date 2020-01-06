package plan

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUser GetUser
func GetUser(ctx *gin.Context) {
	TestTask(ctx.Request.Context())
	ctx.String(http.StatusOK, "user oks")

}
