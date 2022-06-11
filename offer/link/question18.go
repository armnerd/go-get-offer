package link

// 删除链表的节点
func Question18() string {
	var head *ListNode
	for _, v := range [4]int{9, 1, 5, 2} {
		node := ListNode{
			Val:  v,
			Next: head,
		}
		head = &node
	}
	deleteNode(head, 5)
	return "done!"
}

func deleteNode(head *ListNode, val int) *ListNode {
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
