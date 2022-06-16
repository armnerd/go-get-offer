package dynamicprogramming

import (
	"go-get-offer/offer/dynamicProgramming/question85"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

type Problem85DataSet struct {
	Array        []int
	ExpectResult []int
}

func TestProblem85(t *testing.T) {
	dataSet := []Problem85DataSet{
		{
			[]int{1, -2, 3, 10, -4, 7, 2, -5},
			[]int{3, 10, -4, 7, 2},
		},
		{
			[]int{1},
			[]int{1},
		},
		{
			[]int{1, 2, -3, 4, -1, 1, -3, 2},
			[]int{1, 2, -3, 4, -1, 1},
		},
		{
			[]int{-2, -1},
			[]int{-1},
		},
	}
	convey.Convey("连续子数组的最大和(二)", t, func() {
		for _, v := range dataSet {
			res := question85.FindGreatestSumOfSubArray(v.Array)
			convey.So(res, convey.ShouldResemble, v.ExpectResult)
		}
	})
}
