package leettree

import (
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 还原树
func BuildTreeFromArray(row []string) *TreeNode {
	data := [][]string{}
	level := 1
	count := 1
	temp := row
	// 分层整理
	for len(temp) > 0 {
		data = append(data, temp[:count])
		temp = temp[count:]
		level += 1
		count *= 2
	}
	// 创建节点
	nodes := [][]*TreeNode{}
	for i := 0; i < len(data); i++ {
		line := data[i]
		nodeLine := make([]*TreeNode, 0)
		for j := 0; j < len(line); j++ {
			var node *TreeNode
			if line[j] != "null" {
				val, _ := strconv.Atoi(line[j])
				node = &TreeNode{
					Val: val,
				}
			}
			nodeLine = append(nodeLine, node)
		}
		nodes = append(nodes, nodeLine)
	}
	// 从上到下连接节点
	for i := 0; i < len(nodes)-1; i++ {
		line := nodes[i]
		below := nodes[i+1]
		for j := 0; j < len(line); j++ {
			if line[j] != nil {
				line[j].Left = below[j*2]
				line[j].Right = below[j*2+1]
			}
		}
	}
	// 返回根节点
	return nodes[0][0]
}

// 序列化树
func SerializeATree(pRoot *TreeNode) (result []string) {
	if pRoot == nil {
		return
	}
	nodeToBeMachined := []*TreeNode{pRoot} // 待处理的一层节点
	bottom := false
	for !bottom {
		nextLevelNodes := make([]*TreeNode, 0)
		bottom = true
		for _, item := range nodeToBeMachined {
			if item == nil {
				result = append(result, "null")
			} else {
				result = append(result, strconv.Itoa(item.Val))
				// 只要有不是 nil 的节点就说明没到底
				bottom = false
				if item.Left != nil {
					nextLevelNodes = append(nextLevelNodes, item.Left)
				} else if item.Right != nil {
					nextLevelNodes = append(nextLevelNodes, nil)
				}
				// 只有一个子节点时才会有 null
				if item.Right != nil {
					nextLevelNodes = append(nextLevelNodes, item.Right)
				} else if item.Left != nil {
					nextLevelNodes = append(nextLevelNodes, nil)
				}
			}
		}
		nodeToBeMachined = nextLevelNodes
	}
	return
}
