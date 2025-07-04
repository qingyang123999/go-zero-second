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
svc下是上下文文件  依赖注入。比如mysql gorm  redis 等其他服务
type就是type的放置位置

### 生成api项目：    
##### 1.先编写user.api文件:               
##### 2.验证文件是否正确:       
$ C:\Users\A\go\src\go-zero-second\user\api> goctl api validate --api .\user.api        
##### 3.生成项目到当前目录下:  【这个命令不会改变 logic 和 handler 下面的已经生成的文件，需要删除再重新生成。但是会改变types下面的文件】
$ C:\Users\A\go\src\go-zero-second\user\api> goctl api go  --api  .\user.api --dir .  
##### 4.根据 api文件(user.api文件)生成 swagger 文档到当前目录下:  
$ C:\Users\A\go\src\go-zero-second\user\api> goctl api swagger --api ./user.api  --dir .  

#### 使用goctl template init 指令   
goctl template init 用于初始化模板    
然后会在 Templates are generated in C:\Users\A\.goctl\1.8.3, edit on your risk!  生成项目的一系列的模版   
比如要修改生成的 handler目录下的模板 就在这个目录下找到handler.tpl这个模板进行修改。   
如： 		
package {{.PkgName}}

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"api/common/response"
	{{.ImportPackages}}
)

{{if .HasDoc}}{{.Doc}}{{end}}
func {{.HandlerName}}(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		{{if .HasRequest}}var req types.{{.RequestType}}
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		{{end}}l := {{.LogicName}}.New{{.LogicType}}(r.Context(), svcCtx)
		{{if .HasResp}}resp, {{end}}err := l.{{.Call}}({{if .HasRequest}}&req{{end}})
		//if err != nil {
			//httpx.ErrorCtx(r.Context(), w, err)
		//} else {
			//{{if .HasResp}}httpx.OkJsonCtx(r.Context(), w, resp){{else}}httpx.Ok(w){{end}}
		//}
		{{if .HasResp}}response.Response(r, w, resp, err){{else}}response.Response(r, w, nil, err){{end}}  
	}
} 

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

looklook实战项目‌

GitHub：https://github.com/Mikaelemmmm/go-zero-looklook
星标数：4.8k+
特点：全技术栈微服务最佳实践，包含Kafka/Prometheus等集成1
‌惺惺社交平台‌

GitHub：https://github.com/cherish-chat/xxim-server
星标数：2.8k+
特点：开源IM系统，支持私有化部署1
simple-admin后台系统‌

GitHub：https://github.com/suyuan32/simple-admin-core
星标数：1.8k+
特点：RBAC权限管理系统，集成Ent框架1
zero-admin电商系统‌

GitHub：https://github.com/feihua/zero-admin
星标数：1.2k+
特点：前后端分离的电商后台，含支付和订单模块7
awesome-zero生态集‌

GitCode：https://gitcode.com/gh_mirrors/awesome-zero
特点：官方推荐的扩展组件和案例集合