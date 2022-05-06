package main

//https://leetcode.com/problems/search-a-2d-matrix/

func searchMatrix(matrix [][]int, target int) bool {
	rl := 0
	rr := len(matrix) - 1
	l := 0
	r := len(matrix[0]) - 1
	row := 0
	//行の先頭とtarget比較 行特定
	for rl <= rr {
		mid := rl + (rr-rl)/2
		if matrix[mid][r] < target {
			rl = mid + 1
		} else if matrix[mid][0] > target {
			rr = mid - 1
		} else {
			row = mid
			break
		}
	}
	//特定した行で列比較
	for l <= r {
		mid := l + (r-l)/2
		if matrix[row][mid] < target {
			l = mid + 1
		} else if matrix[row][mid] > target {
			r = mid - 1
		} else {
			return true
		}
	}
	return false
}
