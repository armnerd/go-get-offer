package leetcode

import (
	"go-get-offer/leetcode/leettree"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

type Problem111DataSet struct {
	Root         []int
	ExpectResult int
}

func TestProblem111(t *testing.T) {
	dataSet := []Problem111DataSet{
		{
			[]int{3, 9, 20, NULL, NULL, 15, 7},
			2,
		},
		{
			[]int{2, NULL, 3, NULL, 4, NULL, 5, NULL, 6},
			5,
		},
	}
	convey.Convey("二叉树的最小深度", t, func() {
		for _, v := range dataSet {
			root := leettree.Ints2TreeNode(v.Root)
			res := leettree.MinDepth(root)
			convey.So(res, convey.ShouldEqual, v.ExpectResult)
		}
	})
}
