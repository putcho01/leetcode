package main

// Binary Search
//https://leetcode.com/problems/search-in-rotated-sorted-array/

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] >= nums[left] {
			if target <= nums[mid] && target >= nums[left] {
				right = mid
			} else {
				left = mid + 1
			}
		} else {
			if target > nums[mid] && target < nums[left] {
				left = mid + 1
			} else {
				right = mid
			}
		}
	}

	if left >= 0 && left <= len(nums)-1 && nums[left] == target {
		return left
	}
	if right >= 0 && right <= len(nums)-1 && nums[right] == target {
		return right
	}
	return -1
}
