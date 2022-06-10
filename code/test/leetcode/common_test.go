package leetcode

import (
	"code/leetcode/leettree"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestTree(t *testing.T) {
	convey.Convey("树的基本操作", t, func() {
		data := [][]string{
			{"1", "2", "3", "null", "5", "null", "4"},
			// {"6", "3", "5", "null", "2", "0", "null", "null", "1"},
			// {"3", "null", "2", "null", "1"},
		}
		for _, row := range data {
			root := leettree.BuildTreeFromArray(row)
			res := leettree.SerializeATree(root)
			convey.So(res, convey.ShouldResemble, row)
		}
	})
}
