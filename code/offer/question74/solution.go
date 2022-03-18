package question74

// 和为S的连续正数序列
func Run() string {
	FindContinuousSequence(9)
	return "done!"
}

func FindContinuousSequence(sum int) [][]int {
	res := [][]int{}
	l, r := 1, 1
	tmp := 0
	for l <= sum/2 {
		if tmp < sum {
			tmp += r
			r++
		} else if tmp > sum {
			tmp -= l
			l++
		} else {
			ans := []int{}
			for k := l; k < r; k++ {
				ans = append(ans, k)
			}
			res = append(res, ans)
			tmp -= l
			l++
		}
	}
	return res
}
