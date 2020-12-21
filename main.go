package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
	"yy-price-api/router"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	R := gin.New()
	router.LoadMiddleware(R) // 中间件
	router.LoadRoutes(R)
	listen(R)
}
func listen(R *gin.Engine) {
	PORT := os.Getenv("PORT")
	s := &http.Server{
		Addr:              ":" + PORT,
		Handler:           R,
		TLSConfig:         nil,
		ReadTimeout:       time.Second * 10,
		ReadHeaderTimeout: 0,
		WriteTimeout:      time.Second * 10,
		IdleTimeout:       0,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}
	err := s.ListenAndServe()
	if err != nil {
		fmt.Println("release端口出错了，请检查！" + PORT)
	}

}
