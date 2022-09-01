package main

import "fmt"

func TestMaxProfit() {
	fmt.Println(maxProfit([]int{3, 3, 5, 0, 0, 3, 1, 4}))
}

func maxProfit(prices []int) int {

	s1 := -prices[0]
	s2 := 0
	s3 := -prices[0]
	s4 := 0

	for _, price := range prices {
		s1 = max(s1, -price)
		s2 = max(s2, s1+price)
		s3 = max(s3, s2-price)
		s4 = max(s4, s3+price)
	}

	return s4

}

//