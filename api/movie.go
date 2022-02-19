package api

import (
	"fmt"
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
	score := GetAvgScore(c)
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
		"score":score,
	})
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