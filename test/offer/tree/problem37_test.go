package offer

import (
	"go-get-offer/offer/tree"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

type Problem37DataSet struct {
	Root []int
}

func TestProblem37(t *testing.T) {
	dataSet := []Problem37DataSet{
		{
			[]int{1, 2, 3, NULL, NULL, 6, 7},
		},
		{
			[]int{8, 6, 10, 5, 7, 9, 11},
		},
	}
	convey.Convey("序列化二叉树", t, func() {
		for _, v := range dataSet {
			root := tree.Ints2TreeNode(v.Root)
			resStr := tree.Serialize(root)
			rootTree := tree.Deserialize(resStr)
			convey.So(tree.Tree2ints(rootTree), convey.ShouldResemble, v.Root)
		}
	})
}
