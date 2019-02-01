package buildtreeforminorderandpostorder

/*
 * @lc app=leetcode.cn id=106 lang=golang
 *
 * [106] 从中序与后序遍历序列构造二叉树
 *
 * https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/description/
 *
 * algorithms
 * Medium (56.94%)
 * Total Accepted:    4.7K
 * Total Submissions: 8.2K
 * Testcase Example:  '[9,3,15,20,7]\n[9,15,7,20,3]'
 *
 * 根据一棵树的中序遍历与后序遍历构造二叉树。
 *
 * 注意:
 * 你可以假设树中没有重复的元素。
 *
 * 例如，给出
 *
 * 中序遍历 inorder = [9,3,15,20,7]
 * 后序遍历 postorder = [9,15,7,20,3]
 *
 * 返回如下的二叉树：
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 *
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {

	return tree(inorder, postorder, 0, len(postorder)-1, len(inorder)-1)
}

func tree(inorder []int, postorder []int, inStart int, postEnd int, inEnd int) *TreeNode {

	if inStart > inEnd || postEnd < 0 {
		return nil
	}

	root := &TreeNode{postorder[postEnd], nil, nil}
	var inIndex int
	for i := inStart; i <= inEnd; i++ {
		if root.Val == inorder[i] {
			inIndex = i
			break
		}
	}

	root.Left = tree(inorder, postorder, inStart, postEnd-(inEnd-inIndex)-1, inIndex-1)
	root.Right = tree(inorder, postorder, inIndex+1, postEnd-1, inEnd)
	return root
}
