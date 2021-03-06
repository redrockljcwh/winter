package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"static-server/model"
	"static-server/service"
	"static-server/tool"
	"strconv"
	"time"
)

func PostComment(c *gin.Context){
	var comment model.Comment
	var err error
	movieidstr:=c.Param("id")
	comment.Movieid,_ = strconv.Atoi(movieidstr)
	comment.Username = getting(c)
	comment.Content = c.PostForm("comment")
	nowtime:=time.Now()
	nowtimestr := nowtime.Format("2002-02-02")
    comment.Date = nowtimestr
	scorestr := c.PostForm("score")
	scoreint,_ := strconv.Atoi(scorestr)
	comment.Score=float64(scoreint)
	if err != nil{
		fmt.Println(err)
		tool.RespErrorWithDate(c,"非法评分")
		return
	}
	if comment.Score < 2 || comment.Score > 10{
		tool.RespErrorWithDate(c,"非法评分")
		return
	}
	comment.Username = getting(c)
	check,err := service.CheckUserForComment(comment)
	if err!=nil{
		fmt.Println(err)
		tool.RespInternalError(c)
		c.Abort()
		return
	}
	if check{
		tool.RespErrorWithDate(c,"您已对本电影发表过评论了。不可重复评论。")
		return
	}
	err = service.PostComment(comment)
    if err != nil {
		tool.RespInternalError(c)
		fmt.Println(err)
		c.Abort()
		return
	}
	c.Abort()
	return

}
 func GetTheComments(c *gin.Context){
	 movieidstr:=c.Param("id")
     movieid,err:=strconv.Atoi(movieidstr)
	 if err != nil {
		 tool.RespErrorWithDate(c,"请输入合法电影id")
		 return
	 }
	 comments,err:=service.GetCommentsByMovieId(movieid)
	 if err != nil {
		 tool.RespErrorWithDate(c,"请输入合法电影id")
		 return
	 }
	for _, comment := range comments{
		 c.JSON(http.StatusOK,gin.H{
			 "id":comment.Id,
			 "movieid":movieid,
			 "username":comment.Username,
			 "content":comment.Content,
			 "date":comment.Date,
			 "score":comment.Score,
		 })
	 }
	 return
 }