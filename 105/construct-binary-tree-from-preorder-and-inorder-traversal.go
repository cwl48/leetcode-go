package buildtree

/*
 * @lc app=leetcode.cn id=105 lang=golang
 *
 * [105] 从前序与中序遍历序列构造二叉树
 *
 * https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/
 *
 * algorithms
 * Medium (54.73%)
 * Total Accepted:    7.6K
 * Total Submissions: 13.8K
 * Testcase Example:  '[3,9,20,15,7]\n[9,3,15,20,7]'
 *
 * 根据一棵树的前序遍历与中序遍历构造二叉树。
 *
 * 注意:
 * 你可以假设树中没有重复的元素。
 *
 * 例如，给出
 *
 * 前序遍历 preorder = [3,9,20,15,7]
 * 中序遍历 inorder = [9,3,15,20,7]
 *
 * 返回如下的二叉树：
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 */

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {

	return tree(preorder, inorder, 0, 0, len(inorder)-1)
}

func tree(preorder []int, inorder []int, preStart int, inStart int, inEnd int) *TreeNode {

	if len(preorder)-1 < preStart || inEnd < inStart {
		return nil
	}

	// 根节点
	root := &TreeNode{preorder[preStart], nil, nil}
	// 找出根节点在中序遍历中的位置
	var index int
	for i := inStart; i <= inEnd; i++ {
		if inorder[i] == root.Val {
			index = i
			break
		}
	}

	// 构造左子树
	root.Left = tree(preorder, inorder, preStart+1, inStart, index-1)
	// 构造右子树
	root.Right = tree(preorder, inorder, preStart+(index-inStart)+1, index+1, inEnd)
	return root
}
