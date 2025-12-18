
// https://leetcode.com/problems/set-mismatch/?envType=problem-list-v2&envId=dsa-linear-shoal-array-ii

func findErrorNums(nums []int) []int {
  s := len(nums)
	numSet := make(map[int]bool)
	var duplicate int
	var missing int

	for _, num := range nums {
		if numSet[num] {
			duplicate = num
		} else {
			numSet[num] = true
		}
	}

	for i := 1; i <= s; i++ {
		if !numSet[i] {
			missing = i
			break
		}
	}

	return []int{duplicate, missing}
}