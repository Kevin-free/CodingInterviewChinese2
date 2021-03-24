package _6_PrintListInReversedOrder

import (
	"fmt"
	"testing"
)

func Test06(t *testing.T) {
	input := &ListNode{1,&ListNode{3,&ListNode{2, nil}}}
	output := reversePrint(input)
	fmt.Println(output)
}