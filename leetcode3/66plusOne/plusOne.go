package main

func main() {

}

func plusOne(digits []int) []int {
	ans := make([]int, len(digits)+1)
	last := 1
	for i := len(digits) - 1; i >= 0; i-- {
		num := digits[i] + last
		if num >= 10 {
			last = num / 10
			num = num - 10
		} else {
			last = 0
		}
		ans[i+1] = num
	}
	if last > 0 {
		ans[0] = last
	}
	if ans[0] > 0 {
		return ans
	}
	return ans[1:]
}
