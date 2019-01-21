package searchmatrix

/*
 * @lc app=leetcode.cn id=74 lang=golang
 *
 * [74] 搜索二维矩阵
 *
 * https://leetcode-cn.com/problems/search-a-2d-matrix/description/
 *
 * algorithms
 * Medium (32.03%)
 * Total Accepted:    4.8K
 * Total Submissions: 15K
 * Testcase Example:  '[[1,3,5,7],[10,11,16,20],[23,30,34,50]]\n3'
 *
 * 编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：
 *
 *
 * 每行中的整数从左到右按升序排列。
 * 每行的第一个整数大于前一行的最后一个整数。
 *
 *
 * 示例 1:
 *
 * 输入:
 * matrix = [
 * ⁠ [1,   3,  5,  7],
 * ⁠ [10, 11, 16, 20],
 * ⁠ [23, 30, 34, 50]
 * ]
 * target = 3
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入:
 * matrix = [
 * ⁠ [1,   3,  5,  7],
 * ⁠ [10, 11, 16, 20],
 * ⁠ [23, 30, 34, 50]
 * ]
 * target = 13
 * 输出: false
 *
 */
func searchMatrix(matrix [][]int, target int) bool {

	raw := -1
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	if n == 0 {
		return false
	}
	for i := 0; i < m; i++ {
		if matrix[i][0] == target {
			return true
		} else if matrix[i][0] > target {
			raw = i - 1
			break
		}
	}
	// 不存在则在最后行中查找
	if raw == -1 {
		raw = m - 1
	}
	// raw行寻找target
	mid, low, high := 0, 0, n-1
	for low <= high {
		mid = low + (high-low)/2

		// 相等则直接返回
		if matrix[raw][mid] == target {
			return true
		} else if matrix[raw][mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false
}
