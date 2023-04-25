package main

import (
	"blogweb-gin/database"
	"blogweb-gin/routers"
)

func main() {
	database.InitMysql()
	router := routers.InitRouter()
	//静态资源
	router.Static("/static", "./static")
	err := router.Run(":8080")
	if err != nil {
		return
	}
}
