package threeSumClosest

import (
	"sort"
)

/*
	思路:这题的思路和15题3sum差不多，只是结果变成最接近某个数，而不是等于0，并且只需要返回最小的结果，斌不需要筛除相同的结果，所以该题在细节上的处理并没有之前那道题多。
	1.先排序，给结果一个初始值
	2.循环遍历数组，对当前下标索引之后的元素进行首尾遍历，即low,high
	3.当nums[i]+nums[low]+nums[high]-target的绝对值小于result-target的绝对值，则把当前的sum赋值给result
	4.当sum>target,则需要把sum的值减小，使得sum更接近target,则要把high--,当sum<target,low++
*/
func threeSumClosest(nums []int, target int) int {

	sort.Ints(nums)
	arrLen := len(nums)
	result := nums[0] + nums[1] + nums[arrLen-1]
	var sum int //三个数组相加的值
	if len(nums) < 3 {
		return 0
	}

	for i := 0; i < arrLen; i++ {
		low, high := i+1, arrLen-1
		for low < high {
			sum = nums[i] + nums[low] + nums[high]
			if sum > target {
				high--
			} else {
				low++
			}
			if abs(sum-target) < abs(result-target) {
				result = sum
			}
		}
	}

	return result

}

//求绝对值
func abs(a int) int {
	if a < 0 {
		return -a
	}
	if a == 0 {
		return 0
	}
	return a
}
