package generatematrix

/*
 * @lc app=leetcode.cn id=59 lang=golang
 *
 * [59] 螺旋矩阵 II
 *
 * https://leetcode-cn.com/problems/spiral-matrix-ii/description/
 *
 * algorithms
 * Medium (68.89%)
 * Total Accepted:    4.8K
 * Total Submissions: 6.9K
 * Testcase Example:  '3'
 *
 * 给定一个正整数 n，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的正方形矩阵。
 *
 * 示例:
 *
 * 输入: 3
 * 输出:
 * [
 * ⁠[ 1, 2, 3 ],
 * ⁠[ 8, 9, 4 ],
 * ⁠[ 7, 6, 5 ]
 * ]
 *
 */
func generateMatrix(n int) [][]int {

	arr := make([][]int, n)
	for i := range arr {
		arr[i] = make([]int, n)
	}

	rowStart, colStart, rowEnd, colEnd := 0, 0, n-1, n-1
	cnt := 1
	for rowStart <= rowEnd && colStart <= colEnd {

		for i := colStart; i <= colEnd; i++ {
			arr[rowStart][i] = cnt
			cnt++

		}
		rowStart++

		for i := rowStart; i <= rowEnd; i++ {
			arr[i][colEnd] = cnt
			cnt++
		}
		colEnd--

		for i := colEnd; i >= colStart; i-- {
			arr[rowEnd][i] = cnt
			cnt++
		}
		rowEnd--

		for i := rowEnd; i >= rowStart; i-- {
			arr[i][colStart] = cnt
			cnt++
		}
		colStart++
	}

	return arr
}
