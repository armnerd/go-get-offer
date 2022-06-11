package question10

import "fmt"

// 斐波那契数列
func Run() string {
	fmt.Println(Fibonacci(4))
	return "done!"
}

func Fibonacci(n int) int {
	f1 := 0
	f2 := 1
	if n == 0 {
		return f1
	} else if n == 1 {
		return f2
	}
	for i := 1; i < n; i++ {
		tmp := f2
		f2 = f1 + f2
		f1 = tmp
	}
	return f2
}
