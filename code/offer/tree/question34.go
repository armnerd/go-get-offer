package tree

// 二叉树中和为某一值的路径(二)
func Question34() string {
	var pre = []int{1, 2, 3, 4, 5, 6, 7}
	var vin = []int{3, 2, 4, 1, 6, 5, 7}
	root := reConstructBinaryTree(pre, vin)
	FindPath(root, 4)
	return "done!"
}

var paths [][]int
var path []int

func FindPath(root *TreeNode, expectNumber int) [][]int {
	dfs(root, expectNumber)
	return paths
}

func dfs(root *TreeNode, expectNumber int) {
	if root == nil {
		return
	}

	path = append(path, root.Val) // 先将 root 节点放入路径集合【入栈】
	expectNumber -= root.Val      // 同时期望值减去 root 节点的值
	if expectNumber == 0 && root.Left == nil && root.Right == nil {
		// 防止 Go 骚操作
		tmp := make([]int, len(path))
		copy(tmp, path)
		paths = append(paths, tmp)
	}

	// 向左向右搜索
	dfs(root.Left, expectNumber)
	dfs(root.Right, expectNumber)

	// 回溯当然不能当前节点算作内【出栈】
	path = path[:len(path)-1]
}
