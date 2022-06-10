package leetcode

import (
	"code/leetcode/leetcode54"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

type Problem54DataSet struct {
	Matrix       [][]int
	ExpectResult []int
}

func TestProblem54(t *testing.T) {
	dataSet := []Problem54DataSet{
		{
			[][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			[]int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			[][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
			},
			[]int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
	}
	convey.Convey("螺旋矩阵", t, func() {
		for _, v := range dataSet {
			res := leetcode54.SpiralOrder(v.Matrix)
			convey.So(res, convey.ShouldResemble, v.ExpectResult)
		}
	})
}
