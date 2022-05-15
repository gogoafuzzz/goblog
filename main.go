package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化 Gin 实例
	r := gin.Default()

	r.Use(gin.Logger(), gin.Recovery())

	r.NoRoute(func(c *gin.Context) {
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			c.String(http.StatusNotFound, "頁面返回 404")
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error_code": 404, "error_message": "路由未定義，請確認 URL 及請求方法是否正確！"})
		}
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Hello": "World!",
		})
	})

	// 运行服务
	r.Run(":8888")
}
