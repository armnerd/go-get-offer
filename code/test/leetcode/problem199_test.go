package leetcode

import (
	"code/leetcode/leettree"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

type Problem199DataSet struct {
	Root         []string
	ExpectResult []int
}

func TestProblem199(t *testing.T) {
	dataSet := []Problem199DataSet{
		{
			[]string{"1", "2", "3", "null", "5", "null", "4"},
			[]int{1, 3, 4},
		},
		{
			[]string{"1", "null", "3"},
			[]int{1, 3},
		},
	}
	convey.Convey("二叉树的右视图", t, func() {
		for _, v := range dataSet {
			root := leettree.BuildTreeFromArray(v.Root)
			res := leettree.RightSideView(root)
			convey.So(res, convey.ShouldResemble, v.ExpectResult)
		}
	})
}
