package question66

// 构建乘积数组
func Run() string {
	A := []int{1, 2, 3, 4, 5}
	Multiply(A)
	multiply(A)
	return "done!"
}

func Multiply(A []int) []int {
	res := []int{}
	for k := range A {
		one := 1
		for i := 0; i < len(A); i++ {
			if i != k {
				one *= A[i]
			}
		}
		res = append(res, one)
	}
	return res
}

func multiply(A []int) []int {
	res := make([]int, len(A))
	// 先把左边的乘积搞定
	res[0] = 1 // 初始化首位的左边
	for i := 1; i < len(A); i++ {
		res[i] = res[i-1] * A[i-1]
	}
	right := 1 // 最后一位的右边
	// 从倒数第二个开始
	for i := len(A) - 2; i >= 0; i-- {
		right *= A[i+1]
		res[i] *= right // 左右相乘
	}
	return res
}
