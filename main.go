package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 处理跨域请求
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	r.Use(cors.New(config))

	// 部署前端静态网站
	r.Use(static.Serve("/", static.LocalFile("./static", false)))

	// 部署后端API
	r.GET("/api/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": "200",
			"data": "Ok",
		})
	})

	r.Run()
}
