package tree

import (
	"go-get-offer/offer/tree"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

type Problem78DataSet struct {
	Root         []int
	ExpectResult [][]int
}

func TestProblem78(t *testing.T) {
	dataSet := []Problem78DataSet{
		{
			[]int{1, 2, 3, NULL, NULL, 4, 5},
			[][]int{
				{1},
				{2, 3},
				{4, 5},
			},
		},
		{
			[]int{8, 6, 10, 5, 7, 9, 11},
			[][]int{
				{8},
				{6, 10},
				{5, 7, 9, 11},
			},
		},
		{
			[]int{1, 2, 3, 4, 5},
			[][]int{
				{1},
				{2, 3},
				{4, 5},
			},
		},
	}
	convey.Convey("把二叉树打印成多行", t, func() {
		for _, v := range dataSet {
			root := tree.Ints2TreeNode(v.Root)
			res := tree.Print(root)
			convey.So(res, convey.ShouldResemble, v.ExpectResult)
		}
	})
}
