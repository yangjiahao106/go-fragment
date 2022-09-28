package main

import "fmt"

func TestExchange() {
	fmt.Println(exchange([]int{1, 2, 3, 4, 5, 6, 7, 8}))
}

func exchange(nums []int) []int {
	l, r := 0, len(nums)-1
	for l < r {
		for l < r && nums[l]%2 == 1 {
			l++
		}
		for l < r && nums[r]%2 == 0 {
			r--
		}
		nums[l], nums[r] = nums[r], nums[l]
	}
	return nums
}
