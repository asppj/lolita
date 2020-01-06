package article

import (
	"net/http"
	"t-mk-opentrace/interval/material/dao"

	"github.com/gin-gonic/gin"
	"github.com/siddontang/go/bson"
)

var companyID = "P00000000043"
var companyDB = "wp_data_20190612"

// ShowArticle 显示文章列表
func ShowArticle(ctx *gin.Context) {
	c := dao.ArticleService{
		CompanyID: companyID,
		CompanyDB: companyDB,
		Limit:     10,
	}
	query := bson.M{}
	resData, err := c.List(ctx.Request.Context(), query)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	cnt, err := c.ListEs(ctx.Request.Context(), companyID)
	res := struct {
		Data     interface{} `json:"data"`
		Total    int64       `json:"total"`
		Page     int64       `json:"page"`
		PageSize int64       `json:"pageSize"`
	}{}
	res.Data = resData
	res.Total = cnt
	res.Page = 1
	res.PageSize = 10
	ctx.JSON(http.StatusOK, res)
}
