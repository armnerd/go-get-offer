package queueandstack

import (
	"go-get-offer/offer/queueAndStack/question31"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

type Problem31DataSet struct {
	PushV        []int
	PopV         []int
	ExpectResult bool
}

func TestProblem31(t *testing.T) {
	dataSet := []Problem31DataSet{
		{
			[]int{1, 2, 3, 4, 5},
			[]int{4, 5, 3, 2, 1},
			true,
		},
		{
			[]int{1, 2, 3, 4, 5},
			[]int{4, 3, 5, 1, 2},
			false,
		},
	}
	convey.Convey("栈的压入、弹出序列", t, func() {
		for _, v := range dataSet {
			res := question31.IsPopOrder(v.PushV, v.PopV)
			convey.So(res, convey.ShouldEqual, v.ExpectResult)
		}
	})
}
