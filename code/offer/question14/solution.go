package question14

import "fmt"

// 剪绳子
func Run() string {
	res := cutRope(8)
	fmt.Println(res)
	return "done!"
}

func cutRope(number int) int {
	if number == 2 {
		return 1
	} else if number == 3 {
		return 2
	}
	f := make([]int, number+1)
	for i := 1; i <= 4; i++ {
		f[i] = i
	}
	for i := 5; i <= number; i++ {
		for j := 1; j < i; j++ {
			f[i] = max(f[i], j*f[i-j])
		}
	}
	return f[number]
}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}
