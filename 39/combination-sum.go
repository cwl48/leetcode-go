package combinationsum39

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 *
 * https://leetcode-cn.com/problems/combination-sum/description/
 *
 * algorithms
 * Medium (61.23%)
 * Total Accepted:    10.2K
 * Total Submissions: 16.6K
 * Testcase Example:  '[2,3,6,7]\n7'
 *
 * 给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
 *
 * candidates 中的数字可以无限制重复被选取。
 *
 * 说明：
 *
 *
 * 所有数字（包括 target）都是正整数。
 * 解集不能包含重复的组合。
 *
 *
 * 示例 1:
 *
 * 输入: candidates = [2,3,6,7], target = 7,
 * 所求解集为:
 * [
 * ⁠ [7],
 * ⁠ [2,2,3]
 * ]
 *
 *
 * 示例 2:
 *
 * 输入: candidates = [2,3,5], target = 8,
 * 所求解集为:
 * [
 * [2,2,2,2],
 * [2,3,3],
 * [3,5]
 * ]
 *
 */
func combinationSum(candidates []int, target int) [][]int {

	res := make([][]int, 0)
	tmp := make([]int, 0)
	sort.Ints(candidates)
	dfs(candidates, target, 0, &res, tmp)
	return res
}

func dfs(candidates []int, target int, index int, res *[][]int, tmp []int) {

	if target < 0 {
		return
	}

	if target == 0 {
		c := make([]int, len(tmp))
		copy(c, tmp)
		*res = append(*res, c)
		return
	}

	for i := index; i < len(candidates); i++ {

		if candidates[i] > target {
			return
		}
		tmp = append(tmp, candidates[i])
		dfs(candidates, target-candidates[i], i, res, tmp)
		tmp = tmp[:len(tmp)-1]
	}

}
