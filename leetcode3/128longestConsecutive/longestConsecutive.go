package main

func main() {

}

func longestConsecutive(nums []int) int {
	mem := make(map[int]int, len(nums))
	for _, nu := range nums {
		mem[nu]++
	}
	ans := 0
	for _, nu := range nums {
		if mem[nu-1] > 0 { // 说明已计算过
			continue
		}

		// 说明从 nu开始，可能连续
		cur, lev := nu, 0
		for mem[cur] > 0 {
			cur = cur + 1
			lev = lev + 1
		}
		if lev > ans {
			ans = lev
		}
	}
	return ans
}
