package main

import (
	"fmt"
	complex "test_mod/ArrayDemo"
	"test_mod/mathClass"
)

func main() {
	result := mathClass.JieChen(15)
	fmt.Printf("%d 的阶乘是 %d\n", 15, result)
	var i int
	for i = 0; i < 10; i++ {
		fmt.Printf("%d\t", mathClass.Fibonacci(i))
	}
}
func init() {
	fmt.Println("init me first!")
	complex.TestMap()
}
