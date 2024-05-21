package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(validIPAddress("20EE:FGb8:85a3:0:0:8A2E:0370:7334"))
}

func validIPAddress(queryIP string) string {
	if validaIPV6(queryIP) {
		return "IPv6"
	}
	if validaIPV4(queryIP) {
		return "IPv4"
	}
	return "Neither"

}

func validaIPV4(queryIP string) bool {
	split := strings.Split(queryIP, ".")
	if len(split) != 4 {
		return false
	}
	for _, ch := range split {
		n, err := strconv.Atoi(ch)
		if err != nil {
			return false
		}
		if strconv.Itoa(n) != ch {
			return false
		}
		if n < 0 || n > 255 {
			return false
		}
	}
	return true
}

func validaIPV6(queryIP string) bool {
	split := strings.Split(queryIP, ":")
	if len(split) != 8 {
		return false
	}
	for _, ch := range split {
		ch = strings.ToLower(ch)
		if len(ch) > 4 || len(ch) <= 0 {
			return false
		}
		for _, b := range ch {
			if (b >= '0' && b <= '9') || (b >= 'a' && b <= 'f') {
				continue
			} else {
				return false
			}
		}
	}
	return true
}
