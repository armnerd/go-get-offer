package link

import (
	"go-get-offer/offer/link"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

type Problem24DataSet struct {
	List         []int
	ExpectResult []int
}

func TestProblem24(t *testing.T) {
	dataSet := []Problem24DataSet{
		{
			[]int{1, 2, 3},
			[]int{3, 2, 1},
		},
	}
	convey.Convey("反转链表", t, func() {
		for _, v := range dataSet {
			list := link.Ints2List(v.List)
			res := link.ReverseList(list)
			convey.So(link.List2Ints(res), convey.ShouldResemble, v.ExpectResult)
		}
	})
}
