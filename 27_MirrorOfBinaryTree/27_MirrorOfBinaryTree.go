package _7_MirrorOfBinaryTree

import "CodingInterview2/structures"

type TreeNode = structures.TreeNode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mirrorTree(root *TreeNode) *TreeNode {
	return doMirror(root)
}

func doMirror(node *TreeNode) *TreeNode {
	// end case
	if node == nil {
		return nil
	}
	// tmp := node.Left
	// node.Left = node.Right
	// node.Right = tmp
	// opt: size
	node.Left, node.Right = node.Right, node.Left
	doMirror(node.Left)
	doMirror(node.Right)
	return node
}