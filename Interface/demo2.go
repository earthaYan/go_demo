package Interface

import "fmt"

type Human interface {
	name() string
	age() int
}
type Woman struct {
}

type Man struct {
}

func (woman Woman) name() string {
	return "Jin Yawei"
}
func (woman Woman) age() int {
	return 23
}
func (man Man) name() string {
	return "liweibin"
}
func (man Man) age() int {
	return 27
}

func TestSexAndAge() {
	var human Human
	human = new(Woman)
	fmt.Println(human.name())
	fmt.Println(human.age())
	human = new(Man)
	fmt.Println(human.name())
	fmt.Println(human.age())
}
