package tree

// 二叉树中和为某一值的路径(三)
func Question84() string {
	tree := []int{1, 2, 3, 4, 5, 4, 3, NULL, NULL, -1}
	FindPath3(Ints2TreeNode(tree), 6)
	return "done!"
}

var Count = 0

func FindPath3(root *TreeNode, sum int) int {
	if root == nil {
		return Count
	}
	Dfs(root, sum)
	FindPath3(root.Left, sum)
	FindPath3(root.Right, sum)
	return Count
}

func Dfs(root *TreeNode, sum int) {
	if root == nil {
		return
	}
	sum -= root.Val
	if sum == 0 {
		Count++
	}
	Dfs(root.Left, sum)
	Dfs(root.Right, sum)
}
