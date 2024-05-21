package main

func main() {

}

func findMaximumXOR(nums []int) int {
	ans := 0
	return ans
}

// bi表示只考虑多少位
func find(a, b int, bi int) int {
	return (a >> bi) ^ (b >> bi)
}
