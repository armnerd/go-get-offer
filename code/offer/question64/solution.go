package question64

import "fmt"

// 求1+2+3+...+n
func Run() string {
	res := Sum_Solution(5)
	fmt.Println(res)
	return "done!"
}

func Sum_Solution(n int) int {
	if n == 1 {
		return n
	}
	n += Sum_Solution(n - 1)
	return n
}
