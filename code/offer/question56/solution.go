package question56

// 数组中只出现一次的两个数字
func Run() string {
	array := []int{1, 4, 1, 6}
	FindNumsAppearOnce(array)
	return "done!"
}

func FindNumsAppearOnce(data []int) []int {
	if nil == data || len(data) < 2 {
		return []int{}
	}
	// 先将全部数进行异或运算，得出最终结果
	var result int
	for _, v := range data {
		result ^= v
	}
	// 找到那个可以充当分组去进行与运算的数
	// 从最低位开始找起
	mask := 1
	for result&mask == 0 {
		mask <<= 1
	}
	// 进行分组，分成两组，转换为两组 求出现一次的数字 去求
	var num1, num2 int
	for _, v := range data {
		if v&mask == 0 {
			num1 ^= v
		} else {
			num2 ^= v
		}
	}
	// 输出时按非降序排列
	if num1 > num2 {
		num1, num2 = num2, num1
	}
	return []int{num1, num2}
}
