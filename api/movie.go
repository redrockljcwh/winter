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
		c.Abort()
		return
	}
	movie,err := service.SelectMovie(id)
	if err!=nil{
		tool.RespInternalError(c)
		c.Abort()
		return
	}
	score := GetAvgScore(c)
	idstr1 := movie.Main_performer[0:5]
	id1,err:=strconv.Atoi(idstr1)
	idstr2 := movie.Main_performer[6:]
	id2,err := strconv.Atoi(idstr2)
	star1,err:=service.SelectStarById(id1)
	star2,err:=service.SelectStarById(id2)
	c.JSON(http.StatusOK,gin.H{
		"name":movie.Name,
		"director":movie.Director,
		"writer":movie.Writer,
		"main":star1.Name+star2.Name,
		"type":movie.Type,
		"country":movie.Country,
		"language":movie.Language,
		"date":movie.Date,
		"length":movie.Length,
		"stuff":movie.Stuff,//演职员
		"picnum":movie.PicNum,
		"score":score,
	})
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
func GetAvgScore(c *gin.Context)(score float64){
	movieidstr:=c.Param("id")
	movieid,err:=strconv.Atoi(movieidstr)
	if err != nil {
		tool.RespErrorWithDate(c,"请输入合法电影id")
		return 0
	}
	comments,err:=service.GetCommentsByMovieId(movieid)
	if err != nil {
		tool.RespErrorWithDate(c,"请输入合法电影id")
		return 0
	}
	tool.RespSuccessful(c)
	var scores float64
	for _, comment := range comments{
	 scores = scores + comment.Score
	}
	length := len(comments)
	l := float64(length)
	avgscore1:=scores/l
	avgscore := FloatRound(avgscore1,1)
	return avgscore
}
func FloatRound(f float64, n int) float64 {
	format := "%." + strconv.Itoa(n) + "f"
	res, _ := strconv.ParseFloat(fmt.Sprintf(format, f), 64)
	return res
}//此函数来自https://studygolang.com/articles/12927
func SearchByMovieName(c *gin.Context){
	name:=c.Param("name")
	movies,err:=service.SearchMovieByName(name)
	if err != nil{
		tool.RespErrorWithDate(c,"搜索的内容不存在")
		return
	}
	tool.RespSuccessful(c)
	for _,movie := range movies{
		c.JSON(http.StatusOK,gin.H{
			"movieid":movie.Id,
			"name":movie.Name,
		})

	}
	return
}