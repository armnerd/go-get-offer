package link

import (
	"fmt"
	"time"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReviewAll(pHead *ListNode) {
	for {
		if pHead == nil {
			break
		}
		fmt.Println(pHead.Val)
		time.Sleep(time.Duration(1) * time.Second)
		pHead = pHead.Next
	}
}
