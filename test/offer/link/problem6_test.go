package link

import (
	"go-get-offer/offer/link"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

type Problem6DataSet struct {
	List         []int
	ExpectResult []int
}

func TestProblem7(t *testing.T) {
	dataSet := []Problem6DataSet{
		{
			[]int{1, 2, 3},
			[]int{3, 2, 1},
		},
		{
			[]int{67, 0, 24, 58},
			[]int{58, 24, 0, 67},
		},
	}
	convey.Convey("从尾到头打印链表", t, func() {
		for _, v := range dataSet {
			list := link.Ints2List(v.List)
			res := link.PrintListFromTailToHead(list)
			convey.So(res, convey.ShouldResemble, v.ExpectResult)
		}
	})
}
