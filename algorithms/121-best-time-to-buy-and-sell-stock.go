package main

//https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

func maxProfit(prices []int) int {
	profit, min := 0, prices[0]

	for i := 1; i < len(prices); i++ {
		if prices[i]-min > profit {
			profit = prices[i] - min
		} else if prices[i] < min {
			min = prices[i]
		}
	}

	return profit
}
