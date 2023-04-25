package routers

import (
	"blogweb-gin/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	router := gin.Default()
	router.LoadHTMLGlob("views/*")

	// 注册：
	router.GET("/register", controllers.RegisterGet)
	router.POST("/register", controllers.RegisterPost)

	// 登录
	router.GET("/login", controllers.LoginGet)
	router.POST("/login", controllers.LoginPost)

	return router

}
