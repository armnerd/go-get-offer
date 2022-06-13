package link

// 合并两个排序的链表
func Question25() string {
	list1 := Ints2List([]int{1, 3, 5})
	list2 := Ints2List([]int{2, 4, 6})
	Merge(list1, list2)
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
