package question4

import "fmt"

// 二维数组中的查找
func Run() string {
	var target = 7
	array := [][]int{{1, 2, 8, 9}, {2, 4, 9, 12}, {4, 7, 10, 13}, {6, 8, 11, 15}}
	fmt.Println(Find(target, array))
	return "done!"
}

func Find(target int, array [][]int) bool {
	// 参数检查
	var m, n int
	if m = len(array); m == 0 {
		return false
	}
	if n = len(array[0]); n == 0 {
		return false
	}

	// 行与列
	r := 0
	c := n - 1
	for {
		// 循环结束条件
		if r >= m || c < 0 {
			break
		}

		// 二分查找
		if target == array[r][c] {
			return true
		} else if target > array[r][c] {
			r++
		} else {
			c--
		}
	}
	return false
}
