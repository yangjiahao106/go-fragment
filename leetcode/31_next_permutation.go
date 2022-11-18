package main

import (
	"fmt"
	"sort"
)

/*
1 2 3
1 3 2
3 1 2

*/

func TestNextPermutation() {

	nums := []int{1, 2, 3}
	for i := 0; i < 10; i++ {
		nextPermutation(nums)
		fmt.Println(nums)
	}

}

func nextPermutation(nums []int) {
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i-1] < nums[i] {
			for j := len(nums) - 1; j >= i; j-- {
				if nums[j] > nums[i-1] {
					nums[i-1], nums[j] = nums[j], nums[i-1]
					tmp := nums[i:]
					sort.Slice(tmp, func(i, j int) bool {
						return tmp[i] < tmp[j]
					})
					return
				}
			}
		}
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
}


