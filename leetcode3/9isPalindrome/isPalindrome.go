package main

func main() {

}

func isPalindrome(x int) bool {

	if x < 0 {
		return false
	}

	num := 0
	pre := x
	for pre > 0 {
		b := pre % 10
		num = num*10 + b
		pre = pre / 10
	}

	return num == x
}
