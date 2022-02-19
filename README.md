# winter
（前端伙伴不写了所以我的文档很不讲究=.=）
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
Main_performer string，      
此处主演上传两个人，上传格式为演员id+";"+演员id
Type string,
Country string,
Language string,
Date string,
Length string,
Stuff string,
Id int,
Writer string,
PicNum int  

    进行传参时，id之间以分号隔开。传参时间格式为0000-00-00（年月日)

POST /api/douban/movie/:id
url :id,获取movie的信息（PS.又新增了查看电影评分功能）      
    POST/api/douban/movie/:id/comment    上传评论
GET/api/douban/movie/:id/comment 查看评论

**上传个人介绍**  
POST /api/douban/user
  通过token得到请求的用户名。
SelfInfo string 不多于一百字。     
**查看用户个人页**
通过url传参，将要访问的用户主页的所有者的:username填入url   
GET /api/douban/user/:username        
"username":user.Username,       
"name":user.Name,       
"selfInfo":user.SelfInfo,
**查看影人**
POST /api/douban/star/:id
影人id为五位，从10001开始    **影人页面尚未完善。**
**增加了管理员权限，用户名为admin时有权限进行添加电影**    
POST /api/douban/movie/admin "写不完了，具体格式看看代码吧。。。"    
**按照类型查找电影**类似按照电影id查找影评，对数据库进行多行查找     
**增加了电影名搜索**        
GET /api/douban/movie/search/:name返回电影名和id
**图片**
为影人和剧照插入图片的傻瓜式解决办法，     
每个影人、电影有自己的id，在数据库中存储该单位的图片数量，      
图片命名规则为id+00+图片编号（从1开始自动递增,编号为1的为封面）     
存储时按照此规则存储，调用图片时则按照对应名称依次调出（我觉得可以用for循环和picNum这两个数字来解决。）
        
图片存储在项目目录中。电影图片相对路径/moviePics   
影人图片路径/starPics

