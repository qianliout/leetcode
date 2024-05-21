package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(countAndSay(5))
	fmt.Println(say("1211"))
}

func countAndSay(n int) string {
	if n == 0 {
		return ""
	}
	if n == 1 {
		return "1"
	}
	left := countAndSay(n - 1)
	return say(left)
}

func say(left string) string {
	ss := []byte(left)
	if len(ss) == 0 {
		return ""
	}
	start := ss[0] - '0'
	cnt := 1
	ans := make([]string, 0)
	for i := 1; i < len(ss); i++ {
		if ss[i] == ss[i-1] {
			cnt++
			continue
		}

		ans = append(ans, fmt.Sprintf("%d%d", cnt, start))
		start = ss[i] - '0'
		cnt = 1
	}
	ans = append(ans, fmt.Sprintf("%d%d", cnt, start))
	return strings.Join(ans, "")
}
