package main

func main() {

}

func removeDuplicateLetters(s string) string {
	ss := []byte(s)
	lastIndex := make(map[byte]int)
	for i, by := range ss {
		lastIndex[by] = i
	}

	window := make([]byte, 0)
	visit := make(map[byte]bool)
	for i := 0; i < len(ss); i++ {
		ch := ss[i]
		if visit[ss[i]] {
			continue
		}
		for len(window) > 0 && window[len(window)-1] > ch && lastIndex[window[len(window)-1]] > i {
			last := window[len(window)-1]
			visit[last] = false
			window = window[:len(window)-1]
		}
		window = append(window, ch)
		visit[ch] = true
	}
	return string(window)
}
