package link

// 反转链表
func Question24() string {
	var next *ListNode
	for _, v := range [3]int{3, 2, 1} {
		node := ListNode{
			Val:  v,
			Next: next,
		}
		next = &node
	}
	ReverseListRaw(next)
	ReverseList(next)
	return "done!"
}

func ReverseList(pHead *ListNode) *ListNode {
	var pre *ListNode
	var cur *ListNode = pHead
	var nex *ListNode
	for {
		if cur == nil {
			break
		}
		nex = cur.Next
		cur.Next = pre
		pre = cur
		cur = nex
	}
	return pre
}

func ReverseListRaw(pHead *ListNode) *ListNode {
	var reverse *ListNode
	for {
		if pHead == nil {
			break
		}
		node := ListNode{
			Val:  pHead.Val,
			Next: reverse,
		}
		reverse = &node
		pHead = pHead.Next
	}
	return reverse
}
