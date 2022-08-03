package main

//https://leetcode.com/problems/reshape-the-matrix/

func matrixReshape(mat [][]int, r int, c int) [][]int {
	matR := len(mat)
	matC := len(mat[0])
	if matR*matC != r*c {
		return mat
	}
	result := make([][]int, r)
	for i := 0; i < r; i++ {
		result[i] = make([]int, c)
		for j := 0; j < c; j++ {
			elemNum := i*c + j
			matI := elemNum / matC
			matJ := elemNum % matC
			result[i][j] = mat[matI][matJ]
		}
	}
	return result
}
