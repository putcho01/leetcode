package main

import "sort"

// https://leetcode.com/problems/majority-element/?envType=study-plan-v2&envId=top-interview-150
func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}
