package leetcode

import (
	"go-get-offer/leetcode/leetcode46"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

type Problem46DataSet struct {
	Nums         []int
	ExpectResult [][]int
}

func TestProblem46(t *testing.T) {
	dataSet := []Problem46DataSet{
		{
			[]int{1, 2, 3},
			[][]int{
				{1, 2, 3},
				{1, 3, 2},
				{2, 1, 3},
				{2, 3, 1},
				{3, 1, 2},
				{3, 2, 1},
			},
		},
		{
			[]int{0, 1},
			[][]int{
				{0, 1},
				{1, 0},
			},
		},
		{
			[]int{1},
			[][]int{
				{1},
			},
		},
	}
	convey.Convey("全排列", t, func() {
		for _, v := range dataSet {
			tree := leetcode46.Permute(v.Nums)
			convey.So(tree, convey.ShouldResemble, v.ExpectResult)
		}
	})
}
