package main

import (
	"fmt"
	"sort"
)

func TestMerge() {

	ans := merge([][]int{{1, 88}, {5, 8}, {6,6}})
	fmt.Println(ans)
}

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	ans := make([][]int, 0)

	l := 0
	for r := 1; r < len(intervals); r++ {
		if intervals[r][0] <= intervals[l][1] {
			intervals[l][1] = max(intervals[l][1], intervals[r][1])
		} else {
			ans = append(ans, intervals[l])
			l = r
		}
	}

	ans = append(ans, intervals[l])

	return ans
}

// better
func merge2(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	ans := make([][]int, 0)

	for i := 0; i < len(intervals); i++ {
		if len(ans) > 0 && ans[len(ans)-1][1] >= intervals[i][0] {
			ans[len(ans)-1][1] = max(ans[len(ans)-1][1], intervals[i][1])
		} else {
			ans = append(ans, intervals[i])
		}
	}

	return ans
}
