package leetcode

import (
	"go-get-offer/leetcode/leetcode56"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

type Problem56DataSet struct {
	Intervals    [][]int
	ExpectResult [][]int
}

func TestProblem56(t *testing.T) {
	dataSet := []Problem56DataSet{
		{
			[][]int{
				{1, 3},
				{2, 6},
				{8, 10},
				{15, 18},
			},
			[][]int{
				{1, 6},
				{8, 10},
				{15, 18},
			},
		},
		{
			[][]int{
				{1, 4},
				{4, 5},
			},
			[][]int{
				{1, 5},
			},
		},
	}
	convey.Convey("合并区间", t, func() {
		for _, v := range dataSet {
			res := leetcode56.Merge(v.Intervals)
			convey.So(res, convey.ShouldResemble, v.ExpectResult)
		}
	})
}
