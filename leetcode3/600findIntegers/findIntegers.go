package main

func main() {

}

func findIntegers(n int) int {
	mem := make(map[int]int)
	return dfs(n, mem)
}

func dfs(n int, mem map[int]int) int {
	if n == 0 {
		return 1
	}
	if v, ok := mem[n]; ok {
		return v
	}
	res := dfs((n-1)>>2, mem) + dfs(n>>1, mem)
	mem[n] = res
	return res
}
