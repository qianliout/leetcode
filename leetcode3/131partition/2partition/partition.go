package main

func main() {

}

func partition(s string) [][]string {
	ans := make([][]string, 0)
	dfs([]byte(s), 0, []string{}, &ans)
	return ans
}

func dfs(ss []byte, start int, path []string, ans *[][]string) {
	if start == len(ss) {
		*ans = append(*ans, append([]string{}, path...))
		return
	}
	for i := start; i < len(ss); i++ {
		if check(ss[start : i+1]) {
			path = append(path, string(ss[start:i+1]))
			dfs(ss, i+1, path, ans)
			path = path[:len(path)-1]
		}
	}
}

func check(word []byte) bool {
	le, ri := 0, len(word)-1
	for le < ri {
		if word[le] != word[ri] {
			return false
		}
		le++
		ri--
	}
	return true
}
