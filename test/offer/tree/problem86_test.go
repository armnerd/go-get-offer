package tree

import (
	"go-get-offer/offer/tree"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

type Problem86DataSet struct {
	Root         []int
	O1           int
	O2           int
	ExpectResult int
}

func TestProblem86(t *testing.T) {
	dataSet := []Problem86DataSet{
		{
			[]int{3, 5, 1, 6, 2, 0, 8, NULL, NULL, 7, 4},
			5,
			1,
			3,
		},
		{
			[]int{3, 5, 1, 6, 2, 0, 8, NULL, NULL, 7, 4},
			2,
			7,
			2,
		},
	}
	convey.Convey("在二叉树中找到两个节点的最近公共祖先", t, func() {
		for _, v := range dataSet {
			root := tree.Ints2TreeNode(v.Root)
			res := tree.LowestCommonAncestor(root, v.O1, v.O2)
			convey.So(res, convey.ShouldEqual, v.ExpectResult)
		}
	})
}
