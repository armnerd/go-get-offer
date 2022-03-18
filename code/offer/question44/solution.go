package question44

import (
	"fmt"
	"strconv"
)

// 数字序列中某一位的数字
func Run() string {
	res := findNthDigit(10)
	fmt.Println(res)
	return "done!"
}

func findNthDigit(n int) int {
	start, digit, count := 1, 1, 9
	for n > count {
		n -= count
		start *= 10
		digit += 1
		count = start * digit * 9
	}
	num := start + (n-1)/digit
	res, _ := strconv.Atoi(string(strconv.Itoa(num)[(n-1)%digit]))
	return res
}
