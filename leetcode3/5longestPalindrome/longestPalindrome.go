package main

func main() {

}

func longestPalindrome(s string) string {
	ss := []byte(s)
	if len(ss) <= 1 {
		return s
	}
	ans := ""
	for i := 0; i < len(ss); i++ {
		r := expand(ss, i, i)
		if len(r) > len(ans) {
			ans = r
		}
		r1 := expand(ss, i, i+1)
		if len(r1) > len(ans) {
			ans = r1
		}
	}
	return ans
}

func expand(ss []byte, l, r int) string {
	for l >= 0 && r < len(ss) {
		if ss[l] == ss[r] {
			l--
			r++
		} else {
			break
		}
	}

	return string(ss[l+1 : r])
}
