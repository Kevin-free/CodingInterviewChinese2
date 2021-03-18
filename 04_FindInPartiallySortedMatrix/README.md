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

因为数组是从左到右，从上到下递增的，所以我们可以选择从右上或者左下查找。

……

这里我实现的是从左下（相信你可以类比写出右上的）

## 代码

```go
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
```


![](http://wesub.ifree258.top/bottomPic.png)