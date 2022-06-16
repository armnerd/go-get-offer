package question69

import "fmt"

// 跳台阶
func Run() string {
	fmt.Println(jumpFloor1(5))
	fmt.Println(jumpFloor2(5))
	fmt.Println(jumpFloor3(5))
	fmt.Println(JumpFloor(5))
	return "done!"
}

// 纯递归
func jumpFloor1(number int) int {
	if number <= 1 {
		return 1
	}
	return jumpFloor1(number-1) + jumpFloor1(number-2)
}

// 备忘录
func jumpFloor2(number int) int {
	note := make([]int, number)
	return fib(number, note)
}

func fib(number int, note []int) int {
	if number <= 1 {
		return 1
	}

	if note[number-1] != 0 {
		return note[number-1]
	}

	res := fib(number-1, note) + fib(number-2, note)
	note[number-1] = res
	return res
}

// 备忘录升级为动态规划, 不用递归的过程, 直接从子树求得答案, 过程是从下往上
func jumpFloor3(number int) int {
	if number <= 1 {
		return 1
	}
	note := make([]int, number)
	note[0] = 1
	note[1] = 2
	for i := 2; i < number; i++ {
		note[i] = note[i-1] + note[i-2]
	}
	return note[number-1]
}

// 究极动态规划
func JumpFloor(number int) int {
	if number <= 1 {
		return 1
	}
	var a = 1
	var b = 1
	var c int
	for i := 2; i <= number; i++ {
		c = a + b
		a = b
		b = c
	}
	return c
}
