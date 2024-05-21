package main

import (
	"fmt"
)

func main() {
	fmt.Println((-5) % (-2))
	fmt.Println(subarraysDivByK([]int{4, 5, 0, -2, -3, 1}, 5))
}

// timeout
func subarraysDivByK1(nums []int, k int) int {
	preSum := make([]int, len(nums)+1)
	for i, ch := range nums {
		preSum[i+1] = preSum[i] + ch
	}
	ans := 0
	for i := 1; i < len(preSum); i++ {
		for j := 0; j < i; j++ {
			if (preSum[i]-preSum[j])%k == 0 {
				ans++
			}
		}
	}

	return ans
}

func subarraysDivByK(nums []int, k int) int {
	pos := make(map[int]int)
	pos[0] = 1
	ans, sum := 0, 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		// modulus := sum % k 这样写会出错，为啥呢
		// 当被除数为负数时取模结果为负数，需要纠正:  (-5) % (-2) = -1
		modulus := (sum%k + k) % k
		ans += pos[modulus]
		pos[modulus]++
	}

	return ans
}
