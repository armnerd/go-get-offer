package leettree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func reConstructBinaryTree(pre []int, vin []int) *TreeNode {
	if len(pre) == 0 {
		return nil
	}
	// 找到根节点
	root_val := pre[0]
	root_node := &TreeNode{
		Val: root_val,
	}

	// 切开中序
	var index int
	for index = range vin {
		if vin[index] == root_val {
			break
		}
	}

	// 大树拆小树递归处理
	root_node.Left = reConstructBinaryTree(pre[1:1+index], vin[:index])
	root_node.Right = reConstructBinaryTree(pre[1+index:], vin[index+1:])

	return root_node
}
