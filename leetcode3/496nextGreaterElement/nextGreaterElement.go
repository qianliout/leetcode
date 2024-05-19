package main

func main() {

}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	hash := make(map[int]int)
	stark := make([]int, 0)
	for i := len(nums2) - 1; i >= 0; i-- {
		for len(stark) > 0 && stark[len(stark)-1] <= nums2[i] {
			stark = stark[:len(stark)-1]
		}
		if len(stark) == 0 {
			hash[nums2[i]] = -1
		} else {
			hash[nums2[i]] = stark[len(stark)-1]
		}
		stark = append(stark, nums2[i])
	}
	ans := make([]int, 0)
	for _, ch := range nums1 {
		ans = append(ans, hash[ch])
	}
	return ans
}
