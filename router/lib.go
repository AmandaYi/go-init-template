package router

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//处理跨域问题

func useCors(r *gin.Engine) gin.HandlerFunc{
	//初始化一个跨域配置
	config := cors.DefaultConfig()
	config.AllowCredentials = true
	config.AllowMethods = []string{"GET", "POST", "DELETE", "HEAD"}
	config.AllowOriginFunc = func(origin string) bool {
		fmt.Println(origin)
		//config.AllowOrigins = []string{origin,"http://127.0.0.1"}
		//switch origin {
		//case "daishu888.com":
		//}
		////先允许全部的域名通过
		//
		return true
	}
	return cors.New(config)

}
//加载中间件
func LoadMiddleware(r *gin.Engine) {
	r.Use(gin.Logger(), gin.Recovery(), useCors(r)) // 日志 // 容错 // 加载跨域
}
