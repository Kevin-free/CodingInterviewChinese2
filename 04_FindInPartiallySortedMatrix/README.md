# [剑指 Offer 04. 二维数组中的查找](https://leetcode-cn.com/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof/)


## 题目

在一个 n * m 的二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。请完成一个高效的函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。

 
示例:

现有矩阵 matrix 如下：

```
[
  [1,   4,  7, 11, 15],
  [2,   5,  8, 12, 19],
  [3,   6,  9, 16, 22],
  [10, 13, 14, 17, 24],
  [18, 21, 23, 26, 30]
]
```
给定 target = 5，返回 true。

给定 target = 20，返回 false。


限制：

`0 <= n <= 1000`

`0 <= m <= 1000`


## 解题思路

因为数组是从左到右，从上到下递增的，所以我们可以选择

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