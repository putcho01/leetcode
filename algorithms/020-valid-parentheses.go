package main

//https://leetcode.com/problems/valid-parentheses/

func isValid(s string) bool {
	stack := make([]rune, 0)
	m := map[rune]rune{
		'{': '}',
		'[': ']',
		'(': ')',
	}

	for _, c := range s {
		if char, ok := m[c]; ok {
			stack = append(stack, char)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != c {
				return false
			}
			stack = stack[:len(stack)-1]
		}

	}

	return len(stack) == 0
}
