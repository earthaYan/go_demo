package main 
import (
	"fmt"
	"test_mod/mathClass"
)

func main(){
	fmt.Println("Hello, World!")
	sum:=mathClass.Add(1,2)
	sub:=mathClass.Sub(3,2)
	fmt.Println(sum)
	fmt.Println(sub)
}
func init(){
	fmt.Println("init me first!")
}