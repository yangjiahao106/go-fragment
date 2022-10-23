package main

import "fmt"

func TestSearch() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 7))
}

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		m := (l-r)/2 + r
		if nums[m] == target {
			return m
		}

		if nums[m] > l {
			if nums[l] <= target && target < nums[m] {
				r = m - 1
			} else {
				l = m + 1
			}
		} else {
			if nums[m] < target && target <= nums[r] {
				l = m + 1
			} else {
				r = m - 1
			}
		}
	}

	return -1
}
