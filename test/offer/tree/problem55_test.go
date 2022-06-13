package tree

import (
	"go-get-offer/offer/tree"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

type Problem55DataSet struct {
	Root         []int
	ExpectResult int
}

func TestProblem55(t *testing.T) {
	dataSet := []Problem55DataSet{
		{
			[]int{1, 2, 3, 4, 5, NULL, 6, NULL, NULL, 7},
			4,
		},
	}
	convey.Convey("二叉树的深度", t, func() {
		for _, v := range dataSet {
			root := tree.Ints2TreeNode(v.Root)
			res := tree.TreeDepth(root)
			convey.So(res, convey.ShouldEqual, v.ExpectResult)
		}
	})
}
