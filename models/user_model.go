package models

import (
	"blogweb-gin/database"
	"fmt"
)

type User struct {
	Id         int
	Username   string
	Password   string
	Status     int // 0 正常状态， 1删除
	Createtime int64
}

func InsertUser(user User) (int64, error) {
	return database.ModifyDB("insert into users(username,password,status,createtime) values (?,?,?,?)",
		user.Username, user.Password, user.Status, user.Createtime)
}

// QueryUserWightCon 按条件查询
func QueryUserWightCon(con string) int {
	sql := fmt.Sprintf("select id from users %s", con)
	fmt.Println(sql)
	row := database.QueryRowDB(sql)
	id := 0
	err := row.Scan(&id)
	if err != nil {
		return 0
	}
	return id
}

// QueryWithUsername 根据用户名查询id
func QueryWithUsername(username string) int {
	sql := fmt.Sprintf("where username='%s'", username)
	return QueryUserWightCon(sql)
}

// QueryUserWithParam 根据用户名和密码，查询id
// 主要用于登录的时候从数据库中查询用户名和密码的id，然后利用 id 是否大于 0 来判断用户是否注册
func QueryUserWithParam(username, password string) int {
	sql := fmt.Sprintf("where username='%s' and password='%s'", username, password)
	return QueryUserWightCon(sql)
}
