package main

import "strconv"

//https://leetcode.com/problems/palindrome-number/

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	st := strconv.Itoa(x)
	//文字列長さ取得
	len := len(st)
	// length/2でfor文流して先頭と最後取得して比較
	for i := 0; i < len/2; i++ {
		head := st[i : i+1]
		bottom := st[len-i-1 : len-i]
		if head != bottom {
			return false
		}
	}
	return true
}
