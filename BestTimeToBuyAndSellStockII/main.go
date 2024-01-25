package main

import "fmt"

func maxProfit(prices []int) int {
	length := len(prices)
	profit := 0
	for i := 0; i < length-1; i++ {
		if prices[i] < prices[i+1] {
			profit += prices[i+1] - prices[i]
		}
	}
	return profit
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	profit := maxProfit(prices)
	fmt.Println(profit)
}
