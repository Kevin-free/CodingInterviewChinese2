package _9_MoreThanHalfNumber

func majorityElement(nums []int) int {
	// 1. sort
	// if nums == nil || len(nums) == 0 {
	//     return -1
	// }
	// sort.Ints(nums)
	// return nums[len(nums)/2]

	// 2. HashMap
	// m := make(map[int]int, 0)
	// half := len(nums)>>1
	// for _, num := range nums {
	//     m[num]++
	//     if m[num] > half {
	//         return num
	//     }
	// }
	// return -1

	// 3. Moore
	vote := 0
	x := 0
	for _, num := range nums {
		if vote == 0 {
			x = num
		}
		if x == num {
			vote++
		}else {
			vote--
		}
	}
	return x
}
