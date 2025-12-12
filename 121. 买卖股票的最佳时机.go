package main

// 贪心 / 动态规划（O(n)
func maxProfit(prices []int) int {
	minPrice := int(^uint(0) >> 1)
	maxProfit := 0
	for _, price := range prices {
		// 找出最小的price
		minPrice = min(price, minPrice)
		// 找出最大的收益
		maxProfit = max(maxProfit, price-minPrice)
	}
	return maxProfit
}
