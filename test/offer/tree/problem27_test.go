package tree

import (
	"go-get-offer/offer/tree"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

type Problem27DataSet struct {
	Root         []int
	ExpectResult []int
}

func TestProblem27(t *testing.T) {
	dataSet := []Problem27DataSet{
		{
			[]int{8, 6, 10, 5, 7, 9, 11},
			[]int{8, 10, 6, 11, 9, 7, 5},
		},
	}
	convey.Convey("二叉树的镜像", t, func() {
		for _, v := range dataSet {
			root := tree.Ints2TreeNode(v.Root)
			res := tree.Mirror(root)
			convey.So(tree.Tree2ints(res), convey.ShouldResemble, v.ExpectResult)
		}
	})
}
