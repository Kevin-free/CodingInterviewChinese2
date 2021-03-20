package _3_DuplicationInArray

import (
	"fmt"
	"testing"
)

type question3 struct {
	para3
	ans3
}

// para 是参数
// one 代表第一个参数
type para3 struct {
	nums []int
}

// ans 是答案
// one 代表第一个答案
type ans3 struct {
	one int
}

func Test0301(t *testing.T) {

	qs := []question3{
		{
			para3{[]int{2, 3, 1, 0, 2, 5, 3}},
			ans3{2},
		},
		{
			para3{[]int{2, 3, 1, 3}},
			ans3{3},
		},
		{
			para3{[]int{2, 1, 3, 0, 4}},
			ans3{-1},
		},
		{
			para3{[]int{}},
			ans3{-1},
		},
		{
			para3{nil},
			ans3{-1},
		},

		// 如需多个测试，可以复制上方元素。
	}

	fmt.Printf("------------------------Offer Problem 3------------------------\n")

	for _, q := range qs {
		_, p := q.ans3, q.para3
		fmt.Printf("【input】:%v       【output】:%v\n", p, findRepeatNumber3(p.nums))
	}
	fmt.Printf("\n\n\n")

	//input := []int{2, 3, 1, 0, 2, 5, 3}
	//output := findRepeatNumber3(input)
	//println(output)
	//test01()
}
