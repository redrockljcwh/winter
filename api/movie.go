package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"static-server/service"
	"static-server/tool"
	"strconv"
)

func GetMovie(c *gin.Context){
	idstring:=c.Param("id")
	id,err1 := strconv.Atoi(idstring)
	if err1 != nil {
		tool.RespInternalError(c)
		c.Abort()
		return
	}
	movie,err := service.SelectMovie(id)
	if err!=nil{
		tool.RespInternalError(c)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"name":movie.Name,
		"director":movie.Director,
		"writer":movie.Writer,
		"main":movie.Main_performer,
		"type":movie.Type,
		"country":movie.Country,
		"language":movie.Language,
		"date":movie.Date,
		"length":movie.Length,
		"stuff":movie.Stuff,//演职员
		"picnum":movie.PicNum,
	})
	return

}