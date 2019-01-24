package sortedsquares

import (
	"math"
	"sort"
)

/*
 * @lc app=leetcode id=977 lang=golang
 *
 * [977] Squares of a Sorted Array
 *
 * https://leetcode.com/problems/squares-of-a-sorted-array/description/
 *
 * algorithms
 * Easy (76.32%)
 * Total Accepted:    9.3K
 * Total Submissions: 12.2K
 * Testcase Example:  '[-4,-1,0,3,10]'
 *
 * Given an array of integers A sorted in non-decreasing order, return an array
 * of the squares of each number, also in sorted non-decreasing order.
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: [-4,-1,0,3,10]
 * Output: [0,1,9,16,100]
 *
 *
 *
 * Example 2:
 *
 *
 * Input: [-7,-3,2,3,11]
 * Output: [4,9,9,49,121]
 *
 *
 *
 *
 * Note:
 *
 *
 * 1 <= A.length <= 10000
 * -10000 <= A[i] <= 10000
 * A is sorted in non-decreasing order.
 *
 *
 *
 *
 */
func sortedSquares(A []int) []int {

	// 先平方
	for i, v := range A {
		r := int(math.Pow(float64(v), 2))
		A[i] = r
	}
	sort.Ints(A)

	return A
}
