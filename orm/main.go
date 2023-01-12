package main

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

type User struct {
	Id   int
	Name string `orm:"size(100)"`
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
func init() {
	// 注册驱动
	orm.RegisterDriver("postgres", orm.DRPostgres)
	// 设置默认数据库
	connStr := "postgres://postgres:123456@116.204.108.126:3306/test?sslmode=disable"
	orm.RegisterDataBase("default", "postgres", connStr, 30)
	// 注册定义的model
	orm.RegisterModel(new(User))
	//  创建table
	orm.RunSyncdb("default", false, true)
}
func main() {
	o := orm.NewOrm()
	user := User{Name: "amy"}
	// 插入表
	id, err := o.Insert(&user)
	checkErr(err)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)
	// 修改表
	user.Name = "bob"
	num, err := o.Update(&user)
	checkErr(err)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)
	// 读取
	u := User{Id: user.Id}
	err = o.Read(&u)
	fmt.Printf("ERR: %v\n", err)
	// 删除表
	num, err = o.Delete(&u)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)
}
