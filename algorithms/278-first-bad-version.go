package main

// Binary Search
// https://leetcode.com/problems/first-bad-version/

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */
func firstBadVersion(n int) int {
	left, right := 0, n
	for left <= right {
		middle := (left + right) / 2
		if !isBadVersion(middle-1) && isBadVersion(middle) {
			return middle
		}
		if isBadVersion(middle) {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return -1
}
