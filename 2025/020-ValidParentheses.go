package main

// Note: 後で確認する
func isValid(s string) bool {
	stack := make([]rune, 0)
	mapping := map[rune]rune{
		'{': '}',
		'(': ')',
		'[': ']',
	}
	for _, c := range s {
		if char, ok := mapping[c]; ok {
			stack = append(stack, char)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != c {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}
	return len(stack) == 0
}
