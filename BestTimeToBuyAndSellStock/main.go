package main

import "fmt"

func maxProfit(prices []int) int {
	length := len(prices)
	profit := 0
	minPrice := prices[0]
	for i := 0; i < length; i++ {
		if minPrice > prices[i] {
			minPrice = prices[i]
		} else if prices[i]-minPrice > profit {
			profit = prices[i] - minPrice
		}
	}
	return profit
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	profit := maxProfit(prices)
	fmt.Println(profit)
}
