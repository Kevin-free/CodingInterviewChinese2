package _316

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
