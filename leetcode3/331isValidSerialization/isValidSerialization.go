package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isValidSerialization("9,3,4,#,#,1,#,#,2,#,6,#,#"))
	fmt.Println(isValidSerialization("#,#,#"))
}

func isValidSerialization(preorder string) bool {
	if preorder == "" {
		return true
	}
	data := strings.Split(preorder, ",")
	if len(data) == 0 || (len(data) == 1 && data[0] == "#") {
		return true
	}

	slot := 1
	for i := 0; i < len(data); i++ {
		if slot <= 0 {
			return false
		}
		slot -= 1
		if data[i] != "#" {
			slot += 2
		}
	}
	return slot == 0
}
