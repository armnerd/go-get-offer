package leetcode

import (
	"code/leetcode/leettree"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

type Problem654DataSet struct {
	Nums         []int
	ExpectResult []int
}

func TestProblem654(t *testing.T) {
	dataSet := []Problem654DataSet{
		{
			[]int{3, 2, 1, 6, 0, 5},
			[]int{6, 3, 5, NULL, 2, 0, NULL, NULL, 1},
		},
		{
			[]int{3, 2, 1},
			[]int{3, NULL, 2, NULL, 1},
		},
	}
	convey.Convey("最大二叉树", t, func() {
		for _, v := range dataSet {
			res := leettree.ConstructMaximumBinaryTree(v.Nums)
			tree := leettree.Tree2ints(res)
			convey.So(tree, convey.ShouldResemble, v.ExpectResult)
		}
	})
}
