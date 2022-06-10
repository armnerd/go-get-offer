package leetcode

import (
	"code/leetcode/leettree"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

type Problem199DataSet struct {
	Root         []int
	ExpectResult []int
}

func TestProblem199(t *testing.T) {
	dataSet := []Problem199DataSet{
		{
			[]int{1, 2, 3, NULL, 5, NULL, 4},
			[]int{1, 3, 4},
		},
		{
			[]int{1, NULL, 3},
			[]int{1, 3},
		},
	}
	convey.Convey("二叉树的右视图", t, func() {
		for _, v := range dataSet {
			root := leettree.Ints2TreeNode(v.Root)
			res := leettree.RightSideView(root)
			convey.So(res, convey.ShouldResemble, v.ExpectResult)
		}
	})
}
