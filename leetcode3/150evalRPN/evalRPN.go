package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(evalRPN([]string{"2", "1", "+", "3", "*"}))
}

func evalRPN(tokens []string) int {
	numS := make([]int, 0)
	num := 0
	for _, to := range tokens {
		if to == "+" || to == "-" || to == "*" || to == "/" {
			if len(numS) < 2 {
				return num
			}
			last := numS[len(numS)-1]
			sec := numS[len(numS)-2]
			numS = numS[:len(numS)-2]

			switch to {
			case "+":
				numS = append(numS, last+sec)
			case "-":
				numS = append(numS, sec-last)
			case "*":
				numS = append(numS, last*sec)
			case "/":
				if last == 0 {
					return num
				}
				numS = append(numS, sec/last)
			}
			continue
		}
		if nu, err := strconv.Atoi(to); err == nil {
			numS = append(numS, nu)
		}
	}
	if len(numS) > 0 {
		return numS[0]
	}
	return 0
}
