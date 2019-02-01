package yanghuisanjiao

/*
 * @lc app=leetcode.cn id=119 lang=golang
 *
 * [119] 杨辉三角 II
 *
 * https://leetcode-cn.com/problems/pascals-triangle-ii/description/
 *
 * algorithms
 * Easy (52.52%)
 * Total Accepted:    8.7K
 * Total Submissions: 16.5K
 * Testcase Example:  '3'
 *
 * 给定一个非负索引 k，其中 k ≤ 33，返回杨辉三角的第 k 行。
 *
 *
 *
 * 在杨辉三角中，每个数是它左上方和右上方的数的和。
 *
 * 示例:
 *
 * 输入: 3
 * 输出: [1,3,3,1]
 *
 *
 * 进阶：
 *
 * 你可以优化你的算法到 O(k) 空间复杂度吗？
 *
 */
func getRow(rowIndex int) []int {

	size := 1
	var res []int
	var tmp []int
	k := 0
	for i := 0; i < size; i++ {

		if i == size-1 {
			tmp = append(tmp, 1)
			res = tmp
			size = len(res) + 1
			i = -1
			if k == rowIndex {
				return res
			}
			k++
			tmp = []int{}
		} else if i == 0 {
			tmp = append(tmp, 1)
		} else {
			tmp = append(tmp, res[i-1]+res[i])
		}
	}
	return res
}
