package tree

import (
	"go-get-offer/offer/tree"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

type Problem54DataSet struct {
	Root         []int
	K            int
	ExpectResult int
}

func TestProblem54(t *testing.T) {
	dataSet := []Problem54DataSet{
		{
			[]int{5, 3, 7, 2, 4, 6, 8},
			3,
			4,
		},
	}
	convey.Convey("二叉搜索树的第k个结点", t, func() {
		for _, v := range dataSet {
			root := tree.Ints2TreeNode(v.Root)
			res := tree.KthNode(root, v.K)
			convey.So(res, convey.ShouldEqual, v.ExpectResult)
		}
	})
}
