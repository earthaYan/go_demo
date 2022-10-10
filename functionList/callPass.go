package functionList

func SwapByReferPass(x *int,y *int){
	var temp int
	temp = *x
	*x=*y
	*y=temp
}