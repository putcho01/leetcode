package main

func singleNumber(nums []int) int {
	result := 0
	for _, num := range nums {
		result ^= num
	}
	return result
}

// 例: nums = [4, 1, 2, 1, 2]
// 4 ^ 1 ^ 2 ^ 1 ^ 2
// = 4 ^ (1 ^ 1) ^ (2 ^ 2)
// = 4 ^ 0 ^ 0
// = 4
