package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// passfile是用来报告测试成绩是否及格
	fmt.Print("请输入成绩：")
	reader := bufio.NewReader(os.Stdin)
	std, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	grade, err := strconv.ParseFloat(strings.TrimSpace(std), 64)
	if err != nil {
		log.Fatal(err)
	}
	var status string
	if grade == 100 {
		status = "满分"
	} else if grade >= 60 {
		status = "及格"
	} else {
		status = "不及格"
	}
	fmt.Println("分数为", grade, ",",status)

}
