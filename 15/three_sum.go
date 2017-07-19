package threeSum

import (
	"sort"
)

/*
	写下思路
	要从一个数组中获取三个数，使得x+y+z=0
	先对数组排序，for循环遍历数组，用0-nums【i】得到sum
	再对 数组i下标之后的数进行首尾遍历，当[low]+[high]==sum,则i,low,high就是一个答案
	当low和high是一个答案:对于low之后和【low】值相等的数进行跳过或者high之前和[high]值相等的数进行跳过
	如果不是答案，当low+high>sum,则low--,否则high++
*/
func threeSum(nums []int) [][]int {

	array := make([][]int, 0, 0)
	arr := make([]int, 0, 0)
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		low, high := i+1, len(nums)-1
		sum := 0 - nums[i]
		if i == 0 || i > 0 && nums[i] != nums[i-1] {
			for low < high {
				if sum == nums[low]+nums[high] {
					arr = []int{nums[i], nums[low], nums[high]}
					array = append(array, arr)
					for low < high && nums[low] == nums[low+1] {
						low++
					}
					for low < high && nums[high] == nums[high-1] {
						high--
					}
					low++
					high--
				} else if sum > nums[low]+nums[high] {
					low++
				} else {
					high--
				}
			}
		}

	}
	return array
}
