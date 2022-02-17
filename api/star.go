package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"static-server/service"
	"static-server/tool"
	"strconv"
)

func GetStarById(c *gin.Context){
	idstr:=c.Param("id")
	id,err:=strconv.Atoi(idstr)
	if err!=nil{
		tool.RespErrorWithDate(c,"输入的id非法")
		return
	}
	star,err:=service.SelectStarById(id)
	if err!=nil{
		fmt.Println(err)
	    tool.RespInternalError(c)
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"id":star.Id,
		"name":star.Name,
		"birth":star.Birth,
		"birthplace":star.Birthplace,
		"work":star.Work,
		"introduction":star.Introduction,
		"picNum":star.PicNum,
	})
	return
}
