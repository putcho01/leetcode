https://leetcode.com/problems/max-consecutive-ones/?envType=problem-list-v2&envId=dsa-linear-shoal-array-i

// Given a binary array nums, return the maximum number of consecutive 1's in the array.

func findMaxConsecutiveOnes(nums []int) int {
	current := 0
	maxCount := 0
  for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			current = 0
			continue
		}
		current++
		if current > maxCount {
			maxCount = current
		}
	}
	return maxCount
}

// メモ
// 配列を左から順に見る
// 1 が続いている間はカウントを増やす
// 0 が来たらカウントをリセット
// その途中で 今までの最大値 を更新する