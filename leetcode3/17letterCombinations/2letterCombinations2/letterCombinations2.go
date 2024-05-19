package main

func main() {

}

func letterCombinations2(digits string) []string {
	if digits == "" {
		return []string{}
	}
	digitMap := map[byte][]byte{
		'2': []byte("abc"),
		'3': []byte("def"),
		'4': []byte("ghi"),
		'5': []byte("jkl"),
		'6': []byte("mno"),
		'7': []byte("pqrs"),
		'8': []byte("tuv"),
		'9': []byte("wxyz"),
	}
	ans := make([]string, 0)

	dfs(digitMap, []byte(digits), 0, []byte{}, &ans)
	return ans
}

func dfs(digitMap map[byte][]byte, digits []byte, start int, path []byte, ans *[]string) {
	if start == len(digits) {
		*ans = append(*ans, string(path))
		return
	}
	cc := digitMap[digits[start]]
	for i := range cc {
		path = append(path, cc[i])
		dfs(digitMap, digits, start+1, path, ans)
		path = path[:len(path)-1]
	}
}
