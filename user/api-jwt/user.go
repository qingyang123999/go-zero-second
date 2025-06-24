package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"

	"api-jwt/internal/config"
	"api-jwt/internal/handler"
	"api-jwt/internal/svc"

	"api-jwt/common/response"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	// JwtUnauthorizedResult jwt验证失败的回调 用户信息鉴权失败
	server := rest.MustNewServer(c.RestConf, rest.WithUnauthorizedCallback(JwtUnauthorizedResult))
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

// JwtUnauthorizedResult jwt验证失败的回调 用户信息鉴权失败
func JwtUnauthorizedResult(w http.ResponseWriter, r *http.Request, err error) {
	fmt.Println(err) // 具体的错误，没带token，token过期？伪造token？
	if err != nil {
		httpx.WriteJson(w, http.StatusOK, response.Body{Code: 10087, Msg: "鉴权失败，" + err.Error()})
	} else {
		httpx.WriteJson(w, http.StatusOK, response.Body{Code: 10087, Msg: "鉴权失败"})
	}

}
