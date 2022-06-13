package link

import (
	"go-get-offer/offer/link"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

type Problem25DataSet struct {
	List1        []int
	List2        []int
	ExpectResult []int
}

func TestProblem25(t *testing.T) {
	dataSet := []Problem25DataSet{
		{
			[]int{1, 3, 5},
			[]int{2, 4, 6},
			[]int{1, 2, 3, 4, 5, 6},
		},
		{
			[]int{-1, 2, 4},
			[]int{1, 3, 4},
			[]int{-1, 1, 2, 3, 4, 4},
		},
	}
	convey.Convey("合并两个排序的链表", t, func() {
		for _, v := range dataSet {
			list1 := link.Ints2List(v.List1)
			list2 := link.Ints2List(v.List2)
			res := link.Merge(list1, list2)
			convey.So(link.List2Ints(res), convey.ShouldResemble, v.ExpectResult)
		}
	})
}
