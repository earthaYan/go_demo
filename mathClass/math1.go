package mathClass

func Add(x, y int) int {
	return x + y
}
func JieChen(num int) (result int) {
	if num > 0 {
		result = num * JieChen((num - 1))
		return result
	}
	return 1
}
func Fibonacci(num int) int {
	if num < 2 {
		return num
	}
	return Fibonacci(num-2) + Fibonacci(num-1)
}
