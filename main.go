package main

import (
	"static-server/api"
	"static-server/dao"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	dao.InitDB()
	r := gin.Default()

	// 处理跨域请求
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	r.Use(cors.New(config))

	// 部署前端静态网站
	r.Use(static.Serve("/", static.LocalFile("./static", false)))

	// 部署后端API
r.POST("/api/douban/register",api.Register)
	r.POST("/api/douban/Login",api.Login)
	r.Run()
}
