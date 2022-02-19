package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"static-server/model"
	"static-server/service"
	"static-server/tool"
	"time"
)
var Jwtkey = []byte("redrock")
var Str string
func Register(c *gin.Context){
	user:=model.User{}
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username==""&&password==""{
		tool.RespErrorWithDate(c,"用户名或密码不能为空")
		return
	}
	if len(username)<6{
		tool.RespErrorWithDate(c,"用户名不能小于6位。")
	}
	check,err:=service.IsRepeatUsername(username)
	if err!=nil{
		fmt.Println(err)
		tool.RespErrorWithDate(c,"服务器错误")
		c.Abort()
		return
	}
	if err==nil&&check {
		tool.RespErrorWithDate(c,"用户名已被注册")
		return
	}
    user.Username,user.Password = username,password
	err = service.Register(user)
	if err!=nil{
		fmt.Println(err)
		tool.RespErrorWithDate(c,"服务器错误")
		c.Abort()
	}
	tool.RespSuccessfulWithDate(c,"注册成功，请进行登录")
	return
}
func Login(c *gin.Context){
	username:=c.PostForm("username")
	password:=c.PostForm("password")
	if username == ""||password == ""{
		tool.RespErrorWithDate(c,"用户名或密码不能为空")
		return
	}
	check,err:=service.IsPasswordCorrect(username,password)
	if err!=nil{
		fmt.Println(err)
		tool.RespInternalError(c)
		c.Abort()
		return
	}
	if !check{
		tool.RespErrorWithDate(c,"用户名或密码错误")
		return
	}

	expireTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), //过期时间
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// fmt.Println(token)
	tokenString, err := token.SignedString(Jwtkey)
	if err != nil {
		fmt.Println(err)
	}
	Str = tokenString
	c.JSON(200, gin.H{"token": tokenString})
	//得到token后将token写入header，以此进行操作
	fmt.Println(tokenString)
	return
}
func ChangePassword(c *gin.Context){
	username:=getting(c)
    check,err:=service.IsRepeatUsername(username)
	if err!=nil {
		tool.RespInternalError(c)
		c.Abort()
		return
	}
	if !check{
		tool.RespInternalError(c)
		return
	}
	newpassword:=c.PostForm("newpassword")
	err = service.ChangePassword(username,newpassword)
	if err != nil{
		tool.RespInternalError(c)
	}
	tool.RespSuccessfulWithDate(c,"密码修改成功")
	return
}
func UpdateSelfinfo(c *gin.Context){
	selfInfo:=c.PostForm("SelfInfo")
	if len(selfInfo)>100{
		tool.RespSuccessfulWithDate(c,"不可大于100个字。")
		c.Abort()
		return
	}
	username:=getting(c)
    err:=service.UpdateSelfInfo(username,selfInfo)
	if err != nil{
		tool.RespInternalError(c)
		c.Abort()
		return
	}
	tool.RespSuccessful(c)
}
func UserInfo(c *gin.Context){
	username:=c.Param("username")
	user,err:=service.UserInfo(username)
	if err != nil {
		tool.RespInternalError(c)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"username":user.Username,
		"name":user.Name,
		"selfInfo":user.SelfInfo,
	})
	tool.RespSuccessful(c)
	return
}


//解析token

