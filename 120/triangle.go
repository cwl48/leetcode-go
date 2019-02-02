package minimumtotal

/*
 * @lc app=leetcode.cn id=120 lang=golang
 *
 * [120] 三角形最小路径和
 *
 * https://leetcode-cn.com/problems/triangle/description/
 *
 * algorithms
 * Medium (55.18%)
 * Total Accepted:    6.6K
 * Total Submissions: 11.9K
 * Testcase Example:  '[[2],[3,4],[6,5,7],[4,1,8,3]]'
 *
 * 给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。
 *
 * 例如，给定三角形：
 *
 * [
 * ⁠    [2],
 * ⁠   [3,4],
 * ⁠  [6,5,7],
 * ⁠ [4,1,8,3]
 * ]
 *
 *
 * 自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。
 *
 * 说明：
 *
 * 如果你可以只使用 O(n) 的额外空间（n 为三角形的总行数）来解决这个问题，那么你的算法会很加分。
 *
 */
func minimumTotal(triangle [][]int) int {

	m := len(triangle)
	// dp
	for i := m - 1; i >= 0; i-- {
		size := len(triangle[i])
		for j := 0; j < size-1; j++ {
			triangle[i-1][j] += min(triangle[i][j], triangle[i][j+1])
		}
	}

	return triangle[0][0]
}

func min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}
