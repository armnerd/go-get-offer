package question43

import "fmt"

// 整数中1出现的次数（从1到n整数中1出现的次数）
func Run() string {
	res := NumberOf1Between1AndN_Solution(13)
	fmt.Println(res)
	return "done!"
}

func NumberOf1Between1AndN_Solution(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		num := i
		for num > 0 {
			if num == 1 || num%10 == 1 {
				res++
			}
			num /= 10
		}
	}
	return res
}
