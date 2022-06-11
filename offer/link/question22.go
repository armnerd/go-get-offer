package link

// 链表中倒数最后k个结点
func Question22() string {
	var pHead *ListNode
	for _, v := range [5]int{5, 4, 3, 2, 1} {
		node := ListNode{
			Val:  v,
			Next: pHead,
		}
		pHead = &node
	}
	FindKthToTail(pHead, 2)
	return "done!"
}

func FindKthToTail(pHead *ListNode, k int) *ListNode {
	var fast *ListNode = pHead
	var slow *ListNode = pHead
	for i := 0; i < k; i++ {
		if fast == nil {
			var empty *ListNode
			return empty
		}
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}
