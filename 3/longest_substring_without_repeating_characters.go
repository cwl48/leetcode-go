package lengthoflongestsubstring

import (
	"strings"
)

/*
 * [3] Longest Substring Without Repeating Characters
 *
 * https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/description/
 *
 * algorithms
 * Medium (24.22%)
 * Total Accepted:    26.5K
 * Total Submissions: 109.4K
 * Testcase Example:  '"abcabcbb"'
 *
 * 给定一个字符串，找出不含有重复字符的最长子串的长度。
 *
 * 示例 1:
 *
 * 输入: "abcabcbb"
 * 输出: 3
 * 解释: 无重复字符的最长子串是 "abc"，其长度为 3。
 *
 *
 * 示例 2:
 *
 * 输入: "bbbbb"
 * 输出: 1
 * 解释: 无重复字符的最长子串是 "b"，其长度为 1。
 *
 *
 * 示例 3:
 *
 * 输入: "pwwkew"
 * 输出: 3
 * 解释: 无重复字符的最长子串是 "wke"，其长度为 3。
 * 请注意，答案必须是一个子串，"pwke" 是一个子序列 而不是子串。
 *
 *
 */
func lengthOfLongestSubstring(s string) int {

	if s == "" {
		return 0
	}
	max := 0
	current := string(s[0])

	for _, v := range s[1:] {

		i := strings.LastIndex(current, string(v))

		if i < 0 {
			current += string(v)
		} else {
			l := len(current)
			if l > max {
				max = l
			}
			current = current[i+1:] + string(v)
		}

	}
	if max >= len(current) {
		return max
	}
	return len(current)
}
