package main

import "fmt"

// https://mp.weixin.qq.com/s/pCRscwKqQdYYN7M1Sia7xA

func canFinish(numCourses int, prerequisites [][]int) bool {

	rely := map[int][]int{} // 课程 to 依赖的课程
	grade := map[int]int{}  // 课程 to 入度

	for _, p := range prerequisites {
		rely[p[0]] = append(rely[p[0]], p[1])
		grade[p[1]] += 1 // 被依赖的 入度加1
	}

	queue := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if grade[i] == 0 { // 从入度为0 的开始学习
			queue = append(queue, i)
		}
	}

	count := 0
	for len(queue) > 0 {
		c := queue[0]
		queue = queue[1:]
		count++

		for _, v := range rely[c] {
			grade[v] -= 1
			if grade[v] == 0 {
				queue = append(queue, v)
			}
		}
	}

	if count == numCourses {
		return true
	}

	return false
}

// DFS
func canFinishV2(numCourses int, prerequisites [][]int) bool {
	var (
		edges   = make([][]int, numCourses)
		visited = make([]int, numCourses)
		result  []int
		valid   = true
		dfs     func(u int)
	)

	dfs = func(u int) {
		visited[u] = 1
		for _, v := range edges[u] {
			if visited[v] == 0 {
				dfs(v)
				if !valid {
					return
				}
			} else if visited[v] == 1 {
				valid = false
				return
			}
		}
		visited[u] = 2
		result = append(result, u)
	}

	for _, info := range prerequisites {
		edges[info[1]] = append(edges[info[1]], info[0])
	}

	for i := 0; i < numCourses && valid; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
	}

	return valid
}

func TestCanFinish() {
	fmt.Println(canFinish(3, [][]int{{1, 0}, {2, 1}}))
}
