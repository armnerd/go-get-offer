package link

// 删除链表的节点
func Question18() string {
	link := Ints2List([]int{2, 5, 1, 9})
	DeleteNode(link, 5)
	return "done!"
}

func DeleteNode(head *ListNode, val int) *ListNode {
	var cur *ListNode = head
	var pre *ListNode = head
	for {
		if cur == nil {
			break
		}
		if cur.Val == val {
			if cur == head {
				head = head.Next
			}
			pre.Next = cur.Next
			break
		}
		pre = cur
		cur = cur.Next
	}
	return head
}
