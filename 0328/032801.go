package _328

import (
	"CodingInterview2/structures"
	"sort"
)

// 腾讯面试题
// 一个单链表，假设第一个节点我们定为下标为1，第二个为2，依次类推，同时假设下标为奇数的结点是升序排序，偶数的结点是降序排序，如何让整个链表升序？

type ListNode = structures.ListNode

func sortOddDoubleList(list *ListNode) *ListNode {
	sortedList := &ListNode{}
	if list == nil {
		return sortedList
	}
	// 使用两个数组分别存储
	inc := make([]int, 0) // store inc
	dec := make([]int, 0) // store dec
	i := 1
	// 1.分别存储 O(N)
	for list != nil {
		if i%2 == 1 { // 奇数
			inc = append(inc, list.Val)
			i++
		} else if i%2 == 0 { //偶数
			dec = append(dec, list.Val)
			i++
		}
		list = list.Next
	}
	// 2.合并两个数组为一个升序数组 O(N)
	sortedArr := mergerTwoSortedArr(inc, dec)
	// 3.将数组转为链表 O(N)
	tmp := sortedList
	for _, v := range sortedArr {
		tmp.Next = &ListNode{Val: v}
		tmp = tmp.Next
	}

	return sortedList.Next

	//return structures.Ints2List(sortedArr)
}

// 合并两个数组为一个升序数组
func mergerTwoSortedArr(arr1, arr2 []int) []int {
	// 将 dec:arr2 升序 O(n*log(n))
	sort.Sort(sort.IntSlice(arr2))

	res := make([]int, 0)

	i, j := 0, 0
	// 合并两个有序数组 O(N)
	for ; i < len(arr1) && j < len(arr2); {
		if arr1[i] < arr2[j] {
			res = append(res, arr1[i])
			i++
		} else {
			res = append(res, arr2[j])
			j++
		}
	}
	// 尾部处理：一个数组还未遍历完，另一个数组已经遍历完了
	if i != len(arr1) {
		for ; i < len(arr1); i++ {
			res = append(res, arr1[i])
		}
	}
	if j != len(arr2) {
		for ; j < len(arr2); j++ {
			res = append(res, arr2[j])
		}
	}

	return res
}
