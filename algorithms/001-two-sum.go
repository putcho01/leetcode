package main

//https://leetcode.com/problems/two-sum/

func twoSum(nums []int, target int) []int {
	indexMap := make(map[int]int)
	for currIndex, currNum := range nums {
		if requiredIdx, isPresent := indexMap[target-currNum]; isPresent {
			// 1. 現在の値 currNum に対して補数 complement = target - currNum が既にマップに存在するかを確認。
			// 2. 存在すれば、その補数のインデックス requiredIdx と現在のインデックス currIndex を返す（これで和が target になる）。
			return []int{requiredIdx, currIndex}
		}
		// indexMap[currNum] = currIndex
		// 補数チェック後に現在の値をマップに登録。これにより同じ要素を二度使うことを避けつつ、次の要素がこの値を補数として使えるようにする。
		indexMap[currNum] = currIndex
	}
	return []int{}
}

// 例（nums = [2,7,11,15], target = 9）

// indexMap = {}
// i=0, num=2: complement=7 は無い → map[2]=0
// i=1, num=7: complement=2 が map にあり requiredIdx=0 → return [0,1]
// 時間計算量 O(n)、空間計算量 O(n)。