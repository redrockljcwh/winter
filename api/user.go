package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
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
	check,err:=service.IsRepeatUsername(username)
	if err!=nil{
		fmt.Println(err)
		tool.RespErrorWithDate(c,"服务器错误")
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
	}
	tool.RespSuccessfulWithDate(c,"注册成功，请进行登录")
	return
}
func Login(c *gin.Context){
	username:=c.PostForm("username")
	password:=c.PostForm("password")
	if username == ""&&password == ""{
		tool.RespErrorWithDate(c,"用户名或密码不能为空")
		return
	}
	check,err:=service.IsPasswordCorrect(username,password)
	if err!=nil{
		fmt.Println(err)
		tool.RespInternalError(c)
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
	fmt.Println(tokenString)
}



//解析token

