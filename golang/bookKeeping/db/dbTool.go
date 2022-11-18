package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

const (
	userName = "root"
	password = "990518Chz"
	ip       = "bj-cdb-mw00usue.sql.tencentcdb.com"
	port     = "59576"
	dbName   = "testProj"
)

var DB *sql.DB

func InitDB() {
	//构建连接， "用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	DB, _ = sql.Open("mysql", path)
	DB.SetConnMaxLifetime(100)
	DB.SetMaxIdleConns(10)
	if err := DB.Ping(); err != nil {
		fmt.Println(">> open db fail")
		return
	}
	fmt.Println(">> db connect success")
}
