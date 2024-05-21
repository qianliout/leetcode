package main

func main() {

}

func firstUniqChar(s string) int {
	last := make(map[int32]int)
	for _, ch := range s {
		last[ch]++
	}
	for i, ch := range s {
		if last[ch] == 1 {
			return i
		}
	}
	return -1
}
