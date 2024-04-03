package main

func main() {

}

func grayCode(n int) []int {
	stark := make([]int, 0)
	for i := 0; i < 1<<n; i++ {
		stark = append(stark, i^(i>>1))
	}
	return stark
}
