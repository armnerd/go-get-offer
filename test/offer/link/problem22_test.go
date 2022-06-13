package link

import (
	"go-get-offer/offer/link"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

type Problem22DataSet struct {
	List         []int
	Pos          int
	ExpectResult []int
}

func TestProblem22(t *testing.T) {
	dataSet := []Problem22DataSet{
		{
			[]int{1, 2, 3, 4, 5},
			2,
			[]int{4, 5},
		},
	}
	convey.Convey("链表中倒数最后k个结点", t, func() {
		for _, v := range dataSet {
			list := link.Ints2List(v.List)
			res := link.FindKthToTail(list, v.Pos)
			convey.So(link.List2Ints(res), convey.ShouldResemble, v.ExpectResult)
		}
	})
}
