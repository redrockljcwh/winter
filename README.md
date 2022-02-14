# winter
key-要求-数据类型-说明  
**注册**  
POST /api/douban/register   
form-data
username 必选且大于6位 string 用户名  
password 必选且大于六位 string 密码  
**登录**
POST /api/douban/login  
同注册     
在发送登录请求之后返回jwtToken，    
后续进行请求时，将token写入请求头的Authorization，进行认证。
Token的有效时间为7*24H。   
**修改密码**    
登录后
form-data
newpassword 要求同password string 新密码  
**电影**
form-data增加了Movie的模型     
Name string,
Director string,
Main_performer string,
Type string,
Country string,
Language string,
Date string,
Length string,
Stuff string,
Id int,
Writer string,
PicNum int
    
POST /api/douban/movie
url query电影id,获取movie的信息

**图片**
为影人和剧照插入图片的傻瓜式解决办法，     
每个影人、电影有自己的id，在数据库中存储该单位的图片数量，      
图片命名规则为id+00+图片编号（从1开始自动递增,编号为1的为封面）     
存储时按照此规则存储，调用图片时则按照对应名称依次调出（我觉得可以用for循环和picNum这两个数字来解决。）
        
图片存储在项目目录中。电影图片相对路径/moviePics   
影人图片路径/starPics

