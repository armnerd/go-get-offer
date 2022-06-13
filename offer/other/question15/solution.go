package question15

import "fmt"

// 二进制中1的个数
func Run() string {
	for _, v := range []int{1, 10, -1} {
		res := NumberOf1(v)
		fmt.Printf("----------%v----------\n", res)
	}
	return "done!"
}

func NumberOf1(n int) int {
	cnt := 0
	// 对负数的处理
	if n < 0 {
		// 获取负数的补码
		n = n & 0xffffffff // 11111111111111111111111111111111
	}
	// fmt.Printf("%b\n", 0xffffffff) // 16进制: f == 1111
	for n != 0 {
		n = n & (n - 1)
		cnt++
	}
	return cnt
}
