package merge56

import "sort"

/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 *
 * https://leetcode-cn.com/problems/merge-intervals/description/
 *
 * algorithms
 * Medium (33.78%)
 * Total Accepted:    8.5K
 * Total Submissions: 25K
 * Testcase Example:  '[[1,3],[2,6],[8,10],[15,18]]'
 *
 * 给出一个区间的集合，请合并所有重叠的区间。
 *
 * 示例 1:
 *
 * 输入: [[1,3],[2,6],[8,10],[15,18]]
 * 输出: [[1,6],[8,10],[15,18]]
 * 解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
 *
 *
 * 示例 2:
 *
 * 输入: [[1,4],[4,5]]
 * 输出: [[1,5]]
 * 解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。
 *
 */
// Definition for an interval.
type Interval struct {
	Start int
	End   int
}

func merge(intervals []Interval) []Interval {

	if len(intervals) == 0 {
		return make([]Interval, 0)
	}

	sort.Sort(Intervals(intervals))

	res := make([]Interval, 0)
	res = append(res, intervals[0])
	for _, v := range intervals[1:] {
		replace := false
		if v.Start <= res[len(res)-1].Start && v.End >= res[len(res)-1].Start {
			res[len(res)-1].Start = v.Start
			replace = true
		}
		if v.Start <= res[len(res)-1].End && v.End >= res[len(res)-1].End {
			res[len(res)-1].End = v.End
			replace = true
		}
		if v.Start >= res[len(res)-1].Start && v.End <= res[len(res)-1].End {
			replace = true
		}
		if !replace {
			res = append(res, v)
		}
	}
	return res
}

type Intervals []Interval

func (m Intervals) Less(i, j int) bool { return m[i].Start < m[j].Start }
func (m Intervals) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m Intervals) Len() int           { return len(m) }
