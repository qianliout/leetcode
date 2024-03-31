package main

func main() {

}

func removeDuplicates(nums []int) int {
	start := 0
	for _, nu := range nums {
		if nu == nums[start] {
			continue
		}
		start++
		nums[start] = nu
	}
	return start + 1
}

func removeDuplicates2(nums []int) int {
	start := 0
	for _, nu := range nums {
		if nu != nums[start] {
			start++
			nums[start] = nu
		}

	}
	return start + 1
}
