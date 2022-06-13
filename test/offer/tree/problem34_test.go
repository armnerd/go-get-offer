package tree

import (
	"go-get-offer/offer/tree"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

type Problem34DataSet struct {
	Root         []int
	ExpectNumber int
	ExpectResult [][]int
}

func TestProblem34(t *testing.T) {
	dataSet := []Problem34DataSet{
		{
			[]int{10, 5, 12, 4, 7},
			22,
			[][]int{
				{10, 5, 7},
				{10, 12},
			},
		},
		{
			[]int{10, 5, 12, 4, 7},
			15,
			[][]int{},
		},
		{
			[]int{2, 3},
			0,
			[][]int{},
		},
		{
			[]int{1, 3, 4},
			7,
			[][]int{},
		},
	}
	convey.Convey("二叉树中和为某一值的路径(二)", t, func() {
		for _, v := range dataSet {
			root := tree.Ints2TreeNode(v.Root)
			res := tree.FindPath(root, v.ExpectNumber)
			convey.So(res, convey.ShouldResemble, v.ExpectResult)
		}
	})
}
