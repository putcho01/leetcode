package main

// Binary Search
//https://leetcode.com/problems/guess-number-higher-or-lower/
/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	left, right := 1, n
	for left <= right {
		mid := (left + right) / 2
		res := guess(mid)
		if res == 0 {
			return mid
		} else if res == -1 {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
