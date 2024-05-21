package main

func main() {

}

func palindrome(ss []byte) bool {
	if len(ss) == 0 {
		return true
	}
	l, r := 0, len(ss)-1
	for l < r {
		if ss[l] != ss[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func shortestPalindrome(s string) string {
	ss1 := []byte(s)
	ss2 := reverse(ss1)
	notSame := findNotSame(ss2)
	return string(append(notSame, ss1...))
}

func reverse(ss []byte) []byte {
	ans := make([]byte, 0)
	for i := len(ss) - 1; i >= 0; i-- {
		ans = append(ans, ss[i])
	}
	return ans
}

// 从后往前看，第一次回文串的地方
// 返回去除后面回文串的部分
func findNotSame(ss []byte) []byte {
	i := 0
	for i < len(ss) {
		if !palindrome(ss[i:]) {
			i++
		} else {
			break
		}
	}
	return ss[:i]
}
