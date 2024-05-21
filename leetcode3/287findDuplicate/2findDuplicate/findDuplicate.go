package main

func main() {

}

func findDuplicate(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for nums[i]-1 < len(nums) && nums[i]-1 >= 0 && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}

	for i, ch := range nums {
		if i+1 != ch {
			return ch
		}
	}
	return len(nums)
}
