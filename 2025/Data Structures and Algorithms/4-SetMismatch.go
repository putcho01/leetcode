
// https://leetcode.com/problems/set-mismatch/?envType=problem-list-v2&envId=dsa-linear-shoal-array-ii

func findErrorNums(nums []int) []int {
  s := len(nums)
	numSet := make(map[int]bool)
	var duplicate int
	var missing int

	for _, num := range nums {
		if numSet[num] {
			// すでに出現している値の場合、重複した値として記録
			duplicate = num
		} else {
			// 出現済みの値を記録
			numSet[num] = true
		}
	}

	// 1 から s までの数値の中で出現していない値を探す
	// それが欠損した値となる
	for i := 1; i <= s; i++ {
		if !numSet[i] {
			missing = i
			break
		}
	}

	return []int{duplicate, missing}
}