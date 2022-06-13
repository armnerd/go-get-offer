package question65

import "fmt"

// 不用加减乘除做加法
func Run() string {
	res := Add(1, 2)
	fmt.Println(res)
	return "done!"
}

func Add(num1 int, num2 int) int {
	/*
		num2 初始如果等于 0，那么结果就是 num1
		进入循环后，就是进位的值，进位的值为 0 了，加法结束了
	*/
	for num2 != 0 {
		// 相加各位的值，不算进位，就相当于各位做异或操作
		tmp := num1 ^ num2
		// 计算进位值，相当于各位做与操作后，再向左移一位得到
		num2 = (num1 & num2) << 1 // num2 被当做放进位的容器
		num1 = tmp                // num1 被当做放相加各位的值的容器
	}
	return num1
}
