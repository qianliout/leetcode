package main

func main() {

}

func rand10() int {
	for {
		a := (rand7()-1)*7 + rand7() - 1
		if a >= 1 && a <= 10 {
			return a
		}
	}
}

func rand7() int {
	return 0
}
