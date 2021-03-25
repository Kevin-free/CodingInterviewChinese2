package _4_ReverseList

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var pre *ListNode
	for head != nil {
		next := head.Next // 保存next，防断裂
		head.Next = pre   // 指向前
		pre = head        //pre
		head = next
	}
	return pre
}
