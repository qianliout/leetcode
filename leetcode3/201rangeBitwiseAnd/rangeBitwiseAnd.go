package main

func main() {

}

// 结果正确，但是会 timeout
func rangeBitwiseAnd2(left int, right int) int {
	if left > right {
		left, right = right, left
	}
	if left == right {
		return left
	}
	ans := left
	for i := left + 1; i <= right; i++ {
		ans = ans & i
	}
	return ans
}
func rangeBitwiseAnd(left int, right int) int {
	if left > right {
		left, right = right, left
	}
	move := 0
	for left != right {
		left = left >> 1
		right = right >> 1
		move++
	}

	return left << move
}
