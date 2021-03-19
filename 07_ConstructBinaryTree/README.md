# [剑指 Offer 07. 重建二叉树](https://leetcode-cn.com/problems/zhong-jian-er-cha-shu-lcof/)


## 题目

输入某二叉树的前序遍历和中序遍历的结果，请重建该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。


例如，给出

前序遍历 preorder = [3,9,20,15,7]
中序遍历 inorder = [9,3,15,20,7]
返回如下的二叉树：

```
    3
   / \
  9  20
    /  \
   15   7
```

限制：

`0 <= 节点个数 <= 5000`


## 解题思路

前序遍历性质： 节点按照 [ 根节点 | 左子树 | 右子树 ] 排序。

中序遍历性质： 节点按照 [ 左子树 | 根节点 | 右子树 ] 排序。

![0701.png](https://pic.leetcode-cn.com/1616160519-slmuaR-0701.png)

根据以上性质，可得出以下推论：

1. 前序遍历的首元素 为 树的根节点 root 的值。

2. 在中序遍历中搜索根节点 root 的索引 ，可将 中序遍历 划分为 [ 左子树 | 根节点 | 右子树 ] 。

3. 根据中序遍历中的左 / 右子树的节点数量，可将 前序遍历 划分为 [ 根节点 | 左子树 | 右子树 ] 。

所以我们可以通过递归来构建二叉树。

> 为了提升效率，还使用来哈希表 map 存储中序遍历的值与索引的映射，查找操作的时间复杂度为 O(1)


## 代码

```go
//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var indexMap map[int]int // map[value]index

func buildTree(preorder []int, inorder []int) *TreeNode {
	// special judgment
	if preorder == nil || inorder == nil || len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	// init & set map
	indexMap = make(map[int]int, 0)
	for i, v := range inorder {
		indexMap[v] = i
	}

	return buildCore(preorder, inorder, 0, len(preorder)-1, 0, len(inorder)-1)
}

func buildCore(preorder, inorder []int, preorderLeft, preorderRight, inorderLeft, inorderRight int) *TreeNode {
	if preorderLeft > preorderRight {
		return nil
	}

	// 前序遍历的第一个节点是根节点
	preorderRoot := preorderLeft
	// 中序遍历中根节点的位置
	inorderRoot := indexMap[preorder[preorderRoot]]
	// 设置根节点
	root := &TreeNode{Val: preorder[preorderRoot]}
	// 得到左子树中的节点个数
	leftSubtreeNum := inorderRoot - inorderLeft
	// 构建左子树
	// 先序遍历中的左子树元素为[preorder_left+1, preorder_left+left_subTree_num];中序遍历中的左子树元素为[inorder_left, inorder_root-1]
	root.Left = buildCore(preorder, inorder, preorderLeft+1, preorderLeft+leftSubtreeNum, inorderLeft, inorderRoot-1)
	// 构建右子树
	// 先序遍历中的右子树元素为[preorder_left+left_subTree_num+1, preorder_right];中序遍历中的右子树元素为[inorder_root+1, inorder_right]
	root.Right = buildCore(preorder, inorder, preorderLeft+leftSubtreeNum+1, preorderRight, inorderRoot+1, inorderRight)
	// 返回树
	return root
}
```


![](http://wesub.ifree258.top/bottomPic.png)