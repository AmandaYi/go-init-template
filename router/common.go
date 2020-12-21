package router

import (
	"github.com/gin-gonic/gin"
	 "yy-price-api/router/priceSource"
)

func LoadRoutes(r *gin.Engine) {
	// UI报价中心，PC端 APP端。
	loadPriceRoutes(r)
}

// UI报价中心，PC端 APP端。
func loadPriceRoutes(r *gin.Engine){
	version := "v1" // 当前的api的版本
	// 数据中心
	priceSourceRouter := r.Group("/priceSource/"+version + "/price/source")
	priceSourceRouter.GET("/user/:id", priceSource.User.Get)
}

