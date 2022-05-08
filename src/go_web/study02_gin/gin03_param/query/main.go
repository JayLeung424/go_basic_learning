package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// 获取浏览器传来的query参数
	r.GET("/query", func(context *gin.Context) {
		// byKey
		query := context.Query("query")
		// 如果取不到 就用默认
		defaultQuery := context.DefaultQuery("defaultQuery", "default")
		// 判断状态
		queryStatus, ok := context.GetQuery("queryStatus")
		if !ok {
			queryStatus = "queryStatus"
		}
		context.JSON(http.StatusOK, gin.H{
			"query":        query,
			"defaultQuery": defaultQuery,
			"queryStatus":  queryStatus,
		})
	})

	r.Run(":9001")
}
