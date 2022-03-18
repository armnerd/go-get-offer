package tree

// 二叉树的下一个结点
func Question8() string {
	var pre = []int{1, 2, 3, 4, 5, 6, 7}
	var vin = []int{3, 2, 4, 1, 6, 5, 7}
	root := reConstructBinaryTree(pre, vin)
	LowestCommonAncestor(root, 3, 4)
	return "done!"
}

func GetNext(pNode *TreeLinkNode) *TreeLinkNode {
	if pNode == nil {
		return pNode
	}

	if pNode.Right != nil {
		pNode = pNode.Right
		for pNode.Left != nil {
			pNode = pNode.Left
		}
		return pNode
	}

	for pNode.Next != nil {
		if pNode.Next.Left == pNode {
			return pNode.Next
		}
		// 向上回溯
		pNode = pNode.Next
	}

	return nil
}
