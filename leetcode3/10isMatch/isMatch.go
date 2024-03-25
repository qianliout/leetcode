package main

func main() {

}

func isMatch(s string, p string) bool {
	ss, pp := []byte(s), []byte(p)
	return math(ss, pp)
}

func math(ss, pp []byte) bool {
	if len(pp) == 0 {
		return len(ss) == 0
	}
	// 不能加这一个判断
	// if len(ss) == 0 {
	// 	return len(pp) == 0 || (len(pp) == 1 && pp[0] == '*')
	// }
	first := len(ss) > 0 && (ss[0] == pp[0] || pp[0] == '.')

	if len(pp) >= 2 && pp[1] == '*' {
		return math(ss, pp[2:]) || first && math(ss[1:], pp)
	}

	return first && math(ss[1:], pp[1:])
}
