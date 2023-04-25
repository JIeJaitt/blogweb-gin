package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

func InitMysql() {
	fmt.Println("InitMysql...")
	if db == nil {
		db, _ = sql.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/blogweb_gin?charset=utf8")
		err := db.Ping()
		if err != nil {
			log.Fatalf("ping database failed: %v", err)
		}
		CreateTableWithUser()
		CreateTableWithArticle()
		CreateTableWithAlbum()
	}
}

func ModifyDB(sql string, args ...interface{}) (int64, error) {
	result, err := db.Exec(sql, args...)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return count, nil
}

func QueryRowDB(sql string) *sql.Row {
	return db.QueryRow(sql)
}

// CreateTableWithAlbum --------图片--------
func CreateTableWithAlbum() {
	sql := `create table if not exists album(
		id int(4) primary key auto_increment not null,
		filepath varchar(255),
		filename varchar(64),
		status int(4),
		createtime int(10)
		);`
	_, err := ModifyDB(sql)
	if err != nil {
		return
	}
}

func CreateTableWithArticle() {
	sql := `create table if not exists article(
		id int(4) primary key auto_increment not null,
		title varchar(30),
		author varchar(20),
		tags varchar(30),
		short varchar(255),
		content longtext,
		createtime int(10)
		);`
	_, err := ModifyDB(sql)
	if err != nil {
		return
	}
}

func CreateTableWithUser() {
	sql := `CREATE TABLE IF NOT EXISTS users(
    id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
    username VARCHAR(64),
    password VARCHAR(64),
    status INT(4),
    createtime INT(10)
    );`
	_, err := ModifyDB(sql)
	if err != nil {
		return
	}
}
