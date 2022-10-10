package main

import (
	"fmt"
	"test_mod/functionList"
)

func main() {
	nextNumber := functionList.GetSequence()
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	nextNumber1 := functionList.GetSequence()
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())
}
func init() {
	fmt.Println("init me first!")

}
