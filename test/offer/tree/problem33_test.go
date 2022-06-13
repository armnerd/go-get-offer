package tree

import (
	"go-get-offer/offer/tree"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

type Problem33DataSet struct {
	Root         []int
	ExpectResult bool
}

func TestProblem33(t *testing.T) {
	dataSet := []Problem33DataSet{
		{
			[]int{1, 3, 2},
			true,
		},
		{
			[]int{3, 1, 2},
			false,
		},
		{
			[]int{5, 7, 6, 9, 11, 10, 8},
			true,
		},
	}
	convey.Convey("二叉搜索树的后序遍历序列", t, func() {
		for _, v := range dataSet {
			res := tree.VerifySquenceOfBST(v.Root)
			convey.So(res, convey.ShouldResemble, v.ExpectResult)
		}
	})
}
