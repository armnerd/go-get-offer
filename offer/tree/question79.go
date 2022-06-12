package tree

// 判断是不是平衡二叉树
func Question79() string {
	tree := []int{1, 2, 3, 4, 5, 6, 7}
	IsBalanced_Solution(Ints2TreeNode(tree))
	return "done!"
}

func IsBalanced_Solution(pRoot *TreeNode) bool {
	return postOrder(pRoot) != -1
}

func postOrder(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 如果不满足平衡二叉树的定义，则返回-1，并且如果左子树不满足条件了，直接返回-1，右子树也是如此，相当于剪枝，加速结束递归
	left := postOrder(root.Left)
	if left == -1 {
		return -1
	}
	right := postOrder(root.Right)
	if right == -1 {
		return -1
	}
	// 计算高差
	sub := 0
	max := 0
	if left > right {
		sub = left - right
		max = left
	} else {
		sub = right - left
		max = right
	}
	if sub > 1 {
		return -1
	}
	// 较长子树高度加上自己的高度
	return max + 1
}
