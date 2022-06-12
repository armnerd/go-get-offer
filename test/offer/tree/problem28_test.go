package offer

import (
	"go-get-offer/offer/tree"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

type Problem28DataSet struct {
	Root         []int
	ExpectResult bool
}

func TestProblem28(t *testing.T) {
	dataSet := []Problem28DataSet{
		{
			[]int{1, 2, 2, 3, 4, 4, 3},
			true,
		},
		{
			[]int{8, 6, 9, 5, 7, 7, 5},
			false,
		},
	}
	convey.Convey("对称的二叉树", t, func() {
		for _, v := range dataSet {
			root := tree.Ints2TreeNode(v.Root)
			res := tree.IsSymmetrical(root)
			convey.So(res, convey.ShouldEqual, v.ExpectResult)
		}
	})
}
