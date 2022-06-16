package question71

import "fmt"

// 跳台阶扩展问题
func Run() string {
	res := JumpFloorII(3)
	fmt.Println(res)
	return "done!"
}

func JumpFloorII(number int) int {
	if number <= 1 {
		return 1
	}
	var a = 1
	var b = 0
	for i := 2; i <= number; i++ {
		// b = a * 2
		b = a << 1 // 跟上边是一回事儿
		a = b
	}
	return b
}
