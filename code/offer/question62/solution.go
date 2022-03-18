package question62

import "fmt"

// 孩子们的游戏(圆圈中最后剩下的数)
func Run() string {
	res := LastRemaining_Solution(5, 3)
	fmt.Println(res)
	return "done!"
}

func LastRemaining_Solution(n int, m int) int {
	if n < 1 || m < 1 {
		return -1
	}

	last := 0
	for i := 2; i <= n; i++ {
		last = (last + m) % i
	}
	return last
}
