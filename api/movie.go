package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"static-server/model"
	"static-server/service"
	"static-server/tool"
	"strconv"
)

func GetMovie(c *gin.Context){
	idstring:=c.Param("id")
	id,err1 := strconv.Atoi(idstring)
	if err1 != nil {
		tool.RespInternalError(c)
		fmt.Println(1)
		c.Abort()
		return
	}
	fmt.Println(1)
	movie,err := service.SelectMovie(id)
	if err!=nil{
		tool.RespInternalError(c)
		fmt.Println(2)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"name":movie.Name,
		"director":movie.Director,
		"type":movie.Type,
		"country":movie.Country,
		"language":movie.Language,
		"date":movie.Date,
	})
	tool.RespSuccessful(c)
	return

}
func InsertMovie(c *gin.Context){
	username:=getting(c)
	if username!="admin"{
		tool.RespErrorWithDate(c,"您没有权限访问此网页")
		c.Abort()
		return
	}
	var movie model.Movie
	movie.Name=c.PostForm("name")
	movie.Type=c.PostForm("type")
	movie.Date=c.PostForm("date")

	err:=service.InsertMovie(movie)
	if err != nil{
		tool.RespInternalError(c)
	}
	return
}
func SearchByMovieName(c *gin.Context){
	name:=c.Param("name")
	movies,err:=service.SearchMovieByName(name)
	if err != nil{
		tool.RespErrorWithDate(c,"搜索的内容不存在")
		return
	}
	for _,movie := range movies{
		c.JSON(http.StatusOK,gin.H{
			"movieid":movie.Id,
			"name":movie.Name,
		})

	}
	return
}