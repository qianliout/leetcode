package main

func main() {

}

func longestPalindrome(s string) int {
	exit := make(map[byte]int)
	for _, ch := range []byte(s) {
		exit[ch]++
	}
	maxOdd := 0
	ans := 0
	for _, cnt := range exit {
		if cnt&1 == 0 {
			ans += cnt
		} else {
			if maxOdd < cnt {
				maxOdd = cnt
			}
			ans += cnt - 1
		}
	}
	if maxOdd > 0 {
		ans++
	}
	return ans
}
