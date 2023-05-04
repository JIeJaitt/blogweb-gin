package controllers

import (
	"blogweb-gin/models"
	"blogweb-gin/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginGet(c *gin.Context) {
	// 返回登录页面 HTML:login.html
	c.HTML(http.StatusOK, "login.html", gin.H{"title": "登录页"})
}

// LoginPost 处理登录的 POST 请求
func LoginPost(c *gin.Context) {
	// 获取表单信息
	username := c.PostForm("username")
	password := c.PostForm("password")
	fmt.Println("username:", username, ",password:", password)

	// 根据用户名和密码去数据库查询
	id := models.QueryUserWithParam(username, utils.MD5(password))
	fmt.Println("id:", id)
	if id > 0 {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "登录成功"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "登录失败"})
	}
}
