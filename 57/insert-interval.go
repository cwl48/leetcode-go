package insert57

/*
 * @lc app=leetcode.cn id=57 lang=golang
 *
 * [57] 插入区间
 *
 * https://leetcode-cn.com/problems/insert-interval/description/
 *
 * algorithms
 * Hard (32.79%)
 * Total Accepted:    3.1K
 * Total Submissions: 9.3K
 * Testcase Example:  '[[1,3],[6,9]]\n[2,5]'
 *
 * 给出一个无重叠的 ，按照区间起始端点排序的区间列表。
 *
 * 在列表中插入一个新的区间，你需要确保列表中的区间仍然有序且不重叠（如果有必要的话，可以合并区间）。
 *
 * 示例 1:
 *
 * 输入: intervals = [[1,3],[6,9]], newInterval = [2,5]
 * 输出: [[1,5],[6,9]]
 *
 *
 * 示例 2:
 *
 * 输入: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
 * 输出: [[1,2],[3,10],[12,16]]
 * 解释: 这是因为新的区间 [4,8] 与 [3,5],[6,7],[8,10] 重叠。
 *
 *
 */

// Interval Definition for an interval.
type Interval struct {
	Start int
	End   int
}

func insert(intervals []Interval, newInterval Interval) []Interval {

	if intervals == nil || len(intervals) == 0 {
		return []Interval{newInterval}
	}

	added := false
	res := make([]Interval, 0, len(intervals)+1)
	for _, interval := range intervals {
		if added || interval.End < newInterval.Start {
			res = append(res, interval)
		} else if newInterval.End < interval.Start {
			res = append(res, newInterval, interval)
			added = true
		} else {
			start := interval.Start
			end := interval.End
			if newInterval.Start < interval.Start {
				start = newInterval.Start
			}
			if newInterval.End > interval.End {
				end = newInterval.End
			}
			newInterval.Start = start
			newInterval.End = end
		}
	}

	if !added {
		res = append(res, newInterval)
	}

	return res
}
