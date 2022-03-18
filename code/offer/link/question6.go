package link

// 从尾到头打印链表
func Question6() {
	var raw = []int{67, 0, 24, 58}
	var head, next *ListNode
	for i := len(raw) - 1; i >= 0; i-- {
		node := &ListNode{
			Val: raw[i],
		}
		if next != nil {
			node.Next = next
		}
		next = node
		if i == 0 {
			head = node
		}
	}
	printListFromTailToHead(head)
}

func printListFromTailToHead(head *ListNode) []int {
	var tmp = make([]int, 0)
	// 参数判断
	if head == nil {
		return tmp
	}

	// 压入切片
	for {
		tmp = append(tmp, head.Val)
		if head = head.Next; head == nil {
			break
		}
	}
	var res = make([]int, 0, len(tmp))

	// 反转数组
	for i := len(tmp) - 1; i >= 0; i-- {
		res = append(res, tmp[i])
	}
	return res
}
