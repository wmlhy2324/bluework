package funcpackage

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type users struct {
	Username string `db:"username"`
	Pwd      string `db:"pwd"`
}

var db *sql.DB

func initDB() (err error) {

	dns := "root:112304@tcp(127.0.0.1:3306)/lihaoyu"
	db, err = sql.Open("mysql", dns)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("数据库连接成功")
	db.SetMaxIdleConns(2)
	return
}

// 添加用户
func Addusers(username, pwd string) {
	initDB()
	defer db.Close()
	sqlStr := "INSERT into users(username,pwd) VALUES (?,?);"
	_, err := db.Exec(sqlStr, username, pwd)
	if err != nil {
		return
	}

}

// 查询用户是否存在
func Query(username string) bool {
	initDB()
	defer db.Close()
	sqlStr := "SELECT * from users WHERE username=?;"
	rowobj := db.QueryRow(sqlStr, username)
	var u1 users
	err := rowobj.Scan(&u1.Username, &u1.Pwd)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("未找到对应的行")
			return false
		} else {
			fmt.Println("查询出错:", err.Error())
			return false
		}
	}
	fmt.Printf("%#v\n", u1.Username)
	if u1.Username == "" {
		return false
	}
	return true
}

// 检查密码是否正确
func Checkpwd(username string) string {
	initDB()
	defer db.Close()
	sqlStr := "SELECT pwd from users WHERE username=?;"
	rowobj := db.QueryRow(sqlStr, username)
	var u1 users
	err := rowobj.Scan(&u1.Pwd)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("未找到对应的行")
			return "" // 返回空字符串或其他适当的值，表示未找到密码
		} else {
			fmt.Println("查询出错:", err.Error())
			return "" // 返回空字符串或其他适当的值，表示查询出错
		}
	}
	fmt.Printf("%#v\n", u1.Pwd)
	return u1.Pwd
}

// 修改密码
func Updatepwd(username, originpwd, changepwd string) bool {
	//优先连接数据库
	initDB()
	defer db.Close()
	//查询得到原来的密码
	sqlStr := "SELECT pwd from users where username=?;"
	rowObj := db.QueryRow(sqlStr, username)
	var u1 users
	rowObj.Scan(&u1.Pwd)
	if originpwd != u1.Pwd {
		fmt.Println("密码错误")
		return false
	}
	sqlStr1 := "UPDATE users set pwd=? where username=?;"
	_, err := db.Exec(sqlStr1, changepwd, username)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
