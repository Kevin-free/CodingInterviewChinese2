package _328

import (
	"CodingInterview2/structures"
	"fmt"
	"testing"
)

type question032801 struct {
	para032801
	ans032801
}

// para 是参数
// one 代表第一个参数
type para032801 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans032801 struct {
	one []int
}

func Test_Problem032801(t *testing.T) {

	qs := []question032801{

		//{
		//	para032801{[]int{}},
		//	ans032801{[]int{}},
		//},

		{
			para032801{[]int{1, 8, 3, 6, 5, 4, 7, 2, 9, 0}},
			ans032801{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
		},
		// 如需多个测试，可以复制上方元素。
	}

	fmt.Printf("------------------------Leetcode Problem 032801------------------------\n")

	//for _, q := range qs {
	//	_, p := q.ans032801, q.para032801
	//	fmt.Printf("【input】:%v       【output】:%v\n", p, structures.List2Ints(sortOddDoubleList(structures.Ints2List(p.one))))
	//}
	for _, q := range qs {
		_, p := q.ans032801, q.para032801
		fmt.Printf("【input】:%v       【output】:%v\n", p, structures.List2Ints(sortOddDoubleList2(structures.Ints2List(p.one))))
	}
	fmt.Printf("\n\n\n")
}
