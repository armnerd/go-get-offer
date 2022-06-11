package tree

// 二叉搜索树的后序遍历序列
func Question33() string {
	var sequence = []int{1, 3, 2}
	VerifySquenceOfBST(sequence)
	return "done!"
}

func VerifySquenceOfBST(sequence []int) bool {
	if len(sequence) == 0 {
		return false
	}
	return check(sequence, 0, len(sequence)-1)
}

func check(sequence []int, start int, end int) bool {
	if start >= end {
		return true
	}
	// 找到根
	var root = sequence[end]
	var border = end - 1
	// 从右向左扫描，比根大就继续前移，直到找到界限
	for border >= 0 && sequence[border] > root {
		border--
	}
	// 如果符合条件的话边界往前的应该都小于根才对
	for i := start; i <= border; i++ {
		if sequence[i] > root {
			return false
		}
	}
	// 分治检查左右
	return check(sequence, start, border) && check(sequence, border+1, end-1)
}
