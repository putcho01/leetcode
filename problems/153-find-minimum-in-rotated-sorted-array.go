package main

//https://leetcode.com/problems/find-minimum-in-rotated-sorted-array

func findMin(nums []int) int {
	l := 0
	r := len(nums) - 1
	mid := l + (r-l)/2

	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] > nums[r] {
			l = mid + 1
		} else if nums[mid] < nums[r] {
			r = mid
		} else {
			return nums[mid]
		}
	}

	return nums[mid]
}
