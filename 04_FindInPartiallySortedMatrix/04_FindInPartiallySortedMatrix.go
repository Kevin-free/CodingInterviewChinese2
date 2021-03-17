package _4_FindInPartiallySortedMatrix

func findNumberIn2DArray(matrix [][]int, target int) bool {
	// ⚠️特殊判断
	if matrix == nil || len(matrix) == 0 {
		return false
	}

	row, col := len(matrix), len(matrix[0])
	l, r := row-1, 0
	// ⚠️ l >= 0 (index can be zero)
	for l >= 0 && r < col {
		if matrix[l][r] < target {
			r++
		} else if matrix[l][r] > target {
			l--
		} else {
			return true
		}
	}
	return false
}
