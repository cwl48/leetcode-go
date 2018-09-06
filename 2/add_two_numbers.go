package addtwonumbers

/*
 * [2] Add Two Numbers
 *
 * https://leetcode-cn.com/problems/add-two-numbers/description/
 *
 * algorithms
 * Medium (29.06%)
 * Total Accepted:    31.3K
 * Total Submissions: 107.6K
 * Testcase Example:  '[2,4,3]\n[5,6,4]'
 *
 * 给定两个非空链表来表示两个非负整数。位数按照逆序方式存储，它们的每个节点只存储单个数字。将两数相加返回一个新的链表。
 *
 * 你可以假设除了数字 0 之外，这两个数字都不会以零开头。
 *
 * 示例：
 *
 * 输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
 * 输出：7 -> 0 -> 8
 * 原因：342 + 465 = 807
 *
 *
 */

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var p1, p2 = l1, l2
	sum := 0
	var result = &ListNode{Val: 0}
	var current = result
	for p1 != nil || p2 != nil {

		if p1 != nil {
			sum += p1.Val
			p1 = p1.Next
		}
		if p2 != nil {
			sum += p2.Val
			p2 = p2.Next
		}
		current.Next = &ListNode{Val: sum % 10}
		sum = sum / 10
		current = current.Next
	}
	if sum == 1 {
		current.Next = &ListNode{Val: 1}
	}
	return result.Next
}
