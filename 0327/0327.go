package _327

import "CodingInterview2/structures"

type ListNode = structures.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func kthToLast(head *ListNode, k int) int {
	if head == nil {
		return -1
	}
	res := head
	for i := 0; i < k; i++ {
		head = head.Next
	}
	for head != nil {
		head = head.Next
		res = res.Next
	}
	return res.Val
}
