package link

import (
	"go-get-offer/offer/link"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

type Problem76DataSet struct {
	List         []int
	ExpectResult []int
}

func TestProblem76(t *testing.T) {
	dataSet := []Problem76DataSet{
		{
			[]int{1, 2, 3, 3, 4, 4, 5},
			[]int{1, 2, 5},
		},
		{
			[]int{1, 1, 1, 8},
			[]int{8},
		},
	}
	convey.Convey("删除链表中重复的结点", t, func() {
		for _, v := range dataSet {
			list := link.Ints2List(v.List)
			res := link.DeleteDuplication(list)
			convey.So(link.List2Ints(res), convey.ShouldResemble, v.ExpectResult)
		}
	})
}
