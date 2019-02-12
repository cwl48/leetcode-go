package spiralorder

/*
 * @lc app=leetcode.cn id=54 lang=golang
 *
 * [54] 螺旋矩阵
 *
 * https://leetcode-cn.com/problems/spiral-matrix/description/
 *
 * algorithms
 * Medium (33.21%)
 * Total Accepted:    7.7K
 * Total Submissions: 23.2K
 * Testcase Example:  '[[1,2,3],[4,5,6],[7,8,9]]'
 *
 * 给定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。
 *
 * 示例 1:
 *
 * 输入:
 * [
 * ⁠[ 1, 2, 3 ],
 * ⁠[ 4, 5, 6 ],
 * ⁠[ 7, 8, 9 ]
 * ]
 * 输出: [1,2,3,6,9,8,7,4,5]
 *
 *
 * 示例 2:
 *
 * 输入:
 * [
 * ⁠ [1, 2, 3, 4],
 * ⁠ [5, 6, 7, 8],
 * ⁠ [9,10,11,12]
 * ]
 * 输出: [1,2,3,4,8,12,11,10,9,5,6,7]
 *
 *
 */
func spiralOrder(matrix [][]int) []int {

	m, n := len(matrix), 0
	if m > 0 {
		n = len(matrix[0])
	}
	res := make([]int, 0)
	startRow, startCol, endRow, endCol := 0, n-1, m-1, 0

	for startRow <= endRow && startCol >= endCol {

		for i := endCol; i <= startCol; i++ {
			res = append(res, matrix[startRow][i])
		}
		startRow++
		for i := startRow; i <= endRow; i++ {
			res = append(res, matrix[i][startCol])
		}
		startCol--

		if startRow <= endRow {
			for i := startCol; i >= endCol; i-- {
				res = append(res, matrix[endRow][i])
			}
		}
		endRow--
		if endCol <= startCol {
			for i := endRow; i >= startRow; i-- {
				res = append(res, matrix[i][endCol])
			}
		}
		endCol++
	}

	return res
}
