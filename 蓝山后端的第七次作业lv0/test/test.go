package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //导入包但不使用，init()
)

type emp1 struct {
	ID           int
	WORK_NO      string
	NAME         string
	GENDER       string
	AGE          string
	ID_CARD      string
	WORK_ADDRESS string
	ENTRY_DATE   string
}

var db *sql.DB

func query(n int) {
	sqlStr := "SELECT ID,WORK_NO,NAME,GENDER,AGE,ID_CARD,WORK_ADDRESS,ENTRY_DATE FROM emp1 WHERE ID>?;"
	rows, err := db.Query(sqlStr, n)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u1 emp1
		rows.Scan(&u1.ID, &u1.WORK_NO, &u1.NAME, &u1.GENDER, &u1.AGE, &u1.ID_CARD, &u1.WORK_ADDRESS, &u1.ENTRY_DATE)
		fmt.Printf("u1:%#v\n", u1)
	}
}

//Go连接Mysql示例
func main() {
	var err error
	//数据库
	//用户名:密码啊@tcp(ip:端口)/数据库的名字
	dsn := "root:112304@tcp(127.0.0.1:3306)/lihaoyu"
	//连接数据集
	db, err = sql.Open("mysql", dsn) //open不会检验用户名和密码
	if err != nil {
		fmt.Printf("dsn:%s invalid,err:%v\n", dsn, err)
		return
	}
	err = db.Ping() //尝试连接数据库
	if err != nil {
		fmt.Printf("open %s faild,err:%v\n", dsn, err)
		return
	}
	db.SetMaxIdleConns(10)
	fmt.Println("连接数据库成功~")
	query(3)
}
