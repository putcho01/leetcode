package main

// sliding window
//https://leetcode.com/problems/find-all-anagrams-in-a-string/

func findAnagrams(s string, p string) []int {
	m := make(map[[26]byte]int)
	temp := [26]byte{}

	var res []int
	left, right := 0, len(p)

	for _, val := range p {
		temp[val-'a']++
		print(val - 'a')
	}

	m[temp]++

	for right <= len(s) {
		temp = [26]byte{}
		for _, ch := range s[left:right] {
			temp[ch-'a']++
		}
		if _, found := m[temp]; found {
			res = append(res, left)
		}
		left++
		right++
	}
	return res
}
