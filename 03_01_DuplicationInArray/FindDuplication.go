package _3_01_DuplicationInArray

import "sort"

//Author：Kevin

//https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/
//剑指 Offer 03. 数组中重复的数字

//找出数组中重复的数字。
//在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。
//数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。
//请找出数组中任意一个重复的数字。

//方法一：先排序在查重
//时间复杂度：O(nlogn) 用于排序
//空间复杂度：O(1)
func findRepeatNumber1(nums []int) int {
	//特殊判断
	if nums == nil || len(nums) == 0 {
		return -1 //不符合，返回-1
	}
	for _, num := range nums {
		if num < 0 || num > len(nums)-1 {
			return -1
		}
	}

	sort.Ints(nums) // 实现是快排
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			// 前后元素相等
			return nums[i]
		}
	}
	return -1
}

//方法二：哈希表
//时间复杂度：O(n)
//空间复杂度：O(n)
//特殊判断
func findRepeatNumber2(nums []int) int {
	//特殊判断
	if nums == nil || len(nums) == 0 {
		return -1
	}
	for _, num := range nums {
		if num < 0 || num > len(nums)-1 {
			return -1
		}
	}

	m := make(map[int]int, 0)
	for i, v := range nums {
		if _, exist := m[v]; exist {
			// 存在该数
			return v
		}
		// 不存在
		m[v] = i
	}
	return -1
}

//方法三：原地置换
//时间复杂度：O(n）
//空间复杂度：O(1)
func findRepeatNumber3(nums []int) int {
	//特殊判断
	if nums == nil || len(nums) == 0 {
		return -1
	}
	for _, num := range nums {
		if num < 0 || num > len(nums)-1 {
			return -1
		}
	}

	for i := 0; i < len(nums); i++ {
		if i != nums[i] {
			if nums[i] == nums[nums[i]] {
				return nums[i]
			}
			// swap nums[i] and nums[nums[i]]
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}
	return -1
}

//方法四：数组标记
//时间复杂度：O(n)
//空间复杂度：O(n)
func findRepeatNumber4(nums []int) int {
	var sign [100000]bool
	for _,num := range nums {
		if sign[num] {return num}
		sign[num] = true
	}
	return 0
}