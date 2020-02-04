package plan

import (
	"fmt"

	"github.com/asppj/t-go-opentrace/ext/api"
	"github.com/asppj/t-go-opentrace/ext/errors"
	"github.com/gin-gonic/gin"
)

// GetUser GetUser
func GetUser(ctx *gin.Context) {
	// res := api.DefaultRes()
	resData := "正常数据"
	err := errors.NewWithMsg(fmt.Errorf("customize"), "自定义错误", errors.StatusBadRequest)
	api.Send(ctx, resData, err)
}
