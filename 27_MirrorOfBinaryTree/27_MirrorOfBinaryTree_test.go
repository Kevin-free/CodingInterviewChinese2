package _7_MirrorOfBinaryTree

import (
	"CodingInterview2/structures"
	"fmt"
	"testing"
)


type question27 struct {
	para27
	ans27
}

// para 是参数
// one 代表第一个参数
type para27 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans27 struct {
	one []int
}

func Test_Problem27(t *testing.T) {

	qs := []question27{

		{
			para27{[]int{}},
			ans27{[]int{}},
		},

		{
			para27{[]int{1}},
			ans27{[]int{1}},
		},

		{
			para27{[]int{4, 2, 7, 1, 3, 6, 9}},
			ans27{[]int{4, 7, 2, 9, 6, 3, 1}},
		},
	}

	fmt.Printf("------------------------Offer Problem 27------------------------\n")

	for _, q := range qs {
		_, p := q.ans27, q.para27
		fmt.Printf("【input】:%v      ", p)
		root := structures.Ints2TreeNode(p.one)
		fmt.Printf("【output】:%v      \n", structures.Tree2Preorder(mirrorTree(root)))
	}
	fmt.Printf("\n\n\n")
}
