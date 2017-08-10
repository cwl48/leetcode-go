package search

func search(nums []int, target int) int {

	//二分法查找某个数
	low, high := 0, len(nums)-1
	var mid int
	for low <= high {
		mid = (high-low)/2 + low
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < nums[high] {
			if nums[mid] < target && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		} else {
			if nums[mid] > target && target >= nums[low] {
				high = mid - 1
			} else {
				low = low + 1
			}
		}
	}
	return -1
}
