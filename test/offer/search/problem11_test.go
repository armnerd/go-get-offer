package search

import (
	"go-get-offer/offer/search/question11"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

type Problem11DataSet struct {
	RotateArray  []int
	ExpectResult int
}

func TestProblem11(t *testing.T) {
	dataSet := []Problem11DataSet{
		{
			[]int{1, 2, 8, 9},
			1,
		},
		{
			[]int{3, 100, 200, 3},
			3,
		},
	}
	convey.Convey("旋转数组的最小数字", t, func() {
		for _, v := range dataSet {
			res := question11.MinNumberInRotateArray(v.RotateArray)
			convey.So(res, convey.ShouldEqual, v.ExpectResult)
		}
	})
}
