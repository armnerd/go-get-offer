package link

// 链表中倒数最后k个结点
func Question22() string {
	list := Ints2List([]int{1, 2, 3, 4, 5})
	FindKthToTail(list, 2)
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
