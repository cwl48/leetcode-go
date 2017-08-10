package SumOfTwoIntegers

func getSum(a int, b int) int {

	var _a int
	var _b int
	for b != 0 {
		_a = a ^ b
		_b = (a & b) << 1
		a = _a
		b = _b
	}
	return a

}
