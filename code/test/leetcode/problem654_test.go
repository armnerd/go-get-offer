package leetcode

import (
	"code/leetcode/leettree"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

type Problem654DataSet struct {
	Nums         []int
	ExpectResult []string
}

func TestProblem654(t *testing.T) {
	dataSet := []Problem654DataSet{
		{
			[]int{3, 2, 1, 6, 0, 5},
			[]string{"6", "3", "5", "null", "2", "0", "null", "null", "1"},
		},
		{
			[]int{3, 2, 1},
			[]string{"3", "null", "2", "null", "1"},
		},
	}
	convey.Convey("最大二叉树", t, func() {
		for _, v := range dataSet {
			res := leettree.ConstructMaximumBinaryTree(v.Nums)
			tree := leettree.SerializeATree(res)
			convey.So(tree, convey.ShouldResemble, v.ExpectResult)
		}
	})
}
