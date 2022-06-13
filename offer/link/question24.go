package link

// 反转链表
func Question24() string {
	list := Ints2List([]int{1, 2, 3})
	ReverseListRaw(list)
	ReverseList(list)
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
