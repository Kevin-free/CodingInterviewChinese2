# [剑指 Offer 27. 二叉树的镜像](https://leetcode-cn.com/problems/er-cha-shu-de-jing-xiang-lcof/)

简单｜二叉树、递归

## 题目

请完成一个函数，输入一个二叉树，该函数输出它的镜像。

例如输入：
```
     4
   /   \
  2     7
 / \   / \
1   3 6   9
```

镜像输出：
```
     4
   /   \
  7     2
 / \   / \
9   6 3   1
```

示例 1：
```
输入：root = [4,2,7,1,3,6,9]
输出：[4,7,2,9,6,3,1]
```

备注:
这个问题是受到 Max Howell 的 原问题 启发的 ：

谷歌：我们90％的工程师使用您编写的软件(Homebrew)，但是您却无法在面试时在白板上写出翻转二叉树这道题，这太糟糕了。

ps. 所以，如果你做出来了，快去应聘谷歌吧 /doge

## 解题思路

对于某个子树，如
```
     4
   /   \
  7     2
```
我们得到其镜像只需将左右子树交换，即`node.Left, node.Right = node.Right, node.Left`，

所以，对于整课树，我们可以用递归来解决，从根节点开始，先递归调用反转根节点的左孩子，然后递归调用反转根节点的右孩子。

当然，我们别忽略了根节点，即 node == nil 的情况，这也是递归终止的情况，我们需要返回 nil。



## 代码

```go
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
```


![](http://wesub.ifree258.top/bottomPic.png)