package main 
import (
	"fmt"
	"test_mod/mathClass"
	"test_mod/Format"
)

func main(){
	sum:=mathClass.Add(1,2)
	sub:=mathClass.Sub(3,2)
	fmt.Println(sum)
	fmt.Println(sub)
}
func init(){
	fmt.Println("init me first!")
	Format.FormatVariable()
	Format.GetArea()
	Format.GetIOTA()
}