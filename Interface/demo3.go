package Interface

import "fmt"

type MobilePhone interface {
	call()
	call2()
}
type ApplePhone struct {
	id            int
	name          string
	category_id   int
	category_name string
}

func (test ApplePhone) call() {
	fmt.Println("这是第一个类的第一个接口回调函数 结构体数据：", ApplePhone{id: 0, name: "浅笑12", category_id: 4, category_name: "分类名称"})
	fmt.Println(test.name)
}
func (test ApplePhone) call2() {
	fmt.Println("这是一个类的第二个接口回调函数call2", ApplePhone{id: 1, name: "浅笑", category_id: 4, category_name: "分类名称"})
	fmt.Println(test.category_name)
}

type Android struct {
	member_id       int
	member_balance  float32
	member_sex      bool
	member_nickname string
}

func (test Android) call() {
	fmt.Println("这是第二个类的第一个接口回调函数call", Android{member_id: 22, member_balance: 15.23, member_sex: false, member_nickname: "浅笑18"})
}

func (test Android) call2() {
	fmt.Println("这是第二个类的第二个接口回调函数call2", Android{member_id: 44, member_balance: 100, member_sex: true, member_nickname: "陈超"})
}
func TestInterfaceDemo() {
	var phoneD MobilePhone
	phoneD = ApplePhone{id: 0, name: "浅笑2", category_id: 4, category_name: "分类名称"}
	phoneD.call()
	phoneD.call2()

	phoneD = new(Android)
	phoneD.call()
	phoneD.call2()
}
