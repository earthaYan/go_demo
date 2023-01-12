package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
func useMysql() {
	// 打开一个注册过的数据库驱动,第二个参数为data source name
	db, err := sql.Open("mysql", "root:123456@tcp(116.204.108.126:3306)/test?charset=utf8")
	checkErr(err)
	// 插入数据
	// 返回准备要执行的sql操作，然后返回准备完毕的执行状态
	stmt, err := db.Prepare("INSERT INTO userinfo SET username=?,department=?,created=?")
	checkErr(err)
	// 执行stmt准备好的SQL语句
	// 传入的参数都是=?对应的数据
	res, err := stmt.Exec("amy", "研发部门", "2012-12-09")
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)
	// 修改数据
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)
	res, err = stmt.Exec("astaxieupdate", id)
	checkErr(err)
	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)
	// 查询数据
	// 直接执行Sql返回Rows结果
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)
	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}
	// 删除数据
	stmt, err = db.Prepare("delete from userinfo where uid=?")
	checkErr(err)
	res, err = stmt.Exec(id)
	checkErr(err)
	affect, err = res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)
	db.Close() //关闭数据库连接
}
func usePostgreSQL() {
	connStr := "postgres://postgres:123456@116.204.108.126:3306/test?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	checkErr(err)
	// 插入数据
	stmt, err := db.Prepare("INSERT INTO userinfo(username,department,created) VALUES($1,$2,$3) RETURNING uid")
	checkErr(err)
	res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
	checkErr(err)
	var lastInsertId int
	err = db.QueryRow("INSERT INTO userinfo(username,department,created) VALUES($1,$2,$3) returning uid;", "astaxie", "研发部门", "2012-12-09").Scan(&lastInsertId)
	checkErr(err)
	fmt.Println("最后插入id =", lastInsertId, res)
	// 修改数据
	stmt, err = db.Prepare("update userinfo set username=$1 where uid=$2")
	checkErr(err)
	res, err = stmt.Exec("astaxieupdate", lastInsertId)
	checkErr(err)
	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)

	// 查看数据
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)
	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}
	// 删除数据
	stmt, err = db.Prepare("delete from userinfo where uid=$1")
	checkErr(err)
	res, err = stmt.Exec(lastInsertId)
	checkErr(err)
	affect, err = res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)
	db.Close()
}
func main() {
	usePostgreSQL()
}
