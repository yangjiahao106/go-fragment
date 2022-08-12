package main

import (
	"fmt"
	"math"
)

// 前缀和 + 二分查找
func minSubArrayLen(target int, nums []int) int {
	ans := math.MaxInt32
	sum := 0
	sums := make([]int, len(nums))
	for idx, v := range nums {
		sum += v
		sums[idx] = sum
		if sum-target >= 0 {
			l := binarySearch(sums[:idx], sum-target)
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

// 滑动窗口
func minSubArrayLen2(target int, nums []int) int {
	ans := math.MaxInt32
	sum := 0
	l := 0
	for r, v := range nums {
		sum += v
		if sum >= target {
			for l < r && sum-nums[l] >= target {
				sum -= nums[l]
				l += 1
			}
			if r-l+1 < ans {
				ans = r - l + 1
			}
		}
	}
	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}

// 滑动窗口
func minSubArrayLen3(target int, nums []int) int {
	ans := math.MaxInt32
	sum := 0
	l := 0
	for r, v := range nums {
		sum += v
		for sum >= target {
			if r-l+1 < ans {
				ans = r - l + 1
			}
			sum -= nums[l]
			l += 1

		}
	}

	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}
