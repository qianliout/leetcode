package main

func main() {

}

func BucketSort(nums []int) {
	if len(nums) <= 1 {
		return
	}

	mi, ma := nums[0], nums[0]
	for _, ch := range nums {
		if ch < mi {
			mi = ch
		}
		if ch > ma {
			ma = ch
		}
	}
	// 确定一个桶中

}
