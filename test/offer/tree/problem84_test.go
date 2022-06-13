package tree

import (
	"go-get-offer/offer/tree"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

type Problem84DataSet struct {
	Root         []int
	Sum          int
	ExpectResult int
}

func TestProblem84(t *testing.T) {
	dataSet := []Problem84DataSet{
		{
			[]int{1, 2, 3, 4, 5, 4, 3, NULL, NULL, -1},
			6,
			3,
		},
		{
			[]int{0, 1},
			1,
			2,
		},
		{
			[]int{1, NULL, 2, NULL, 3},
			3,
			2,
		},
	}
	convey.Convey("二叉树中和为某一值的路径(三)", t, func() {
		for _, v := range dataSet {
			tree.Count = 0
			root := tree.Ints2TreeNode(v.Root)
			res := tree.FindPath3(root, v.Sum)
			convey.So(res, convey.ShouldEqual, v.ExpectResult)
		}
	})
}
