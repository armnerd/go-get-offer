package tree

import (
	"go-get-offer/offer/tree"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

type Problem68DataSet struct {
	Root         []int
	P            int
	Q            int
	ExpectResult int
}

func TestProblem68(t *testing.T) {
	dataSet := []Problem68DataSet{
		{
			[]int{7, 1, 12, 0, 4, 11, 14, NULL, NULL, 3, 5},
			1,
			12,
			7,
		},
		{
			[]int{7, 1, 12, 0, 4, 11, 14, NULL, NULL, 3, 5},
			12,
			11,
			12,
		},
	}
	convey.Convey("二叉搜索树的最近公共祖先", t, func() {
		for _, v := range dataSet {
			root := tree.Ints2TreeNode(v.Root)
			res := tree.LowestCommonAncestor68(root, v.P, v.Q)
			convey.So(res, convey.ShouldEqual, v.ExpectResult)
		}
	})
}
