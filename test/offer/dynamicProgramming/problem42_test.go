package dynamicprogramming

import (
	"go-get-offer/offer/dynamicProgramming/question42"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

type Problem42DataSet struct {
	Array        []int
	ExpectResult int
}

func TestProblem42(t *testing.T) {
	dataSet := []Problem42DataSet{
		{
			[]int{1, -2, 3, 10, -4, 7, 2, -5},
			18,
		},
		{
			[]int{2},
			2,
		},
		{
			[]int{-10},
			-10,
		},
	}
	convey.Convey("连续子数组的最大和", t, func() {
		for _, v := range dataSet {
			res := question42.FindGreatestSumOfSubArray(v.Array)
			convey.So(res, convey.ShouldEqual, v.ExpectResult)
		}
	})
}
