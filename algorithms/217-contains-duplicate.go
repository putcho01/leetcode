package main

//
// https://leetcode.com/problems/contains-duplicate/
func containsDuplicate(nums []int) bool {
	q := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		q[nums[i]]++
		if q[nums[i]] > 1 {
			return true
		}
	}

	return false
}
