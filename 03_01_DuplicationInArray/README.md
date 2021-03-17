# [剑指 Offer 03. 数组中重复的数字](https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/)


## 题目

在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。

示例 1：
```
输入：
[2, 3, 1, 0, 2, 5, 3]
[1, 3, 2, ...
[1
输出：2 或 3 
```

限制：
`2 <= n <= 100000`


## 解题思路

解决这个问题虽然不难想到方法，比如：

一、先把输入的数组排序，从排序的数组中找出重复的数字就是一件很容易的事情，只需要从头到尾扫描排序后的数组。
一个长度为 n 的数组排序需要 O(nlogn) 的时间。

实现如方法 `findRepeatNumber1` 。

二、我们还可以利用哈希表来解决这个问题。从头到尾扫描数组的每一个数，用 O(1) 的时候来判断哈希表中是否包含这个数，如果哈希表中没有这个数，就把它加入哈希表；如果哈希表已经存在这个数，就找到来一个重复的数。
这个算法的时间复杂度是 O(n)，但它提高时间效率的代价是需要一个大小为 O(n) 的哈希表。

实现如方法 `findRepeatNumber2` 。

那我们再想想有没有空间复杂度为 O(1) 的算法呢？

三、需要注意到数组中的数字都在 0～n-1 的范围内。如果这个数组中没有重复的数字，那么数组排序之后数字 v 将出现在下标为 i 的位置。
由于数组中有重复的数字，有些位置可能存在多个数字，同时有些位置可能没有数字。

那么我们重排这个数组，从头到尾依次扫描数组中的每个数字。扫描到下标为 i 的数字时，我们首先比较这个数（用 v 表示）是不是等与 i 。
若果是，则接着扫描下一个数字；
如果不是，则拿它和第 v 个数字进行比较。
    若果它和第 v 个数字相等，就找到了一个重复的数字（该数字在下标为 i 和 v 的位置都出现了）；
    如果它和第 v 个数不相等，就把第 i 个数字和第 v 个数字交换，把 v 放到属于它的位置。
接下来再重复这个比较、交换的过程，直到找到一个重复的数字。

这里以数组 {2, 3, 1, 0, 2, 5, 3} 为例来分析找到重复数字的步骤。
数组的第 0 个数字（从 0 开始计数，和数组的下标保持一致）是 2 ，与它的下标不想等，于是把它和下标为 2 的数字 1 交换，
交换之后的数组是 {1, 3, 2, 0, 2, 5, 3} 。
此时第 0 个数字是 1 ，仍然与它的下标不相等，同样把它和下标为 1 的数字 3 交换，
交换之后的数组是 {3, 1, 2, 0, 2, 5, 3} 。
接下来继续交换第 0 个数字 3 和第 3 个数字 0，
交换之后的数组是 {0, 1, 2, 3, 2, 5, 3} 。
此时第 0 个数字是 0 ，接着扫描下一个数字。在接下来的几个数中，下标为 1，2，3 的数分别为 1，2，3，它们的下标和数字相等，因此继续扫描其后的数字。
接下来扫描下标为 4 的数字 2，由于它的数字和下标不相等，再比较它和下标为 2 的数字，注意到此时数组中下标为 2 的数也是 2 ，也就是数字 2 在下标 2 和 下标 4 的两个位置都出现了，因此找到了一个重复的数字。

实现如方法 `findRepeatNumber3` 。


## 代码

```go
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
```

```go
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
```

```go
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
```

```go
//方法四是提交记录看到 beat 99% 的，学习了：）
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
```

![0304.png](https://pic.leetcode-cn.com/1615707932-Bpveov-0304.png)

![](http://wesub.ifree258.top/bottomPic.png)