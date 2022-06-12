package offer

import (
	"go-get-offer/offer/tree"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

type Problem32DataSet struct {
	Root         []int
	ExpectResult []int
}

func TestProblem32(t *testing.T) {
	dataSet := []Problem32DataSet{
		{
			[]int{8, 6, 10, NULL, NULL, 2, 1},
			[]int{8, 6, 10, 2, 1},
		},
		{
			[]int{5, 4, NULL, 3, NULL, 2, NULL, 1},
			[]int{5, 4, 3, 2, 1},
		},
	}
	convey.Convey("从上往下打印二叉树", t, func() {
		for _, v := range dataSet {
			root := tree.Ints2TreeNode(v.Root)
			res := tree.PrintFromTopToBottom(root)
			convey.So(res, convey.ShouldResemble, v.ExpectResult)
		}
	})
}
