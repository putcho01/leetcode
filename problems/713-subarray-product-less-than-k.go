package main

// sliding window
// https://leetcode.com/problems/subarray-product-less-than-k/

func numSubarrayProductLessThanK(nums []int, k int) int {
	counter := 0
	left := 0

	for right := 0; right < len(nums); right++ {

		currentProduct := 1
		left = right + 1
		currentProduct *= nums[right]

		if currentProduct < k {
			counter++
		}

		for currentProduct < k && left < len(nums) {

			currentProduct *= nums[left]
			left++

			if currentProduct < k {
				counter++
			}

		}

	}

	return counter

}
