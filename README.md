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
##### 3.生成项目到当前目录下:  【这个命令不会改变 logic 和 handler 下面的文件】
$ C:\Users\A\go\src\go-zero-second\user\api> goctl api go  --api  .\user.api --dir .  
##### 4.根据 api文件(user.api文件)生成 swagger 文档到当前目录下:  
$ C:\Users\A\go\src\go-zero-second\user\api> goctl api swagger --api ./user.api  --dir .  


### 框架整理
1.响应如何封装?  
2.统一api前级  
3.用户信息接口应该要进行jwt 验证  
4.api文档  

## go-zero官方项目‌

GitHub地址：https://github.com/zeromicro/go-zero
星标数：27k+
简介：go-zero核心框架，包含完整的微服务开发工具链和最佳实践24
zero-admin电商系统‌

GitHub地址：https://github.com/zjtomoon-cc/zero-admin
星标数：1.2k+（Fork自feihua/zero-admin）
简介：基于go-zero实现的电商后台系统，包含前后端分离架构和Docker部署方案7
awesome-zero生态项目集‌

GitCode地址：https://gitcode.com/gh_mirrors/awesome-zero
简介：整理了基于go-zero的优质项目集合，涵盖中间件集成和业务场景案例2
go-zero实战案例库‌

GitHub趋势项目（参考2025年6月数据）：
部分新兴项目如tensorzero（Rust/Go混合架构）和HumanSystemOptimization（健康类微服务）近期获得较高关注910
建议优先考察官方项目和awesome-zero生态集，其中zero-admin电商系统具备完整的业务实现参考