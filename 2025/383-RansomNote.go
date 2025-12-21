// https://leetcode.com/problems/ransom-note/?envType=study-plan-v2&envId=top-interview-150
// 🔺

package main

func canConstruct(ransomNote string, magazine string) bool {
	stock := make(map[int]int)

	// magazine の文字をカウント
	for i := 0; i < len(magazine); i++ {
		stock[int(magazine[i]-'a')]++ // aは基準点
	}

	// ransomNote の文字を消費
	for i := 0; i < len(ransomNote); i++ {
		idx := int(ransomNote[i] - 'a')
		stock[idx]--
		if stock[idx] < 0 {
			return false
		}
	}
	return true
}
