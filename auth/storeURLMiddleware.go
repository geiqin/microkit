package auth

import (
	"github.com/geiqin/gotools/helper"
	"github.com/gin-gonic/gin"
)

//处理URL带店铺参数中间件, url中带 :storeId
func StoreURLMiddleware(ctx *gin.Context) {
	strId := ctx.Param("storeId")
	storeId := helper.StringToInt64(strId)
	if ctx.Keys == nil {
		ctx.Keys = make(map[string]interface{})
	}
	if storeId > 0 {
		ctx.Keys["store_id"] = strId
		ctx.Header("Auth-Store-Id", strId)
	}
	ctx.Next()
}
