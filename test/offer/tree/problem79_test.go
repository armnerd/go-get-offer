package offer

import (
	"go-get-offer/offer/tree"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

type Problem79DataSet struct {
	Root         []int
	ExpectResult bool
}

func TestProblem79(t *testing.T) {
	dataSet := []Problem79DataSet{
		{
			[]int{1, 2, 3, 4, 5, 6, 7},
			true,
		},
	}
	convey.Convey("判断是不是平衡二叉树", t, func() {
		for _, v := range dataSet {
			root := tree.Ints2TreeNode(v.Root)
			res := tree.IsBalanced_Solution(root)
			convey.So(res, convey.ShouldEqual, v.ExpectResult)
		}
	})
}
