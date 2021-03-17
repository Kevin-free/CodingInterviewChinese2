package _4_FindInPartiallySortedMatrix

import "testing"

func Test04(t *testing.T) {
	var nums = [][]int{{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30}}

	var res bool
	res = findNumberIn2DArray(nums, 5)
	println(res)
	res = findNumberIn2DArray(nums, 20)
	println(res)
}
