/*
You are given an integer array prices where prices[i] is the price of a given stock on the ith day.

On each day, you may decide to buy and/or sell the stock.
You can only hold at most one share of the stock at any time.
However, you can sell and buy the stock multiple times on the same day, ensuring you never hold more than one share of the stock.

Find and return the maximum profit you can achieve.
*/

// Input: prices = [7,1,5,3,6,4]
// 7
package main

import "fmt"

func main() {
	stocks := []int{1, 1, 1, 1, 1, 10, 11, 2, 2, 2, 8}

	fmt.Println(maxProfit(stocks))
}

func maxProfit(prices []int) int {
	profit := 0
	buy := 0

	for i := 0; i < len(prices)-1; i++ {
		if prices[i] < prices[i+1] {
			buy = prices[i]
			profit += prices[i+1] - buy
		}
	}
	return profit
}
