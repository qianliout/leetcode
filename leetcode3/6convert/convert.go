package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3))
	fmt.Println(convert("PAYPALISHIRING", 4))
	fmt.Println(convert("ab", 1))
}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	add, idx := 1, 0
	ans := make(map[int][]byte)
	ss := []byte(s)
	for i := 0; i < len(ss); i++ {
		if ans[idx] == nil {
			ans[idx] = make([]byte, 0)
		}

		ans[idx] = append(ans[idx], ss[i])
		idx = idx + add
		if idx == numRows-1 || idx == 0 {
			add = -add
		}
	}
	res := make([]string, 0)
	for i := 0; i < numRows; i++ {
		if len(ans[i]) > 0 {
			res = append(res, string(ans[i]))
		}
	}

	return strings.Join(res, "")
}
