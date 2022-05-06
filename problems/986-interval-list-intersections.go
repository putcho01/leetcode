package main

//https://leetcode.com/problems/interval-list-intersections/

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	result := make([][]int, 0)
	for i, j := 0, 0; i < len(firstList) && j < len(secondList); {
		start, end := getInterval(firstList[i], secondList[j])
		if start <= end {
			result = append(result, []int{start, end})
		}
		if firstList[i][1] == end {
			i++
		}
		if secondList[j][1] == end {
			j++
		}
	}
	return result

}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func getInterval(A []int, B []int) (int, int) {
	start := max(A[0], B[0])
	end := min(A[1], B[1])
	return start, end
}
