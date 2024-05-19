package main

func main() {

}

func minDeletionSize(strs []string) int {
	if len(strs) <= 1 {
		return 0
	}
	n := len(strs[0])
	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < len(strs)-1; j++ {
			if strs[j][i] > strs[j+1][i] {
				ans++
				break
			}
		}
	}
	return ans
}
