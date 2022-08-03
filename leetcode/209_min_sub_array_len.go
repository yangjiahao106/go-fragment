package main

import (
	"fmt"
	"math"
)

func minSubArrayLen(target int, nums []int) int {
	ans := math.MaxInt32
	sum := 0
	sums := make([]int, len(nums))
	for idx, v := range nums {
		sum += v
		sums[idx] = sum
		fmt.Println(sums)
		if sum-target >= 0 {
			l := binarySearch(sums[:idx], sum-target)
			fmt.Println(l)
			if idx-l < ans {
				ans = idx - l
			}
			if ans == 1 {
				return ans
			}
		}
	}

	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}

func binarySearch(nums []int, target int) int {
	fmt.Println("search", nums, target)
	l, r := 0, len(nums)-1
	for l <= r {
		m := (l + r) / 2
		if nums[m] == target {
			return m
		}
		if nums[m] < target {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return r
}
