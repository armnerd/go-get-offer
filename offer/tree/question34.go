package tree

// 二叉树中和为某一值的路径(二)
func Question34() string {
	tree := []int{10, 5, 12, 4, 7}
	FindPath(Ints2TreeNode(tree), 22)
	return "done!"
}

func FindPath(root *TreeNode, expectNumber int) [][]int {
	var paths = make([][]int, 0)
	var path = make([]int, 0)
	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, expectNumber int) {
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
	dfs(root, expectNumber)
	return paths
}
