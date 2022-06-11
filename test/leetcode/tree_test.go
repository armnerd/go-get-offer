package leetcode

import (
	"go-get-offer/leetcode/leettree"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

var NULL = -1 << 63

func TestTree(t *testing.T) {
	convey.Convey("leetcode树操作", t, func() {
		data := [][]int{
			{1, 2, 3, NULL, 5, NULL, 4},
			{6, 3, 5, NULL, 2, 0, NULL, NULL, 1},
			{3, NULL, 2, NULL, 1},
			{1, 2, 3, NULL, NULL, 4, 5},
			{1, NULL, 2, 3},
			{5, 4, 7, 3, NULL, 2, NULL, -1, NULL, 9},
		}
		for _, row := range data {
			root := leettree.Ints2TreeNode(row)
			res := leettree.Tree2ints(root)
			convey.So(res, convey.ShouldResemble, row)
		}
	})
}
