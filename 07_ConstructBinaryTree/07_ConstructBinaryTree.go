package _7_ConstructBinaryTree

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
