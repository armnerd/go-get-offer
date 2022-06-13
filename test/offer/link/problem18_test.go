package link

import (
	"go-get-offer/offer/link"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

type Problem18DataSet struct {
	List         []int
	Pos          int
	ExpectResult []int
}

func TestProblem18(t *testing.T) {
	dataSet := []Problem18DataSet{
		{
			[]int{2, 5, 1, 9},
			5,
			[]int{2, 1, 9},
		},
		{
			[]int{2, 5, 1, 9},
			1,
			[]int{2, 5, 9},
		},
	}
	convey.Convey("删除链表的节点", t, func() {
		for _, v := range dataSet {
			list := link.Ints2List(v.List)
			res := link.DeleteNode(list, v.Pos)
			convey.So(link.List2Ints(res), convey.ShouldResemble, v.ExpectResult)
		}
	})
}
