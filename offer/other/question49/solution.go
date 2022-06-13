package question49

// 丑数
func Run() string {
	GetUglyNumber_Solution(7)
	return "done!"
}

func GetUglyNumber_Solution(index int) int {
	if 0 >= index {
		return 0
	}
	// 折三个变量[能成大事]
	var i2, i3, i5 int
	res := make([]int, index)
	res[0] = 1 // 第一个丑数为 1
	for i := 1; i < index; i++ {
		res[i] = min3(res[i2]*2, res[i3]*3, res[i5]*5)
		if res[i] == res[i2]*2 {
			i2++
		}
		if res[i] == res[i3]*3 {
			i3++
		}
		if res[i] == res[i5]*5 {
			i5++
		}
	}
	return res[index-1]
}

func min3(num1, num2, num3 int) int {
	number := min(num1, num2)
	return min(num3, number)
}

func min(num1, num2 int) int {
	if num1 < num2 {
		return num1
	}
	return num2
}
