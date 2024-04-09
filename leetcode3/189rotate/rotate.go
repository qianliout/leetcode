package main

func main() {

}

func rotate2(nums []int, k int) {

	if len(nums) <= 1 || k == 0 {
		return
	}

	k = k % len(nums)
	if k == 0 {
		return
	}
	pre := nums[:len(nums)-k]
	after := make([]int, 0)

	after = append(after, nums[len(nums)-k:]...)
	after = append(after, pre...)

	for i := range after {
		nums[i] = after[i]
	}
}

func rotate(nums []int, k int) {
	if len(nums) <= 1 || k == 0 {
		return
	}

	k = k % len(nums)
	if k == 0 {
		return
	}
	rotateAll(nums)
	rotateAll(nums[:k])
	rotateAll(nums[k:])
}

func rotateAll(nums []int) {
	l, r := 0, len(nums)-1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}
