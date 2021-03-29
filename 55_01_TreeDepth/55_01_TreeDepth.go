package _5_01_TreeDepth

import (
	"CodingInterview2/structures"
	"CodingInterview2/util"
)

type TreeNode = structures.TreeNode

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return util.MaxInt(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
