package tree

import (
	"go-get-offer/offer/tree"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

type Problem26DataSet struct {
	Root1        []int
	Root2        []int
	ExpectResult bool
}

func TestProblem26(t *testing.T) {
	dataSet := []Problem26DataSet{
		{
			[]int{8, 8, 7, 9, 2, NULL, NULL, NULL, NULL, 4, 7},
			[]int{8, 9, 2},
			true,
		},
		{
			[]int{1, 2, 3, 4, 5},
			[]int{2, 4},
			true,
		},
		{
			[]int{1, 2, 3},
			[]int{3, 1},
			false,
		},
	}
	convey.Convey("树的子结构", t, func() {
		for _, v := range dataSet {
			root1 := tree.Ints2TreeNode(v.Root1)
			root2 := tree.Ints2TreeNode(v.Root2)
			res := tree.HasSubtree(root1, root2)
			convey.So(res, convey.ShouldEqual, v.ExpectResult)
		}
	})
}
