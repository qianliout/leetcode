package main

func main() {

}

func majorityElement(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res, cnt := nums[0], 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == res {
			cnt++
			continue
		}
		cnt--
		if cnt == 0 {
			res = nums[i]
			cnt = 1
		}
	}
	return res
}
