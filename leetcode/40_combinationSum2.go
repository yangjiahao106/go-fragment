package main

import (
	"fmt"
	"sort"
)

func TestCombinationSum2() {
	//  [10,1,2,7,6,1,5], target = 8,

	ans := combinationSum2([]int{10, 1, 1, 2, 2, 7, 6, 1, 5}, 18)
	fmt.Println(ans)
}

func combinationSum2(candidates []int, target int) [][]int {
	ans := make([][]int, 0)
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})

	var dfs func(s int, path []int, left int)
	dfs = func(s int, path []int, left int) {
		if left == 0 {
			ans = append(ans, append([]int{}, path...))
		}

		for i := s; i < len(candidates); i++ {
			if candidates[i] > left {
				break
			}
			if i > s && candidates[i-1] == candidates[i] {
				continue
			}
			dfs(i+1, append(path, candidates[i]), left-candidates[i])
		}
	}

	dfs(0, []int{}, target)
	return ans
}
