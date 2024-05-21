package main

func main() {

}

func minSubarray(nums []int, p int) int {
	sum := make([]int, len(nums)+1)
	for i := range nums {
		sum[i+1] = (sum[i] + nums[i] + p) % p
	}
	x := sum[len(nums)] // 所有数据相加除以 p 的余数

	if x == 0 {
		return 0
	}
	ans := len(nums)

	last := make(map[int]int)
	// 设所有元素的和为 x，去掉的元素和为 y,要使 x−y 能被 p 整除，这等价于满足,y 和 x同余，
	// last[i]=k 表示余数为 i 的前缀和上一次出现的位置是 j
	for i, ch := range sum {
		last[ch] = i
		// 同余原理
		if j, ok := last[(ch-x+p)%p]; ok {
			ans = min(ans, i-j)
		}
	}
	if ans == len(nums) {
		return -1
	}
	return ans
}
