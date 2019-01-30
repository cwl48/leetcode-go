package largestrectanglearea

import "container/list"

/*
 * @lc app=leetcode.cn id=84 lang=golang
 *
 * [84] 柱状图中最大的矩形
 *
 * https://leetcode-cn.com/problems/largest-rectangle-in-histogram/description/
 *
 * algorithms
 * Hard (35.09%)
 * Total Accepted:    3.5K
 * Total Submissions: 10K
 * Testcase Example:  '[2,1,5,6,2,3]'
 *
 * 给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
 *
 * 求在该柱状图中，能够勾勒出来的矩形的最大面积。
 *
 *
 *
 *
 *
 * 以上是柱状图的示例，其中每个柱子的宽度为 1，给定的高度为 [2,1,5,6,2,3]。
 *
 *
 *
 *
 *
 * 图中阴影部分为所能勾勒出的最大矩形面积，其面积为 10 个单位。
 *
 *
 *
 * 示例:
 *
 * 输入: [2,1,5,6,2,3]
 * 输出: 10
 *
 */
func largestRectangleArea(heights []int) int {

	// list 模拟构建栈
	stack := list.New()

	for i := 0; i < len(heights); i++ {
		last := stack.Back()
		if last == nil || last.Value.(int) <= heights[i] {
			stack.PushBack(heights[i])
		}
		if last != nil && last.Value.(int) > heights[i] {
			break
		}
	}
	result := 0
	if stack.Len() > 0 {
		result = stack.Back().Value.(int)
		for i := stack.Len(); i < len(heights); i++ {

			times := 1
			last := stack.Back()
			for last != nil && last.Value.(int) >= heights[i] {
				result = max(last.Value.(int)*times, result)
				// 出栈
				stack.Remove(last)
				last = stack.Back()
				times++
			}
			for ; times > 1; times-- {
				stack.PushBack(heights[i])
			}
			stack.PushBack(heights[i])
		}
		var n *list.Element
		t := 0
		for i := stack.Front(); i != nil; i = n {
			result = max(i.Value.(int)*(stack.Len()-t), result)
			n = i.Next()
			t++
		}
	}
	return result
}

func max(x int, y int) int {
	if x >= y {
		return x
	}
	return y
}
