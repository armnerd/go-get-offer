package tree

// 二叉树中和为某一值的路径(三)
func Question84() string {
	var pre = []int{1, 2, 3, 4, 5, 6, 7}
	var vin = []int{3, 2, 4, 1, 6, 5, 7}
	root := reConstructBinaryTree(pre, vin)
	findPath(root, 3)
	return "done!"
}

var count = 0

func findPath(root *TreeNode, sum int) int {
	if root == nil {
		return count
	}
	Dfs(root, sum)
	findPath(root.Left, sum)
	findPath(root.Right, sum)
	return count
}

func Dfs(root *TreeNode, sum int) {
	if root == nil {
		return
	}
	sum -= root.Val
	if sum == 0 {
		count++
	}
	Dfs(root.Left, sum)
	Dfs(root.Right, sum)
}
