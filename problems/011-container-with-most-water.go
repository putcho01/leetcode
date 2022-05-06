package main

func maxArea(height []int) int {
	// two pointer algorithm
	n := len(height)

	var l, result int
	r := n - 1

	for l < r {
		// Select the minimum of left and right pointer heights for area calculation
		currArea := min(height[l], height[r]) * (r - l)
		// Select the maximum of result so far and current area result
		result = max(result, currArea)

		// If the current l height is less than r height,
		// there is no use selecting l for next iteration because, width decreases, so area cant increase
		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
