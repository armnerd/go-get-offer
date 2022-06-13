package link

// 删除链表中重复的结点
func Question76() string {
	list := Ints2List([]int{1, 2, 3, 3, 4, 4, 5})
	DeleteDuplication(list)
	return "done!"
}

func DeleteDuplication(pHead *ListNode) *ListNode {
	var vhead *ListNode = &ListNode{
		Val:  -1,
		Next: pHead,
	}
	var pre *ListNode = vhead
	var cur *ListNode = pHead
	for {
		if cur == nil {
			break
		}
		if cur.Next != nil && cur.Val == cur.Next.Val {
			cur = cur.Next
			for {
				if cur.Next == nil || cur.Val != cur.Next.Val {
					break
				}
				cur = cur.Next
			}
			cur = cur.Next
			pre.Next = cur
		} else {
			pre = cur
			cur = cur.Next
		}

	}
	return vhead.Next
}
