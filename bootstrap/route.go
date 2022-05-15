package bootstrap

import (
	"goblog/routes"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func SetupRoute(router *gin.Engine) {
	// Register global middleware
	registerGlobalMiddleware(router)
	// Register API routes
	routes.RegisterAPIRoutes(router)
	// Register 404 handler
	setup404Handler(router)
}

func registerGlobalMiddleware(router *gin.Engine) {
	router.Use(
		gin.Logger(),
		gin.Recovery(),
	)
}

func setup404Handler(router *gin.Engine) {
	router.NoRoute(func(c *gin.Context) {
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			c.String(http.StatusNotFound, "頁面返回 404")
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error_code": 404, "error_message": "路由未定義，請確認 URL 及請求方法是否正確！"})
		}
	})
}
