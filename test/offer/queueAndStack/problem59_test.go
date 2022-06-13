package queueandstack

import (
	"go-get-offer/offer/queueAndStack/question59"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

type Problem59DataSet struct {
	Num          []int
	Size         int
	ExpectResult []int
}

func TestProblem59(t *testing.T) {
	dataSet := []Problem59DataSet{
		{
			[]int{2, 3, 4, 2, 6, 2, 5, 1},
			3,
			[]int{4, 4, 6, 6, 6, 5},
		},
		{
			[]int{9, 10, 9, -7, -3, 8, 2, -6},
			5,
			[]int{10, 10, 9, 8},
		},
		{
			[]int{1, 2, 3, 4},
			3,
			[]int{3, 4},
		},
	}
	convey.Convey("滑动窗口的最大值", t, func() {
		for _, v := range dataSet {
			res := question59.MaxInWindows(v.Num, v.Size)
			convey.So(res, convey.ShouldResemble, v.ExpectResult)
		}
	})
}
