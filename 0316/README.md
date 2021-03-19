# [面试 offer]


## 题目

查找第一个值等于给定值的元素。

说明 : 数组有序且可能包含重复元素，输出目标元素下标

举例

1,3,4,5,8,8,8,11,18

第一个值等于3，输出1

第一个值等于8，输出4

没找到返回-1


## 解题思路

当然我们很容易想到直接扫描查找，但是 O(N) 的时间复杂度，这种方法显然不能让面试官满意。但是面试但时候也可以先提一嘴，为后续的优化做铺垫。

对于查找，我们通常要想到时间复杂度为 O(logN) 的[二分查找](https://mp.weixin.qq.com/s/-svAlOT4jNezExYAIegjYQ)。并且这题满意有序的前提。

但是这题与普通二分查找不同的是重复、找第一个。

所以所以我们可以使用寻找左侧边界的[二分查找](https://mp.weixin.qq.com/s/-svAlOT4jNezExYAIegjYQ)

与二分查找的不同在于，寻找左侧边界的二分查找在找到目标值时并不直接返回，而是将右侧边界收紧以锁定左侧边界。

如果目标值不存在，找出的是比它大的数的左边界，所以最后需要额外判断左边界值是否等于目标值。

## 代码

```go
func searchLeft(nums []int, target int) int {
	// 特殊判断
	if nums == nil || len(nums) == 0 {
		return -1
	}
	left := 0
	right := len(nums) - 1

	for left <= right {
		//mid := (left + right)/2 // 简单写法
		//mid := left + (right-left)/2 // 防止溢出
		mid := left + (right-left)>>1 // 位运算，高效写法
		if nums[mid] == target {
			right = mid - 1 // diff with binary search
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}

	// 如果目标值不存在，找出的是比它大的数的左边界
	if nums[left] == target {
		return left
	}else {
		return -1
	}
}
```

![罗翔老师都觉得合适的不得了 :D ](http://wesub.ifree258.top/gifBottom.gif)

![](http://wesub.ifree258.top/bottomPic.png)