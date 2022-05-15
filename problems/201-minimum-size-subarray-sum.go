package main

import "math"

// sliding window
//https://leetcode.com/problems/minimum-size-subarray-sum/

func minSubArrayLen(s int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	ans := math.MaxInt32
	sum := 0
	start := 0

	for end := 0; end < len(nums); end++ {
		sum += nums[end]
		for sum >= s {
			ans = min(end-start+1, ans)
			sum -= nums[start]
			start++
		}
	}

	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
