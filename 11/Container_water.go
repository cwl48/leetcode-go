package ContainerWater

/*
	写下思路
	1. 两个游标从数组两端不断往中间靠拢
	2. 计算 游标和对应的数组的乘积是否最大 即不断调整x,y  使得×*y的值最大
	3. y指代数组值，即为高度,x是两游标的差，即x轴的大小。
	4. 当所有的值都遍历之后，返回x*y最大值

*/
func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}
	var (
		maxValue int               //最大的容积
		arrMax   int               //数组中对应的最大值
		left     int               //左边的游标
		right    = len(height) - 1 //右边的游标
	)

	for right > left {
		var val int
		if height[right] < height[left] {
			arrMax = height[right]
		} else {
			arrMax = height[left]
		}
		val = arrMax * (right - left)
		if val > maxValue {
			maxValue = val
		}
		if height[right] > height[left] {
			left++
		} else {
			right--
		}
	}

	return maxValue
}
