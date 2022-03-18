package question17

import (
	"math"
)

// 打印从1到最大的n位数
func Run() string {
	printNumbers(1)
	return "done!"
}

func printNumbers(n int) []int {
	limit := math.Pow(10, float64(n))
	res := []int{}
	for i := 1; i < int(limit); i++ {
		res = append(res, i)
	}
	return res
}
