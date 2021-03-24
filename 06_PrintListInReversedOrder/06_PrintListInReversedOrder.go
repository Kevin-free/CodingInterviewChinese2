package _6_PrintListInReversedOrder

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}
	stack := make([]int, 0)
	for head != nil {
		stack = append(stack, head.Val)
		head = head.Next
		// }
		// if len(stack) == 0 {
		//     return nil
		// }
		// res := make([]int, 0)
		// for i := 0; i < len(stack); i++ {
		//     res[i] = stack[len(stack)-1-i]
		// }
	}

	//sort.Reverse(sort.IntSlice(stack))

	// reverse
	for i, j := 0, len(stack)-1; i < j; {
		stack[i], stack[j] = stack[j], stack[i]
		i++
		j--
	}
	return stack
}
