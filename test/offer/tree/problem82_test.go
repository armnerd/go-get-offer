package tree

import (
	"go-get-offer/offer/tree"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

type Problem82DataSet struct {
	Root         []int
	Sum          int
	ExpectResult bool
}

func TestProblem82(t *testing.T) {
	dataSet := []Problem82DataSet{
		{
			[]int{5, 4, 8, 1, 11, NULL, 9, NULL, NULL, 2, 7},
			22,
			true,
		},
		{
			[]int{1, 2},
			0,
			false,
		},
		{
			[]int{1, 2},
			3,
			true,
		},
	}
	convey.Convey("二叉树中和为某一值的路径(一)", t, func() {
		for _, v := range dataSet {
			root := tree.Ints2TreeNode(v.Root)
			res := tree.HasPathSum(root, v.Sum)
			convey.So(res, convey.ShouldEqual, v.ExpectResult)
		}
	})
}
