package main

func main() {

}

func singleNumber(nums []int) int {
	ans := 0
	for i := 0; i < 64; i++ {
		cnt := 0
		for _, ch := range nums {
			cnt += (ch >> i) & 1
		}
		if cnt%3 == 1 {
			ans = ans + (1 << i)
		}
	}
	return ans
}
