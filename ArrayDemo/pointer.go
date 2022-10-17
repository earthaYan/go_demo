package complex

import "fmt"

func GetAddress() {
	var a int = 20 /* 声明实际变量 */
	var ip *int    /* 声明指针变量 */

	ip = &a /* 指针变量的存储地址 */

	fmt.Printf("a 变量的地址是: %x\n", &a)

	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip)

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip)
}

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func TestStruct() {
	var Book1 Books
	Book1 = Books{title: "Go 语言", author: "www.runoob.com", subject: "Go 语言教程", book_id: 6495407}
	fmt.Printf("Book 1 title : %s\n", Book1.title)
}

func TestRange() {
	var slice1 []int = make([]int, 5, 10)
	slice2 := []int{1, 2, 3}
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	slice3 := pow[:len(pow)-1]
	fmt.Print(slice1, slice2, slice3, cap(slice3), len(slice3))
	for key := range slice3 {
		fmt.Printf("key is: %d\n", key)
	}
}
func TestMap() {
	var countryCapitalMap map[string]string /*创建集合 */
	countryCapitalMap = make(map[string]string)
	countryCapitalMap["France"] = "巴黎"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India"] = "新德里"
	countryCapitalMap["American"] = "华盛顿"
	testDelete(countryCapitalMap, "American")
	delete(countryCapitalMap, "American")
	testDelete(countryCapitalMap, "American")

}
func testDelete(countryCapitalMap map[string]string, capitalName string) {
	capital, ok := countryCapitalMap[capitalName]
	if ok {
		fmt.Println(capitalName, "的首都是", capital)
	} else {
		fmt.Println(capitalName, "的首都不存在")
	}
}
