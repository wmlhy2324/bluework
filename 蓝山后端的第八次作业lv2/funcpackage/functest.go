package funcpackage

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gomodule/redigo/redis"
)

type users struct {
	Username string `db:"username"`
	Pwd      string `db:"pwd"`
}

var pool *redis.Pool
var db *sql.DB

// 连接到redis
func initpool() {
	pool = &redis.Pool{
		MaxIdle:     8,
		MaxActive:   0,
		IdleTimeout: 300,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		}}
}

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

// 将用户的信息存入redis,并设置过期时间
func Cacheuser(username, pwd string) bool {
	initpool()
	conn := pool.Get()
	defer conn.Close()
	_, err := conn.Do("hset", "users", username, pwd)
	if err != nil {
		fmt.Println("设置数据失败,错误是", err)
		return false
	}
	_, err = conn.Do("expire", "users", 60)
	if err != nil {
		fmt.Println("过期时间设置失败,错误是", err)
		return false
	}
	return true
}

// 查询用户在redis中是否存在
func Queryredis(username string) bool {
	initpool()
	conn := pool.Get()
	value, err := conn.Do("hget", "users", username)

	if err != nil {
		return false
	}
	if value != nil {
		return true
	}
	//不存在返回false
	return false
}

// 通过用户获取密码信息,从redis里面获取
func Userinfo(username string) bool {
	initpool()
	conn := pool.Get()

	pwd, err := redis.String(conn.Do("hget", "users", username))
	if err != nil {
		fmt.Print("获取有误错误是=", err)
		return false
	}
	if pwd == "" {
		fmt.Println("缓存中没有这个用户")
		return false
	}
	fmt.Println("成功从redis中获取用户")
	return true
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

// 查询得到密码
func Chectpwd(username string) string {
	initDB()
	defer db.Close()
	sqlStr := "SELECT pwd from users where username=?;"
	rowObj := db.QueryRow(sqlStr, username)
	var u1 users
	rowObj.Scan(&u1.Pwd)
	return u1.Pwd
}

//给文章点赞
func Likearticle(username string) bool {
	flag := Query(username)
	if flag == false {
		fmt.Println("用户不存在")
		return flag
	}

	initpool()
	conn := pool.Get()
	_, err := conn.Do("SADD", "article", username)
	if err != nil {
		fmt.Println("添加错误,错误是=", err)
		return false
	}
	count, err := conn.Do("scard", "article")
	if err != nil {
		fmt.Println("查询点赞个数失败")

	}
	fmt.Println("点赞数为", count)
	return true
}
