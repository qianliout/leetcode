package main

import (
	"fmt"
)

func main() {
	fmt.Println(isMatch("aa", "a"))
	fmt.Println(isMatch("aa", "a*"))
	fmt.Println(isMatch("ab", ".b"))
	fmt.Println(isMatch("aab", "c*a*b"))
	fmt.Println(isMatch("ab", ".*c"))
}

/*
'.' 匹配任意单个字符
'*' 匹配零个或多个前面的那一个元素
*/
func isMatch(s string, p string) bool {
	// 字符串的匹配一定要加这个
	s = " " + s
	p = " " + p
	return match([]byte(s), []byte(p))
}

func match(ss, pp []byte) bool {
	if len(pp) == 0 {
		return len(ss) == 0
	}
	// 先查看一下第一个字符
	// 这里 len(ss) > 0 容易忘记
	// 因为 ss="",pp="*"，也是可以的，所以 ss=="" 时还不能直接返回
	first := len(ss) > 0 && (ss[0] == pp[0] || pp[0] == '.')
	// 如果P第二个字符是 *，比如 p =="a*",会对第0号的匹配有影响，
	// 为啥在这一层做，而不等到下一层发出是*才做处理呢，因为如果等到下一层再做处理，就得要记录上一层的信息
	if len(pp) >= 2 && pp[1] == '*' {
		// 匹配0个
		a := match(ss, pp[2:])
		// 匹配多个,在这一层只匹配一个，如果还可以再匹配，就等到下一层做
		// 首先，ss[0]和pp[0]要匹配
		// b := match(ss[1:], pp) && first // 这种写法有问题，因为首先要 first==true 才行
		b := first && match(ss[1:], pp)
		return a || b
	}

	return first && match(ss[1:], pp[1:])
}
