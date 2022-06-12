package offer

import (
	"go-get-offer/offer/tree"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

var NULL = -1 << 63

type Problem7DataSet struct {
	Pre          []int
	Vin          []int
	ExpectResult []int
}

func TestProblem7(t *testing.T) {
	dataSet := []Problem7DataSet{
		{
			[]int{1, 2, 4, 7, 3, 5, 6, 8},
			[]int{4, 7, 2, 1, 5, 3, 8, 6},
			[]int{1, 2, 3, 4, NULL, 5, 6, NULL, 7, NULL, NULL, 8},
		},
		{
			[]int{1},
			[]int{1},
			[]int{1},
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7},
			[]int{3, 2, 4, 1, 6, 5, 7},
			[]int{1, 2, 5, 3, 4, 6, 7},
		},
	}
	convey.Convey("重建二叉树", t, func() {
		for _, v := range dataSet {
			res := tree.ReConstructBinaryTree(v.Pre, v.Vin)
			convey.So(tree.Tree2ints(res), convey.ShouldResemble, v.ExpectResult)
		}
	})
}
