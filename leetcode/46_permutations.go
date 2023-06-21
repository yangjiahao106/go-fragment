package main

import "fmt"

func TestPermute() {

	ans := permute([]int{1, 2, 3})
	fmt.Println(ans)

}

// 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
func permute(nums []int) [][]int {
	res := make([][]int, 0)
	visited := make([]bool, len(nums))
	path := make([]int, 0)
	var dfs func()
	dfs = func() {
		if len(path) == len(nums) {
			res = append(res, append([]int{}, path...))
			return
		}
		for i, v := range nums {
			if visited[i] {
				continue
			}
			path = append(path, v)
			visited[i] = true
			dfs()
			visited[i] = false
			path = path[:len(path)-1]
		}
	}
	dfs()
	return res
}
