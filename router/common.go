package router

import "github.com/gin-gonic/gin"

func LoadRoutes(r *gin.Engine) {
	// UI报价中心，PC端 APP端。
	loadApiRoutes(r)
}

func loadApiRoutes(r *gin.Engine){
	version := "v1" // 当前的api的版本

	//列表中心
	api := r.Group("/api/"+version)
	api.GET("/list",  )



	//客户中心


}