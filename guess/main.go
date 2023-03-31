// 猜数字挑战
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)       //播种随机数生成器
	target := rand.Intn(101) //生成1-100之间的随机整数
	fmt.Println("目标为0-100之间的任意一个整数", target)
	fmt.Println("你能猜到么？")
	reader := bufio.NewReader(os.Stdin) //创建一个bufio.Reader，有读取键盘输入的权限
	fmt.Print("请给出你的猜测数字：")
	var success = false //设置为默认打印失败
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Print("你还有", 10-guesses, "次机会,")
		fmt.Print("请选择：")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		if guess < target {
			fmt.Println("你的数字小于目标数")
		} else if guess > target {
			fmt.Println("你的数字大于目标数")
		} else {
			success = true
			fmt.Println("猜中了！！！")
			break
		}
	}
	if !success {
		fmt.Println("你已失败，正确的数字是", target)
	}
}
