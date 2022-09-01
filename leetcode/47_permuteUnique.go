package main

import (
	"fmt"
	"sort"
)

func TestPermuteUnique() {
	ans := permuteUnique2([]int{2,2,1,1})
	fmt.Println(ans)
}

func permuteUnique(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	visit := make([]bool, len(nums))
	ans := make([][]int, 0)
	path := make([]int, 0)

	var helper func()
	helper = func() {
		if len(path) == len(nums) {
			ans = append(ans, append([]int{}, path...))
		}
		for k, v := range nums {
			if visit[k] {
				continue
			}
			if k > 0 && nums[k] == nums[k-1] && !visit[k-1] {
				continue
			}
			visit[k] = true
			path = append(path, v)
			helper()
			visit[k] = false
			path = path[:len(path)-1]
		}
	}

	helper()
	return ans
}



func permuteUnique2(nums []int) [][]int {
	dp := make([][]int, 0)
	dp = append(dp, []int{})

	for _, num := range nums {
		dp2 := make([][]int, 0)
		for _, path := range dp {
			for i := 0; i <= len(path); i++ {
				temp := append([]int{}, path[:i]...)
				temp = append(temp, num)
				temp = append(temp, path[i:]...)
				dp2 = append(dp2, temp)
				if i < len(path) && path[i] == num {
					break
				}
			}
		}
		fmt.Println(dp2)
		dp = dp2
	}

	return dp
}
