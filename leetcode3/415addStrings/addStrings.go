package main

import (
	"fmt"
)

func main() {
	fmt.Println(addStrings("11", "123"))
	fmt.Println(addStrings("456", "77"))
}

func addStrings(num1 string, num2 string) string {
	if len(num1) < len(num2) {
		return addStrings(num2, num1)
	}
	n1, n2 := []byte(num1), []byte(num2)
	sub := len(n1) - len(n2)
	add := 0
	res := make([]byte, len(n1)+1)

	for i := len(n1) - 1; i >= 0; i-- {
		ans := int(n1[i]-'0') + add
		if i-sub >= 0 {
			ans += int(n2[i-sub] - '0')
		}
		if ans >= 10 {
			add = 1
			ans = ans - 10
		} else {
			add = 0
		}
		res[i+1] = byte(ans + '0')
	}
	if add == 0 {
		return string(res[1:])
	}
	res[0] = byte(add + '0')
	return string(res)
}
