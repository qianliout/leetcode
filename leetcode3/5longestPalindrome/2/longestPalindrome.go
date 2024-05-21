package main

func main() {

}

func longestPalindrome(s string) string {
	ans := ""
	ss := []byte(s)
	for i := range ss {
		a := expand(ss, i, i)
		if len(ans) < len(a) {
			ans = a
		}
		b := expand(ss, i, i+1)
		if len(ans) < len(b) {
			ans = b
		}
	}
	return ans
}

func expand(ss []byte, start1, start2 int) string {
	if start2 >= len(ss) {
		return ""
	}
	if start1 < 0 {
		return ""
	}
	for start1 >= 0 && start2 <= len(ss)-1 {
		if ss[start2] != ss[start1] {
			break
		}
		start1--
		start2++
	}

	return string(ss[start1+1 : start2])
}
