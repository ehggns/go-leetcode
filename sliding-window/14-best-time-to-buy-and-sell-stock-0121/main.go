package main

// maxProfit calculates the maximum profit that can be made from buying and selling a stock once.
// It uses a sliding window approach to track the minimum price seen so far and the maximum profit.
//
// The function takes an array of prices where prices[i] represents the price of the stock on day i.
// It returns the maximum profit that can be achieved by buying on one day and selling on a later day.
// If no profit can be made, it returns 0.
//
// Time Complexity: O(n) where n is the length of the prices array
// Space Complexity: O(1)
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	// minPrice is initialized to the first price
	profit, minPrice := 0, prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else {
			profit = max(profit, prices[i]-minPrice) // current profit = prices[i] - minPrice
		}
	}
	return profit
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	result := maxProfit(prices)
	println(result) // Output: 5
}