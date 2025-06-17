# go-zero-second

### 项目目录文件介绍
#### api
etc下相当于具体的项目配置数据  config下相当于将配置引入项目  
handler 下相当于有controller和路由配置  
logic下就是逻辑文件  
svc下是上下文文件  
type就是type的放置位置
#### rpc
etc下相当于具体的项目配置数据  config下相当于将配置引入项目  
server 下相当于有controller和路由配置  
client 下相当于客户端controller  
logic下就是逻辑文件  
svc下是上下文文件  
type就是type的放置位置

### 生成api项目：    
##### 1.先编写user.api文件:               
##### 2.验证文件是否正确:       
$ C:\Users\A\go\src\go-zero-second\user\api> goctl api validate --api .\user.api        
##### 3.生成项目到当前目录下:  
$ C:\Users\A\go\src\go-zero-second\user\api> goctl api go  --api  .\user.api --dir .  
##### 4.根据 api文件(user.api文件)生成 swagger 文档到当前目录下:  
$ C:\Users\A\go\src\go-zero-second\user\api> goctl api swagger --api ./user.api  --dir .  
