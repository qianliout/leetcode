package main

func main() {

}

// 进制的思想
func rand10() int {
	for {
		n := (rand7()-1)*7 + rand7() - 1
		if n >= 1 && n <= 10 {
			return n
		}
	}
}

func rand7() int {
	return 0
}
