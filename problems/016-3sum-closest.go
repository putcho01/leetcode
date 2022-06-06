package main

import "sort"

//https://leetcode.com/problems/3sum-closest/

func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}
	minDiff := 10001
	sort.Ints(nums)
	// iterate over nums arr
	for i := 0; i < len(nums); i++ {
		// skip duplicates
		if i > 0 && nums[i] == nums[i-1] {
			i++
		}
		// binary search
		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[l] + nums[r] + nums[i]
			if abs(target-sum) < abs(target-minDiff) {
				minDiff = sum
			}
			if sum < target {
				l++
			} else if sum > target {
				r--
			} else {
				return target
			}
		}
	}
	return minDiff
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
