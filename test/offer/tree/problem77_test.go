package tree

import (
	"go-get-offer/offer/tree"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

type Problem77DataSet struct {
	Root         []int
	ExpectResult [][]int
}

func TestProblem77(t *testing.T) {
	dataSet := []Problem77DataSet{
		{
			[]int{1, 2, 3, NULL, NULL, 4, 5},
			[][]int{
				{1},
				{3, 2},
				{4, 5},
			},
		},
		{
			[]int{8, 6, 10, 5, 7, 9, 11},
			[][]int{
				{8},
				{10, 6},
				{5, 7, 9, 11},
			},
		},
		{
			[]int{1, 2, 3, 4, 5},
			[][]int{
				{1},
				{3, 2},
				{4, 5},
			},
		},
	}
	convey.Convey("按之字形顺序打印二叉树", t, func() {
		for _, v := range dataSet {
			root := tree.Ints2TreeNode(v.Root)
			res := tree.SnakePrint(root)
			convey.So(res, convey.ShouldResemble, v.ExpectResult)
		}
	})
}
