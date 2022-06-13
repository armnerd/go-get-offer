package link

import (
	"go-get-offer/offer/link"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

type Problem23DataSet struct {
	List         []int
	Pos          int
	ExpectResult int
}

func TestProblem23(t *testing.T) {
	dataSet := []Problem23DataSet{
		{
			[]int{1, 2, 3, 4, 5},
			2,
			3,
		},
	}
	convey.Convey("链表中环的入口结点", t, func() {
		for _, v := range dataSet {
			list := link.Ints2ListWithCycle(v.List, v.Pos)
			res := link.EntryNodeOfLoop(list)
			convey.So(res.Val, convey.ShouldEqual, v.ExpectResult)
		}
	})
}
