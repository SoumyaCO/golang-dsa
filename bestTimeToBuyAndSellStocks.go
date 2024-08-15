package main

/*
* Best Time to buys and Sell Stock
* [Leetcode 121]
 */

import "math"

func maxProfit(prices []int) int {
	min := math.MaxUint32
	profit := 0

	for _, value := range prices {
		if value > min {
			if value-min > profit {
				profit = value - min
			}

		} else {
			min = value
		}
	}
	return profit
}
