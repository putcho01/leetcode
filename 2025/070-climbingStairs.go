package main

// https://leetcode.com/problems/climbing-stairs/?envType=study-plan-v2&envId=top-interview-150
// 動的計画法 (DP法)

func climbStairs(n int) int {
	if n <= 1 {
		return 1
	}

	dp := make([]int, n+1)

	// 初期値
	dp[0] = 1
	dp[1] = 1
	// dp[i] をi 段目に到達する方法の数とする
	// dp[i] = dp[i-1] + dp[i-2]
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}
