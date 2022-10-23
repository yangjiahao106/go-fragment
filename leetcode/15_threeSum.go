package main

import (
	"fmt"
	"sort"
)

func TestThreeSum() {

	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -1, 0, 1, 1}))

}

func threeSum(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	var towSum func(nums []int, target int) [][]int
	towSum = func(nums []int, target int) [][]int {
		res := make([][]int, 0)
		if len(nums) < 2 {
			return res
		}
		l, r := 0, len(nums)-1
		for l < r {
			if nums[l]+nums[r] == target {
				res = append(res, []int{nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] { // 去重， 易错：先跳过重复的，再执行下面的l, r = l+1, r-1
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l, r = l+1, r-1

			} else if nums[l]+nums[r] < target {
				l++
			} else {
				r--
			}
		}
		return res
	}

	fmt.Println(towSum([]int{-1, -1, -1, 1, 1, 1}, 0))

	res := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i] > 0 { // 枝减
			break
		}
		subRes := towSum(nums[i+1:], - nums[i])
		for j := 0; j < len(subRes); j++ {
			res = append(res, append([]int{nums[i]}, subRes[j]...))
		}
	}

	return res
}
