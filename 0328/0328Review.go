package _328

// 一个单链表，假设第一个节点我们定为下标为1，第二个为2，依次类推，同时假设下标为奇数的结点是升序排序，偶数的结点是降序排序，如何让整个链表升序？
/**
 * 单链表奇数递增偶数递减，使之升序
 * 分三步：
 * 1.拆分成2个链表
 * 2.对逆序的链表反转
 * 3.合并2个链表
 */

// 16：50
func sortOddDoubleList2(list *ListNode) *ListNode {
	// special deal
	res := &ListNode{}
	if list == nil {
		return res
	}
	// 1。分成奇偶两个链表
	odd, double := part2List(list)

	// 2。对降序的偶链表升序
	sortedDouble := sort2IncList(double)

	// 3。合并两个升序链表
	res = mergeTowSortedList(odd, sortedDouble)

	return res
}

// end 17:38

// 16：50
// 按奇偶拆分成两个链表
func part2List(list *ListNode) (*ListNode, *ListNode) {
	l1 := list
	l2 := list.Next
	next := &ListNode{}
	copyNode := &ListNode{}
	for l1 != nil {
		next = l1.Next.Next
		copyNode = l1.Next
		l1.Next = next
		if next == nil {
			copyNode.Next = nil
		} else {
			copyNode.Next = next.Next
		}
		l1 = next
	}
	return list, l2

	//odd := list
	//double := list.Next
	//oddNext := &ListNode{}    // 存放奇数的下一个节点
	//doubleNext := &ListNode{} // 存放偶数的一下个节点
	//for odd != nil {
	//	oddNext = odd.Next.Next
	//	doubleNext = odd.Next
	//	odd.Next = oddNext
	//	if oddNext == nil {
	//		doubleNext.Next = nil
	//	} else {
	//		doubleNext.Next = oddNext.Next
	//	}
	//	odd = oddNext
	//}
	//return odd, double
}

// end 17:10

// pause 17:12
// 反转链表（剑指offer 24）
func sort2IncList(list *ListNode) *ListNode {
	//pre := &ListNode{} //can not use this: ListNode{0, nil}
	var pre *ListNode // use this: nil
	var next *ListNode
	for list != nil {
		next = list.Next // 保存next，防止断链
		list.Next = pre  // 反转
		pre = list       // pre 后移
		list = next      // list 后移
	}
	return pre
}

// end 17:22 cost 10min

// 17:23
// 合并两个排序的链表（剑指offer 25）
func mergeTowSortedList(list1, list2 *ListNode) *ListNode {
	//var dummyHead *ListNode // can not! dummyHead is nil
	dummyHead := &ListNode{}
	cur := dummyHead // cur & dummyHead is same address, so use cur get all value
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next = &ListNode{Val: list1.Val}
			list1 = list1.Next
		} else {
			cur.Next = &ListNode{Val: list2.Val}
			list2 = list2.Next
		}
		cur = cur.Next
	}
	if list1 != nil {
		cur.Next = list1
	}
	if list2 != nil {
		cur.Next = list2
	}
	return dummyHead.Next // dummyHead.Next is the head

	//res := &ListNode{}
	//for list1 != nil && list2 != nil {
	//	if list1.Val < list2.Val {
	//		res.Next = &ListNode{Val: list1.Val}
	//		list1 = list1.Next
	//	} else {
	//		res.Next = &ListNode{Val: list2.Val}
	//		list2 = list2.Next
	//	}
	//}
	//// deal
	//if list1 != nil {
	//	res.Next = list1
	//}
	//if list2 != nil {
	//	res.Next = list2
	//}
	//return res
}

// end 15:35 10min
