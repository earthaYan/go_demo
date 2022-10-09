package Format
import (
	"fmt"
)
func FormatString(){
	// %d 表示整型数字，%s 表示字符串
	var stock_code int=123
	const end_date string="2020-12-31"
	var url string="Code=%d&endDate=%s"
	var target_url=fmt.Sprintf(url,stock_code,end_date)
	fmt.Println(target_url)
}
func FormatVariable(){
	var i ,m int
	var f float64
	var b bool
	var s string
	var name1, name2, name3 int
name1, name2, name3 = 1, 2,'a'
	fmt.Printf("零值:%d %v %v %v %q\n",i,m,f,b,s)
	fmt.Printf("赋值:%d %d %v \n",name1, name2, name3)
	fmt.Printf("地址：%d\n",&i)
}
func GetArea(){
	const LENGTH int=10
	const WIDTH int=5
	var area int
	const a,b,c=1,false,"str"

	area=LENGTH*WIDTH
	fmt.Printf("面积为 : %d", area)
	println()
	println(a, b, c)  
}
func GetIOTA(){
	const (
		a = iota   //0
		b          //1
		c          //2
		d = "ha"   //独立值，iota += 1
		e          //"ha"   iota += 1
		f = 100    //iota +=1
		g          //100  iota +=1
		h = iota   //7,恢复计数
		i          //8
	)	
	fmt.Printf("ITOA:%v %v %v %q %q %v %v %v %v\n",a,b,c,d,e,f,g,h,i)
}