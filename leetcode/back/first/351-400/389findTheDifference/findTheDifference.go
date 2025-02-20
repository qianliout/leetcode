package main

func main() {

}

/*
给定两个字符串 s 和 t，它们只包含小写字母。
字符串 t 由字符串 s 随机重排，然后在随机位置添加一个字母。
请找出在 t 中被添加的字母。
示例:
输入：
s = "abcd"
t = "abcde"
输出：
e
解释：
'e' 是那个被添加的字母。
*/

func findTheDifference(s string, t string) byte {
	if len(s) == 0 {
		return []byte(t)[0]
	}
	sMap := make(map[byte]int)
	tMap := make(map[byte]int)

	for _, v := range []byte(s) {
		sMap[v]++
	}

	for _, v := range []byte(t) {
		tMap[v]++
	}
	for k, v := range tMap {
		if v-sMap[k] >= 1 {
			return k
		}
	}
	return 'a'
}
