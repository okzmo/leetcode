package main

import "fmt"

func candy(ratings []int) int {
	candies := []int{}
	for range ratings {
		candies = append(candies, 1)
	}
	for i := 1; i < len(ratings); i++ {
		if ratings[i-1] < ratings[i] {
			candies[i] += candies[i-1]
		}
	}
	fmt.Println(candies)
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i+1] < ratings[i] && candies[i] <= candies[i+1] {
			candies[i] = candies[i+1] + 1
		}
	}
	fmt.Println(candies)
	res := 0
	for _, v := range candies {
		res += v
	}
	return res
}

func main() {
	ratings := []int{1, 6, 10, 8, 7, 3, 2}
	fmt.Println(candy(ratings))
}
