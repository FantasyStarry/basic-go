package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	//静态路由
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello Go")
	})
	//参数路由
	r.GET("/user/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		ctx.String(http.StatusOK, "hello, 这是参数路由"+name)
	})
	r.GET("/views/*.html", func(ctx *gin.Context) {
		page := ctx.Param(".html")
		ctx.String(http.StatusOK, "hello, 只是通配符路由"+page)
	})
	r.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
}
