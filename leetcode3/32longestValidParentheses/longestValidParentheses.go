package main

func main() {

}

func longestValidParentheses(s string) int {
	// 具体做法是我们始终保持栈底元素为当前已经遍历过的元素中「最后一个没有被匹配的右括号的下标」，
	// 这样的做法主要是考虑了边界条件的处理，栈里其他元素维护左括号的下标：
	/*
		对于遇到的每个 "(" ，我们将它的下标放入栈中
		对于遇到的每个 ") ，我们先弹出栈顶元素表示匹配了当前右括号：
		如果栈为空，说明当前的右括号为没有被匹配的右括号，我们将其下标放入栈中来更新我们之前提到的「最后一个没有被匹配的右括号的下标」
		如果栈不为空，当前右括号的下标减去栈顶元素即为「以该右括号为结尾的最长有效括号的长度」
	*/
	stark := []int{-1}
	ss := []byte(s)
	ans := 0

	for i, ch := range ss {
		if ch == '(' {
			stark = append(stark, i)
			continue
		}

		stark = stark[:len(stark)-1]
		if len(stark) == 0 {
			stark = append(stark, i)
		}
		if len(stark) > 0 && ans < i-stark[len(stark)-1] {
			ans = i - stark[len(stark)-1]
		}
	}

	return ans
}
