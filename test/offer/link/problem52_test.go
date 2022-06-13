package link

import (
	"go-get-offer/offer/link"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

type Problem52DataSet struct {
	List1  []int
	List2  []int
	Common []int
}

func TestProblem52(t *testing.T) {
	dataSet := []Problem52DataSet{
		{
			[]int{1, 2, 3},
			[]int{4, 5},
			[]int{6, 7},
		},
	}
	convey.Convey("两个链表的第一个公共结点", t, func() {
		for _, v := range dataSet {
			common := link.Ints2List(v.Common)
			list1 := link.Ints2List(v.List1)
			temp1 := list1
			for {
				if temp1.Next == nil {
					temp1.Next = common
					break
				}
				temp1 = temp1.Next
			}
			list2 := link.Ints2List(v.List2)
			temp2 := list2
			for {
				if temp2.Next == nil {
					temp2.Next = common
					break
				}
				temp2 = temp2.Next
			}
			res := link.FindFirstCommonNode(list1, list2)
			convey.So(link.List2Ints(res), convey.ShouldResemble, v.Common)
		}
	})
}
