package link

// 合并两个排序的链表
func Question25() string {
	var one *ListNode
	for _, v := range [3]int{5, 3, 1} {
		node := ListNode{
			Val:  v,
			Next: one,
		}
		one = &node
	}
	var two *ListNode
	for _, v := range [3]int{6, 4, 2} {
		node := ListNode{
			Val:  v,
			Next: two,
		}
		two = &node
	}
	Merge(one, two)
	return "done!"
}

func Merge(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	result := ListNode{
		Val:  -1,
		Next: nil,
	}
	cur := &result
	for {
		if pHead1 == nil || pHead2 == nil {
			break
		}
		if pHead1.Val >= pHead2.Val {
			cur.Next = pHead2
			pHead2 = pHead2.Next
		} else {
			cur.Next = pHead1
			pHead1 = pHead1.Next
		}
		cur = cur.Next
	}
	if pHead1 != nil {
		cur.Next = pHead1
	} else if pHead2 != nil {
		cur.Next = pHead2
	}
	return result.Next
}
