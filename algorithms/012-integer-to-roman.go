package algorithms

import "strings"

// https://leetcode.com/problems/integer-to-roman/

func intToRoman(num int) string {
	vals := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	res := ""
	for i, val := range vals {
		res += strings.Repeat(romans[i], num/val)
		num %= val
	}
	return res
}
