package question70

import "fmt"

// 矩形覆盖
func Run() string {
	number := 3
	res := rectCover(number)
	fmt.Println(res)
	return "done!"
}

func rectCover(number int) int {
	if number == 0 || number == 1 || number == 2 {
		return number
	}
	first := 1
	second := 2
	res := 0
	for i := 3; i <= number; i++ {
		res = first + second
		first = second
		second = res
	}
	return res
}
