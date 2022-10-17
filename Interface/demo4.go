package Interface

import "fmt"

type Phone2 interface {
	call() string
}
type Android2 struct {
	brand string
}
type IPhone2 struct {
	version string
}

func (android Android2) call() string {
	return "I am Android " + android.brand
}

func (iPhone IPhone2) call() string {
	return "I am iPhone " + iPhone.version
}
func printCall(p Phone2) {
	fmt.Println(p.call() + ", I can call you!")
}
func TestInterfaceDemo2() {
	var vivo = Android2{brand: "Vivo"}
	var hw = Android2{"HuaWei"}
	i7 := IPhone2{version: "7 Plus"}
	ix := IPhone2{"X"}
	printCall(vivo)
	printCall(hw)
	printCall(i7)
	printCall(ix)
}
