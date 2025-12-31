package main

// Note: 後で確認する
func isValid(s string) bool {
	stack := make([]rune, 0)
	mapping := map[rune]rune{
		'{': '}',
		'(': ')',
		'[': ']',
	}
	// 文字列を1文字ずつ処理
	for _, c := range s {
		if char, ok := mapping[c]; ok {
			stack = append(stack, char)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != c {
				return false
			} else {
				// 一致 → pop
				stack = stack[:len(stack)-1]
			}
		}
	}
	// 全部見終わって、スタックが空なら全部対応取れてる → true
	// 残ってたら閉じられてない括弧がある → false
	return len(stack) == 0
}
