package nextPermutation

/*
	思路:这题的题意比较隐晦啊，网上找了别人的例子才明白题目的意思。。。。，知道题目的意思后，解题简单了很多。题意：就是把数组每个元素组合起来成为一个数字，比如[1,3,1,3,1],则为13131,然后用数组中的元素重新排列，得到比13131大一点的数，即13311，答案就是它.

	实现过程:
	1.倒序遍历数组，如果nums[i-1]<nums[i],则i-1该位置需要调换值，我们设index=i-1 2.然后再次倒序遍历，如果找到nums[i]>nums[index],调换两个数字，并且把index下标后的数字升序排列，得到答案。
	3.题目还要求没有比原始值大的下一个排列，则需要把该原始序列升序排列，即3,2,1，没有比他更大的数，则需要返回1,2,3
*/
func nextPermutation(nums []int) []int {

	arrLen := len(nums)
	if arrLen < 2 {
		return nums
	}
	index := -1
	for i := arrLen - 1; i > 0; i-- {
		if nums[i-1] < nums[i] {
			index = i - 1
			break
		}
	}
	if index == -1 {
		reserve(nums)
		return nums
	}

	for i := arrLen - 1; i > 0; i-- {
		if nums[i] > nums[index] {
			nums[i], nums[index] = nums[index], nums[i]
			reserve(nums[index+1:])
			break
		}
	}

	return nums
}

func reserve(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
