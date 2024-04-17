package main

func main() {

}

func maxCoins(nums []int) int {
	nums = append(append([]int{1}, nums...), 1)

	dp := make([][]int, len(nums))
	for i := range dp {
		dp[i] = make([]int, len(nums))
	}

	for i := len(nums) - 1; i >= 0; i-- {
		for j := i + 1; j < len(nums); j++ {

		}
	}
	return 0
}
