# winter
key-要求-数据类型-说明  
**注册**  
POST /api/douban/register   
username 必选且大于6位 string 用户名  
password 必选且大于六位 string 密码  
**登录**    
POST /api/douban/login  
同注册     
在发送登录请求之后返回jwtToken，    
后续进行请求时，将token写入请求头的Authorization，进行认证。
Token的有效时间为7*24H。
**退出登录**
将Authorization内容删除。